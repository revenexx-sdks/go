package models

import (
    "encoding/json"
    "errors"
)

// FormPostSubmitAction One post-submit action. `webhook` POSTs `{form,
// source, data}` to `url`; `entity` writes the mapped fields into another
// app's entity; `event` is a no-op, because `form.submitted` already carries
// it.
type FormPostSubmitAction struct {
    // Entity actions: the app that owns the target entity, e.g. 'crm'.
    App string `json:"app"`
    // Disabled actions are skipped. An action with no flag is not run.
    Enabled bool `json:"enabled"`
    // Entity actions: the entity to write, e.g. 'contacts'.
    Entity string `json:"entity"`
    // Entity actions: which submitted value becomes which column — `{"source":
    // "email", "target": "email"}` reads `data.email` and writes it to the
    // target's `email`.
    Mapping []FormActionMapping `json:"mapping"`
    // Webhook actions: the HTTP method. Defaults to POST.
    Method string `json:"method"`
    // Entity actions: an explicit route to POST to, instead of the one built from
    // `app` and `entity`.
    Path string `json:"path"`
    // Which action this is: 'webhook', 'entity' or 'event'.
    Type string `json:"type"`
    // Webhook actions: where to POST. It is called with an 8 second timeout and
    // its answer is not shown to the visitor.
    Url string `json:"url"`

    // Used by Decode() method
    data []byte
}

func (model FormPostSubmitAction) New(data []byte) *FormPostSubmitAction {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *FormPostSubmitAction) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}