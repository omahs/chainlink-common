package loop

import (
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvConfig_parse(t *testing.T) {
	cases := []struct {
		name        string
		envVars     map[string]string
		expectError bool

		expectedDatabaseURL                    string
		expectedDatabaseIdleInTxSessionTimeout time.Duration
		expectedDatabaseLockTimeout            time.Duration
		expectedDatabaseQueryTimeout           time.Duration
		expectedDatabaseLogSQL                 bool
		expectedDatabaseMaxOpenConns           int
		expectedDatabaseMaxIdleConns           int

		expectedPrometheusPort         int
		expectedTracingEnabled         bool
		expectedTracingCollectorTarget string
		expectedTracingSamplingRatio   float64
		expectedTracingTLSCertPath     string
	}{
		{
			name: "All variables set correctly",
			envVars: map[string]string{
				envDatabaseURL:                    "postgres://user:password@localhost:5432/db",
				envDatabaseIdleInTxSessionTimeout: "42s",
				envDatabaseLockTimeout:            "8m",
				envDatabaseQueryTimeout:           "7s",
				envDatabaseLogSQL:                 "true",
				envDatabaseMaxOpenConns:           "9999",
				envDatabaseMaxIdleConns:           "8080",

				envPromPort:                 "8080",
				envTracingEnabled:           "true",
				envTracingCollectorTarget:   "some:target",
				envTracingSamplingRatio:     "1.0",
				envTracingTLSCertPath:       "internal/test/fixtures/client.pem",
				envTracingAttribute + "XYZ": "value",
			},
			expectError: false,

			expectedDatabaseURL:                    "postgres://user:password@localhost:5432/db",
			expectedDatabaseIdleInTxSessionTimeout: 42 * time.Second,
			expectedDatabaseLockTimeout:            8 * time.Minute,
			expectedDatabaseQueryTimeout:           7 * time.Second,
			expectedDatabaseLogSQL:                 true,
			expectedDatabaseMaxOpenConns:           9999,
			expectedDatabaseMaxIdleConns:           8080,

			expectedPrometheusPort:         8080,
			expectedTracingEnabled:         true,
			expectedTracingCollectorTarget: "some:target",
			expectedTracingSamplingRatio:   1.0,
			expectedTracingTLSCertPath:     "internal/test/fixtures/client.pem",
		},
		{
			name: "CL_DATABASE_URL parse error",
			envVars: map[string]string{
				envDatabaseURL: "wrong-db-url",
			},
			expectError: true,
		},
		{
			name: "CL_PROMETHEUS_PORT parse error",
			envVars: map[string]string{
				envPromPort: "abc",
			},
			expectError: true,
		},
		{
			name: "TRACING_ENABLED parse error",
			envVars: map[string]string{
				envPromPort:       "8080",
				envTracingEnabled: "invalid_bool",
			},
			expectError: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.envVars {
				t.Setenv(k, v)
			}

			var config EnvConfig
			err := config.parse()

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				} else {
					if config.DatabaseURL.String() != tc.expectedDatabaseURL {
						t.Errorf("Expected Database URL %s, got %s", tc.expectedDatabaseURL, config.DatabaseURL)
					}
					if config.DatabaseIdleInTxSessionTimeout != tc.expectedDatabaseIdleInTxSessionTimeout {
						t.Errorf("Expected Database idle in tx session timeout %s, got %s", tc.expectedDatabaseIdleInTxSessionTimeout, config.DatabaseIdleInTxSessionTimeout)
					}
					if config.DatabaseLockTimeout != tc.expectedDatabaseLockTimeout {
						t.Errorf("Expected Database lock timeout %s, got %s", tc.expectedDatabaseLockTimeout, config.DatabaseLockTimeout)
					}
					if config.DatabaseQueryTimeout != tc.expectedDatabaseQueryTimeout {
						t.Errorf("Expected Database query timeout %s, got %s", tc.expectedDatabaseQueryTimeout, config.DatabaseQueryTimeout)
					}
					if config.DatabaseLogSQL != tc.expectedDatabaseLogSQL {
						t.Errorf("Expected Database log sql %t, got %t", tc.expectedDatabaseLogSQL, config.DatabaseLogSQL)
					}
					if config.DatabaseMaxOpenConns != tc.expectedDatabaseMaxOpenConns {
						t.Errorf("Expected Database max open conns %d, got %d", tc.expectedDatabaseMaxOpenConns, config.DatabaseMaxOpenConns)
					}
					if config.DatabaseMaxIdleConns != tc.expectedDatabaseMaxIdleConns {
						t.Errorf("Expected Database max idle conns %d, got %d", tc.expectedDatabaseMaxIdleConns, config.DatabaseMaxIdleConns)
					}

					if config.PrometheusPort != tc.expectedPrometheusPort {
						t.Errorf("Expected Prometheus port %d, got %d", tc.expectedPrometheusPort, config.PrometheusPort)
					}
					if config.TracingEnabled != tc.expectedTracingEnabled {
						t.Errorf("Expected tracingEnabled %v, got %v", tc.expectedTracingEnabled, config.TracingEnabled)
					}
					if config.TracingCollectorTarget != tc.expectedTracingCollectorTarget {
						t.Errorf("Expected tracingCollectorTarget %s, got %s", tc.expectedTracingCollectorTarget, config.TracingCollectorTarget)
					}
					if config.TracingSamplingRatio != tc.expectedTracingSamplingRatio {
						t.Errorf("Expected tracingSamplingRatio %f, got %f", tc.expectedTracingSamplingRatio, config.TracingSamplingRatio)
					}
					if config.TracingTLSCertPath != tc.expectedTracingTLSCertPath {
						t.Errorf("Expected tracingTLSCertPath %s, got %s", tc.expectedTracingTLSCertPath, config.TracingTLSCertPath)
					}
				}
			}
		})
	}
}

func TestEnvConfig_AsCmdEnv(t *testing.T) {
	envCfg := EnvConfig{
		DatabaseURL:            &url.URL{Scheme: "postgres", Host: "localhost:5432", User: url.UserPassword("user", "password"), Path: "/db"},
		PrometheusPort:         9090,
		TracingEnabled:         true,
		TracingCollectorTarget: "http://localhost:9000",
		TracingSamplingRatio:   0.1,
		TracingTLSCertPath:     "some/path",
		TracingAttributes:      map[string]string{"key": "value"},
	}
	got := map[string]string{}
	for _, kv := range envCfg.AsCmdEnv() {
		pair := strings.SplitN(kv, "=", 2)
		require.Len(t, pair, 2)
		got[pair[0]] = pair[1]
	}

	assert.Equal(t, "postgres://user:password@localhost:5432/db", got[envDatabaseURL])
	assert.Equal(t, strconv.Itoa(9090), got[envPromPort])
	assert.Equal(t, "true", got[envTracingEnabled])
	assert.Equal(t, "http://localhost:9000", got[envTracingCollectorTarget])
	assert.Equal(t, "0.1", got[envTracingSamplingRatio])
	assert.Equal(t, "some/path", got[envTracingTLSCertPath])
	assert.Equal(t, "value", got[envTracingAttribute+"key"])
}
