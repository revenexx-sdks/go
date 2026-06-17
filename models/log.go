package models

import (
    "encoding/json"
    "errors"
)

// Log Model
type Log struct {
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
    // Device brand name.
    DeviceBrand string `json:"deviceBrand"`
    // Device model name.
    DeviceModel string `json:"deviceModel"`
    // Device name.
    DeviceName string `json:"deviceName"`
    // Event name.
    Event string `json:"event"`
    // IP session in use when the session was created.
    Ip string `json:"ip"`
    // API mode when event triggered.
    Mode string `json:"mode"`
    // Operating system code name. View list of [available
    // options](https://github.com/appwrite/appwrite/blob/master/docs/lists/os.json).
    OsCode string `json:"osCode"`
    // Operating system name.
    OsName string `json:"osName"`
    // Operating system version.
    OsVersion string `json:"osVersion"`
    // Log creation date in ISO 8601 format.
    Time string `json:"time"`
    // User Email.
    UserEmail string `json:"userEmail"`
    // User ID.
    UserId string `json:"userId"`
    // User Name.
    UserName string `json:"userName"`

    // Used by Decode() method
    data []byte
}

func (model Log) New(data []byte) *Log {
    model.data = data
    return &model
}

func (model *Log) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}