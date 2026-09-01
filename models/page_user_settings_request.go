package models

import (
    "encoding/json"
    "errors"
)

// PageUserSettingsRequest The preferences to store for the calling user.
type PageUserSettingsRequest struct {
    // The whole preferences bag — replaced, not merged, so send all of it. Its
    // keys vary by the editor build and this app reads none of them. Null or
    // omitted stores `{}`, which is how a user resets their editor.
    Settings interface{} `json:"settings"`

    // Used by Decode() method
    data []byte
}

func (model PageUserSettingsRequest) New(data []byte) *PageUserSettingsRequest {
    model.data = data
    return &model
}

func (model *PageUserSettingsRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}