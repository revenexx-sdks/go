package models

import (
    "encoding/json"
    "errors"
)

// PrincipalResolveRequest model.
type PrincipalResolveRequest struct {
    // The contact the caller is acting for.
    ContactId string `json:"contact_id"`

    // Used by Decode() method
    data []byte
}

func (model PrincipalResolveRequest) New(data []byte) *PrincipalResolveRequest {
    model.data = data
    return &model
}

func (model *PrincipalResolveRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}