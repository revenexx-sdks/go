package models

import (
    "encoding/json"
    "errors"
)

// StockLevelCreateRequest A stock row tracks an item: 'product_id' or 'sku'.
type StockLevelCreateRequest struct {
    // The location this balance is held at — a `locations` row of this tenant
    // (GET /inventories/locations). There is ONE stock row per (location, item):
    // the same SKU in three warehouses is three rows, and what a storefront shows
    // is their sum (POST /inventories/availability). Deleting the location
    // deletes its stock rows with it. It has to exist already (GET
    // /inventories/locations); an id no location carries is answered 400 by the
    // foreign key, not 404.
    LocationId string `json:"location_id"`
    // Free-form data the tenant keeps on this stock row, and ONE key this app
    // reads: `backorder`. A literal boolean `true` there opts this item into
    // backorders while `backorder_policy` is 'allow_per_sku' — anything else,
    // including the string "true", does not, and the reservation is refused with
    // 422. That is how a merchant backorders the supplier-stocked half of a
    // catalogue without promising the rest.
    Metadata interface{} `json:"metadata"`
    // The product this row tracks, as the products app knows it. A row tracks a
    // `product_id` or a `sku` — the database insists on at least one (CHECK
    // `product_id is not null or sku is not null`) — and matching is exact: a
    // row keyed by SKU is not found by product id.
    ProductId string `json:"product_id"`
    // The available quantity at or below which this row belongs on the
    // replenishment worklist (GET /inventories/reorder-alerts). Null falls back
    // to the `reorder_point_default` setting, so replenishment works without a
    // threshold per SKU; 0 never alerts, which is how one row opts out.
    ReorderPoint float64 `json:"reorder_point"`
    // The article number this row tracks when there is no product id, which is
    // the normal case for an ERP-stocked catalogue. Exact match, and the identity
    // every stock call may use instead of a uuid.
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model StockLevelCreateRequest) New(data []byte) *StockLevelCreateRequest {
    model.data = data
    return &model
}

func (model *StockLevelCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}