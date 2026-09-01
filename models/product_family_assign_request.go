package models

import (
    "encoding/json"
    "errors"
)

// ProductFamilyAssignRequest Name the family either way — `family_id` wins
// when both are sent. The family has to exist already; this route assigns
// one, it does not create one.
type ProductFamilyAssignRequest struct {
    // Alternative to family_id — a `families.code` this tenant holds, from `GET
    // /products/families`. No example: a code is tenant data, and any value
    // published here names a family somebody does not have.
    FamilyCode string `json:"family_code"`
    // The family to assign.
    FamilyId string `json:"family_id"`

    // Used by Decode() method
    data []byte
}

func (model ProductFamilyAssignRequest) New(data []byte) *ProductFamilyAssignRequest {
    model.data = data
    return &model
}

func (model *ProductFamilyAssignRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}