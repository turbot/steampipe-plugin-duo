package duo

import (
	"context"

	"github.com/duosecurity/duo_api_golang/admin"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDuoIntegration(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "duo_integration",
		Description: "Integrations in the Duo account.",
		List: &plugin.ListConfig{
			Hydrate: listIntegration,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("integration_key"),
			Hydrate:    getIntegration,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "integration_key", Type: proto.ColumnType_STRING, Description: "Integration ID."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The integration's name."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Integration type, e.g. 1password, okta."},
			// Other columns
			{Name: "adminapi_admins", Type: proto.ColumnType_INT, Description: "1 if the integration has been granted permission for Administrators methods; otherwise 0. Only applicable to Admin API integrations."},
			{Name: "adminapi_info", Type: proto.ColumnType_INT, Description: "1 if the integration has been granted permission for Account Info methods; otherwise 0. Only applicable to Admin API integrations."},
			{Name: "adminapi_integrations", Type: proto.ColumnType_INT, Description: "1 if the integration has been granted permission for Integrations methods; otherwise 0. Only applicable to Admin API integrations."},
			{Name: "adminapi_read_log", Type: proto.ColumnType_INT, Description: "1 if the integration has been granted permission for Logs methods; otherwise 0. Only applicable to Admin API integrations."},
			{Name: "adminapi_read_resource", Type: proto.ColumnType_INT, Description: "1 if the integration has been granted permission to retrieve objects like users, phones, and hardware tokens; otherwise 0. Only applicable to Admin API integrations."},
			{Name: "adminapi_settings", Type: proto.ColumnType_INT, Description: "1 if the integration has been granted permission for Settings methods; otherwise 0. Only applicable to Admin API integrations."},
			{Name: "adminapi_write_resource", Type: proto.ColumnType_INT, Description: "1 if the integration has been granted permission to modify objects like users, phones, and hardware tokens; otherwise 0. Only applicable to Admin API integrations."},
			{Name: "frameless_auth_prompt_enabled", Type: proto.ColumnType_INT, Description: "1 if the integration has been updated to support Duo Universal Prompt, otherwise 0. Only appears for a given integration after Duo makes the frameless prompt available for that application, and the value is set to 1 automatically when Duo detects a frameless authentication for the integration."},
			{Name: "greeting", Type: proto.ColumnType_STRING, Description: "Voice greeting read before the authentication instructions to users who authenticate with a phone callback."},
			{Name: "groups_allowed", Type: proto.ColumnType_JSON, Description: "A list of groups, as group IDs, that are allowed to authenticate with the integration. If empty, all groups are allowed."},
			{Name: "networks_for_api_access", Type: proto.ColumnType_STRING, Description: "A comma-separated list of IP addresses, IP ranges, or CIDRs specifying the networks allowed to access this API integration. Only returned for Accounts API and Admin API integrations."},
			{Name: "notes", Type: proto.ColumnType_STRING, Description: "Description of the integration."},
			{Name: "policy_key", Type: proto.ColumnType_STRING, Description: "The identifying policy key for the custom policy attached to the integration. Not shown if no policy attached to the integration."},
			{Name: "prompt_v4_enabled", Type: proto.ColumnType_INT, Description: "1 if Duo Universal Prompt is activated for the application, otherwise 0. Only appears for a given integration when frameless_auth_prompt_enabled is 1 (value set automatically when Duo detects a frameless authentication for the integration)."},
			{Name: "secret_key", Type: proto.ColumnType_STRING, Description: "Secret used when configuring systems to use this integration."},
			{Name: "self_service_allowed", Type: proto.ColumnType_JSON, Description: "1 if users may use self-service from this integration's 2FA prompt to update authentication devices, otherwise false (default)."},
			{Name: "username_normalization_policy", Type: proto.ColumnType_STRING, Description: "This controls whether or not usernames should be altered before trying to match them to a user account. One of: None, Simple."},
		},
	}
}

func listIntegration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_integration.listIntegration", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	offset := uint64(0)
	for {
		result, err := client.GetIntegrations(admin.Limit(defaultLimit), admin.Offset(offset))
		rerr := resultToError(result.StatResult, err)
		if rerr != nil {
			plugin.Logger(ctx).Error("duo_integration.listIntegration", "query_error", rerr)
			return nil, rerr
		}
		for _, i := range result.Response {
			d.StreamListItem(ctx, i)
		}
		if result.Metadata.NextOffset == "" {
			break
		}
		offsetInt, err := result.Metadata.NextOffset.Int64()
		if err != nil {
			plugin.Logger(ctx).Error("duo_integration.listIntegration", "offset_error", err)
			return nil, err
		}
		offset = uint64(offsetInt)
	}
	return nil, nil
}

func getIntegration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	integrationKey := d.EqualsQuals["integration_key"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_integration.getIntegration", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	result, err := client.GetIntegration(integrationKey)
	rerr := resultToError(result.StatResult, err)
	if rerr != nil {
		plugin.Logger(ctx).Error("duo_integration.getIntegration", "query_error", rerr, "integration_key", integrationKey)
		return nil, rerr
	}
	return result.Response, nil
}
