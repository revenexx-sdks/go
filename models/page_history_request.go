package models

import (
    "encoding/json"
    "errors"
)

// PageHistoryRequest Where to put the undo pointer.
type PageHistoryRequest struct {
    // The position in the mutation log to materialize at. `-1` undoes everything;
    // the last position redoes everything. Values outside the log are clamped
    // rather than refused.
    Index int `json:"index"`
    // Which language the returned state should be resolved for.
    Langcode string `json:"langcode"`

    // Used by Decode() method
    data []byte
}

func (model PageHistoryRequest) New(data []byte) *PageHistoryRequest {
    model.data = data
    return &model
}

func (model *PageHistoryRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}