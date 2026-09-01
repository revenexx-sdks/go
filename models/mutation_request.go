package models

import (
    "encoding/json"
    "errors"
)

// MutationRequest One change to the page.
type MutationRequest struct {
    // Which language the returned state should be resolved for. Not the language
    // the change is written in — that lives in the payload.
    Langcode string `json:"langcode"`
    // The arguments of that change; the keys depend on the plugin (`add` takes `{
    // bundle, hostEntityType, hostEntityUuid, hostField }`, `move` takes `{ uuid,
    // preceedingUuid }`, and so on). Anything non-deterministic in it — new
    // uuids, a library item's tree, a copied subtree — is resolved once here
    // and stored, so replaying the log is deterministic forever.
    Payload interface{} `json:"payload"`
    // Which kind of change this is — `add`, `move`, `delete`, `duplicate`,
    // `update_field_value`, `update_options`, … An id this app does not
    // implement is refused with 400 rather than stored, because the log has to
    // replay.
    Plugin string `json:"plugin"`

    // Used by Decode() method
    data []byte
}

func (model MutationRequest) New(data []byte) *MutationRequest {
    model.data = data
    return &model
}

func (model *MutationRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}