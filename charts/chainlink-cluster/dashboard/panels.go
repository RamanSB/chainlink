package dashboard

import (
	"github.com/K-Phoen/grabana/dashboard"
	"github.com/K-Phoen/grabana/gauge"
	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/stat"
	"github.com/K-Phoen/grabana/table"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/grabana/timeseries"
	"github.com/K-Phoen/grabana/timeseries/axis"
)

func (m *Dashboard) addMainPanels() {
	var balancePanelSpanSize float32 = 4
	var panelsIncluded []row.Option
	var goVersionLegend string = "version"

	if m.platform == "kubernetes" {
		balancePanelSpanSize = 3
		goVersionLegend = "exported_version"
	}

	globalInfoPanels := []row.Option{
		row.WithStat(
			"App Version",
			stat.DataSource(m.PrometheusDataSourceName),
			stat.Text(stat.TextValueAndName),
			stat.Orientation(stat.OrientationVertical),
			stat.TitleFontSize(12),
			stat.ValueFontSize(20),
			stat.Span(2),
			stat.Text("name"),
			stat.Height("100px"),
			stat.WithPrometheusTarget(
				`version{`+m.panelOption.labelFilter+`=~"$instance"}`,
				prometheus.Legend("{{version}}"),
			),
		),
		row.WithStat(
			"Go Version",
			stat.DataSource(m.PrometheusDataSourceName),
			stat.Text(stat.TextValueAndName),
			stat.Orientation(stat.OrientationVertical),
			stat.TitleFontSize(12),
			stat.ValueFontSize(20),
			stat.Span(2),
			stat.Text("name"),
			stat.Height("100px"),
			stat.WithPrometheusTarget(
				`go_info{`+m.panelOption.labelFilter+`=~"$instance"}`,
				prometheus.Legend("{{"+goVersionLegend+"}}"),
			),
		),
		row.WithStat(
			"Uptime in days",
			stat.DataSource(m.PrometheusDataSourceName),
			stat.Text(stat.TextValueAndName),
			stat.Orientation(stat.OrientationVertical),
			stat.TitleFontSize(12),
			stat.ValueFontSize(20),
			stat.Span(2),
			stat.Height("100px"),
			stat.WithPrometheusTarget(
				`uptime_seconds{`+m.panelOption.labelFilter+`=~"$instance"} / 86400`,
				prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
			),
		),
		row.WithStat(
			"ETH Balance",
			stat.DataSource(m.PrometheusDataSourceName),
			stat.Text(stat.TextValueAndName),
			stat.Orientation(stat.OrientationVertical),
			stat.TitleFontSize(12),
			stat.ValueFontSize(20),
			stat.Span(balancePanelSpanSize),
			stat.Height("100px"),
			stat.Decimals(2),
			stat.WithPrometheusTarget(
				`eth_balance{`+m.panelOption.labelFilter+`=~"$instance"}`,
				prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{account}}"),
			),
		),
		row.WithStat(
			"Solana Balance",
			stat.DataSource(m.PrometheusDataSourceName),
			stat.Text(stat.TextValueAndName),
			stat.Orientation(stat.OrientationVertical),
			stat.TitleFontSize(12),
			stat.ValueFontSize(20),
			stat.Span(balancePanelSpanSize),
			stat.Height("100px"),
			stat.Decimals(2),
			stat.WithPrometheusTarget(
				`solana_balance{`+m.panelOption.labelFilter+`=~"$instance"}`,
				prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{account}}"),
			),
		),
	}

	additionalPanels := []row.Option{
		row.WithTimeSeries(
			"Service Components Health",
			timeseries.Span(12),
			timeseries.Height("200px"),
			timeseries.DataSource(m.PrometheusDataSourceName),
			timeseries.WithPrometheusTarget(
				`health{`+m.panelOption.labelFilter+`=~"$instance"}`,
				prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{service_id}}"),
			),
		),
		row.WithTimeSeries(
			"ETH Balance",
			timeseries.Span(6),
			timeseries.Height("200px"),
			timeseries.DataSource(m.PrometheusDataSourceName),
			timeseries.Axis(
				axis.Unit(""),
				axis.Decimals(2),
			),
			timeseries.WithPrometheusTarget(
				`eth_balance{`+m.panelOption.labelFilter+`=~"$instance"}`,
				prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{account}}"),
			),
		),
		row.WithTimeSeries(
			"SOL Balance",
			timeseries.Span(6),
			timeseries.Height("200px"),
			timeseries.DataSource(m.PrometheusDataSourceName),
			timeseries.Axis(
				axis.Unit(""),
				axis.Decimals(2),
			),
			timeseries.WithPrometheusTarget(
				`solana_balance{`+m.panelOption.labelFilter+`=~"$instance"}`,
				prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{account}}"),
			),
		),
	}

	panelsIncluded = append(panelsIncluded, globalInfoPanels...)
	if m.platform == "kubernetes" {
		panelsIncluded = append(panelsIncluded, row.WithStat(
			"Pod Restarts",
			stat.Span(2),
			stat.Height("100px"),
			stat.DataSource(m.PrometheusDataSourceName),
			stat.WithPrometheusTarget(
				`sum(increase(kube_pod_container_status_restarts_total{pod=~"$instance.*", namespace=~"${namespace}"}[$__rate_interval])) by (pod)`,
				prometheus.Legend("{{pod}}"),
			),
		))
	}
	panelsIncluded = append(panelsIncluded, additionalPanels...)

	opts := []dashboard.Option{
		dashboard.Row(
			"Global health",
			panelsIncluded...,
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addKubePanels() {
	opts := []dashboard.Option{
		dashboard.Row(
			"Pod health",
			row.WithStat(
				"Pod Restarts",
				stat.Span(4),
				stat.Height("200px"),
				stat.DataSource(m.PrometheusDataSourceName),
				stat.SparkLine(),
				stat.SparkLineYMin(0),
				stat.WithPrometheusTarget(
					`sum(increase(kube_pod_container_status_restarts_total{pod=~"$instance.*", namespace=~"${namespace}"}[$__rate_interval])) by (pod)`,
					prometheus.Legend("{{pod}}"),
				),
			),
			row.WithStat(
				"OOM Events",
				stat.Span(4),
				stat.Height("200px"),
				stat.DataSource(m.PrometheusDataSourceName),
				stat.SparkLine(),
				stat.SparkLineYMin(0),
				stat.WithPrometheusTarget(
					`sum(container_oom_events_total{pod=~"$instance.*", namespace=~"${namespace}"}) by (pod)`,
					prometheus.Legend("{{pod}}"),
				),
			),
			row.WithStat(
				"OOM Killed",
				stat.Span(4),
				stat.Height("200px"),
				stat.DataSource(m.PrometheusDataSourceName),
				stat.SparkLine(),
				stat.SparkLineYMin(0),
				stat.WithPrometheusTarget(
					`kube_pod_container_status_last_terminated_reason{reason="OOMKilled", pod=~"$instance.*", namespace=~"${namespace}"}`,
					prometheus.Legend("{{pod}}"),
				),
			),
			row.WithTimeSeries(
				"CPU Usage",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`sum(node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{pod=~"$instance.*", namespace=~"${namespace}"}) by (pod)`,
					prometheus.Legend("{{pod}}"),
				),
			),
			row.WithTimeSeries(
				"Memory Usage",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("bytes"),
					axis.Label("Memory"),
					axis.SoftMin(0),
				),
				timeseries.WithPrometheusTarget(
					`sum(container_memory_rss{pod=~"$instance.*", namespace=~"${namespace}", container!=""}) by (pod)`,
					prometheus.Legend("{{pod}}"),
				),
			),
			row.WithTimeSeries(
				"Receive Bandwidth",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Bps"),
					axis.SoftMin(0),
				),
				timeseries.WithPrometheusTarget(
					`sum(irate(container_network_receive_bytes_total{pod=~"$instance.*", namespace=~"${namespace}"}[$__rate_interval])) by (pod)`,
					prometheus.Legend("{{pod}}"),
				),
			),
			row.WithTimeSeries(
				"Transmit Bandwidth",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Bps"),
					axis.SoftMin(0),
				),
				timeseries.WithPrometheusTarget(
					`sum(irate(container_network_transmit_bytes_total{pod=~"$instance.*", namespace=~"${namespace}"}[$__rate_interval])) by (pod)`,
					prometheus.Legend("{{pod}}"),
				),
			),
			row.WithTimeSeries(
				"Average Container Bandwidth by Namespace: Received",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Bps"),
					axis.SoftMin(0),
				),
				timeseries.WithPrometheusTarget(
					`avg(irate(container_network_receive_bytes_total{pod=~"$instance.*", namespace=~"${namespace}"}[$__rate_interval])) by (pod)`,
					prometheus.Legend("{{pod}}"),
				),
			),
			row.WithTimeSeries(
				"Average Container Bandwidth by Namespace: Transmitted",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Bps"),
					axis.SoftMin(0),
				),
				timeseries.WithPrometheusTarget(
					`avg(irate(container_network_transmit_bytes_total{pod=~"$instance.*", namespace=~"${namespace}"}[$__rate_interval])) by (pod)`,
					prometheus.Legend("{{pod}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addLogPollerPanels() {
	opts := []dashboard.Option{
		dashboard.Row("LogPoller",
			row.Collapse(),
			row.WithTimeSeries(
				"LogPoller RPS",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`avg(sum(rate(log_poller_query_duration_count{`+m.panelOption.labelFilter+`=~"$instance", evmChainID=~"$evmChainID"}[$__rate_interval])) by (query, instance)) by (query)`,
					prometheus.Legend("{{query}}"),
				),
				timeseries.WithPrometheusTarget(
					`avg(sum(rate(log_poller_query_duration_count{`+m.panelOption.labelFilter+`=~"$instance", evmChainID=~"$evmChainID"}[$__rate_interval]))) by (instance)`,
					prometheus.Legend("Total"),
				),
			),
			row.WithTimeSeries(
				"LogPoller Logs Number Returned",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`log_poller_query_dataset_size{`+m.panelOption.labelFilter+`=~"$instance", evmChainID=~"$evmChainID"}`,
					prometheus.Legend("{{query}} : {{type}}"),
				),
			),
			row.WithTimeSeries(
				"LogPoller Average Logs Number Returned",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`avg(log_poller_query_dataset_size{`+m.panelOption.labelFilter+`=~"$instance", evmChainID=~"$evmChainID"}) by (query)`,
					prometheus.Legend("{{query}}"),
				),
			),
			row.WithTimeSeries(
				"LogPoller Max Logs Number Returned",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`max(log_poller_query_dataset_size{`+m.panelOption.labelFilter+`=~"$instance", evmChainID=~"$evmChainID"}) by (query)`,
					prometheus.Legend("{{query}}"),
				),
			),
			row.WithTimeSeries(
				"LogPoller Logs Number Returned by Chain",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`max(log_poller_query_dataset_size{`+m.panelOption.labelFilter+`=~"$instance"}) by (evmChainID)`,
					prometheus.Legend("{{evmChainID}}"),
				),
			),
			row.WithTimeSeries(
				"LogPoller Queries Duration Avg",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`(sum(rate(log_poller_query_duration_sum{`+m.panelOption.labelFilter+`=~"$instance", evmChainID=~"$evmChainID"}[$__rate_interval])) by (query) / sum(rate(log_poller_query_duration_count{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (query)) / 1e6`,
					prometheus.Legend("{{query}}"),
				),
			),
			row.WithTimeSeries(
				"LogPoller Queries Duration p99",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`histogram_quantile(0.99, sum(rate(log_poller_query_duration_bucket{`+m.panelOption.labelFilter+`=~"$instance", evmChainID=~"$evmChainID"}[$__rate_interval])) by (le, query)) / 1e6`,
					prometheus.Legend("{{query}}"),
				),
			),
			row.WithTimeSeries(
				"LogPoller Queries Duration p95",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`histogram_quantile(0.95, sum(rate(log_poller_query_duration_bucket{`+m.panelOption.labelFilter+`=~"$instance", evmChainID=~"$evmChainID"}[$__rate_interval])) by (le, query)) / 1e6`,
					prometheus.Legend("{{query}}"),
				),
			),
			row.WithTimeSeries(
				"LogPoller Queries Duration p90",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`histogram_quantile(0.95, sum(rate(log_poller_query_duration_bucket{`+m.panelOption.labelFilter+`=~"$instance", evmChainID=~"$evmChainID"}[$__rate_interval])) by (le, query)) / 1e6`,
					prometheus.Legend("{{query}}"),
				),
			),
			row.WithTimeSeries(
				"LogPoller Queries Duration Median",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`histogram_quantile(0.5, sum(rate(log_poller_query_duration_bucket{`+m.panelOption.labelFilter+`=~"$instance", evmChainID=~"$evmChainID"}[$__rate_interval])) by (le, query)) / 1e6`,
					prometheus.Legend("{{query}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addFeedsJobsPanels() {
	opts := []dashboard.Option{
		dashboard.Row("Feeds Jobs",
			row.Collapse(),
			row.WithTimeSeries(
				"Feeds Job Proposal Requests",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`feeds_job_proposal_requests{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Feeds Job Proposal Count",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`feeds_job_proposal_count{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addMailboxPanels() {
	opts := []dashboard.Option{
		dashboard.Row("Mailbox",
			row.Collapse(),
			row.WithTimeSeries(
				"Mailbox Load Percent",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`mailbox_load_percent{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{ name }}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addPromReporterPanels() {
	opts := []dashboard.Option{
		dashboard.Row("Prom Reporter",
			row.Collapse(),
			row.WithTimeSeries(
				"Unconfirmed Transactions",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Tx"),
				),
				timeseries.WithPrometheusTarget(
					`unconfirmed_transactions{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Unconfirmed TX Age",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Sec"),
				),
				timeseries.WithPrometheusTarget(
					`max_unconfirmed_tx_age{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Unconfirmed TX Blocks",
				timeseries.Span(4),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Blocks"),
				),
				timeseries.WithPrometheusTarget(
					`max_unconfirmed_blocks{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addTxManagerPanels() {
	opts := []dashboard.Option{
		dashboard.Row("TX Manager",
			row.Collapse(),
			row.WithTimeSeries(
				"TX Manager Time Until TX Broadcast",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`tx_manager_time_until_tx_broadcast{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"TX Manager Num Gas Bumps",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`tx_manager_num_gas_bumps{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"TX Manager Num Gas Bumps Exceeds Limit",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`tx_manager_gas_bump_exceeds_limit{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"TX Manager Num Confirmed Transactions",
				timeseries.Span(3),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`tx_manager_num_confirmed_transactions{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"TX Manager Num Successful Transactions",
				timeseries.Span(3),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`tx_manager_num_successful_transactions{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"TX Manager Num Reverted Transactions",
				timeseries.Span(3),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`tx_manager_num_tx_reverted{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"TX Manager Num Fwd Transactions",
				timeseries.Span(3),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`tx_manager_fwd_tx_count{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"TX Manager Num Transactions Attempts",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`tx_manager_tx_attempt_count{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"TX Manager Time Until TX Confirmed",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`tx_manager_time_until_tx_confirmed{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"TX Manager Block Until TX Confirmed",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`tx_manager_blocks_until_tx_confirmed{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addHeadTrackerPanels() {
	opts := []dashboard.Option{
		dashboard.Row("Head tracker",
			row.Collapse(),
			row.WithTimeSeries(
				"Head tracker current head",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Block"),
				),
				timeseries.WithPrometheusTarget(
					`head_tracker_current_head{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Head tracker very old head",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Block"),
				),
				timeseries.WithPrometheusTarget(
					`head_tracker_very_old_head{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Head tracker heads received",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Block"),
				),
				timeseries.WithPrometheusTarget(
					`head_tracker_heads_received{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Head tracker connection errors",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Block"),
				),
				timeseries.WithPrometheusTarget(
					`head_tracker_connection_errors{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addDatabasePanels() {
	opts := []dashboard.Option{
		// DB Metrics
		dashboard.Row("DB Connection Metrics (App)",
			row.Collapse(),
			row.WithTimeSeries(
				"DB Connections",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Conn"),
				),
				timeseries.WithPrometheusTarget(
					`db_conns_max{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Max"),
				),
				timeseries.WithPrometheusTarget(
					`db_conns_open{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Open"),
				),
				timeseries.WithPrometheusTarget(
					`db_conns_used{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Used"),
				),
				timeseries.WithPrometheusTarget(
					`db_conns_wait{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Wait"),
				),
			),
			row.WithTimeSeries(
				"DB Wait Count",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`db_wait_count{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"DB Wait Time",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Sec"),
				),
				timeseries.WithPrometheusTarget(
					`db_wait_time_seconds{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addSQLQueryPanels() {
	opts := []dashboard.Option{
		dashboard.Row(
			"SQL Query",
			row.Collapse(),
			row.WithTimeSeries(
				"SQL Query Timeout Percent",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("percent"),
				),
				timeseries.WithPrometheusTarget(
					`histogram_quantile(0.9, sum(rate(sql_query_timeout_percent_bucket{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (le))`,
					prometheus.Legend("p90"),
				),
				timeseries.WithPrometheusTarget(
					`histogram_quantile(0.95, sum(rate(sql_query_timeout_percent_bucket{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (le))`,
					prometheus.Legend("p95"),
				),
				timeseries.WithPrometheusTarget(
					`histogram_quantile(0.99, sum(rate(sql_query_timeout_percent_bucket{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (le))`,
					prometheus.Legend("p99"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addLogsPanels() {
	opts := []dashboard.Option{
		dashboard.Row("Logs Metrics",
			row.Collapse(),
			row.WithTimeSeries(
				"Logs Counters",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`log_panic_count{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - panic"),
				),
				timeseries.WithPrometheusTarget(
					`log_fatal_count{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - fatal"),
				),
				timeseries.WithPrometheusTarget(
					`log_critical_count{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - critical"),
				),
				timeseries.WithPrometheusTarget(
					`log_warn_count{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - warn"),
				),
				timeseries.WithPrometheusTarget(
					`log_error_count{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - error"),
				),
			),
			row.WithTimeSeries(
				"Logs Rate",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`sum(rate(log_panic_count{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval]))`,
					prometheus.Legend("panic"),
				),
				timeseries.WithPrometheusTarget(
					`sum(rate(log_fatal_count{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval]))`,
					prometheus.Legend("fatal"),
				),
				timeseries.WithPrometheusTarget(
					`sum(rate(log_critical_count{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval]))`,
					prometheus.Legend("critical"),
				),
				timeseries.WithPrometheusTarget(
					`sum(rate(log_warn_count{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval]))`,
					prometheus.Legend("warn"),
				),
				timeseries.WithPrometheusTarget(
					`sum(rate(log_error_count{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval]))`,
					prometheus.Legend("error"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addEVMPoolLifecyclePanels() {
	opts := []dashboard.Option{
		dashboard.Row(
			"EVM Pool Lifecycle",
			row.Collapse(),
			row.WithTimeSeries(
				"EVM Pool Highest Seen Block",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Block"),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_highest_seen_block{`+m.panelOption.labelFilter+`=~"$instance", evmChainID="${evmChainID}"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"EVM Pool Num Seen Blocks",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Block"),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_num_seen_blocks{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"EVM Pool Node Polls Total",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Block"),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_polls_total{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"EVM Pool Node Polls Failed",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Block"),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_polls_failed{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"EVM Pool Node Polls Success",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Block"),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_polls_success{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addEVMPoolRPCNodePanels() {
	opts := []dashboard.Option{
		dashboard.Row(
			"EVM Pool RPC Node Metrics (App)",
			row.Collapse(),
			row.WithTimeSeries(
				"EVM Pool RPC Node Calls Success Rate",
				timeseries.Span(7),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
					axis.Label("%"),
					axis.SoftMin(0),
					axis.SoftMax(100),
				),
				timeseries.WithPrometheusTarget(
					`sum(increase(evm_pool_rpc_node_calls_success{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) / sum(increase(evm_pool_rpc_node_calls_total{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) * 100`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{evmChainID}} - {{nodeName}}"),
				),
			),
			row.WithGauge(
				"EVM Pool RPC Node Calls Success Rate",
				gauge.Span(5),
				gauge.Height("200px"),
				gauge.DataSource(m.PrometheusDataSourceName),
				gauge.Unit("percentunit"),
				gauge.WithPrometheusTarget(
					`sum(increase(evm_pool_rpc_node_calls_success{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) / sum(increase(evm_pool_rpc_node_calls_total{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName)`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{evmChainID}} - {{nodeName}}"),
				),
				gauge.AbsoluteThresholds([]gauge.ThresholdStep{
					{Color: "#ff0000"},
					{Color: "#ffa500", Value: float64Ptr(0.8)},
					{Color: "#00ff00", Value: float64Ptr(0.9)},
				}),
			),
			// issue when value is 0
			row.WithTimeSeries(
				"EVM Pool RPC Node Dials Success Rate",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
					axis.Label("%"),
					axis.SoftMin(0),
					axis.SoftMax(100),
				),
				timeseries.WithPrometheusTarget(
					`sum(increase(evm_pool_rpc_node_dials_success{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) / sum(increase(evm_pool_rpc_node_dials_total{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) * 100`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{evmChainID}} - {{nodeName}}"),
				),
			),
			// issue when value is 0
			row.WithTimeSeries(
				"EVM Pool RPC Node Dials Failure Rate",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
					axis.Label("%"),
					axis.SoftMin(0),
					axis.SoftMax(100),
				),
				timeseries.WithPrometheusTarget(
					`sum(increase(evm_pool_rpc_node_dials_failed{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) / sum(increase(evm_pool_rpc_node_dials_total{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) * 100`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{evmChainID}} - {{nodeName}}"),
				),
			),
			row.WithTimeSeries(
				"EVM Pool RPC Node Transitions",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_num_transitions_to_alive{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend(""),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_num_transitions_to_in_sync{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend(""),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_num_transitions_to_out_of_sync{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend(""),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_num_transitions_to_unreachable{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend(""),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_num_transitions_to_invalid_chain_id{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend(""),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_num_transitions_to_unusable{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend(""),
				),
			),
			row.WithTimeSeries(
				"EVM Pool RPC Node States",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`evm_pool_rpc_node_states{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{evmChainID}} - {{state}}"),
				),
			),
			row.WithTimeSeries(
				"EVM Pool RPC Node Verifies Success Rate",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
					axis.Label("%"),
					axis.SoftMin(0),
					axis.SoftMax(100),
				),
				timeseries.WithPrometheusTarget(
					`sum(increase(evm_pool_rpc_node_verifies_success{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) / sum(increase(evm_pool_rpc_node_verifies{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) * 100`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{evmChainID}} - {{nodeName}}"),
				),
			),
			row.WithTimeSeries(
				"EVM Pool RPC Node Verifies Failure Rate",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
					axis.Label("%"),
					axis.SoftMin(0),
					axis.SoftMax(100),
				),
				timeseries.WithPrometheusTarget(
					`sum(increase(evm_pool_rpc_node_verifies_failed{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) / sum(increase(evm_pool_rpc_node_verifies{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (`+m.panelOption.labelFilter+`, evmChainID, nodeName) * 100`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{evmChainID}} - {{nodeName}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addEVMRPCNodeLatenciesPanels() {
	opts := []dashboard.Option{
		dashboard.Row(
			"EVM Pool RPC Node Latencies (App)",
			row.Collapse(),
			row.WithTimeSeries(
				"EVM Pool RPC Node Calls Latency 0.95 quantile",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("ms"),
				),
				timeseries.WithPrometheusTarget(
					`histogram_quantile(0.95, sum(rate(evm_pool_rpc_node_rpc_call_time_bucket{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (le, rpcCallName)) / 1e6`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addBlockHistoryEstimatorPanels() {
	opts := []dashboard.Option{
		dashboard.Row("Block History Estimator",
			row.Collapse(),
			row.WithTimeSeries(
				"Gas Updater All Gas Price Percentiles",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`gas_updater_all_gas_price_percentiles{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{ percentile }}"),
				),
			),
			row.WithTimeSeries(
				"Gas Updater All Tip Cap Percentiles",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`gas_updater_all_tip_cap_percentiles{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - {{ percentile }}"),
				),
			),
			row.WithTimeSeries(
				"Gas Updater Set Gas Price",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`gas_updater_set_gas_price{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Gas Updater Set Tip Cap",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`gas_updater_set_tip_cap{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Gas Updater Current Base Fee",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`gas_updater_current_base_fee{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Block History Estimator Connectivity Failure Count",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`block_history_estimator_connectivity_failure_count{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addPipelinePanels() {
	opts := []dashboard.Option{
		dashboard.Row("Pipeline Metrics (Runner)",
			row.Collapse(),
			row.WithTimeSeries(
				"Pipeline Task Execution Time",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Sec"),
				),
				timeseries.WithPrometheusTarget(
					`pipeline_task_execution_time{`+m.panelOption.labelFilter+`=~"$instance"} / 1e6`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} JobID: {{ job_id }}"),
				),
			),
			row.WithTimeSeries(
				"Pipeline Run Errors",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`pipeline_run_errors{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} JobID: {{ job_id }}"),
				),
			),
			row.WithTimeSeries(
				"Pipeline Run Total Time to Completion",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Sec"),
				),
				timeseries.WithPrometheusTarget(
					`pipeline_run_total_time_to_completion{`+m.panelOption.labelFilter+`=~"$instance"} / 1e6`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} JobID: {{ job_id }}"),
				),
			),
			row.WithTimeSeries(
				"Pipeline Tasks Total Finished",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`pipeline_tasks_total_finished{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} JobID: {{ job_id }}"),
				),
			),
		),
		dashboard.Row(
			"Pipeline Metrics (ETHCall)",
			row.Collapse(),
			row.WithTimeSeries(
				"Pipeline Task ETH Call Execution Time",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Sec"),
				),
				timeseries.WithPrometheusTarget(
					`pipeline_task_eth_call_execution_time{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
		dashboard.Row(
			"Pipeline Metrics (HTTP)",
			row.Collapse(),
			row.WithTimeSeries(
				"Pipeline Task HTTP Fetch Time",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Sec"),
				),
				timeseries.WithPrometheusTarget(
					`pipeline_task_http_fetch_time{`+m.panelOption.labelFilter+`=~"$instance"} / 1e6`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Pipeline Task HTTP Response Body Size",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Bytes"),
				),
				timeseries.WithPrometheusTarget(
					`pipeline_task_http_response_body_size{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
		dashboard.Row(
			"Pipeline Metrics (Bridge)",
			row.Collapse(),
			row.WithTimeSeries(
				"Pipeline Bridge Latency",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Sec"),
				),
				timeseries.WithPrometheusTarget(
					`bridge_latency_seconds{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Pipeline Bridge Errors Total",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`bridge_errors_total{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Pipeline Bridge Cache Hits Total",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`bridge_cache_hits_total{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Pipeline Bridge Cache Errors Total",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`bridge_cache_errors_total{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
		dashboard.Row(
			"Pipeline Metrics",
			row.Collapse(),
			row.WithTimeSeries(
				"Pipeline Runs Queued",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`pipeline_runs_queued{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Pipeline Runs Tasks Queued",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`pipeline_task_runs_queued{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addHTTPAPIPanels() {
	opts := []dashboard.Option{
		// HTTP API Metrics
		dashboard.Row(
			"HTTP API Metrics",
			row.Collapse(),
			row.WithTimeSeries(
				"Request Duration p95",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Sec"),
				),
				timeseries.WithPrometheusTarget(
					`histogram_quantile(0.95, sum(rate(service_gonic_request_duration_bucket{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (le, path, method))`,
					prometheus.Legend("{{ method }} {{ path }}"),
				),
			),
			row.WithTimeSeries(
				"Request Total Rate over interval",
				timeseries.Span(12),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`sum(rate(service_gonic_requests_total{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (path, method, code)`,
					prometheus.Legend("{{ method }} {{ path }} {{ code }}"),
				),
			),
			row.WithTimeSeries(
				"Request Size",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Bytes"),
				),
				timeseries.WithPrometheusTarget(
					`avg(rate(service_gonic_request_size_bytes_sum{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval]))/avg(rate(service_gonic_request_size_bytes_count{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval]))`,
					prometheus.Legend("Average"),
				),
			),
			row.WithTimeSeries(
				"Response Size",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("Bytes"),
				),
				timeseries.WithPrometheusTarget(
					`avg(rate(service_gonic_response_size_bytes_sum{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval]))/avg(rate(service_gonic_response_size_bytes_count{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval]))`,
					prometheus.Legend("Average"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addPromHTTPPanels() {
	opts := []dashboard.Option{
		dashboard.Row(
			"PromHTTP Metrics",
			row.Collapse(),
			row.WithGauge("HTTP Request in flight",
				gauge.Span(2),
				gauge.Height("200px"),
				gauge.DataSource(m.PrometheusDataSourceName),
				gauge.WithPrometheusTarget(
					`promhttp_metric_handler_requests_in_flight`,
					prometheus.Legend(""),
				),
			),
			row.WithTimeSeries(
				"HTTP rate",
				timeseries.Span(10),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`sum(rate(promhttp_metric_handler_requests_total{`+m.panelOption.labelFilter+`=~"$instance"}[$__rate_interval])) by (code)`,
					prometheus.Legend("{{ code }}"),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addGoMetricsPanels() {
	opts := []dashboard.Option{
		dashboard.Row(
			"Go Metrics",
			row.Collapse(),
			row.WithTable(
				"Threads",
				table.Span(3),
				table.Height("200px"),
				table.DataSource(m.PrometheusDataSourceName),
				table.WithPrometheusTarget(
					`sum(go_threads{`+m.panelOption.labelFilter+`=~"$instance"}) by (instance)`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}")),
				table.HideColumn("Time"),
				table.AsTimeSeriesAggregations([]table.Aggregation{
					{Label: "AVG", Type: table.AVG},
					{Label: "Current", Type: table.Current},
				}),
			),
			row.WithTimeSeries(
				"Threads",
				timeseries.Span(9),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit(""),
				),
				timeseries.WithPrometheusTarget(
					`sum(go_threads{`+m.panelOption.labelFilter+`=~"$instance"}) by (instance)`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
			),
			row.WithTimeSeries(
				"Heap allocations",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`sum(go_memstats_heap_alloc_bytes{`+m.panelOption.labelFilter+`=~"$instance"}) by (instance)`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
				timeseries.Axis(
					axis.Unit("bytes"),
					axis.Label("Memory"),
					axis.SoftMin(0),
				),
			),
			row.WithStat(
				"Heap Allocations",
				stat.Span(6),
				stat.Height("200px"),
				stat.DataSource(m.PrometheusDataSourceName),
				stat.Unit("bytes"),
				stat.ColorValue(),
				stat.WithPrometheusTarget(`sum(go_memstats_heap_alloc_bytes{`+m.panelOption.labelFilter+`=~"$instance"})`),
				/*stat.AbsoluteThresholds([]stat.ThresholdStep{
					{
						Color: "green",
						Value: nil,
					},
					{
						Color: "orange",
						Value: float64Ptr(6.711e+7),
					},
					{
						Color: "red",
						Value: float64Ptr(1.342e+8),
					},
				}),*/
			),
			row.WithTimeSeries(
				"Memory in Heap",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("bytes"),
					axis.Label("Memory"),
					axis.SoftMin(0),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_heap_alloc_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Alloc"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_heap_sys_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Sys"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_heap_idle_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Idle"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_heap_inuse_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - InUse"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_heap_released_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Released"),
				),
			),
			row.WithTimeSeries(
				"Memory in Off-Heap",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.Axis(
					axis.Unit("bytes"),
					axis.Label("Memory"),
					axis.SoftMin(0),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_mspan_inuse_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Total InUse"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_mspan_sys_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Total Sys"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_mcache_inuse_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Cache InUse"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_mcache_sys_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Cache Sys"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_buck_hash_sys_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Hash Sys"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_gc_sys_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - GC Sys"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_other_sys_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - bytes of memory are used for other runtime allocations"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_next_gc_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Next GC"),
				),
			),
			row.WithTimeSeries(
				"Memory in Stack",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`go_memstats_stack_inuse_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - InUse"),
				),
				timeseries.WithPrometheusTarget(
					`go_memstats_stack_sys_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}} - Sys"),
				),
				timeseries.Axis(
					axis.Unit("bytes"),
					axis.Label("Memory"),
					axis.SoftMin(0),
				),
			),
			row.WithTimeSeries(
				"Total Used Memory",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`go_memstats_sys_bytes{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
				timeseries.Axis(
					axis.Unit("bytes"),
					axis.Label("Memory"),
					axis.SoftMin(0),
				),
			),
			row.WithTimeSeries(
				"Number of Live Objects",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`go_memstats_mallocs_total{`+m.panelOption.labelFilter+`=~"$instance"} - go_memstats_frees_total{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
				timeseries.Axis(
					axis.SoftMin(0),
				),
			),
			row.WithTimeSeries(
				"Rate of Objects Allocated",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`rate(go_memstats_mallocs_total{`+m.panelOption.labelFilter+`=~"$instance"}[1m])`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
				timeseries.Axis(
					axis.SoftMin(0),
				),
			),
			row.WithTimeSeries(
				"Rate of a Pointer Dereferences",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`rate(go_memstats_lookups_total{`+m.panelOption.labelFilter+`=~"$instance"}[1m])`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
				timeseries.Axis(
					axis.Unit("ops"),
					axis.SoftMin(0),
				),
			),
			row.WithTimeSeries(
				"Goroutines",
				timeseries.Span(6),
				timeseries.Height("200px"),
				timeseries.DataSource(m.PrometheusDataSourceName),
				timeseries.WithPrometheusTarget(
					`go_goroutines{`+m.panelOption.labelFilter+`=~"$instance"}`,
					prometheus.Legend("{{"+m.panelOption.labelFilter+"}}"),
				),
				timeseries.Axis(
					axis.SoftMin(0),
				),
			),
		),
	}

	m.opts = append(m.opts, opts...)
}

func (m *Dashboard) addCorePanels() {
	m.addMainPanels()
	m.addLogPollerPanels()
	m.addFeedsJobsPanels()
	m.addMailboxPanels()
	m.addPromReporterPanels()
	m.addTxManagerPanels()
	m.addHeadTrackerPanels()
	m.addDatabasePanels()
	m.addSQLQueryPanels()
	m.addLogsPanels()
	m.addEVMPoolLifecyclePanels()
	m.addEVMPoolRPCNodePanels()
	m.addEVMRPCNodeLatenciesPanels()
	m.addBlockHistoryEstimatorPanels()
	m.addPipelinePanels()
	m.addHTTPAPIPanels()
	m.addPromHTTPPanels()
	m.addGoMetricsPanels()
}

func (m *Dashboard) addKubernetesPanels() {
	m.addKubePanels()
}

func float64Ptr(input float64) *float64 {
	return &input
}
