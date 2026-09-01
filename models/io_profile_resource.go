package models

import (
    "encoding/json"
    "errors"
)

// IoProfileResource A saved profile. Mirrors the controller's presenter
// exactly — there
// are no `created_at` / `updated_at` fields on this resource.
type IoProfileResource struct {
    // 
    App string `json:"app"`
    // 
    ApplyMode string `json:"apply_mode"`
    // 
    CreatedBy string `json:"created_by"`
    // 
    Direction string `json:"direction"`
    // 
    Entity string `json:"entity"`
    // 
    Format IoProfileFormat `json:"format"`
    // 
    Id string `json:"id"`
    // 
    Mapping interface{} `json:"mapping"`
    // `null` means global — offered for every market.
    Markets []string `json:"markets"`
    // 
    Name string `json:"name"`
    // 
    Options interface{} `json:"options"`
    // 
    Vendor string `json:"vendor"`

    // Used by Decode() method
    data []byte
}

func (model IoProfileResource) New(data []byte) *IoProfileResource {
    model.data = data
    return &model
}

func (model *IoProfileResource) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}