package models

import (
    "encoding/json"
    "errors"
)

// Model
type PriceEntriesReplaceRequest struct {
    // The complete new entry set (set semantics).
    Entries []PriceEntryReplaceItem `json:"entries"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntriesReplaceRequest) New(data []byte) *PriceEntriesReplaceRequest {
    model.data = data
    return &model
}

func (model *PriceEntriesReplaceRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}