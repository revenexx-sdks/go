package models

import (
    "encoding/json"
    "errors"
)

// PriceEntriesAdjustResponse What the change did (or would do, on a dry run),
// plus the rounding policy it was computed under — so a dialog can show a
// merchant the before/after before it commits.
type PriceEntriesAdjustResponse struct {
    // Echo of the request: true means nothing was written.
    DryRun bool `json:"dry_run"`
    // Priced entries the filter selected. On-request entries are never counted
    // — a percentage of "ask us" is not a number.
    Matched int `json:"matched"`
    // Decimals the new prices were rounded to before snapping — the tenant’s
    // price_precision.
    Precision int `json:"precision"`
    // The first 50 changes, before and after. `matched` says how many there were
    // in total.
    Preview []PriceAdjustPreviewRow `json:"preview"`
    // true when more than 50 entries changed, so `preview` is a sample rather
    // than the whole set.
    PreviewTruncated bool `json:"preview_truncated"`
    // The price list this answer came out of — enough to link to it or to
    // explain the number to a merchant ("this came from the dealer list").
    PriceList PriceListRef `json:"price_list"`
    // The price ending the results were snapped to — the request’s, or the
    // tenant’s bulk_adjust_rounding where it sent none.
    Rounding string `json:"rounding"`
    // How they landed on the last decimal — the tenant’s rounding_mode.
    RoundingMode string `json:"rounding_mode"`
    // Rows actually written — 0 on a dry run, and a price that came out
    // unchanged is not rewritten.
    Updated int `json:"updated"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntriesAdjustResponse) New(data []byte) *PriceEntriesAdjustResponse {
    model.data = data
    return &model
}

func (model *PriceEntriesAdjustResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}