package corenodecomponents

import (
	"github.com/K-Phoen/grabana/alert"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/observability-lib/utils"
)

var panelId uint32

func BuildDashboard(name string, dataSourceMetric string, alertsTags map[string]string) (dashboard.Dashboard, []*alert.Alert, error) {
	props := Props{
		MetricsDataSource: dataSourceMetric,
		PlatformOpts:      PlatformPanelOpts(),
		AlertsTags:        alertsTags,
	}

	builder := dashboard.NewDashboardBuilder(name).
		Tags([]string{"Core", "Node", "Components"}).
		Refresh("30s").
		Time("now-30m", "now")

	utils.AddVars(builder, vars(props))
	panels, alerts := panelsGeneralInfo(props)
	utils.AddPanels(builder, panels)

	newDashboard, err := builder.Build()

	return newDashboard, alerts, err
}

func vars(p Props) []cog.Builder[dashboard.VariableModel] {
	var variables []cog.Builder[dashboard.VariableModel]

	variables = append(variables,
		utils.IntervalVariable("interval", "Interval", "30s,1m,5m,15m,30m,1h,6h,12h"))
	variables = append(variables,
		utils.QueryVariable(p.MetricsDataSource, "env", "Environment", `label_values(up, env)`, false))
	variables = append(variables,
		utils.QueryVariable(p.MetricsDataSource, "cluster", "Cluster", `label_values(up{env="$env"}, cluster)`, false))
	variables = append(variables,
		utils.QueryVariable(p.MetricsDataSource, "blockchain", "Blockchain", `label_values(up{env="$env", cluster="$cluster"}, blockchain)`, false))
	variables = append(variables,
		utils.QueryVariable(p.MetricsDataSource, "product", "Product", `label_values(up{env="$env", cluster="$cluster", blockchain="$blockchain"}, product)`, false))
	variables = append(variables,
		utils.QueryVariable(p.MetricsDataSource, "network_type", "Network Type", `label_values(up{env="$env", cluster="$cluster", blockchain="$blockchain", product="$product"}, network_type)`, false))
	variables = append(variables,
		utils.QueryVariable(p.MetricsDataSource, "component", "Component", `label_values(up{env="$env", cluster="$cluster", blockchain="$blockchain", network_type="$network_type"}, component)`, false))
	variables = append(variables,
		utils.QueryVariable(p.MetricsDataSource, "service", "Service", `label_values(up{env="$env", cluster="$cluster", blockchain="$blockchain", network_type="$network_type", component="$component"}, service)`, false))
	variables = append(variables,
		utils.QueryVariable(p.MetricsDataSource, "service_id", "Service ID", `label_values(health{cluster="$cluster", blockchain="$blockchain", network_type="$network_type", component="$component", service="$service"}, service_id)`, true))

	return variables
}

