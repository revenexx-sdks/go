package models

import (
    "encoding/json"
    "errors"
)

// CartIoMappingColumn model.
type CartIoMappingColumn struct {
    // The cart or line field, spelled as this app spells it — one of the
    // canonical column names.
    From string `json:"from"`
    // What that field is called on the outside: the CSV header, or the JSON key
    // of the system on the other end.
    To string `json:"to"`

    // Used by Decode() method
    data []byte
}

func (model CartIoMappingColumn) New(data []byte) *CartIoMappingColumn {
    model.data = data
    return &model
}

func (model *CartIoMappingColumn) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}