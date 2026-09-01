package models

import (
    "encoding/json"
    "errors"
)

// PageLibraryItemUpdateRequest Partial update — omitted fields keep their
// current value. Every page that references this item renders the new tree
// the next time it is delivered, which is the whole point of the library and
// the whole risk of editing one.
type PageLibraryItemUpdateRequest struct {
    // The block type this item instantiates. Changing it moves the item to a
    // different part of the picker.
    Bundle string `json:"bundle"`
    // What the item is called in the picker.
    Label string `json:"label"`
    // A block and its whole subtree, serialized. Produced by the editor when a
    // selection is made reusable or saved as a template, and instantiated back
    // into real blocks when one is inserted.
    Tree PageBlockTree `json:"tree"`

    // Used by Decode() method
    data []byte
}

func (model PageLibraryItemUpdateRequest) New(data []byte) *PageLibraryItemUpdateRequest {
    model.data = data
    return &model
}

func (model *PageLibraryItemUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}