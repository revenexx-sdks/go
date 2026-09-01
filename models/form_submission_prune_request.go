package models

import (
    "encoding/json"
    "errors"
)

// FormSubmissionPruneRequest Retention sweep. Previews unless `dry_run` is
// explicitly false.
type FormSubmissionPruneRequest struct {
    // Default TRUE. Nothing is deleted until this is explicitly false.
    DryRun bool `json:"dry_run"`
    // Narrow the sweep to one form.
    FormSlug string `json:"form_slug"`
    // Age threshold. Omit to use the retention floor. A value BELOW the floor is
    // raised to it — the setting is the floor, not a default, and the floor is
    // the LONGEST submission_retention_days configured anywhere in the tenant
    // (see the operation description).
    OlderThanDays int `json:"older_than_days"`
    // Narrow the sweep to one inbox status, e.g. 'spam'.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model FormSubmissionPruneRequest) New(data []byte) *FormSubmissionPruneRequest {
    model.data = data
    return &model
}

func (model *FormSubmissionPruneRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}