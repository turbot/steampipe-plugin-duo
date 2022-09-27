package duo

import (
	"context"

	"github.com/duosecurity/duo_api_golang/admin"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableDuoAdministrativeUnit(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "duo_administrative_unit",
		Description: "AdministrativeUnits in the Duo account.",
		List: &plugin.ListConfig{
			Hydrate: listAdministrativeUnit,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("admin_unit_id"),
			Hydrate:    getAdministrativeUnit,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "admin_unit_id", Type: proto.ColumnType_STRING, Description: "The administrative unit's unique ID."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The administrative unit's name."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The administrative unit's description."},
			{Name: "groups", Type: proto.ColumnType_JSON, Hydrate: getAdministrativeUnit, Description: "The groups assigned to the new administrative unit, listed by group_id."},
			{Name: "integrations", Type: proto.ColumnType_JSON, Hydrate: getAdministrativeUnit, Description: "The groups assigned to the new administrative unit, listed by integration_id."},
			{Name: "restrict_by_groups", Type: proto.ColumnType_BOOL, Description: "Does the administrative unit specify groups?"},
			{Name: "restrict_by_integrations", Type: proto.ColumnType_BOOL, Description: "Does the administrative unit specify integrations?"},
		},
	}
}

func listAdministrativeUnit(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_administrative_unit.listAdministrativeUnit", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	offset := uint64(0)
	for {
		result, err := client.GetAdministrativeUnits(admin.Limit(defaultLimit), admin.Offset(offset))
		rerr := resultToError(result.StatResult, err)
		if rerr != nil {
			plugin.Logger(ctx).Error("duo_administrative_unit.listAdministrativeUnit", "query_error", rerr)
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
			plugin.Logger(ctx).Error("duo_administrative_unit.listAdministrativeUnit", "offset_error", err)
			return nil, err
		}
		offset = uint64(offsetInt)
	}
	return nil, nil
}

func getAdministrativeUnit(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var adminID string
	if h.Item != nil {
		// Get admin ID from row that is being hydrated
		adminID = h.Item.(admin.AdministrativeUnit).AdminUnitID
	} else {
		// Get admin ID from the qualifier on a get
		adminID = d.KeyColumnQuals["admin_unit_id"].GetStringValue()
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_administrative_unit.getAdministrativeUnit", "connection_error", err, "admin_id", adminID)
		return nil, err
	}
	client := admin.New(*conn)

	result, err := client.GetAdministrativeUnit(adminID)
	rerr := resultToError(result.StatResult, err)
	if rerr != nil {
		plugin.Logger(ctx).Error("duo_administrative_unit.getAdministrativeUnit", "query_error", rerr, "admin_id", adminID)
		return nil, rerr
	}
	return result.Response, nil
}
