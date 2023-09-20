// Copyright 2022 Cisco Systems, Inc. and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Flame REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/cisco-open/flame/pkg/openapi/constants"
)

// JobsApiController binds http requests to an api service and writes the service results to the http response
type JobsApiController struct {
	service      JobsApiServicer
	errorHandler ErrorHandler
}

// JobsApiOption for how the controller is set up.
type JobsApiOption func(*JobsApiController)

// WithJobsApiErrorHandler inject ErrorHandler into controller
func WithJobsApiErrorHandler(h ErrorHandler) JobsApiOption {
	return func(c *JobsApiController) {
		c.errorHandler = h
	}
}

// NewJobsApiController creates a default api controller
func NewJobsApiController(s JobsApiServicer, opts ...JobsApiOption) Router {
	controller := &JobsApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the JobsApiController
func (c *JobsApiController) Routes() Routes {
	return Routes{
		{
			"CreateJob",
			strings.ToUpper("Post"),
			"/users/{user}/jobs",
			c.CreateJob,
		},
		{
			"DeleteJob",
			strings.ToUpper("Delete"),
			"/users/{user}/jobs/{jobId}",
			c.DeleteJob,
		},
		{
			"GetJob",
			strings.ToUpper("Get"),
			"/users/{user}/jobs/{jobId}",
			c.GetJob,
		},
		{
			"GetJobStatus",
			strings.ToUpper("Get"),
			"/users/{user}/jobs/{jobId}/status",
			c.GetJobStatus,
		},
		{
			"GetJobs",
			strings.ToUpper("Get"),
			"/users/{user}/jobs",
			c.GetJobs,
		},
		{
			"GetJobsByCompute",
			strings.ToUpper("Get"),
			"/jobs/{computeId}",
			c.GetJobsByCompute,
		},
		{
			"GetTask",
			strings.ToUpper("Get"),
			"/jobs/{jobId}/{taskId}/task",
			c.GetTask,
		},
		{
			"GetTaskInfo",
			strings.ToUpper("Get"),
			"/users/{user}/jobs/{jobId}/tasks/{taskId}",
			c.GetTaskInfo,
		},
		{
			"GetTasksInfo",
			strings.ToUpper("Get"),
			"/users/{user}/jobs/{jobId}/tasks",
			c.GetTasksInfo,
		},
		{
			"UpdateJob",
			strings.ToUpper("Put"),
			"/users/{user}/jobs/{jobId}",
			c.UpdateJob,
		},
		{
			"UpdateJobStatus",
			strings.ToUpper("Put"),
			"/users/{user}/jobs/{jobId}/status",
			c.UpdateJobStatus,
		},
		{
			"UpdateTaskStatus",
			strings.ToUpper("Put"),
			"/jobs/{jobId}/{taskId}/task/status",
			c.UpdateTaskStatus,
		},
	}
}

// CreateJob - Create a new job specification
func (c *JobsApiController) CreateJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params[constants.ParamUser]

	jobSpecParam := JobSpec{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&jobSpecParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertJobSpecRequired(jobSpecParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateJob(r.Context(), userParam, jobSpecParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteJob - Delete job specification
func (c *JobsApiController) DeleteJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params[constants.ParamUser]

	jobIdParam := params[constants.ParamJobID]

	result, err := c.service.DeleteJob(r.Context(), userParam, jobIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetJob - Get a job specification
func (c *JobsApiController) GetJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params[constants.ParamUser]

	jobIdParam := params[constants.ParamJobID]

	result, err := c.service.GetJob(r.Context(), userParam, jobIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetJobStatus - Get job status of a given jobId
func (c *JobsApiController) GetJobStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params[constants.ParamUser]

	jobIdParam := params[constants.ParamJobID]

	result, err := c.service.GetJobStatus(r.Context(), userParam, jobIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetJobs - Get status info on all the jobs owned by user
func (c *JobsApiController) GetJobs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	userParam := params[constants.ParamUser]

	limitParam, err := parseInt32Parameter(query.Get(constants.ParamLimit), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetJobs(r.Context(), userParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetJobsByCompute - Get status info on all the jobs by compute
func (c *JobsApiController) GetJobsByCompute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	computeId := params[constants.ParamComputeID]

	result, err := c.service.GetJobsByCompute(r.Context(), computeId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetTask - Get a job task for a given job and task
func (c *JobsApiController) GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	jobIdParam := params[constants.ParamJobID]

	taskIdParam := params[constants.ParamTaskID]

	keyParam := query.Get(constants.ParamKey)
	result, err := c.service.GetTask(r.Context(), jobIdParam, taskIdParam, keyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}

	mediatype, _, err := mime.ParseMediaType(r.Header.Get("Accept"))
	if err != nil {
		result.Code = http.StatusNotAcceptable
		c.errorHandler(w, r, err, &result)
		return
	}
	if mediatype != "multipart/form-data" {
		result.Code = http.StatusMultipleChoices
		c.errorHandler(w, r, err, &result)
		return
	}

	mw := multipart.NewWriter(w)
	w.Header().Set("Content-Type", mw.FormDataContentType())

	taskMap, ok := result.Body.(map[string][]byte)
	if !ok {
		result.Code = http.StatusNotFound
		c.errorHandler(w, r, err, &result)
		return
	}

	w.Header().Set("Content-Type", mw.FormDataContentType())
	for filename, data := range taskMap {
		fw, err := mw.CreateFormFile(filename, filename)
		if err != nil {
			result.Code = http.StatusInternalServerError
			c.errorHandler(w, r, err, &result)
			return
		}
		if _, err := fw.Write(data); err != nil {
			result.Code = http.StatusInternalServerError
			c.errorHandler(w, r, err, &result)
			return
		}
	}

	if err := mw.Close(); err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}

	w.WriteHeader(result.Code)
}

// GetTaskInfo - Get the info of a task in a job
func (c *JobsApiController) GetTaskInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params[constants.ParamUser]

	jobIdParam := params[constants.ParamJobID]

	taskIdParam := params[constants.ParamTaskID]

	result, err := c.service.GetTaskInfo(r.Context(), userParam, jobIdParam, taskIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetTasksInfo - Get the info of tasks in a job
func (c *JobsApiController) GetTasksInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	userParam := params[constants.ParamUser]

	jobIdParam := params[constants.ParamJobID]

	limitParam, err := parseInt32Parameter(query.Get(constants.ParamLimit), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	genericParam, err := parseBoolParameter(query.Get(constants.ParamGeneric), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	var result ImplResponse
	if genericParam {
		result, err = c.service.GetTasksInfoGeneric(r.Context(), userParam, jobIdParam, limitParam)
	} else {
		result, err = c.service.GetTasksInfo(r.Context(), userParam, jobIdParam, limitParam)
	}
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateJob - Update a job specification
func (c *JobsApiController) UpdateJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	jobSpecParam := JobSpec{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&jobSpecParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	if err := AssertJobSpecRequired(jobSpecParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}

	userParam := params[constants.ParamUser]
	jobIdParam := params[constants.ParamJobID]

	result, err := c.service.UpdateJob(r.Context(), userParam, jobIdParam, jobSpecParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateJobStatus - Update the status of a job
//
//nolint:dupl
func (c *JobsApiController) UpdateJobStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params[constants.ParamUser]

	jobIdParam := params[constants.ParamJobID]

	jobStatusParam := JobStatus{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&jobStatusParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertJobStatusRequired(jobStatusParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateJobStatus(r.Context(), userParam, jobIdParam, jobStatusParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateTaskStatus - Update the status of a task
func (c *JobsApiController) UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	taskStatusParam := TaskStatus{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&taskStatusParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	if err := AssertTaskStatusRequired(taskStatusParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}

	params := mux.Vars(r)
	jobIdParam := params[constants.ParamJobID]
	taskIdParam := params[constants.ParamTaskID]

	result, err := c.service.UpdateTaskStatus(r.Context(), jobIdParam, taskIdParam, taskStatusParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}

	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
