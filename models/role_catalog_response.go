package models

import (
    "encoding/json"
    "errors"
)

// RoleCatalogResponse model.
type RoleCatalogResponse struct {
    // The built-in permission vocabulary, one entry per grant. The authoritative,
    // installed-app-aware list is the platform's permission ledger — this app
    // deliberately does not duplicate it.
    Permissions []interface{} `json:"permissions"`
    // Every role a contact of this tenant can hold, least to most privileged.
    Roles []interface{} `json:"roles"`
    // 'tenant' — the configured mapping answered. 'defaults' — this tenant
    // has no roles yet, or custom_roles_enabled locks the ledger, and the
    // built-ins answered.
    Source string `json:"source"`

    // Used by Decode() method
    data []byte
}

func (model RoleCatalogResponse) New(data []byte) *RoleCatalogResponse {
    model.data = data
    return &model
}

func (model *RoleCatalogResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}