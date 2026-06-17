package models

import (
    "encoding/json"
    "errors"
)

// Model
type IoProfile struct {
    // 
    ApplyMode string `json:"apply_mode"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Direction string `json:"direction"`
    // 
    Entity string `json:"entity"`
    // 
    Format string `json:"format"`
    // 
    Id string `json:"id"`
    // 
    IsTemplate bool `json:"is_template"`
    // 
    Mapping interface{} `json:"mapping"`
    // 
    Name string `json:"name"`
    // 
    Options interface{} `json:"options"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model IoProfile) New(data []byte) *IoProfile {
    model.data = data
    return &model
}

func (model *IoProfile) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}