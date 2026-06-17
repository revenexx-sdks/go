package models

import (
    "encoding/json"
    "errors"
)

// PublishedPageResolvedForOneLanguageNestedBlockTreeWithI18nFallbackAppliedAndScheduledBlocksFiltered
// Model
type DeliveryPage struct {
    // Field name → ordered block list ({ uuid, bundle, props, options, children
    // }).
    Fields interface{} `json:"fields"`
    // 
    Page interface{} `json:"page"`

    // Used by Decode() method
    data []byte
}

func (model DeliveryPage) New(data []byte) *DeliveryPage {
    model.data = data
    return &model
}

func (model *DeliveryPage) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}