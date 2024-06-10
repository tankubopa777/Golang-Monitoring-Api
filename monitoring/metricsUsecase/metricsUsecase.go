package metricsUsecase

import (
	"tansan/monitoring/metricsModel"
	"tansan/monitoring/metricsRepository"
)

type MetricsUsecase interface {
    CreateMetrics(metrics metricsModel.Metrics) error
    FetchMetrics(deviceID string) ([]metricsModel.Metrics, error)
}

type metricsUsecase struct {
    Repo metricsRepository.MetricsRepository
}

func NewMetricsUsecase(repo metricsRepository.MetricsRepository) MetricsUsecase {
    return &metricsUsecase{Repo: repo}
}

func (u *metricsUsecase) CreateMetrics(metrics metricsModel.Metrics) error {
    return u.Repo.SaveMetrics(metrics)
}

func (u *metricsUsecase) FetchMetrics(deviceID string) ([]metricsModel.Metrics, error) {
    return u.Repo.GetMetrics(deviceID)
}
