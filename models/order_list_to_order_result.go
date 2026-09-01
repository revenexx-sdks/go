package models

import (
    "encoding/json"
    "errors"
)

// OrderListToOrderResult model.
type OrderListToOrderResult struct {
    // The list that was ordered. Unchanged by the call — the list stays, so it
    // can be ordered again next month.
    ListId string `json:"list_id"`
    // The orders app's answer, verbatim and unreshaped — the whole created
    // order, whose shape is the orders app's own `Order` schema (GET
    // /v1/orders/{id}) and is deliberately not restated here, because a copy
    // would be the thing that goes stale. `order_id`, `order_number` and `status`
    // are lifted out of it for a client that needs nothing else.
    Order interface{} `json:"order"`
    // The order the orders app created. Null only when that app answered without
    // one, which is a fault worth reporting rather than a normal outcome.
    OrderId string `json:"order_id"`
    // The order number a human quotes, drawn from the tenant's order range by the
    // orders app. It is NOT the id: every orders route addresses an order by
    // uuid.
    OrderNumber string `json:"order_number"`
    // Positions handed to the orders app — the list's count minus `skipped`.
    Positions int `json:"positions"`
    // Positions left out because the catalogue no longer knows their article.
    // Only ever non-empty when 'on_missing_article' is 'skip'.
    Skipped []OrderListSkippedPosition `json:"skipped"`
    // Where the new order stands, as the orders app decided: 'placed' when it was
    // accepted outright, 'pending' when it awaits approval — a contact holding
    // only orders.request, or an order above the tenant's approval threshold.
    // This app does not choose it and cannot override it.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model OrderListToOrderResult) New(data []byte) *OrderListToOrderResult {
    model.data = data
    return &model
}

func (model *OrderListToOrderResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}