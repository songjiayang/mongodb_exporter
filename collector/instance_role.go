package collector

import "github.com/prometheus/client_golang/prometheus"

var (
	instanceRoleGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "instance_role",
		Help:      "The store of mongodb instance role",
	}, []string{"host"})
)

type ReplStatus struct {
	Host     string `bson:"me"`
	IsMaster bool   `bson:"ismaster"`
}

func (replStatus *ReplStatus) Export(ch chan<- prometheus.Metric) {
	ls := prometheus.Labels{
		"host": replStatus.Host,
	}

	if replStatus.IsMaster {
		instanceRoleGauge.With(ls).Set(1)
	} else {
		instanceRoleGauge.With(ls).Set(0)
	}

	instanceRoleGauge.Collect(ch)
}

func (replStatus *ReplStatus) Describe(ch chan<- *prometheus.Desc) {
	instanceRoleGauge.Describe(ch)
}
