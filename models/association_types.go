package models

import (
    "encoding/json"
    "errors"
)

// Model
type AssociationTypes struct {
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    IsQuantified bool `json:"is_quantified"`
    // 
    IsTwoWay bool `json:"is_two_way"`
    // 
    Labels interface{} `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model AssociationTypes) New(data []byte) *AssociationTypes {
    model.data = data
    return &model
}

func (model *AssociationTypes) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}