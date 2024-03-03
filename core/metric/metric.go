package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	prometheus.Registerer
	prometheus.Gatherer

	BulkRegister(...prometheus.Collector) error
}

var (
	// TypeHTTP ...
	TypeHTTP = "http"
	// TypeGRPCUnary ...
	TypeGRPCUnary = "unary"
	// TypeGRPCStream ...
	TypeGRPCStream = "stream"
	// TypeRedis ...
	TypeRedis = "redis"
	TypeGorm  = "gorm"
	// TypeRocketMQ ...
	TypeRocketMQ = "rocketmq"
	// TypeWebsocket ...
	TypeWebsocket = "ws"

	// TypeMySQL ...
	TypeMySQL = "mysql"
	// TypeLocalCache ...
	TypeLocalCache = "localCache"

	// TypeTableStore
	TypeTableStore = "tstore"
	// CodeJob
	CodeJobSuccess = "ok"
	// CodeJobFail ...
	CodeJobFail = "fail"
	// CodeJobReentry ...
	CodeJobReentry = "reentry"

	// CodeCache
	CodeCacheMiss = "miss"
	// CodeCacheHit ...
	CodeCacheHit = "hit"
)

var (
	// ServerHandleCounter ...
	ServerHandleCounter = CounterVecOpts{
		Namespace: DefaultNamespace,
		Name:      "server_handle_total",
		Labels:    []string{"type", "method", "client", "code"},
	}.Build()

	// ServerHandleHistogram ...
	ServerHandleHistogram = HistogramVecOpts{
		Namespace: DefaultNamespace,
		Name:      "server_handle_seconds",
		Labels:    []string{"type", "method", "client"},
	}.Build()

	// ClientHandleCounter ...
	ClientHandleCounter = CounterVecOpts{
		Namespace: DefaultNamespace,
		Name:      "client_handle_total",
		Labels:    []string{"type", "name", "method", "server", "code"},
	}.Build()

	// ClientHandleHistogram ...
	ClientHandleHistogram = HistogramVecOpts{
		Namespace: DefaultNamespace,
		Name:      "client_handle_seconds",
		Labels:    []string{"type", "name", "method", "server"},
	}.Build()

	// JobHandleCounter ...
	JobHandleCounter = CounterVecOpts{
		Namespace: DefaultNamespace,
		Name:      "job_handle_total",
		Labels:    []string{"type", "name", "code"},
	}.Build()

	// JobHandleHistogram ...
	JobHandleHistogram = HistogramVecOpts{
		Namespace: DefaultNamespace,
		Name:      "job_handle_seconds",
		Labels:    []string{"type", "name"},
	}.Build()

	LibHandleHistogram = HistogramVecOpts{
		Namespace: DefaultNamespace,
		Name:      "lib_handle_seconds",
		Labels:    []string{"type", "method", "address"},
	}.Build()
	// LibHandleCounter ...
	LibHandleCounter = CounterVecOpts{
		Namespace: DefaultNamespace,
		Name:      "lib_handle_total",
		Labels:    []string{"type", "method", "address", "code"},
	}.Build()

	LibHandleSummary = SummaryVecOpts{
		Namespace: DefaultNamespace,
		Name:      "lib_handle_stats",
		Labels:    []string{"name", "status"},
	}.Build()

	// CacheHandleCounter ...
	CacheHandleCounter = CounterVecOpts{
		Namespace: DefaultNamespace,
		Name:      "cache_handle_total",
		Labels:    []string{"type", "name", "method", "code"},
	}.Build()

	// CacheHandleHistogram ...
	CacheHandleHistogram = HistogramVecOpts{
		Namespace: DefaultNamespace,
		Name:      "cache_handle_seconds",
		Labels:    []string{"type", "name", "method"},
	}.Build()

	// BuildInfoGauge ...
	BuildInfoGauge = GaugeVecOpts{
		Namespace: DefaultNamespace,
		Name:      "build_info",
		Labels:    []string{"name", "id", "env", "region", "zone", "version", "go_version"},
		// Labels:    []string{"name", "aid", "mode", "region", "zone", "app_version", "jupiter_version", "start_time", "build_time", "go_version"},
	}.Build()

	// LogLevelCounter ...
	LogLevelCounter = NewCounterVec("log_level_total", []string{"name", "lv"})

	// ClientStatsGauge ...
	ClientStatsGauge = GaugeVecOpts{
		Namespace: DefaultNamespace,
		Name:      "client_stats_gauge",
		Labels:    []string{"type", "name", "index"},
	}.Build()

	// CustomizedHandleGauge 业务自定义 gauge.
	CustomizedHandleGauge = NewGaugeVec("customized_handle_gauge", []string{"type"})
	// CustomizedHandleCounter 业务自定义 counter.
	CustomizedHandleCounter = NewCounterVec("customized_handle_total", []string{"type"})
	// CustomizedHandleHistogram 业务自定义 histogram.
	CustomizedHandleHistogram = NewHistogramVec("customized_handle_seconds", []string{"type"})
)

func init() {
	//conf.OnLoaded(func(c *conf.Configuration) {
	//	BuildInfoGauge.WithLabelValues(
	//		pkg.Name(),
	//		pkg.AppID(),
	//		pkg.AppMode(),
	//		pkg.AppRegion(),
	//		pkg.AppZone(),
	//		pkg.AppVersion(),
	//		pkg.GoVersion(),
	//	).Set(float64(time.Now().UnixNano() / 1e6))
	//})
}
