package models

import (
    "encoding/json"
    "errors"
)

// ShippingRateTier model.
type ShippingRateTier struct {
    // When the row was created (UTC).
    CreatedAt string `json:"created_at"`
    // Lower bound of this tier, in the method's matrix measure — kilograms (or
    // whatever the market's `weight_unit` names, converted through its factor)
    // for a weight matrix, items for quantity, money in the method's currency for
    // order_value, and the raw attribute value for 'attribute'. INCLUSIVE: the
    // tier applies from this value upward, and the tier that wins is the one with
    // the highest from_value at or below the measured value, so a measure of
    // exactly 10 is priced by the tier at 10 rather than the one below it. The
    // last tier has no upper bound. Unique per method — a second tier at the
    // same threshold is a 409, because which of the two won would be whatever the
    // database returned first.
    FromValue float64 `json:"from_value"`
    // Row id, assigned by the database on insert.
    Id string `json:"id"`
    // The shipping method this tier prices. Set from the path on every write, so
    // a body that names another method is ignored rather than obeyed. ON DELETE
    // CASCADE: deleting the method deletes its table.
    MethodId string `json:"method_id"`
    // Display order in the matrix editor (default 0; a bulk replace derives it
    // from the array index). Pricing reads from_value, never this.
    Position int `json:"position"`
    // What this tier costs, in the method's currency. Charged in full for the
    // whole consignment — a matrix is a lookup table, not a rate per unit.
    Price float64 `json:"price"`
    // When the row was last written (UTC).
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ShippingRateTier) New(data []byte) *ShippingRateTier {
    model.data = data
    return &model
}

func (model *ShippingRateTier) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}