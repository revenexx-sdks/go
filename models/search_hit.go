package models

import (
    "encoding/json"
    "errors"
)

// SearchHit model.
type SearchHit struct {
    // The matching document; its properties are the collection's own fields.
    Document interface{} `json:"document"`
    // Per-field highlight snippets, keyed by field name.
    Highlight interface{} `json:"highlight"`
    // Relevance score.
    TextMatch int `json:"text_match"`

    // Used by Decode() method
    data []byte
}

func (model SearchHit) New(data []byte) *SearchHit {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *SearchHit) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}