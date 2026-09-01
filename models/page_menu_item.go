package models

import (
    "encoding/json"
    "errors"
)

// PageMenuItem One entry of a navigation menu. Stored verbatim, so a theme
// may carry extra keys of its own alongside these.
type PageMenuItem struct {
    // Sub-entries. This is how a two-level main navigation or a grouped footer is
    // stored.
    Items []interface{} `json:"items"`
    // The words a visitor clicks.
    Label string `json:"label"`
    // Where the entry goes: a page slug this app serves, a path the theme routes,
    // or an absolute URL to somewhere else.
    To string `json:"to"`

    // Used by Decode() method
    data []byte
}

func (model PageMenuItem) New(data []byte) *PageMenuItem {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *PageMenuItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}