package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type AttributeGroupsUpdateRequest struct {
    // 
    Code string `json:"code"`
    // 
    Labels interface{} `json:"labels"`
    // 
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model AttributeGroupsUpdateRequest) New(data []byte) *AttributeGroupsUpdateRequest {
    model.data = data
    return &model
}

func (model *AttributeGroupsUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}