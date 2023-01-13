package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	v1 "k8s.io/api/core/v1"
)


func EmitSubscriptionMetric (oldPod *) (
	subscriptionStatus.WithLabelValues(newCSV.Namespace, newCSV.Name, newCSV.Spec.Version.String(), string(newCSV.Status.Phase), string(newCSV.Status.Reason)).Set(1)
)

var (
	subscriptionStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pod_events",
			Help: "CSV is not installed",
		},
		[]string{NamespaceLabel, NameLabel, PhaseLabel, ReasonLabel},
)

