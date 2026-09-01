package models

import (
    "encoding/json"
    "errors"
)

// IoEntity One importable / exportable entity of an installed app.
type IoEntity struct {
    // 
    App string `json:"app"`
    // 
    Entity string `json:"entity"`
    // Humanised entity name for pickers.
    Label string `json:"label"`
    // The physical table name Baseline provisioned.
    Table string `json:"table"`
    // 
    Vendor string `json:"vendor"`

    // Used by Decode() method
    data []byte
}

func (model IoEntity) New(data []byte) *IoEntity {
    model.data = data
    return &model
}

func (model *IoEntity) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}