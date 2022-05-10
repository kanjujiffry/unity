package statscollector

type StatsCollector struct{}

func New(timeout int) *StatsCollector {
	return &StatsCollector{}
}

func (st *StatsCollector) Record(responseTimeMs ...int) error {
	return nil
}

func (st *StatsCollector) Median() (float64, error) {
	return 0, nil
}

func (st *StatsCollector) Average() (float64, error) {
	return 0, nil
}
