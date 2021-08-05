/*
 * Job REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"

	"wwwin-github.cisco.com/eti/fledge/cmd/controller/database"
	"wwwin-github.cisco.com/eti/fledge/pkg/objects"
)

// JobsApiService is a service that implents the logic for the JobsApiServicer
// This service should implement the business logic for every endpoint for the JobsApi API.
// Include any external packages or services that will be required by this service.
type JobsApiService struct {
}

// NewJobsApiService creates a default api service
func NewJobsApiService() JobsApiServicer {
	return &JobsApiService{}
}

// GetJobs - Get list of all the jobs by the user or based on designId.
func (s *JobsApiService) GetJobs(ctx context.Context, user string, designId string, getType string, limit int32) (objects.ImplResponse, error) {
	jInfo, err := database.GetJobs(user, getType, designId, limit)
	if err != nil {
		return objects.Response(http.StatusInternalServerError, nil), errors.New("get list of jobs request failed")
	}
	return objects.Response(http.StatusOK, jInfo), nil
}
