package models

import (
    "encoding/json"
    "errors"
)

// PageMutationStatusRequest Which entry of the history to switch, and to
// what.
type PageMutationStatusRequest struct {
    // Whether the entry takes part in the replay.
    Enabled bool `json:"enabled"`
    // The position in the mutation log to switch. Unknown positions answer 404.
    Index int `json:"index"`
    // Which language the returned state should be resolved for.
    Langcode string `json:"langcode"`

    // Used by Decode() method
    data []byte
}

func (model PageMutationStatusRequest) New(data []byte) *PageMutationStatusRequest {
    model.data = data
    return &model
}

func (model *PageMutationStatusRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}