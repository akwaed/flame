/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A JobsApiController binds http requests to an api service and writes the service results to the http response
type JobsApiController struct {
	service JobsApiServicer
}

// NewJobsApiController creates a default api controller
func NewJobsApiController(s JobsApiServicer) Router {
	return &JobsApiController{service: s}
}

// Routes returns all of the api route for the JobsApiController
func (c *JobsApiController) Routes() Routes {
	return Routes{
		{
			"ChangeJobSchema",
			strings.ToUpper("Post"),
			"/{user}/jobs/{jobId}/schema/{schemaId}/design/{designId}",
			c.ChangeJobSchema,
		},
		{
			"CreateJob",
			strings.ToUpper("Post"),
			"/{user}/jobs",
			c.CreateJob,
		},
		{
			"DeleteJob",
			strings.ToUpper("Delete"),
			"/{user}/jobs/{jobId}",
			c.DeleteJob,
		},
		{
			"GetJob",
			strings.ToUpper("Get"),
			"/{user}/jobs/{jobId}",
			c.GetJob,
		},
		{
			"GetJobStatus",
			strings.ToUpper("Get"),
			"/{user}/jobs/{jobId}/status",
			c.GetJobStatus,
		},
		{
			"GetJobsStatus",
			strings.ToUpper("Get"),
			"/{user}/jobs/status",
			c.GetJobsStatus,
		},
		{
			"UpdateJob",
			strings.ToUpper("Put"),
			"/{user}/jobs/{jobId}",
			c.UpdateJob,
		},
		{
			"UpdateJobStatus",
			strings.ToUpper("Put"),
			"/{user}/jobs/{jobId}/status",
			c.UpdateJobStatus,
		},
	}
}

// ChangeJobSchema - Change the schema for the given job
func (c *JobsApiController) ChangeJobSchema(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	jobId := params["jobId"]

	schemaId := params["schemaId"]

	designId := params["designId"]

	result, err := c.service.ChangeJobSchema(r.Context(), user, jobId, schemaId, designId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateJob - Create a new job specification
func (c *JobsApiController) CreateJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	jobSpec := &JobSpec{}
	if err := json.NewDecoder(r.Body).Decode(&jobSpec); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.CreateJob(r.Context(), user, *jobSpec)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteJob - Delete job specification
func (c *JobsApiController) DeleteJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	jobId := params["jobId"]

	result, err := c.service.DeleteJob(r.Context(), user, jobId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetJob - Get a job specification
func (c *JobsApiController) GetJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	jobId := params["jobId"]

	result, err := c.service.GetJob(r.Context(), user, jobId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetJobStatus - Get job status of a given jobId
func (c *JobsApiController) GetJobStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	jobId := params["jobId"]

	result, err := c.service.GetJobStatus(r.Context(), user, jobId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetJobsStatus - Get status info on all the jobs owned by user
func (c *JobsApiController) GetJobsStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	user := params["user"]

	limit, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.GetJobsStatus(r.Context(), user, limit)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateJob - Update a job specification
func (c *JobsApiController) UpdateJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	jobId := params["jobId"]

	jobSpec := &JobSpec{}
	if err := json.NewDecoder(r.Body).Decode(&jobSpec); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.UpdateJob(r.Context(), user, jobId, *jobSpec)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateJobStatus - Update a job's status
func (c *JobsApiController) UpdateJobStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	jobId := params["jobId"]

	jobStatus := &JobStatus{}
	if err := json.NewDecoder(r.Body).Decode(&jobStatus); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.UpdateJobStatus(r.Context(), user, jobId, *jobStatus)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
