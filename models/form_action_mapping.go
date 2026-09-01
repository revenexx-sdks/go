package models

import (
    "encoding/json"
    "errors"
)

// FormActionMapping model.
type FormActionMapping struct {
    // The key in the submission `data` — i.e. the `name` of a definition node.
    Source string `json:"source"`
    // The column of the target entity it is written to.
    Target string `json:"target"`

    // Used by Decode() method
    data []byte
}

func (model FormActionMapping) New(data []byte) *FormActionMapping {
    model.data = data
    return &model
}

func (model *FormActionMapping) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}