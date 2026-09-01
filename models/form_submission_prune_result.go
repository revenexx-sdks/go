package models

import (
    "encoding/json"
    "errors"
)

// FormSubmissionPruneResult model.
type FormSubmissionPruneResult struct {
    // Submissions created before this instant match. It is `now -
    // older_than_days`, computed after the retention floor was applied, so it is
    // the honest answer to "what did this call actually consider".
    Cutoff string `json:"cutoff"`
    // How many rows this call actually removed — always 0 on a dry run, and at
    // most the 500-row batch size on a real one.
    Deleted int `json:"deleted"`
    // Whether this call was a preview. True — the default — means nothing was
    // deleted and `matched` is what a real run would take.
    DryRun bool `json:"dry_run"`
    // True when the request asked for a shorter age than the floor allows.
    FloorApplied bool `json:"floor_applied"`
    // How many rows match, ignoring the batch size.
    Matched int `json:"matched"`
    // The threshold actually applied, after the retention floor.
    OlderThanDays float64 `json:"older_than_days"`
    // Matched rows left after this batch — call again. Absent on a dry run,
    // which deletes nothing.
    Remaining int `json:"remaining"`
    // The retention floor this sweep honoured: the LONGEST
    // submission_retention_days configured anywhere in the tenant, baseline or
    // market. Not the value the calling market sees — a tenant-wide sweep has
    // to keep the longest promise anybody was given.
    RetentionDays float64 `json:"retention_days"`
    // The market whose submission_retention_days set the floor — the merchant's
    // own market CODE — or null when the tenant baseline did. It is there so a
    // merchant can see WHY the sweep would not go younger, since the market that
    // bound it is often not the one the request was made from.
    RetentionMarket string `json:"retention_market"`
    // Up to five matching rows (dry runs only) — id, form_slug and created_at,
    // never the submitted data.
    Sample []FormSubmissionPruneSample `json:"sample"`

    // Used by Decode() method
    data []byte
}

func (model FormSubmissionPruneResult) New(data []byte) *FormSubmissionPruneResult {
    model.data = data
    return &model
}

func (model *FormSubmissionPruneResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}