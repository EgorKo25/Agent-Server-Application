package storage

type Gauge float64
type Counter uint64

type Storagier interface {
	GetStats(map[string]any)
	TakeStats() (map[string]Gauge, map[string]Counter)
	TakeThisStat(string) any
}

type MetricStorage struct {
	MetricsGauge   map[string]Gauge
	MetricsCounter map[string]Counter
}

func (m *MetricStorage) GetStats(name string, value any, mType string) {
	if mType == "gauge" {
		(*m).MetricsGauge[name] = value.(Gauge)
	}
	if mType == "counter" {
		(*m).MetricsCounter[name] += value.(Counter)

	}

}
func (m *MetricStorage) TakeStats() (map[string]Gauge, map[string]Counter) {
	return (*m).MetricsGauge, (*m).MetricsCounter
}
func (m *MetricStorage) TakeThisStat(name, mType string) (value any) {
	if mType == "gauge" {
		if _, ok := (*m).MetricsGauge[name]; ok {
			value = (*m).MetricsGauge[name]
			return value
		}
		return nil
	}
	if mType == "counter" {
		if _, ok := (*m).MetricsCounter[name]; ok {
			value = (*m).MetricsCounter[name]
			return value

		}
		return nil
	}

	return nil
}
func NewStorage() *MetricStorage {
	var m MetricStorage
	m.MetricsCounter = make(map[string]Counter)
	m.MetricsGauge = make(map[string]Gauge)
	return &m
}
