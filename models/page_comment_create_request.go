package models

import (
    "encoding/json"
    "errors"
)

// PageCommentCreateRequest A new comment. Send `blockUuids` for a thread
// anchored to blocks, `parentUuid` for a reply.
type PageCommentCreateRequest struct {
    // The blocks this thread is about, so the editor can draw a marker next to
    // them. Leave empty for a comment about the page as a whole.
    BlockUuids []string `json:"blockUuids"`
    // The comment, as editor HTML. `<span data-type="mention" data-id="USER_ID">`
    // is what this app reads to decide whom to notify; `<li data-type="taskItem"
    // data-checked="false">` makes a checkbox the toggle-task route can flip.
    Body string `json:"body"`
    // The root comment this replies to. Omit for a new thread — only roots can
    // be resolved.
    ParentUuid string `json:"parentUuid"`

    // Used by Decode() method
    data []byte
}

func (model PageCommentCreateRequest) New(data []byte) *PageCommentCreateRequest {
    model.data = data
    return &model
}

func (model *PageCommentCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}