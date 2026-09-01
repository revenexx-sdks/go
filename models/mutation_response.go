package models

import (
    "encoding/json"
    "errors"
)

// MutationResponse blökkli MutationResponseLike: whether the call was
// applied, plus the FULL re-materialized editor state — so a client never
// has to re-fetch after a change.
type MutationResponse struct {
    // Everything the blökkli editor runs on, for one page in one language,
    // materialized at the current point of the undo history. The theme adapter
    // maps it 1:1 onto blökkli's MappedState.
    State EditorState `json:"state"`
    // Whether the change was applied.
    Success bool `json:"success"`
    // Why the call was refused, when `success` is false.
    Violations []interface{} `json:"violations"`

    // Used by Decode() method
    data []byte
}

func (model MutationResponse) New(data []byte) *MutationResponse {
    model.data = data
    return &model
}

func (model *MutationResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}