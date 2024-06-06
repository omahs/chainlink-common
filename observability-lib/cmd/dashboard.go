package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/K-Phoen/grabana"
	"github.com/K-Phoen/grabana/alert"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/smartcontractkit/chainlink-common/observability-lib/utils"
	"github.com/smartcontractkit/chainlink-testing-framework/grafana"
)

type Dashboard struct {
	name        string
	token       string
	url         string
	folder      string
	dataSources DataSources
	platform    string
	builder     dashboard.Dashboard
	alerts      []*alert.Alert
}

type DataSources struct {
	Metrics string
	Logs    string
}

func SetDataSources(dataSources []string) DataSources {
	var dataSourcesType DataSources
	for _, datasource := range dataSources {
		if datasource == "Prometheus" || datasource == "prometheus" || datasource == "Thanos" || datasource == "thanos" {
			dataSourcesType.Metrics = datasource
		} else if datasource == "Loki" || datasource == "loki" {
			dataSourcesType.Logs = datasource
		}
	}

	return dataSourcesType
}

func NewDashboard(
	name string,
	token string,
	url string,
	folder string,
	dataSources DataSources,
	platform string,
	builder dashboard.Dashboard,
	alerts ...*alert.Alert,
) *Dashboard {
	return &Dashboard{
		name:        name,
		token:       token,
		url:         url,
		folder:      folder,
		dataSources: dataSources,
		platform:    platform,
		builder:     builder,
		alerts:      alerts,
	}
}

func (d *Dashboard) Deploy() error {
	ctx := context.Background()
	client := grabana.NewClient(
		&http.Client{},
		d.url,
		grabana.WithAPIToken(d.token),
	)

	folder, errFolder := client.FindOrCreateFolder(ctx, d.folder)
	if errFolder != nil {
		return errFolder
	}

	grafanaClient := grafana.NewGrafanaClient(
		d.url,
		d.token,
	)
	resp, _, errPostDashboard := grafanaClient.PostDashboard(grafana.PostDashboardRequest{
		Dashboard: d.builder,
		Overwrite: true,
		FolderID:  int(folder.ID),
	})
	if errPostDashboard != nil {
		return errPostDashboard
	}

	utils.Logger.Info().
		Str("Name", d.name).
		Str("URL", d.url).
		Str("Folder", d.folder).
		Msg("Dashboard deployed")

	boardFromGrafana, _, errGetDashboard := grafanaClient.GetDashboard(*resp.UID)
	if errGetDashboard != nil {
		return fmt.Errorf("error getting dashboard: %w", errGetDashboard)
	}

	datasourcesMap, _, errGetDataSources := grafanaClient.GetDatasources()
	if errGetDataSources != nil {
		return nil
	}

	// Create alerts for the dashboard
	if len(d.alerts) > 0 {
		for i := range d.alerts {
			newAlert := *d.alerts[i]
			newAlert.HookDashboardUID(*boardFromGrafana.Dashboard.Uid)

			if err := client.AddAlert(ctx, folder.UID, newAlert, datasourcesMap); err != nil {
				return fmt.Errorf("error creating alert: %w", err)
			}
		}
		utils.Logger.Info().
			Str("URL", d.url+"/alerting/list").
			Str("Folder", d.folder).
			Msg("Alerts deployed")
	}

	return nil
}

func (d *Dashboard) GetJSON() ([]byte, error) {
	newDashboard := d.builder
	dashboardJSON, err := json.MarshalIndent(newDashboard, "", "  ")
	if err != nil {
		return nil, err
	}

	return dashboardJSON, nil
}
