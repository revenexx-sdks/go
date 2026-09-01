package models

import (
    "encoding/json"
    "errors"
)

// BulkJob A bulk job as returned by `/bulk-jobs`. Note that the row counts
// are
// nested under `counts` — they are not top-level fields — and that the
// response carries no `tenant_id` (the listing envelope does) and no
// `updated_at`.
type BulkJob struct {
    // 
    App string `json:"app"`
    // 
    CorrelationId string `json:"correlation_id"`
    // 
    Counts interface{} `json:"counts"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    CreatedBy string `json:"created_by"`
    // 
    DurationMs int `json:"duration_ms"`
    // 
    Entity string `json:"entity"`
    // 
    ErrorMessage string `json:"error_message"`
    // 
    FinishedAt string `json:"finished_at"`
    // 
    Id string `json:"id"`
    // 
    ProfileId string `json:"profile_id"`
    // Engine-reported progress. For an export this carries the
    // `object_key` and `format` the result is written to.
    Progress interface{} `json:"progress"`
    // 
    StartedAt string `json:"started_at"`
    // 
    Status BulkJobStatus `json:"status"`
    // 
    Type BulkJobType `json:"type"`
    // 
    Vendor string `json:"vendor"`

    // Used by Decode() method
    data []byte
}

func (model BulkJob) New(data []byte) *BulkJob {
    model.data = data
    return &model
}

func (model *BulkJob) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}