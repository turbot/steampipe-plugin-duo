package duo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	duoapi "github.com/duosecurity/duo_api_golang"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

const defaultLimit uint64 = 1000

func connect(_ context.Context, d *plugin.QueryData) (*duoapi.DuoApi, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "duo"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*duoapi.DuoApi), nil
	}

	// Default to using env vars
	apiHostname := os.Getenv("DUO_API_HOSTNAME")
	integrationKey := os.Getenv("DUO_INTEGRATION_KEY")
	secretKey := os.Getenv("DUO_SECRET_KEY")

	// But prefer the config
	duoConfig := GetConfig(d.Connection)
	if duoConfig.APIHostname != nil {
		apiHostname = *duoConfig.APIHostname
	}
	if duoConfig.IntegrationKey != nil {
		integrationKey = *duoConfig.IntegrationKey
	}
	if duoConfig.SecretKey != nil {
		secretKey = *duoConfig.SecretKey
	}
	if apiHostname == "" || integrationKey == "" || secretKey == "" {
		// Credentials not set
		return nil, errors.New("api_hostname, integration_key and secret_key must be configured")
	}

	conn := duoapi.NewDuoApi(
		integrationKey,
		secretKey,
		apiHostname,
		"steampipe-plugin-duo",
	)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func resultToError(result duoapi.StatResult, err error) error {
	if err != nil {
		return err
	}
	if result.Stat == "OK" {
		return nil
	}
	s := result.Stat
	if result.Code != nil {
		s += fmt.Sprintf(" [%d]", *result.Code)
	}
	if result.Message != nil {
		s += fmt.Sprintf(": %s", *result.Message)
	}
	if result.Message_Detail != nil {
		s += fmt.Sprintf(" - %s", *result.Message_Detail)
	}
	return errors.New(s)
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "Resource not found")
}
