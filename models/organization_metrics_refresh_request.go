package models

import (
    "encoding/json"
    "errors"
)

// OrganizationMetricsRefreshRequest model.
type OrganizationMetricsRefreshRequest struct {
    // Anchor for the rolling windows — pass back the value the previous call
    // returned.
    AsOf string `json:"as_of"`
    // Continue an unfinished refresh: the value the previous call returned,
    // verbatim. It is the id of the last organization processed, so only a value
    // this API handed out ever resolves.
    Cursor string `json:"cursor"`
    // Refresh exactly these organizations in one call instead of walking all of
    // them.
    OrganizationIds []string `json:"organization_ids"`

    // Used by Decode() method
    data []byte
}

func (model OrganizationMetricsRefreshRequest) New(data []byte) *OrganizationMetricsRefreshRequest {
    model.data = data
    return &model
}

func (model *OrganizationMetricsRefreshRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}