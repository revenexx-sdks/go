package models

import (
    "encoding/json"
    "errors"
)

// PageScheduleRequest When this working copy should go live.
type PageScheduleRequest struct {
    // The moment to publish at. Stored on the edit state and echoed back
    // normalized to UTC.
    ScheduledAt string `json:"scheduledAt"`

    // Used by Decode() method
    data []byte
}

func (model PageScheduleRequest) New(data []byte) *PageScheduleRequest {
    model.data = data
    return &model
}

func (model *PageScheduleRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}