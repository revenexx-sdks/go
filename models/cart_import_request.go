package models

import (
    "encoding/json"
    "errors"
)

// ImportIntoAnExistingCartTargetCartIdOrANewCartOwnerContactIdSessionKeyRequired
// Model
type CartImportRequest struct {
    // Owner of a newly created cart.
    ContactId string `json:"contact_id"`
    // Raw CSV content (alternative to payload for csv profiles).
    Csv string `json:"csv"`
    // Name for a newly created cart.
    Name string `json:"name"`
    // The import payload: '{cart, items}' object, or a raw JSON/CSV string in the
    // profile's format.
    Payload interface{} `json:"payload"`
    // Import profile to run; ad-hoc import when omitted.
    ProfileId string `json:"profile_id"`
    // Guest owner of a newly created cart.
    SessionKey string `json:"session_key"`
    // Existing active cart to import into.
    TargetCartId string `json:"target_cart_id"`

    // Used by Decode() method
    data []byte
}

func (model CartImportRequest) New(data []byte) *CartImportRequest {
    model.data = data
    return &model
}

func (model *CartImportRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}