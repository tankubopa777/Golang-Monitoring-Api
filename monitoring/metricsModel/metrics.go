package metricsModel

type Metrics struct {
    ID          string  `json:"id"`
    DeviceID    string  `json:"device_id"`
    MetricType  string  `json:"metric_type"`
    Value       float64 `json:"value"`
    Timestamp   int64   `json:"timestamp"`
}
