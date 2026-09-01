package models

import (
    "encoding/json"
    "errors"
)

// CartExportRequest model.
type CartExportRequest struct {
    // Format of an ad-hoc export, read only when no profile_id is sent. 'json'
    // returns the whole `{cart, items}` document, 'csv' the lines alone. Default
    // 'json'.
    Format string `json:"format"`
    // The export profile to run — one of the ids `GET
    // /carts/io/profiles?direction=export` lists. Omit it for an ad-hoc export in
    // the canonical shape, which is what `format` is for.
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