func panelsGeneralInfo(p Props) ([]cog.Builder[dashboard.Panel], []*alert.Alert) {
	var panelsArray []cog.Builder[dashboard.Panel]
	var alertsArray []*alert.Alert

	commonTags := map[string]string{
		"severity":      "info",
		"slack_channel": "team-blockchain-foundation-alerts-info",
		"runbook":       "https://github.com/smartcontractkit/alerts/blob/master/runbooks/blockchain-foundation/core-node-components.md",
		"dashboardUrl":  "https://grafana.ops.prod.cldev.sh/d/c6d19f97-4ead-4917-aa11-cdbe5a3ef6f9/corenodecomponents?orgId=1&refresh=30s",
	}
	maps.Copy(commonTags, p.AlertsTags)

	panelsArray = append(panelsArray, utils.TablePanel(
		utils.Inc(&panelId),
		p.MetricsDataSource,
		"List Nodes",
		"",
		4,
		24,
		1,
		"",
	).WithTarget(prometheus.NewDataqueryBuilder().
		Expr(`max(up{`+p.PlatformOpts.LabelQuery+`}) by (env, cluster, blockchain, product, network_type, network, version, team, component, service)`).
		LegendFormat("").
		Format("table").
		Instant(true),
	))

	panelsArray = append(panelsArray, utils.TimeSeriesPanel(
		utils.Inc(&panelId),
		p.MetricsDataSource,
		"Node Uptime",
		"",
		4,
		24,
		1,
		"percent",
		common.LegendPlacementBottom,
		utils.PrometheusQuery{
			Query:  `100 * up{` + p.PlatformOpts.LabelQuery + `}`,
			Legend: "Team: {{team}} env: {{env}} cluster: {{cluster}} namespace: {{namespace}} job: {{job}} blockchain: {{blockchain}} product: {{product}} networkType: {{network_type}} component: {{component}} service: {{service}}",
		},
	).Min(0).Max(100))

	panelsArray = append(panelsArray, utils.StatPanel(
		utils.Inc(&panelId),
		p.MetricsDataSource,
		"Components Health Avg by Service",
		"",
		4,
		24,
		1,
		"percent",
		common.BigValueColorModeValue,
		common.BigValueGraphModeLine,
		common.BigValueTextModeValueAndName,
		common.VizOrientationVertical,
		utils.PrometheusQuery{
			Query:  `100 * avg(avg_over_time(health{` + p.PlatformOpts.LabelQuery + `service_id=~"${service_id}"}[$interval])) by (service_id, version, service, cluster, env)`,
			Legend: "{{service_id}}",
		},
	).Thresholds(
		dashboard.NewThresholdsConfigBuilder().
			Mode(dashboard.ThresholdsModeAbsolute).
			Steps([]dashboard.Threshold{
				{Value: utils.Float64Ptr(0), Color: "red"},
				{Value: utils.Float64Ptr(80), Color: "orange"},
				{Value: utils.Float64Ptr(99), Color: "green"},
			})),
	)

	panelsArray = append(panelsArray, utils.TimeSeriesPanel(
		utils.Inc(&panelId),
		p.MetricsDataSource,
		"Components Health by Service",
		"",
		6,
		24,
		1,
		"percent",
		common.LegendPlacementRight,
		utils.PrometheusQuery{
			Query:  `100 * (health{` + p.PlatformOpts.LabelQuery + `service_id=~"${service_id}"})`,
			Legend: "{{service_id}}",
		},
	).Min(0).Max(100))

	panelsArray = append(panelsArray, utils.TimeSeriesPanel(
		utils.Inc(&panelId),
		p.MetricsDataSource,
		"Components Health Avg by Service",
		"",
		6,
		24,
		1,
		"percent",
		common.LegendPlacementRight,
		utils.PrometheusQuery{
			Query:  `100 * (avg(avg_over_time(health{` + p.PlatformOpts.LabelQuery + `service_id=~"${service_id}"}[$interval])) by (service_id, version, service, cluster, env))`,
			Legend: "{{service_id}}",
		},
	).Min(0).Max(100))

	alertsArray = append(alertsArray, utils.AlertGrafana(
		panelId,
		"Uptime less than 90% over last 15 minutes on one component in a Node",
		"Component {{$labels.service_id}} uptime in the last 15m is {{$value}}%",
		`env: {{$labels.env}}
cluster: {{$labels.cluster}}
namespace: {{$labels.namespace}}
job: {{$labels.job}}
blockchain: {{$labels.blockchain}}
product: {{$labels.product}}
networkType: {{$labels.network_type}}
component: {{$labels.component}}`,
		"1m",
		"https://github.com/smartcontractkit/alerts/blob/master/runbooks/blockchain-foundation/core-node-components.md",
		commonTags,
		utils.AlertPrometheusQuery{
			Ref:       "A",
			Query:     `100 * avg(avg_over_time(health{service!~".*-preview", env="production"}[15m])) by (service_id, version, service, blockchain, network_type, cluster, env)`,
			Condition: alert.If(alert.Avg, "A", alert.IsBelow(90)),
		},
	))

	return panelsArray, alertsArray
}
