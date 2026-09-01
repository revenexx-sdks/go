package models

import (
    "encoding/json"
    "errors"
)

// MenuUpsertRequest Create or replace the menu identified by menuKey
// (idempotent per tenant). `items` is written wholesale — there is no
// per-entry edit, so send the whole tree every time.
type MenuUpsertRequest struct {
    // The ordered navigation tree. Replaces the stored one completely.
    Items []PageMenuItem `json:"items"`
    // What this menu is called for the people who edit it. Required on a create;
    // an update keeps the label it had when this is left out.
    Label string `json:"label"`
    // The stable slot the theme asks for this menu by. Idempotency is keyed on
    // it: sending an existing key replaces that menu instead of creating a second
    // one.
    MenuKey string `json:"menuKey"`

    // Used by Decode() method
    data []byte
}

func (model MenuUpsertRequest) New(data []byte) *MenuUpsertRequest {
    model.data = data
    return &model
}

func (model *MenuUpsertRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}