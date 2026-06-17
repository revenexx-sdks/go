package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type PageLibraryItemUpdateRequest struct {
    // 
    Bundle string `json:"bundle"`
    // 
    Label string `json:"label"`
    // Serialized block tree ({ bundle, props, props_i18n, options, children }).
    Tree interface{} `json:"tree"`

    // Used by Decode() method
    data []byte
}

func (model PageLibraryItemUpdateRequest) New(data []byte) *PageLibraryItemUpdateRequest {
    model.data = data
    return &model
}

func (model *PageLibraryItemUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}