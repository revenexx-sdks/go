package models

import (
    "encoding/json"
    "errors"
)

// AttributeInteger Model
type AttributeInteger struct {
    // Attribute creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Attribute update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Is attribute an array?
    Array bool `json:"array"`
    // Error message. Displays error generated on failure of creating or deleting
    // an attribute.
    Error string `json:"error"`
    // Attribute Key.
    Key string `json:"key"`
    // Maximum value to enforce for new documents.
    Max int `json:"max"`
    // Minimum value to enforce for new documents.
    Min int `json:"min"`
    // Is attribute required?
    Required bool `json:"required"`
    // Attribute status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Attribute type.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model AttributeInteger) New(data []byte) *AttributeInteger {
    model.data = data
    return &model
}

func (model *AttributeInteger) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}