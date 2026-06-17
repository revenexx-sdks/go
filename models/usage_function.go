package models

import (
    "encoding/json"
    "errors"
)

// UsageFunction Model
type UsageFunction struct {
    // Aggregated number of function builds per period.
    Builds []Metric `json:"builds"`
    // Aggregated number of failed builds per period.
    BuildsFailed []Metric `json:"buildsFailed"`
    // Total aggregated number of failed function builds.
    BuildsFailedTotal int `json:"buildsFailedTotal"`
    // Aggregated number of function builds mbSeconds per period.
    BuildsMbSeconds []Metric `json:"buildsMbSeconds"`
    // Total aggregated sum of function builds mbSeconds.
    BuildsMbSecondsTotal int `json:"buildsMbSecondsTotal"`
    // Aggregated sum of function builds storage per period.
    BuildsStorage []Metric `json:"buildsStorage"`
    // total aggregated sum of function builds storage.
    BuildsStorageTotal int `json:"buildsStorageTotal"`
    // Aggregated number of successful builds per period.
    BuildsSuccess []Metric `json:"buildsSuccess"`
    // Total aggregated number of successful function builds.
    BuildsSuccessTotal int `json:"buildsSuccessTotal"`
    // Aggregated sum of function builds compute time per period.
    BuildsTime []Metric `json:"buildsTime"`
    // Average builds compute time.
    BuildsTimeAverage int `json:"buildsTimeAverage"`
    // Total aggregated sum of function builds compute time.
    BuildsTimeTotal int `json:"buildsTimeTotal"`
    // Total aggregated number of function builds.
    BuildsTotal int `json:"buildsTotal"`
    // Aggregated number of function deployments per period.
    Deployments []Metric `json:"deployments"`
    // Aggregated number of  function deployments storage per period.
    DeploymentsStorage []Metric `json:"deploymentsStorage"`
    // Total aggregated sum of function deployments storage.
    DeploymentsStorageTotal int `json:"deploymentsStorageTotal"`
    // Total aggregated number of function deployments.
    DeploymentsTotal int `json:"deploymentsTotal"`
    // Aggregated number of function executions per period.
    Executions []Metric `json:"executions"`
    // Aggregated number of function mbSeconds per period.
    ExecutionsMbSeconds []Metric `json:"executionsMbSeconds"`
    // Total aggregated sum of function executions mbSeconds.
    ExecutionsMbSecondsTotal int `json:"executionsMbSecondsTotal"`
    // Aggregated number of function executions compute time per period.
    ExecutionsTime []Metric `json:"executionsTime"`
    // Total aggregated sum of function  executions compute time.
    ExecutionsTimeTotal int `json:"executionsTimeTotal"`
    // Total  aggregated number of function executions.
    ExecutionsTotal int `json:"executionsTotal"`
    // The time range of the usage stats.
    Range string `json:"range"`

    // Used by Decode() method
    data []byte
}

func (model UsageFunction) New(data []byte) *UsageFunction {
    model.data = data
    return &model
}

func (model *UsageFunction) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}