package models

import (
    "encoding/json"
    "errors"
)

// Error Uniform error response. The same shape is emitted by the gateway and
// by the apps behind it, so one parser covers the whole API.
type Error struct {
    // Machine-readable discriminator, e.g. not_found, invalid_value,
    // unique_violation.
    Code string `json:"code"`
    // Human-readable message. Was a boolean on gateway-emitted errors before; it
    // is a string everywhere now.
    Error string `json:"error"`
    // Deprecated duplicate of `error`, kept so existing readers keep working.
    // Read `error`.
    Message string `json:"message"`

    // Used by Decode() method
    data []byte
}

func (model Error) New(data []byte) *Error {
    model.data = data
    return &model
}

func (model *Error) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}