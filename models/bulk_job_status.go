package models

import (
    "encoding/json"
    "errors"
)

// BulkJobStatus Lifecycle of a `baseline.bulk_jobs` row:
// `pending → running → completed`, or `partial` (finished with
// `counts.rejected > 0`), `failed`, or `canceled`.
type BulkJobStatus struct {

    // Used by Decode() method
    data []byte
}

func (model BulkJobStatus) New(data []byte) *BulkJobStatus {
    model.data = data
    return &model
}

func (model *BulkJobStatus) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}