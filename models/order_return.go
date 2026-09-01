package models

import (
    "encoding/json"
    "errors"
)

// OrderReturn Goods coming BACK, with their own lifecycle: registered →
// received → completed | rejected. Only completing books anything onto the
// positions; registering and receiving are announcements.
type OrderReturn struct {
    // When the return was settled, stamped by the SERVER. Never taken from the
    // body: a client clock records when a client thinks it acted, not when the
    // goods were booked.
    CompletedAt string `json:"completed_at"`
    // When the return row was written.
    CreatedAt string `json:"created_at"`
    // Primary key of the return. The {rid} segment of the return routes.
    Id string `json:"id"`
    // Free-form data for the caller — the returns portal's own reference.
    // Stored and returned untouched.
    Metadata interface{} `json:"metadata"`
    // The RETURN number — drawn from the tenant's return range, unique per
    // tenant, and a third series alongside orders and delivery notes. What the
    // customer writes on the parcel.
    Number string `json:"number"`
    // The order the goods are coming back from. A return of another order is a
    // 404 on these routes, not a cross-order write.
    OrderId string `json:"order_id"`
    // The positions and quantities this return covers, fixed when it was
    // registered and guarded against the shipped-but-not-yet-returned quantity of
    // each. Entries flagged restock are what the completion reports back for the
    // inventories call.
    Positions []OrderReturnedPosition `json:"positions"`
    // Why the goods are coming back, free text as the customer or the desk stated
    // it. Also what /reject stores when it is given no resolution out of the
    // published set.
    Reason string `json:"reason"`
    // When the goods physically arrived back. Null until POST …/receive — and
    // null forever on a return that was completed straight out of registered,
    // which is allowed.
    ReceivedAt string `json:"received_at"`
    // When the return was announced. Defaults to now.
    RegisteredAt string `json:"registered_at"`
    // When the return was refused. Null unless it was.
    RejectedAt string `json:"rejected_at"`
    // How it ended, in one of the words this app publishes — the settlement
    // words on a completion (refund, partial_refund, replacement, repair,
    // store_credit), the refusal words on a rejection (wear_and_tear,
    // not_returnable); GET /orders/vocabularies/return-resolutions carries both
    // sets with the stage that accepts each. The column carries no database
    // constraint; the ROUTES enforce the set, which is what stopped a client
    // settling returns with a word nobody else knew. On a rejection that named no
    // resolution, the free-text reason is stored here instead — which is the
    // one case a value outside the two sets appears.
    Resolution string `json:"resolution"`
    // Where the return stands: 'registered' = announced, nothing booked;
    // 'received' = the goods are back but not yet settled; 'completed' = settled,
    // and the only transition that books quantity_returned; 'rejected' = refused,
    // nothing booked. The last two are final.
    Status string `json:"status"`
    // When the return last changed — each of its transitions writes it.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model OrderReturn) New(data []byte) *OrderReturn {
    model.data = data
    return &model
}

func (model *OrderReturn) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}