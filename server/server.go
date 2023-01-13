package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func GetListenAndServeFunc() error {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)

	return nil
}
