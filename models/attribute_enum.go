package models

import (
    "encoding/json"
    "errors"
)

// AttributeEnum Model
type AttributeEnum struct {
    // Attribute creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Attribute update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Is attribute an array?
    Array bool `json:"array"`
    // Array of elements in enumerated type.
    Elements []string `json:"elements"`
    // Error message. Displays error generated on failure of creating or deleting
    // an attribute.
    Error string `json:"error"`
    // String format.
    Format string `json:"format"`
    // Attribute Key.
    Key string `json:"key"`
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

func (model AttributeEnum) New(data []byte) *AttributeEnum {
    model.data = data
    return &model
}

func (model *AttributeEnum) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}