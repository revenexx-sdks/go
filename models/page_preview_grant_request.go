package models

import (
    "encoding/json"
    "errors"
)

// PagePreviewGrantRequest How long the link should live.
type PagePreviewGrantRequest struct {
    // Hours until the link expires. Defaults to 72. After that `GET
    // /pages/delivery/preview/{token}` answers 410 rather than 404, so the holder
    // can tell "expired" from "wrong link".
    TtlHours int `json:"ttlHours"`

    // Used by Decode() method
    data []byte
}

func (model PagePreviewGrantRequest) New(data []byte) *PagePreviewGrantRequest {
    model.data = data
    return &model
}

func (model *PagePreviewGrantRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}