package models

import (
    "encoding/json"
    "errors"
)

// BlKkliMutationResponseLikeSuccessFlagPlusTheFullReMaterializedEditorState
// Model
type MutationResponse struct {
    // Full editor state (see pages.editor.state).
    State interface{} `json:"state"`
    // 
    Success bool `json:"success"`
    // 
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