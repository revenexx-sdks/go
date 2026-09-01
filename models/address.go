package models

import (
    "encoding/json"
    "errors"
)

// Address A postal address belonging to an organization or to a contact, used
// for billing or shipping. Ownership is exactly one of the two.
type Address struct {
    // City or town.
    City string `json:"city"`
    // Company line on the label. Often the owning organization's name, but not
    // always — a delivery to a construction site carries the site.
    Company string `json:"company"`
    // Owning person — a personal address only that contact uses. Exactly one of
    // organization_id / contact_id is set.
    ContactId string `json:"contact_id"`
    // ISO 3166-1 alpha-2 country code, exactly two letters. Uppercase by
    // convention; it is what shipping and tax both key off.
    Country string `json:"country"`
    // When the address was created.
    CreatedAt string `json:"created_at"`
    // Primary key of the address.
    Id string `json:"id"`
    // The default address of its owner AND type: one default billing and one
    // default shipping address per owner. Setting it moves the flag off the
    // previous holder.
    IsDefault bool `json:"is_default"`
    // Recipient line on the label — the person or department the parcel is
    // addressed to.
    Name string `json:"name"`
    // Owning company — a company address, shared by everyone in it. Exactly one
    // of organization_id / contact_id is set.
    OrganizationId string `json:"organization_id"`
    // Phone number for the carrier to reach at this address — often a different
    // one from the contact's own.
    Phone string `json:"phone"`
    // State, province or Bundesland. Required by some destinations (US, CA),
    // unused by most European ones.
    Region string `json:"region"`
    // Street and house number, on one line, as the local post expects it.
    Street string `json:"street"`
    // The second address line: building, floor, gate, c/o. Null when there is
    // none.
    Street2 string `json:"street2"`
    // The tenant this row belongs to — the store slug, not an id. Set by the
    // platform from the authenticated context, never by a caller; a write that
    // carries it is ignored, and no request can read another tenant's rows by
    // sending a different one.
    TenantId string `json:"tenant_id"`
    // What the address is FOR — one of the tenant's own address types (GET
    // /customers/address-types), seeded with billing and shipping. A merchant may
    // add their own (a works entrance, a central accounts office) without a
    // release of this app.
    Type string `json:"type"`
    // When any column of this row last changed.
    UpdatedAt string `json:"updated_at"`
    // Postal code, as text — leading zeros are real in most countries.
    Zip string `json:"zip"`

    // Used by Decode() method
    data []byte
}

func (model Address) New(data []byte) *Address {
    model.data = data
    return &model
}

func (model *Address) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}