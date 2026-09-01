package models

import (
    "encoding/json"
    "errors"
)

// FormSettings Everything about a form that is not a field: what the
// storefront renders around the inputs, what happens after a successful
// submit, and who is told about it. Open jsonb, so an unknown key is stored
// and handed back rather than refused — the keys below are the ones
// something actually READS, and each says which reader that is. Null on a
// form nobody has configured, which is not an error: every one of these has a
// fallback.
type FormSettings struct {
    // What the storefront runs after a successful submit, in order. Executed by
    // the cover BFF, not by this API — this app only stores them, and a
    // workflow that wants the same event should listen to `form.submitted`
    // instead.
    Actions []FormPostSubmitAction `json:"actions"`
    // The language the definition itself is written in. Read by the storefront
    // BFF, which overlays `i18n` on top of it.
    DefaultLocale string `json:"default_locale"`
    // Translations for the definition, keyed by language tag and then by field
    // name: `{"en": {"email": {"label": "Email"}}}`. Only `label`, `placeholder`
    // and `help` are overlaid — a translation of anything else is stored and
    // ignored. Applied by the storefront BFF before the definition reaches the
    // browser, so the API always returns the untranslated definition.
    I18n interface{} `json:"i18n"`
    // This form's own notification recipient, read by THIS app at insert. It
    // beats the tenant's `notify_email` setting; null means fall back to the
    // tenant. The storefront never sees it — the BFF hands the browser only the
    // submit label and the success message.
    NotifyEmail string `json:"notify_email"`
    // The submit button caption, read by the storefront. Null falls back to
    // 'Submit'.
    SubmitLabel string `json:"submit_label"`
    // What the visitor reads after a successful submit, read by the storefront.
    // Null falls back to a generic thank-you.
    SuccessMessage string `json:"success_message"`

    // Used by Decode() method
    data []byte
}

func (model FormSettings) New(data []byte) *FormSettings {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *FormSettings) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}