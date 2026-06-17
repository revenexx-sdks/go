package models

import (
    "encoding/json"
    "errors"
)

// Execution Model
type Execution struct {
    // Execution creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Execution ID.
    Id string `json:"$id"`
    // Execution roles.
    Permissions []string `json:"$permissions"`
    // Execution update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Function's deployment ID used to create the execution.
    DeploymentId string `json:"deploymentId"`
    // Resource(function/site) execution duration in seconds.
    Duration float64 `json:"duration"`
    // Function errors. Includes the last 4,000 characters. This will return an
    // empty string unless the response is returned using an API key or as part of
    // a webhook payload.
    Errors string `json:"errors"`
    // Function ID.
    FunctionId string `json:"functionId"`
    // Function logs. Includes the last 4,000 characters. This will return an
    // empty string unless the response is returned using an API key or as part of
    // a webhook payload.
    Logs string `json:"logs"`
    // HTTP request headers as a key-value object. This will return only
    // whitelisted headers. All headers are returned if execution is created as
    // synchronous.
    RequestHeaders []Headers `json:"requestHeaders"`
    // HTTP request method type.
    RequestMethod string `json:"requestMethod"`
    // HTTP request path and query.
    RequestPath string `json:"requestPath"`
    // HTTP response body. This will return empty unless execution is created as
    // synchronous.
    ResponseBody string `json:"responseBody"`
    // HTTP response headers as a key-value object. This will return only
    // whitelisted headers. All headers are returned if execution is created as
    // synchronous.
    ResponseHeaders []Headers `json:"responseHeaders"`
    // HTTP response status code.
    ResponseStatusCode int `json:"responseStatusCode"`
    // The scheduled time for execution. If left empty, execution will be queued
    // immediately.
    ScheduledAt string `json:"scheduledAt"`
    // The status of the function execution. Possible values can be: `waiting`,
    // `processing`, `completed`, `failed`, or `scheduled`.
    Status string `json:"status"`
    // The trigger that caused the function to execute. Possible values can be:
    // `http`, `schedule`, or `event`.
    Trigger string `json:"trigger"`

    // Used by Decode() method
    data []byte
}

func (model Execution) New(data []byte) *Execution {
    model.data = data
    return &model
}

func (model *Execution) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}