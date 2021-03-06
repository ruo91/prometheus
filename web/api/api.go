// Copyright 2013 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	clientmodel "github.com/prometheus/client_golang/model"

	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage/local"
	"github.com/prometheus/prometheus/util/httputil"
	"github.com/prometheus/prometheus/util/route"
)

// MetricsService manages the /api HTTP endpoint.
type MetricsService struct {
	Now         func() clientmodel.Timestamp
	Storage     local.Storage
	QueryEngine *promql.Engine
}

// RegisterHandler registers the handler for the various endpoints below /api.
func (msrv *MetricsService) RegisterHandler(router *route.Router) {
	router.Get("/query", handle("query", msrv.Query))
	router.Get("/query_range", handle("query_range", msrv.QueryRange))
	router.Get("/metrics", handle("metrics", msrv.Metrics))
}

func handle(name string, f http.HandlerFunc) http.HandlerFunc {
	h := httputil.CompressionHandler{
		Handler: f,
	}
	return prometheus.InstrumentHandler(name, h)
}
