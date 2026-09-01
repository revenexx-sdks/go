package models

import (
    "encoding/json"
    "errors"
)

// RolePermissionsRequest model.
type RolePermissionsRequest struct {
    // The complete new set. Duplicates and blanks are ignored; an empty array
    // revokes everything.
    Permissions []string `json:"permissions"`

    // Used by Decode() method
    data []byte
}

func (model RolePermissionsRequest) New(data []byte) *RolePermissionsRequest {
    model.data = data
    return &model
}

func (model *RolePermissionsRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}