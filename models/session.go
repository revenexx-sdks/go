package models

import (
    "encoding/json"
    "errors"
)

// Session Model
type Session struct {
    // Session creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Session ID.
    Id string `json:"$id"`
    // Session update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Client code name. View list of [available
    // options](https://github.com/appwrite/appwrite/blob/master/docs/lists/clients.json).
    ClientCode string `json:"clientCode"`
    // Client engine name.
    ClientEngine string `json:"clientEngine"`
    // Client engine name.
    ClientEngineVersion string `json:"clientEngineVersion"`
    // Client name.
    ClientName string `json:"clientName"`
    // Client type.
    ClientType string `json:"clientType"`
    // Client version.
    ClientVersion string `json:"clientVersion"`
    // Country two-character ISO 3166-1 alpha code.
    CountryCode string `json:"countryCode"`
    // Country name.
    CountryName string `json:"countryName"`
    // Returns true if this the current user session.
    Current bool `json:"current"`
    // Device brand name.
    DeviceBrand string `json:"deviceBrand"`
    // Device model name.
    DeviceModel string `json:"deviceModel"`
    // Device name.
    DeviceName string `json:"deviceName"`
    // Session expiration date in ISO 8601 format.
    Expire string `json:"expire"`
    // Returns a list of active session factors.
    Factors []string `json:"factors"`
    // IP in use when the session was created.
    Ip string `json:"ip"`
    // Most recent date in ISO 8601 format when the session successfully passed
    // MFA challenge.
    MfaUpdatedAt string `json:"mfaUpdatedAt"`
    // Operating system code name. View list of [available
    // options](https://github.com/appwrite/appwrite/blob/master/docs/lists/os.json).
    OsCode string `json:"osCode"`
    // Operating system name.
    OsName string `json:"osName"`
    // Operating system version.
    OsVersion string `json:"osVersion"`
    // Session Provider.
    Provider string `json:"provider"`
    // Session Provider Access Token.
    ProviderAccessToken string `json:"providerAccessToken"`
    // The date of when the access token expires in ISO 8601 format.
    ProviderAccessTokenExpiry string `json:"providerAccessTokenExpiry"`
    // Session Provider Refresh Token.
    ProviderRefreshToken string `json:"providerRefreshToken"`
    // Session Provider User ID.
    ProviderUid string `json:"providerUid"`
    // Secret used to authenticate the user. Only included if the request was made
    // with an API key
    Secret string `json:"secret"`
    // User ID.
    UserId string `json:"userId"`

    // Used by Decode() method
    data []byte
}

func (model Session) New(data []byte) *Session {
    model.data = data
    return &model
}

func (model *Session) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}