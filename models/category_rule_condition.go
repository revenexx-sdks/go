package models

import (
    "encoding/json"
    "errors"
)

// CategoryRuleCondition model.
type CategoryRuleCondition struct {
    // A product column (sku, kind, enabled, family_id, parent_id) or
    // 'attribute:<code>' for the common bucket of attribute_values. An attribute
    // code is [A-Za-z0-9_]+. Locale-/channel-scoped attributes are not supported.
    Field string `json:"field"`
    // How to compare. 'eq'/'neq' are equality, 'gt'/'gte'/'lt'/'lte' order
    // (numerically for a number, as text for a string), 'in' membership,
    // 'contains'/'starts_with'/'ends_with' substring, 'is_empty'/'is_not_empty'
    // presence — those last two take no `value`.
    Operator string `json:"operator"`
    // Comparison value. An array for 'in' — non-empty, at most 200 entries, all
    // of the same type; omitted for 'is_empty'/'is_not_empty'; a non-empty string
    // for 'contains'/'starts_with'/'ends_with'; a string or number for
    // gt/gte/lt/lte. Numbers compare numerically (jsonb), strings as text.
    Value string `json:"value"`

    // Used by Decode() method
    data []byte
}

func (model CategoryRuleCondition) New(data []byte) *CategoryRuleCondition {
    model.data = data
    return &model
}

func (model *CategoryRuleCondition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}