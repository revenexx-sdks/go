package models

import (
    "encoding/json"
    "errors"
)

// OrderListCreateRequest model.
type OrderListCreateRequest struct {
    // Optional initial positions. Every one is validated — and article-checked
    // where `reject_unknown_articles` is on — BEFORE the list row is written,
    // so a rejected position never leaves an empty list behind.
    Items []OrderListItemInput `json:"items"`
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
    // The organization the sharing is scoped to. Null means the list can only
    // ever be the owner's own: `shared` is meaningless without it, because there
    // is no set of people to share with. It is also what the order conversion
    // hands the orders app as the buying organization.
    OrganizationId string `json:"organization_id"`
    // The contact who owns the list. Ownership IS the authorization here: a
    // caller the gateway resolved to a contact sees their own lists plus their
    // organization's shared ones, and may write only their own — unless
    // `shared_lists_editable` opens a shared list to the whole owning
    // organization. Set once at create; no route moves a list to another owner.
    OwnerId string `json:"owner_id"`
    // The owner's display name as it stood when the list was created — a
    // snapshot, so renaming the contact does not rewrite it. Carried so a shared
    // list can say whose it is without a call to the contacts app.
    OwnerName string `json:"owner_name"`
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

func (model OrderListCreateRequest) New(data []byte) *OrderListCreateRequest {
    model.data = data
    return &model
}

func (model *OrderListCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}