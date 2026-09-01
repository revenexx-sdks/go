package models

import (
    "encoding/json"
    "errors"
)

// OrderCustomerRollupResponse model.
type OrderCustomerRollupResponse struct {
    // The anchor the windows were measured from — echoed so a paging caller can
    // pin it.
    AsOf string `json:"as_of"`
    // Where to resume, when `done` is false — the id of the last order this
    // call read. Null once the scan finished. Send it back unchanged, together
    // with the same as_of.
    Cursor string `json:"cursor"`
    // True = the whole set was scanned and this answer is complete. False = the
    // scan hit its time budget: send `cursor` back to continue, and MERGE the
    // parts (every number is additive, min for first_order_at, max for
    // last_order_at, union for currencies).
    Done bool `json:"done"`
    // One row per organization that appeared on a counted order, sorted by id. A
    // company with no counted order is absent — this answer does not carry zero
    // rows.
    Items []OrderCustomerRollup `json:"items"`
    // How many order rows this call read, attributed or not. It is the cost of
    // the call, and on a partial answer the size of the part.
    OrdersScanned int `json:"orders_scanned"`
    // Orders read that carry no organization_id — private and guest orders.
    // They are real revenue and are deliberately not attributed to anybody, so
    // they appear here and in no row of items.
    OrdersWithoutOrganization int `json:"orders_without_organization"`
    // How many rows `items` carries. On a partial answer this counts what THIS
    // part saw, not the whole tenant.
    Organizations int `json:"organizations"`
    // The statuses that were counted, echoed — the default set unless the
    // request named its own.
    Statuses []string `json:"statuses"`
    // The rolling windows the *_30d / *_90d / *_365d numbers were measured over,
    // in days. Echoed so a caller reads the numbers with the right labels instead
    // of hard-coding three of them.
    Windows []int `json:"windows"`

    // Used by Decode() method
    data []byte
}

func (model OrderCustomerRollupResponse) New(data []byte) *OrderCustomerRollupResponse {
    model.data = data
    return &model
}

func (model *OrderCustomerRollupResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}