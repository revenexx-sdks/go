package models

import (
    "encoding/json"
    "errors"
)

// UsageFunctions Model
type UsageFunctions struct {
    // Aggregated number of functions build per period.
    Builds []Metric `json:"builds"`
    // Aggregated number of failed function builds per period.
    BuildsFailed []Metric `json:"buildsFailed"`
    // Total aggregated number of failed function builds.
    BuildsFailedTotal int `json:"buildsFailedTotal"`
    // Aggregated sum of functions build mbSeconds per period.
    BuildsMbSeconds []Metric `json:"buildsMbSeconds"`
    // Total aggregated sum of functions build mbSeconds.
    BuildsMbSecondsTotal int `json:"buildsMbSecondsTotal"`
    // Aggregated sum of functions build storage per period.
    BuildsStorage []Metric `json:"buildsStorage"`
    // total aggregated sum of functions build storage.
    BuildsStorageTotal int `json:"buildsStorageTotal"`
    // Aggregated number of successful function builds per period.
    BuildsSuccess []Metric `json:"buildsSuccess"`
    // Total aggregated number of successful function builds.
    BuildsSuccessTotal int `json:"buildsSuccessTotal"`
    // Aggregated sum of  functions build compute time per period.
    BuildsTime []Metric `json:"buildsTime"`
    // Total aggregated sum of functions build compute time.
    BuildsTimeTotal int `json:"buildsTimeTotal"`
    // Total aggregated number of functions build.
    BuildsTotal int `json:"buildsTotal"`
    // Aggregated number of functions deployment per period.
    Deployments []Metric `json:"deployments"`
    // Aggregated number of  functions deployment storage per period.
    DeploymentsStorage []Metric `json:"deploymentsStorage"`
    // Total aggregated sum of functions deployment storage.
    DeploymentsStorageTotal int `json:"deploymentsStorageTotal"`
    // Total aggregated number of functions deployments.
    DeploymentsTotal int `json:"deploymentsTotal"`
    // Aggregated number of  functions execution per period.
    Executions []Metric `json:"executions"`
    // Aggregated number of functions mbSeconds per period.
    ExecutionsMbSeconds []Metric `json:"executionsMbSeconds"`
    // Total aggregated sum of functions execution mbSeconds.
    ExecutionsMbSecondsTotal int `json:"executionsMbSecondsTotal"`
    // Aggregated number of functions execution compute time per period.
    ExecutionsTime []Metric `json:"executionsTime"`
    // Total aggregated sum of functions  execution compute time.
    ExecutionsTimeTotal int `json:"executionsTimeTotal"`
    // Total  aggregated number of functions execution.
    ExecutionsTotal int `json:"executionsTotal"`
    // Aggregated number of functions per period.
    Functions []Metric `json:"functions"`
    // Total aggregated number of functions.
    FunctionsTotal int `json:"functionsTotal"`
    // Time range of the usage stats.
    Range string `json:"range"`

    // Used by Decode() method
    data []byte
}

func (model UsageFunctions) New(data []byte) *UsageFunctions {
    model.data = data
    return &model
}

func (model *UsageFunctions) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}