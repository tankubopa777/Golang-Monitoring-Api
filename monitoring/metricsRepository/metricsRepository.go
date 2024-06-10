package metricsRepository

import (
	"tansan/monitoring/metricsModel"

	"gorm.io/gorm"
)

type MetricsRepository interface {
    SaveMetrics(metrics metricsModel.Metrics) error
    GetMetrics(deviceID string) ([]metricsModel.Metrics, error)
}

type metricsRepository struct {
    DB *gorm.DB
}

func NewMetricsRepository(db *gorm.DB) MetricsRepository {
    return &metricsRepository{DB: db}
}

func (r *metricsRepository) SaveMetrics(metrics metricsModel.Metrics) error {
    return r.DB.Create(&metrics).Error
}

func (r *metricsRepository) GetMetrics(deviceID string) ([]metricsModel.Metrics, error) {
    var metrics []metricsModel.Metrics
    if err := r.DB.Where("device_id = ?", deviceID).Find(&metrics).Error; err != nil {
        return nil, err
    }
    return metrics, nil
}
