package models

import (
    "encoding/json"
    "errors"
)

// Model
type Families struct {
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    ImageAttribute string `json:"image_attribute"`
    // 
    LabelAttribute string `json:"label_attribute"`
    // 
    Labels interface{} `json:"labels"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Families) New(data []byte) *Families {
    model.data = data
    return &model
}

func (model *Families) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}