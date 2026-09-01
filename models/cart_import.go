package models

import (
    "encoding/json"
    "errors"
)

// CartImport `cart` is the cart as it now stands, totals already recomputed
// — the newly created one, or the target with the imported lines folded in.
type CartImport struct {
    // 
    Cart Cart `json:"cart"`
    // Lines read out of the payload. Identical product lines merge, so the cart
    // may have gained fewer rows than this.
    ImportedLines int `json:"imported_lines"`

    // Used by Decode() method
    data []byte
}

func (model CartImport) New(data []byte) *CartImport {
    model.data = data
    return &model
}

func (model *CartImport) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}