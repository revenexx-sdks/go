package models

import (
    "encoding/json"
    "errors"
)

// ContactPermissions A contact's effective grants, derived from its role on
// every read — nothing here is stored, so a role change can never leave a
// stale grant behind. Carried here so a BFF does not need a second call to
// decide what to render.
type ContactPermissions struct {
    // False while the contact is blocked or its registration is still
    // pending/rejected — it holds the role but must not act on it.
    Active bool `json:"active"`
    // The person these grants belong to. Null when the answer describes nobody
    // — a user with no contact mirrored against it.
    ContactId string `json:"contact_id"`
    // Amount ceiling in the market's currency; null means no ceiling. Only
    // meaningful together with the 'orders.approve' permission.
    OrderApprovalLimit float64 `json:"order_approval_limit"`
    // The organization the role applies inside. Null for a standalone (B2C)
    // contact — a role with no company to hold it in.
    OrganizationId string `json:"organization_id"`
    // What this role may do. Derived from the role — see GET /customers/roles.
    Permissions []string `json:"permissions"`
    // The role this contact holds in its organization, and the only input to
    // `permissions`.
    Role string `json:"role"`

    // Used by Decode() method
    data []byte
}

func (model ContactPermissions) New(data []byte) *ContactPermissions {
    model.data = data
    return &model
}

func (model *ContactPermissions) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}