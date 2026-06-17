package models

import (
    "encoding/json"
    "errors"
)

// PlatformAuthSessionTreatSecretAsACredentialTheTrustedBFFStoresItServerSideHTTPOnlyCookieNeverInTheBrowser
// Model
type AuthSession struct {
    // 
    Id string `json:"$id"`
    // 
    Expire string `json:"expire"`
    // 
    Provider string `json:"provider"`
    // 
    Secret string `json:"secret"`
    // 
    UserId string `json:"userId"`

    // Used by Decode() method
    data []byte
}

func (model AuthSession) New(data []byte) *AuthSession {
    model.data = data
    return &model
}

func (model *AuthSession) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}