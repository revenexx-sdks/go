package models

import (
    "encoding/json"
    "errors"
)

// MenuUpdateRequest Partial update — omitted fields keep their current
// value. `items` is replaced wholesale when sent.
type MenuUpdateRequest struct {
    // The ordered navigation tree. Replaces the stored one completely.
    Items []PageMenuItem `json:"items"`
    // What this menu is called for the people who edit it.
    Label string `json:"label"`

    // Used by Decode() method
    data []byte
}

func (model MenuUpdateRequest) New(data []byte) *MenuUpdateRequest {
    model.data = data
    return &model
}

func (model *MenuUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}