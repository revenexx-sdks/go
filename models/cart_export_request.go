package models

import (
    "encoding/json"
    "errors"
)

// Model
type CartExportRequest struct {
    // Ad-hoc export format (only without profile_id).
    Format string `json:"format"`
    // Export profile to run; ad-hoc JSON/CSV export when omitted.
    ProfileId string `json:"profile_id"`

    // Used by Decode() method
    data []byte
}

func (model CartExportRequest) New(data []byte) *CartExportRequest {
    model.data = data
    return &model
}

func (model *CartExportRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}