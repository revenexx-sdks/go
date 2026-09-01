package models

import (
    "encoding/json"
    "errors"
)

// BulkJobType One value per PE-102 block that moves data.
type BulkJobType struct {

    // Used by Decode() method
    data []byte
}

func (model BulkJobType) New(data []byte) *BulkJobType {
    model.data = data
    return &model
}

func (model *BulkJobType) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}