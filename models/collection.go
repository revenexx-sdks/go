package models

import (
    "encoding/json"
    "errors"
)

// Collection A Typesense collection definition, passed through from
// Typesense. `name` is rewritten back to the tenant's public collection name.
type Collection struct {
    // 
    DefaultSortingField string `json:"default_sorting_field"`
    // 
    EnableNestedFields bool `json:"enable_nested_fields"`
    // 
    Fields []CollectionField `json:"fields"`
    // The public collection name.
    Name string `json:"name"`
    // Documents currently indexed.
    NumDocuments int `json:"num_documents"`

    // Used by Decode() method
    data []byte
}

func (model Collection) New(data []byte) *Collection {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *Collection) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}