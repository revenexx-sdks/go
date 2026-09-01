package models

import (
    "encoding/json"
    "errors"
)

// LibraryItem One reusable block. Every page that references it renders THIS
// tree, so editing the item changes every placement at once.
type LibraryItem struct {
    // The block type this item instantiates. The library picker filters by it, so
    // an item only ever appears where its bundle is allowed. Theme-defined.
    Bundle string `json:"bundle"`
    // When the item entered the library.
    CreatedAt string `json:"created_at"`
    // The user id that made the block reusable.
    CreatedBy string `json:"created_by"`
    // The tombstone. A soft-deleted item is never listed or handed out, and a
    // block still referencing it keeps rendering its own last state rather than
    // breaking.
    DeletedAt string `json:"deleted_at"`
    // The library item id. A block references it to become an instance of the
    // item rather than a copy.
    Id string `json:"id"`
    // What the item is called in the library picker. This is the only thing an
    // editor sees before inserting it, so it carries the whole description.
    Label string `json:"label"`
    // The block and everything under it, serialized. This is the payload: every
    // page that references the item renders THIS tree, so editing it here changes
    // every placement at once.
    Tree PageBlockTree `json:"tree"`
    // When the item last changed — i.e. when every page referencing it last
    // changed with it.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model LibraryItem) New(data []byte) *LibraryItem {
    model.data = data
    return &model
}

func (model *LibraryItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}