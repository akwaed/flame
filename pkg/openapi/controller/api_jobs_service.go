/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package controller

import (
	"context"
	"errors"
	"net/http"

	"wwwin-github.cisco.com/eti/fledge/pkg/openapi"
)

// JobsApiService is a service that implents the logic for the JobsApiServicer
// This service should implement the business logic for every endpoint for the JobsApi API.
// Include any external packages or services that will be required by this service.
type JobsApiService struct {
}

// NewJobsApiService creates a default api service
func NewJobsApiService() openapi.JobsApiServicer {
	return &JobsApiService{}
}

// ChangeJobSchema - Change the schema for the given job
func (s *JobsApiService) ChangeJobSchema(ctx context.Context, user string, jobId string, schemaId string,
	designId string) (openapi.ImplResponse, error) {
	// TODO - update ChangeJobSchema with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	//return Response(201, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ChangeJobSchema method not implemented")
}

// CreateJob - Create a new job specification
func (s *JobsApiService) CreateJob(ctx context.Context, user string, jobSpec openapi.JobSpec) (openapi.ImplResponse, error) {
	// TODO - update CreateJob with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	//return Response(201, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("CreateJob method not implemented")
}

// DeleteJob - Delete job specification
func (s *JobsApiService) DeleteJob(ctx context.Context, user string, jobId string) (openapi.ImplResponse, error) {
	// TODO - update DeleteJob with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("DeleteJob method not implemented")
}

// GetJob - Get a job specification
func (s *JobsApiService) GetJob(ctx context.Context, user string, jobId string) (openapi.ImplResponse, error) {
	// TODO - update GetJob with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, JobSpec{}) or use other options such as http.Ok ...
	//return Response(200, JobSpec{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetJob method not implemented")
}

// GetJobStatus - Get job status of a given jobId
func (s *JobsApiService) GetJobStatus(ctx context.Context, user string, jobId string) (openapi.ImplResponse, error) {
	// TODO - update GetJobStatus with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, JobStatus{}) or use other options such as http.Ok ...
	//return Response(200, JobStatus{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetJobStatus method not implemented")
}

// GetJobsStatus - Get status info on all the jobs owned by user
func (s *JobsApiService) GetJobsStatus(ctx context.Context, user string, limit int32) (openapi.ImplResponse, error) {
	// TODO - update GetJobsStatus with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []JobStatus{}) or use other options such as http.Ok ...
	//return Response(200, []JobStatus{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetJobsStatus method not implemented")
}

// UpdateJob - Update a job specification
func (s *JobsApiService) UpdateJob(ctx context.Context, user string, jobId string,
	jobSpec openapi.JobSpec) (openapi.ImplResponse, error) {
	// TODO - update UpdateJob with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("UpdateJob method not implemented")
}

// UpdateJobStatus - Update a job&#39;s status
func (s *JobsApiService) UpdateJobStatus(ctx context.Context, user string, jobId string,
	jobStatus openapi.JobStatus) (openapi.ImplResponse, error) {
	// TODO - update UpdateJobStatus with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("UpdateJobStatus method not implemented")
}
