package models

import (
    "encoding/json"
    "errors"
)

// Model
type AuthMeResponse struct {
    // 
    Contact Contact `json:"contact"`
    // 
    User interface{} `json:"user"`

    // Used by Decode() method
    data []byte
}

func (model AuthMeResponse) New(data []byte) *AuthMeResponse {
    model.data = data
    return &model
}

func (model *AuthMeResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}