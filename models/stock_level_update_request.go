package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type StockLevelUpdateRequest struct {
    // Owning location.
    LocationId string `json:"location_id"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // Physical stock (default 0).
    OnHand float64 `json:"on_hand"`
    // Tracked product.
    ProductId string `json:"product_id"`
    // 
    ReorderPoint float64 `json:"reorder_point"`
    // Reserved stock (default 0) — normally managed by reserve/release/commit.
    Reserved float64 `json:"reserved"`
    // Tracked SKU (alternative to product_id).
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model StockLevelUpdateRequest) New(data []byte) *StockLevelUpdateRequest {
    model.data = data
    return &model
}

func (model *StockLevelUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}