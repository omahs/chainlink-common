package utils

import (
	"fmt"

	"github.com/K-Phoen/grabana/alert"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/gauge"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/table"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func Inc(p *uint32) uint32 {
	*p++
	return *p
}

func datasourceRef(uid string) dashboard.DataSourceRef {
	return dashboard.DataSourceRef{Uid: &uid}
}

func AddVars(builder *dashboard.DashboardBuilder, items []cog.Builder[dashboard.VariableModel]) {
	for _, item := range items {
		builder.WithVariable(item)
	}
}

func AddPanels(builder *dashboard.DashboardBuilder, items []cog.Builder[dashboard.Panel]) {
	for _, item := range items {
		builder.WithPanel(item)
	}
}

func QueryVariable(datasource string, name string, label string, query string, multi bool) *dashboard.QueryVariableBuilder {
	return dashboard.NewQueryVariableBuilder(name).
		Label(label).
		Query(dashboard.StringOrAny{String: cog.ToPtr[string](query)}).
		Datasource(datasourceRef(datasource)).
		Current(dashboard.VariableOption{
			Selected: cog.ToPtr[bool](true),
			Text:     dashboard.StringOrArrayOfString{ArrayOfString: []string{"All"}},
			Value:    dashboard.StringOrArrayOfString{ArrayOfString: []string{"$__all"}},
		}).
		// Refresh(dashboard.VariableRefreshOnTimeRangeChanged).
		Sort(dashboard.VariableSortAlphabeticalAsc).
		Multi(multi).
		IncludeAll(true)
}

func IntervalVariable(name string, label string, interval string) *dashboard.IntervalVariableBuilder {
	return dashboard.NewIntervalVariableBuilder(name).
		Label(label).
		Values(dashboard.StringOrAny{String: cog.ToPtr[string](interval)}).
		Current(dashboard.VariableOption{
			Selected: cog.ToPtr[bool](true),
			Text:     dashboard.StringOrArrayOfString{ArrayOfString: []string{"All"}},
			Value:    dashboard.StringOrArrayOfString{ArrayOfString: []string{"$__all"}},
		})
}

func prometheusQuery(query string, legend string) *prometheus.DataqueryBuilder {
	return prometheus.NewDataqueryBuilder().
		Expr(query).
		LegendFormat(legend)
}

type PrometheusQuery struct {
	Query  string
	Legend string
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func StatPanel(
	id uint32,
	datasource string,
	title string,
	description string,
	height uint32,
	span uint32,
	decimals float64,
	unit string,
	colorMode common.BigValueColorMode,
	graphMode common.BigValueGraphMode,
	textMode common.BigValueTextMode,
	orientation common.VizOrientation,
	query ...PrometheusQuery,
) *stat.PanelBuilder {
	panel := stat.NewPanelBuilder().
		Id(id).
		Datasource(datasourceRef(datasource)).
		Height(height).
		Span(span).
		Decimals(decimals).
		ReduceOptions(common.NewReduceDataOptionsBuilder().Calcs([]string{"last"})).
		ColorMode(colorMode).
		GraphMode(graphMode).
		TextMode(textMode).
		Title(title).
		Description(description).
		Unit(unit).
		Orientation(orientation)

	for _, q := range query {
		panel.WithTarget(prometheusQuery(q.Query, q.Legend))
	}

	return panel
}

func TimeSeriesPanel(
	id uint32,
	datasource string,
	title string,
	description string,
	height uint32,
	span uint32,
	decimals float64,
	unit string,
	legendPlacement common.LegendPlacement,
	query ...PrometheusQuery,
) *timeseries.PanelBuilder {
	panel := timeseries.NewPanelBuilder().
		Id(id).
		Datasource(datasourceRef(datasource)).
		Height(height).
		Span(span).
		Decimals(decimals).
		Title(title).
		Description(description).
		Unit(unit).
		FillOpacity(2)
	/*Legend(common.NewVizLegendOptionsBuilder().
		Placement(legendPlacement).
		ShowLegend(true).
		IsVisible(true).
		DisplayMode(common.LegendDisplayModeList),
	)*/

	for _, q := range query {
		panel.WithTarget(prometheusQuery(q.Query, q.Legend))
	}

	return panel
}

func GaugePanel(
	id uint32,
	datasource string,
	title string,
	description string,
	height uint32,
	span uint32,
	decimals float64,
	unit string,
	query ...PrometheusQuery,
) *gauge.PanelBuilder {
	panel := gauge.NewPanelBuilder().
		Id(id).
		Datasource(datasourceRef(datasource)).
		Height(height).
		Span(span).
		Decimals(decimals).
		Title(title).
		Description(description).
		Unit(unit)

	for _, q := range query {
		panel.WithTarget(prometheusQuery(q.Query, q.Legend))
	}

	return panel
}

func TablePanel(
	id uint32,
	datasource string,
	title string,
	description string,
	height uint32,
	span uint32,
	decimals float64,
	unit string,
	query ...PrometheusQuery,
) *table.PanelBuilder {
	panel := table.NewPanelBuilder().
		Id(id).
		Datasource(datasourceRef(datasource)).
		Height(height).
		Span(span).
		Decimals(decimals).
		Title(title).
		Description(description).
		Unit(unit)

	for _, q := range query {
		panel.WithTarget(prometheusQuery(q.Query, q.Legend))
	}

	return panel
}

type AlertPrometheusQuery struct {
	Ref       string
	Query     string
	Condition alert.Option
}

func AlertGrafana(
	panelID uint32,
	alertName string,
	summary string,
	description string,
	forTimeRange string,
	runbookUrl string,
	tags map[string]string,
	query ...AlertPrometheusQuery,
) *alert.Alert {

	options := []alert.Option{
		alert.Summary(summary),
		alert.Description(description),
		alert.EvaluateEvery("30s"),
		alert.For(forTimeRange),
		alert.OnNoData(alert.NoDataEmpty),
		alert.OnExecutionError(alert.ErrorAlerting),
		alert.Runbook(runbookUrl),
		alert.Tags(tags),
	}

	for _, q := range query {
		options = append(options, alert.WithPrometheusQuery(q.Ref, q.Query))
		if q.Condition != nil {
			options = append(options, q.Condition)
		}
	}

	newAlert := alert.New(alertName, options...)
	newAlert.HookPanelID(fmt.Sprintf("%d", panelID))

	return newAlert
}
