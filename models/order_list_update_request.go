package models

import (
    "encoding/json"
    "errors"
)

// OrderListUpdateRequest Partial update — rename, visibility or kind.
// Positions go through the items routes, and the owner cannot be changed.
type OrderListUpdateRequest struct {
    // List kind — the `code` of one of the tenant's own kinds (GET
    // /orderlists/kinds); defaults to the flagged one, or the market's
    // 'default_kind' setting.
    Kind string `json:"kind"`
    // Free-form data the tenant keeps on the list — an ERP requisition number,
    // a department, whatever an integration needs to recognise the list again.
    // Never read by this app, and never merged: a write replaces the whole
    // document.
    Metadata interface{} `json:"metadata"`
    // What the buyer calls this list. Free text, at least one character, and not
    // unique: two contacts may both keep a "Weekly office supplies". It is also
    // the name a NEW cart gets when POST /orderlists/{id}/cart creates one.
    Name string `json:"name"`
    // Whether the OWNING ORGANIZATION may see this list. False — the default
    // — keeps it private to `owner_id`, and a foreign private list answers 404
    // rather than 403, so an outsider learns nothing from the difference. True
    // lets every contact of `organization_id` READ it, and write it only where
    // the tenant turned on the `shared_lists_editable` setting. A list with no
    // `organization_id` shares with nobody however this is set.
    Shared bool `json:"shared"`

    // Used by Decode() method
    data []byte
}

func (model OrderListUpdateRequest) New(data []byte) *OrderListUpdateRequest {
    model.data = data
    return &model
}

func (model *OrderListUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}