package models

import (
    "encoding/json"
    "errors"
)

// CategoryRulesRequest model.
type CategoryRulesRequest struct {
    // Between 1 and 25 conditions — a rule is a selector, not a query language.
    // An empty list is a 400, not "everything".
    Conditions []CategoryRuleCondition `json:"conditions"`
    // 'all' ANDs every condition (default), 'any' ORs them.
    RuleMatch string `json:"rule_match"`

    // Used by Decode() method
    data []byte
}

func (model CategoryRulesRequest) New(data []byte) *CategoryRulesRequest {
    model.data = data
    return &model
}

func (model *CategoryRulesRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}