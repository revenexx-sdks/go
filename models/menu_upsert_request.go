package models

import (
    "encoding/json"
    "errors"
)

// CreateOrUpdateTheMenuIdentifiedByMenuKeyIdempotentPerTenantItemsIsTheOrderedNavTreeLabelToItems
// Model
type MenuUpsertRequest struct {
    // Ordered menu entries ({ label, to?, items? }).
    Items []interface{} `json:"items"`
    // 
    Label string `json:"label"`
    // Stable menu identifier, e.g. "main", "footer", "account".
    MenuKey string `json:"menuKey"`

    // Used by Decode() method
    data []byte
}

func (model MenuUpsertRequest) New(data []byte) *MenuUpsertRequest {
    model.data = data
    return &model
}

func (model *MenuUpsertRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}