package models

import (
    "encoding/json"
    "errors"
)

// OrderNumberRangesSeeded Which of the three standard codes this call had to
// create and which were already there.
type OrderNumberRangesSeeded struct {
    // The codes that were created just now, with the standard format
    // ORD-/DEL-/RET- and padding 6. Empty on every call after the first.
    Created []string `json:"created"`
    // The codes that were already there and were left EXACTLY as they are — a
    // merchant who changed the prefix or the counter keeps their change. That is
    // what makes this call safe to run again.
    Existing []string `json:"existing"`

    // Used by Decode() method
    data []byte
}

func (model OrderNumberRangesSeeded) New(data []byte) *OrderNumberRangesSeeded {
    model.data = data
    return &model
}

func (model *OrderNumberRangesSeeded) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}