package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type CategoriesUpdateRequest struct {
    // 
    Code string `json:"code"`
    // 
    Labels interface{} `json:"labels"`
    // 
    ParentId string `json:"parent_id"`
    // 
    Path string `json:"path"`
    // 
    Position int `json:"position"`
    // 
    Values interface{} `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model CategoriesUpdateRequest) New(data []byte) *CategoriesUpdateRequest {
    model.data = data
    return &model
}

func (model *CategoriesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}