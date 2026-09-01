package models

import (
    "encoding/json"
    "errors"
)

// PriceEntriesAdjustRequest Change every priced entry of a list at once. Send
// 'percent' OR 'amount', never both. On-request entries are never touched —
// a percentage of "ask us" is not a number.
type PriceEntriesAdjustRequest struct {
    // Absolute change added to every unit price, in the list's currency.
    Amount float64 `json:"amount"`
    // true writes nothing and answers the same preview — what the Cockpit
    // dialog shows before it commits.
    DryRun bool `json:"dry_run"`
    // Relative change in percent: 5 raises by 5 %, -10 cuts by 10 %.
    Percent float64 `json:"percent"`
    // Ending the computed prices snap to (nearest match). Omit to use the
    // tenant's bulk_adjust_rounding setting.
    Rounding string `json:"rounding"`
    // Restrict the change to entries whose SKU starts with this (a prefix,
    // case-sensitive, no wildcards). Entries identified only by product_id never
    // match a prefix. Omit to change the whole list.
    SkuPrefix string `json:"sku_prefix"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntriesAdjustRequest) New(data []byte) *PriceEntriesAdjustRequest {
    model.data = data
    return &model
}

func (model *PriceEntriesAdjustRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}