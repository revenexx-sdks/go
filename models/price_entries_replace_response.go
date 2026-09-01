package models

import (
    "encoding/json"
    "errors"
)

// PriceEntriesReplaceResponse The list as it now stands: everything that was
// there is gone and these are the rows that took its place.
type PriceEntriesReplaceResponse struct {
    // The complete new entry set, as stored — including the ids and timestamps
    // the database filled in.
    Entries []PriceEntry `json:"entries"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntriesReplaceResponse) New(data []byte) *PriceEntriesReplaceResponse {
    model.data = data
    return &model
}

func (model *PriceEntriesReplaceResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}