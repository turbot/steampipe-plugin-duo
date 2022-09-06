package duo

import (
	"context"

	"github.com/duosecurity/duo_api_golang/admin"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableDuoAdministrator(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "duo_administrator",
		Description: "Administrators in the Duo account.",
		List: &plugin.ListConfig{
			Hydrate: listAdministrator,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("admin_id"),
			Hydrate:    getAdministrator,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "admin_id", Type: proto.ColumnType_STRING, Description: "The administrator's ID."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The administrator's email address."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The administrator's full name."},
			// Other columns
			{Name: "admin_units", Type: proto.ColumnType_JSON, Description: "The list of administrative units (by admin_unit_id) to which the admin belongs. For an unrestricted admin, this is an empty list."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "The administrator's creation date as a UNIX timestamp. No creation date shown for administrators created before October 2021."},
			{Name: "hardtoken", Type: proto.ColumnType_JSON, Description: "Information about hardware tokens attached to the administrator, or null if none attached. See Retrieve Hardware Tokens for descriptions of the response values."},
			{Name: "last_login", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "An integer indicating the last time this administrator logged in, as a Unix timestamp, or null if the administrator has not logged in."},
			{Name: "password_change_required", Type: proto.ColumnType_BOOL, Description: "Either true if the administrator must change their password at the next login, or false if no password change is required."},
			{Name: "phone", Type: proto.ColumnType_STRING, Description: "The administrator's phone number."},
			{Name: "restricted_by_admin_units", Type: proto.ColumnType_BOOL, Description: "Is this administrator restricted by an administrative unit assignment? Either true or false. Must be set to true in order to add the admin to an administrative unit using the API."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "The administrator's role. One of: Owner, Administrator, Application Manager, User Manager, Help Desk, Billing, Phishing Manager, or Read-only. Only present in the response if the customer edition includes the Administrative Roles feature."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The administrator account's status. One of: Active (admin can log in to Duo), Disabled (admin prevented from access), Expired (admin blocked from access due to inactivity), or Pending Activation (new admin must complete activation to gain access)."},
		},
	}
}

func listAdministrator(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_administrator.listAdministrator", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	offset := uint64(0)
	for {
		result, err := client.GetAdministrators(admin.Limit(defaultLimit), admin.Offset(offset))
		rerr := resultToError(result.StatResult, err)
		if rerr != nil {
			plugin.Logger(ctx).Error("duo_administrator.listAdministrator", "query_error", rerr)
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
			plugin.Logger(ctx).Error("duo_administrator.listAdministrator", "offset_error", err)
			return nil, err
		}
		offset = uint64(offsetInt)
	}
	return nil, nil
}

func getAdministrator(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	adminID := d.KeyColumnQuals["admin_id"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_administrator.getAdministrator", "connection_error", err, "admin_id", adminID)
		return nil, err
	}
	client := admin.New(*conn)
	result, err := client.GetAdministrator(adminID)
	rerr := resultToError(result.StatResult, err)
	if rerr != nil {
		plugin.Logger(ctx).Error("duo_administrator.getAdministrator", "query_error", rerr, "admin_id", adminID)
		return nil, rerr
	}
	return result.Response, nil
}
