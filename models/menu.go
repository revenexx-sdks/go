package models

import (
    "encoding/json"
    "errors"
)

// Menu One navigation menu of the tenant, addressed by the stable key a theme
// looks it up under.
type Menu struct {
    // When the menu was created.
    CreatedAt string `json:"created_at"`
    // The user id that created the menu.
    CreatedBy string `json:"created_by"`
    // The tombstone. A soft-deleted menu disappears from the renderer
    // immediately.
    DeletedAt string `json:"deleted_at"`
    // The menu row id. Used by the management routes; the renderer addresses a
    // menu by its `menu_key` instead, because that is the thing a theme
    // hard-codes.
    Id string `json:"id"`
    // The ordered navigation tree itself. Stored exactly as it was sent, so the
    // theme and the editor agree on the shape without this app enforcing one.
    Items []PageMenuItem `json:"items"`
    // What this menu is called for the people who edit it. Never rendered in the
    // storefront.
    Label string `json:"label"`
    // The stable name the theme asks for a menu by — `main`, `footer`,
    // `account`. It is what makes seeding idempotent and what a header component
    // looks up; renaming it detaches the menu from the theme slot.
    MenuKey string `json:"menu_key"`
    // When the menu was last replaced. The upsert rewrites `items` wholesale, so
    // this is the timestamp of the whole navigation, not of one entry.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Menu) New(data []byte) *Menu {
    model.data = data
    return &model
}

func (model *Menu) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}