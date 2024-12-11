# Copyright 2022 Cisco Systems, Inc. and its affiliates
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

import asyncio
import atexit
import logging
import sys
from typing import Optional

from flame.backends import backend_provider
from flame.channel import Channel
from flame.common.util import run_async
from flame.config import Config
from flame.selectors import selector_provider

logger = logging.getLogger(__name__)


def custom_excepthook(exc_type, exc_value, exc_traceback):
    """Implement a custom exception hook."""
    logger.critical(
        "Uncaught exception:", exc_info=(exc_type, exc_value, exc_traceback)
    )


sys.excepthook = custom_excepthook


class ChannelManager(object):
    """ChannelManager manages channels and creates a singleton instance."""

    _instance = None

    _config = None
    _job_id = None
    _role = None

    _channels = None

    _loop = None

    _backend = None  # default backend in case there is no per-channel backend
    _backends = dict()  # backend per channel

    def __new__(cls):
        """Create a singleton instance."""
        if cls._instance is None:
            logger.info("Creating a ChannelManager instance")
            cls._instance = super(ChannelManager, cls).__new__(cls)
        return cls._instance

    def __call__(self, config: Config):
        """Initialize instance variables."""
        self._config = config
        self._job_id = self._config.job.job_id
        self._role = self._config.role
        self._task_id = self._config.task_id

        self._channels = {}

        self._setup_backends()

        atexit.register(self.cleanup)

    def _setup_backends(self):
        distinct_backends = {}

        for ch_name, channel in self._config.channels.items():
            sort = channel.backend
            if not sort:
                logger.warning(f"No backend specified for channel {ch_name}.")
                continue

            if sort not in distinct_backends:
                backend = backend_provider.get(sort)
                broker_host = channel.broker_host or self._config.brokers.sort_to_host.get(sort)

                backend.configure(broker_host, self._job_id, self._task_id)

                distinct_backends[sort] = backend

            self._backends[ch_name] = distinct_backends[sort]

        if len(self._backends) == len(self._config.channels):
            return

        sort = self._config.backend
        if sort not in distinct_backends:
            self._backend = backend_provider.get(sort)
            broker_host = self._config.brokers.sort_to_host.get(sort)
            self._backend.configure(broker_host, self._job_id, self._task_id)
        else:
            self._backend = distinct_backends[sort]

    def join_all(self) -> None:
        """join_all ensures that a role joins all of its channels."""
        for ch_name in self._config.channels.keys():
            self.join(ch_name)

    def join(self, name: str) -> bool:
        """Join a channel."""
        if self.is_joined(name):
            return True

        channel_config = self._config.channels[name]

        if self._role == channel_config.pair[0]:
            me = channel_config.pair[0]
            other = channel_config.pair[1]
        else:
            me = channel_config.pair[1]
            other = channel_config.pair[0]

        groupby = channel_config.group_by.groupable_value(
            self._config.group_association.get(name)
        )

        selector = selector_provider.get(
            self._config.selector.sort, **self._config.selector.kwargs
        )

        if name in self._backends:
            backend = self._backends[name]
        else:
            logger.info(f"No backend found for channel {name}; using default backend.")
            backend = self._backend

        self._channels[name] = Channel(
            backend, selector, self._job_id, name, me, other, groupby
        )
        self._channels[name].join()

    def leave(self, name):
        """Leave a channel."""
        if not self.is_joined(name):
            return

        self._channels[name].leave()
        del self._channels[name]

    def get_by_tag(self, tag: str) -> Optional[Channel]:
        """Return a channel object that matches a given function tag."""
        if tag not in self._config.func_tag_map:
            return None

        channel_name = self._config.func_tag_map[tag]
        logger.debug(f"{tag} through {channel_name}")
        return self.get(channel_name)

    def get(self, name: str) -> Optional[Channel]:
        """Return a channel object in a given channel name."""
        if not self.is_joined(name):
            return None

        return self._channels[name]

    def is_joined(self, name):
        """Check if node joined a channel or not."""
        return True if name in self._channels else False

    def cleanup(self):
        """Clean up pending asyncio tasks."""
        logger.debug("Calling cleanup")
        for _, ch in self._channels.items():
            logger.debug(f"Calling leave for channel {ch.name()}")
            ch.leave()

        async def _inner(backend):
            for task in asyncio.all_tasks():
                task.cancel()
                try:
                    await task
                except asyncio.CancelledError:
                    logger.debug(f"Successfully cancelled {task.get_name()}")

            logger.debug("Done cleaning up asyncio tasks")

        if self._backend:
            _ = run_async(_inner(self._backend), self._backend.loop())

        for k, v in self._backends.items():
            _ = run_async(_inner(v), v.loop())