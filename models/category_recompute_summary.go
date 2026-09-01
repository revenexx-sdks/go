package models

import (
    "encoding/json"
    "errors"
)

// CategoryRecomputeSummary model.
type CategoryRecomputeSummary struct {
    // Membership rows inserted with source='rule' by this call.
    Added int `json:"added"`
    // False → the bulk insert was refused and the call fell back to one request
    // per row. A performance fact, not an error.
    Batched bool `json:"batched"`
    // The category this pass belongs to, echoed back — a caller driving several
    // loops keys its state by it.
    CategoryId string `json:"category_id"`
    // The category's code, so a nightly log names something a person recognises.
    Code string `json:"code"`
    // When the pass completed, and what `categories.rules_computed_at` was
    // stamped with. Null while `done` is false.
    ComputedAt string `json:"computed_at"`
    // The product id this call reconciled up to, to hand back on the next one.
    // Null when `done`.
    Cursor string `json:"cursor"`
    // False → this call spent its budget mid-pass. Send `cursor` back to
    // continue; the counters below are THIS call only, so a caller looping to
    // completion sums them itself.
    Done bool `json:"done"`
    // Present instead of the counters when this category failed.
    Error string `json:"error"`
    // Matching products examined by this call.
    Processed int `json:"processed"`
    // Stale rule rows deleted by this call.
    Removed int `json:"removed"`
    // True → the budget ran out before this category was reached; it carries no
    // counters.
    Skipped bool `json:"skipped"`
    // The HTTP status this category WOULD have answered on its own — 400 for a
    // rule that does not compile, 404 for one that vanished mid-run. Null when it
    // succeeded.
    Status int `json:"status"`
    // Products the rule currently selects. Null while `done` is false — the
    // pass has not seen the whole catalog yet, so there is no total to report.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model CategoryRecomputeSummary) New(data []byte) *CategoryRecomputeSummary {
    model.data = data
    return &model
}

func (model *CategoryRecomputeSummary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}