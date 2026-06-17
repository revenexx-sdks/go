package models

import (
    "encoding/json"
    "errors"
)

// Model
type Organization struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    ExternalTeamId string `json:"external_team_id"`
    // 
    Id string `json:"id"`
    // 
    Name string `json:"name"`
    // 
    Settings interface{} `json:"settings"`
    // 
    Status string `json:"status"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    VatId string `json:"vat_id"`

    // Used by Decode() method
    data []byte
}

func (model Organization) New(data []byte) *Organization {
    model.data = data
    return &model
}

func (model *Organization) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}