package duo

import (
	"context"

	"github.com/duosecurity/duo_api_golang/admin"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableDuoAccountSummary(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "duo_account_summary",
		Description: "Get summary info for the Duo account.",
		List: &plugin.ListConfig{
			Hydrate: listAccountSummary,
		},
		// Counts should return zero not null, so change the default transform
		DefaultTransform: transform.FromJSONTag(),
		Columns: []*plugin.Column{
			// Top columns
			{Name: "admin_count", Type: proto.ColumnType_INT, Description: "Current number of admins in the account."},
			{Name: "integration_count", Type: proto.ColumnType_INT, Description: "Current number of integrations in the account."},
			{Name: "user_count", Type: proto.ColumnType_INT, Description: "Current number of users in the account."},
			{Name: "telephony_credits_remaining", Type: proto.ColumnType_INT, Description: "Current total number of telephony credits available in the account. This is the sum of all types of telephony credits."},
			{Name: "user_pending_deletion_count", Type: proto.ColumnType_INT, Description: "Current number of users pending deletion from the account (seen in the Admin Panel's Trash view)."},
		},
	}
}

func listAccountSummary(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_account_summary.getAccountSummary", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	result, err := client.GetAccountInfoSummary()
	rerr := resultToError(result.StatResult, err)
	if rerr != nil {
		plugin.Logger(ctx).Error("duo_account_summary.listAccountSummary", "query_error", rerr)
		return nil, rerr
	}
	d.StreamListItem(ctx, result.Response)
	return nil, nil
}
