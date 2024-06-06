package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/K-Phoen/grabana"
	"github.com/K-Phoen/grabana/alert"
	"github.com/K-Phoen/grabana/alertmanager"
	"github.com/K-Phoen/grabana/alertmanager/slack"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/spf13/cobra"

	atlasdon "github.com/smartcontractkit/chainlink-common/observability-lib/atlas-don"
	corenode "github.com/smartcontractkit/chainlink-common/observability-lib/core-node"
	corenodecomponents "github.com/smartcontractkit/chainlink-common/observability-lib/core-node-components"
	k8sresources "github.com/smartcontractkit/chainlink-common/observability-lib/k8s-resources"
	"github.com/smartcontractkit/chainlink-common/observability-lib/utils"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy Grafana Dashboards, Prometheus Alerts",
	RunE: func(cmd *cobra.Command, args []string) error {
		dataSources, errDataSources := cmd.Flags().GetStringArray("grafana-data-sources")
		if errDataSources != nil || len(dataSources) < 1 {
			return errDataSources
		}

		dataSourcesType := SetDataSources(dataSources)
		name := cmd.Flag("dashboard-name").Value.String()
		platform := cmd.Flag("platform").Value.String()
		url := cmd.Flag("grafana-url").Value.String()
		folder := cmd.Flag("dashboard-folder").Value.String()
		typeDashboard := cmd.Flag("type").Value.String()
		slackWebhook := cmd.Flag("slack-webhook").Value.String()

		alertsTags, errAlertsTags := cmd.Flags().GetStringToString("alerts-tags")
		if errAlertsTags != nil {
			return errors.New("invalid alerts tags")
		}

		var builder dashboard.Dashboard
		var alerts []*alert.Alert
		var err error

		switch typeDashboard {
		case "core-node":
			builder, err = corenode.BuildDashboard(name, dataSourcesType.Metrics, platform)
		case "core-node-components":
			builder, alerts, err = corenodecomponents.BuildDashboard(name, dataSourcesType.Metrics, alertsTags)
		case "core-node-resources":
			builder, err = k8sresources.BuildDashboard(name, dataSourcesType.Metrics, dataSourcesType.Logs)
		case "ocr":
			builder, err = atlasdon.BuildDashboard(name, dataSourcesType.Metrics, platform, typeDashboard)
		case "ocr2":
			builder, err = atlasdon.BuildDashboard(name, dataSourcesType.Metrics, platform, typeDashboard)
		case "ocr3":
			builder, err = atlasdon.BuildDashboard(name, dataSourcesType.Metrics, platform, typeDashboard)
		default:
			return errors.New("invalid dashboard type")
		}
		if err != nil {
			utils.Logger.Error().
				Str("Name", name).
				Str("URL", url).
				Str("Folder", folder).
				Str("Type", typeDashboard).
				Msg("Could not build dashboard")
			return err
		}

		dashboard := NewDashboard(
			name,
			cmd.Flag("grafana-token").Value.String(),
			url,
			folder,
			dataSourcesType,
			platform,
			builder,
			alerts...,
		)
		errDeploy := dashboard.Deploy()
		if errDeploy != nil {
			return errDeploy
		}

		if alertsTags != nil && len(alertsTags) > 0 && slackWebhook != "" {
			var options []alertmanager.Option
			for key, value := range alertsTags {
				title := cases.Title(language.AmericanEnglish, cases.NoLower).String(value)
				options = append(options,
					alertmanager.ContactPoints(
						alertmanager.ContactPoint(
							title,
							slack.Webhook(slackWebhook),
						),
					),
					alertmanager.Routing(
						alertmanager.Policy(title, alertmanager.TagEq(key, value)),
					),
					alertmanager.DefaultContactPoint(title),
				)

				configureAlertManager(url, cmd.Flag("grafana-token").Value.String(), options...)
			}
		}

		return nil
	},
}

func init() {
	DeployCmd.Flags().String("dashboard-name", "", "Name of the dashboard to deploy")
	errName := DeployCmd.MarkFlagRequired("dashboard-name")
	if errName != nil {
		panic(errName)
	}
	DeployCmd.Flags().String("dashboard-folder", "", "Dashboard folder")
	errFolder := DeployCmd.MarkFlagRequired("dashboard-folder")
	if errFolder != nil {
		panic(errFolder)
	}
	DeployCmd.Flags().String("grafana-url", "", "Grafana URL")
	errURL := DeployCmd.MarkFlagRequired("grafana-url")
	if errURL != nil {
		panic(errURL)
	}
	DeployCmd.Flags().String("grafana-token", "", "Grafana API token")
	errToken := DeployCmd.MarkFlagRequired("grafana-token")
	if errToken != nil {
		panic(errToken)
	}
	DeployCmd.Flags().StringArray("grafana-data-sources", []string{"Prometheus"}, "Data sources to add to the dashboard, at least one required")
	DeployCmd.Flags().String("platform", "docker", "Platform where the dashboard is deployed (docker or kubernetes)")
	DeployCmd.Flags().String("type", "core-node", "Dashboard type can be either core-node | core-node-components")
	DeployCmd.Flags().String("slack-webhook", "", "Slack webhook to send alerts")
	DeployCmd.Flags().StringToString("alerts-tags", map[string]string{
		"team": "chainlink-team",
	}, "Alerts tags")
}

func configureAlertManager(url string, token string, alertManagerOptions ...alertmanager.Option) {
	ctx := context.Background()
	client := grabana.NewClient(
		&http.Client{},
		url,
		grabana.WithAPIToken(token),
	)

	manager := alertmanager.New(
		alertManagerOptions...,
	)

	if err := client.ConfigureAlertManager(ctx, manager); err != nil {
		fmt.Printf("Could not configure alerting: %s\n", err)
		os.Exit(1)
	}
}
