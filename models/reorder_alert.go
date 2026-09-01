package models

import (
    "encoding/json"
    "errors"
)

// ReorderAlert model.
type ReorderAlert struct {
    // on_hand − reserved: the figure compared against the reorder point.
    // Alerting on AVAILABLE rather than on_hand is the point of this list — a
    // shelf that looks full but is entirely sold is exactly the row a buyer must
    // see.
    Available float64 `json:"available"`
    // That location's code, resolved for the reader so no second call is needed.
    // Null if the location row could not be read.
    LocationCode string `json:"location_code"`
    // Whether that location is enabled. A DISABLED location still alerts — its
    // stock is invisible to availability, but the goods are real and somebody has
    // to decide. Null if the location row could not be read.
    LocationEnabled bool `json:"location_enabled"`
    // The location holding it.
    LocationId string `json:"location_id"`
    // What is physically there right now, promised units included.
    OnHand float64 `json:"on_hand"`
    // The product this row tracks, null when it is tracked by SKU.
    ProductId string `json:"product_id"`
    // The threshold that was applied to this row — its own, or the tenant
    // default.
    ReorderPoint float64 `json:"reorder_point"`
    // 'row' — the stock row's own threshold. 'default' — the
    // reorder_point_default setting.
    ReorderPointSource string `json:"reorder_point_source"`
    // How much of it is already promised to orders.
    Reserved float64 `json:"reserved"`
    // How far below the point this row has fallen. The list is sorted by it,
    // worst first.
    Shortfall float64 `json:"shortfall"`
    // The article number this row tracks, null when it is tracked by product id.
    Sku string `json:"sku"`
    // The stock row that is low — the id to correct or receive against (POST
    // /inventories/stock/{id}/adjust).
    StockLevelId string `json:"stock_level_id"`

    // Used by Decode() method
    data []byte
}

func (model ReorderAlert) New(data []byte) *ReorderAlert {
    model.data = data
    return &model
}

func (model *ReorderAlert) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}