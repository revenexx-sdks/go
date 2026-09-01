package models

import (
    "encoding/json"
    "errors"
)

// RolePermissionsResponse model.
type RolePermissionsResponse struct {
    // The role that was written.
    Key string `json:"key"`
    // Its complete new set, after de-duplication.
    Permissions []string `json:"permissions"`

    // Used by Decode() method
    data []byte
}

func (model RolePermissionsResponse) New(data []byte) *RolePermissionsResponse {
    model.data = data
    return &model
}

func (model *RolePermissionsResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}