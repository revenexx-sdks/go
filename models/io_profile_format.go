package models

import (
    "encoding/json"
    "errors"
)

// IoProfileFormat Profile source/sink format. `bmecat` is profile-only —
// the ad-hoc
// `/io/imports` and `/io/exports` endpoints do not accept it.
type IoProfileFormat struct {

    // Used by Decode() method
    data []byte
}

func (model IoProfileFormat) New(data []byte) *IoProfileFormat {
    model.data = data
    return &model
}

func (model *IoProfileFormat) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}