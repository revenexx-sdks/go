package models

import (
    "encoding/json"
    "errors"
)

// RolesDefaultsResponse model.
type RolesDefaultsResponse struct {
    // Role keys created by this call.
    Created []string `json:"created"`
    // Role keys that were already there and were left untouched, permissions
    // included.
    Existing []string `json:"existing"`

    // Used by Decode() method
    data []byte
}

func (model RolesDefaultsResponse) New(data []byte) *RolesDefaultsResponse {
    model.data = data
    return &model
}

func (model *RolesDefaultsResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}