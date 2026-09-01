package models

import (
    "encoding/json"
    "errors"
)

// PriceEntriesBulkRequest A chunk of an import. Unlike the replace call it
// never wipes the list.
type PriceEntriesBulkRequest struct {
    // At most 5000 rows per call — send a large book in chunks.
    Entries []PriceEntryReplaceItem `json:"entries"`
    // Default 'upsert': a row naming a rung the list already has (same
    // product/sku AND quantity_min) updates it. 'append' always inserts — a
    // re-run then duplicates the ladder, which is what makes an ambiguous tier
    // table.
    Mode string `json:"mode"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntriesBulkRequest) New(data []byte) *PriceEntriesBulkRequest {
    model.data = data
    return &model
}

func (model *PriceEntriesBulkRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}