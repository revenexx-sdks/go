package models

import (
    "encoding/json"
    "errors"
)

// FormDeleteResult model.
type FormDeleteResult struct {
    // True when the policy is 'archive' and submissions exist — the form was
    // archived, not deleted.
    Archived bool `json:"archived"`
    // The form row was removed — and with it, via the cascade, every submission
    // it had. `submissions` below says how many went, and they are not
    // recoverable.
    Deleted bool `json:"deleted"`
    // The form in the path.
    Id string `json:"id"`
    // The form's status after the call. Only present on the archive branch.
    Status string `json:"status"`
    // How many submissions the form had when the call was weighed — and
    // therefore, when `deleted` is true, how many were deleted with it. The whole
    // inbox, across every market: the cascade is a database operation and takes
    // them all, so an active `X-Revenexx-Market` does not narrow this number.
    Submissions int `json:"submissions"`

    // Used by Decode() method
    data []byte
}

func (model FormDeleteResult) New(data []byte) *FormDeleteResult {
    model.data = data
    return &model
}

func (model *FormDeleteResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}