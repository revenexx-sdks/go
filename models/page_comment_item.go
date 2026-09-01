package models

import (
    "encoding/json"
    "errors"
)

// PageCommentItem One comment, in the shape the editor renders — this is
// not the stored row: the id is `uuid`, the timestamps are
// `created`/`updated` and the author is nested under `user`.
type PageCommentItem struct {
    // The blocks this thread hangs on, so the editor can draw a marker next to
    // them. Empty for a comment about the page as a whole.
    BlockUuids []string `json:"blockUuids"`
    // The comment itself, as editor HTML. @mentions are `<span
    // data-type="mention" data-id="…">` — that is what this app reads to
    // decide whom to notify — and task checkboxes are `<li data-type="taskItem"
    // data-checked="…">`.
    Body string `json:"body"`
    // When the comment was written.
    Created string `json:"created"`
    // The root comment this is a reply to. Absent on a root — and only roots
    // can be resolved.
    ParentUuid string `json:"parentUuid"`
    // Whether the thread was marked done. Replies inherit nothing: resolving is a
    // property of the root.
    Resolved bool `json:"resolved"`
    // When it was last edited. Absent when it never was.
    Updated string `json:"updated"`
    // Who wrote it, or `null` when it was written without an identity.
    User interface{} `json:"user"`
    // The comment id. Every comment route addresses one by it.
    Uuid string `json:"uuid"`

    // Used by Decode() method
    data []byte
}

func (model PageCommentItem) New(data []byte) *PageCommentItem {
    model.data = data
    return &model
}

func (model *PageCommentItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}