package models

import (
    "encoding/json"
    "errors"
)

// ChannelTypeDefaults The same answer for the channel types, which are seeded
// first because the seeded channel carries one.
type ChannelTypeDefaults struct {
    // Channel type codes this call wrote. A fresh tenant gets all 5; a settled
    // one gets none.
    Created []string `json:"created"`
    // Seeded type codes that were already there. Note the consequence of
    // "idempotent" being keyed on the code: a seeded type the merchant
    // deliberately retired is re-created by the next call and comes back under
    // `created`. Types the merchant added themselves are never touched.
    Existing []string `json:"existing"`

    // Used by Decode() method
    data []byte
}

func (model ChannelTypeDefaults) New(data []byte) *ChannelTypeDefaults {
    model.data = data
    return &model
}

func (model *ChannelTypeDefaults) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}