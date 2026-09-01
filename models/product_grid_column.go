package models

import (
    "encoding/json"
    "errors"
)

// ProductGridColumn model.
type ProductGridColumn struct {
    // The key to read out of a row: a column name for the fixed columns, an
    // attribute code for the rest (then it is a key of the row's `attributes`
    // object).
    Code string `json:"code"`
    // The attribute's i18n labels, or a plain title for the fixed columns.
    Label interface{} `json:"label"`
    // Where the value comes from: 'column' is a plain products column,
    // 'attribute' a key inside `attribute_values`, 'resolved' something this
    // route computed (the display name).
    Source string `json:"source"`
    // Which control renders the cell — the same widget vocabulary `GET
    // /products/attribute-schema` uses, so one renderer serves both.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model ProductGridColumn) New(data []byte) *ProductGridColumn {
    model.data = data
    return &model
}

func (model *ProductGridColumn) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}