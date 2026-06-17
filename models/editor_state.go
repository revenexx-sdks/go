package models

import (
    "encoding/json"
    "errors"
)

// TheBlKkliAdapterStatePageTranslationsEditStateMutationLogMaterializedFieldListsMutatedOptionsEntityValuesTextFieldValuesDroppableFieldValuesAndViolations
// Model
type EditorState struct {
    // 
    CurrentUserIsOwner bool `json:"currentUserIsOwner"`
    // 
    DroppableFieldValues []interface{} `json:"droppableFieldValues"`
    // 
    EditState interface{} `json:"editState"`
    // 
    Fields []interface{} `json:"fields"`
    // 
    IgnoredAnalyzeIdentifiers []string `json:"ignoredAnalyzeIdentifiers"`
    // 
    Langcode string `json:"langcode"`
    // 
    MutatedEntity interface{} `json:"mutatedEntity"`
    // 
    MutatedHostOptions interface{} `json:"mutatedHostOptions"`
    // 
    MutatedOptions interface{} `json:"mutatedOptions"`
    // 
    Mutations []interface{} `json:"mutations"`
    // 
    Page interface{} `json:"page"`
    // 
    TextFieldValues []interface{} `json:"textFieldValues"`
    // 
    Translations []interface{} `json:"translations"`
    // 
    Violations []interface{} `json:"violations"`

    // Used by Decode() method
    data []byte
}

func (model EditorState) New(data []byte) *EditorState {
    model.data = data
    return &model
}

func (model *EditorState) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}