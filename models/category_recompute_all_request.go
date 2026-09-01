package models

import (
    "encoding/json"
    "errors"
)

// CategoryRecomputeAllRequest model.
type CategoryRecomputeAllRequest struct {

    // Used by Decode() method
    data []byte
}

func (model CategoryRecomputeAllRequest) New(data []byte) *CategoryRecomputeAllRequest {
    model.data = data
    return &model
}

func (model *CategoryRecomputeAllRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}