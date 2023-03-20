package duo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-duo",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"duo_account_settings":    tableDuoAccountSettings(ctx),
			"duo_account_summary":     tableDuoAccountSummary(ctx),
			"duo_admin_log_record":    tableDuoAdminLogRecord(ctx),
			"duo_administrator":       tableDuoAdministrator(ctx),
			"duo_administrative_unit": tableDuoAdministrativeUnit(ctx),
			"duo_auth_log_record":     tableDuoAuthLogRecord(ctx),
			"duo_group":               tableDuoGroup(ctx),
			"duo_integration":         tableDuoIntegration(ctx),
			"duo_phone":               tableDuoPhone(ctx),
			"duo_token":               tableDuoToken(ctx),
			"duo_user":                tableDuoUser(ctx),
		},
	}
	return p
}
