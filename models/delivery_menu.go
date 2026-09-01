package models

import (
    "encoding/json"
    "errors"
)

// DeliveryMenu One navigation menu, ready to render.
type DeliveryMenu struct {
    // The menu KEY (`main`, `footer`, `account`), not the row id — this is the
    // handle a theme hard-codes.
    Id string `json:"id"`
    // The ordered navigation tree, exactly as it is stored. Render it in order;
    // nesting is `items` inside an entry.
    Items []PageMenuItem `json:"items"`
    // What the menu is called for the people who edit it. A theme rarely renders
    // it.
    Label string `json:"label"`

    // Used by Decode() method
    data []byte
}

func (model DeliveryMenu) New(data []byte) *DeliveryMenu {
    model.data = data
    return &model
}

func (model *DeliveryMenu) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}