package duo

import (
	"context"

	"github.com/duosecurity/duo_api_golang/admin"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDuoGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "duo_group",
		Description: "Groups in the Duo account.",
		List: &plugin.ListConfig{
			Hydrate: listGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("group_id"),
			Hydrate:    getGroup,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "group_id", Type: proto.ColumnType_STRING, Description: "The group's ID."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The group's name. If managed by directory sync, then the name returned here also indicates the source directory."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The group's authentication status. May be one of: active, bypass, disabled."},
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Desc"), Description: "The group's description."},
		},
	}
}

func listGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_group.listGroup", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	offset := uint64(0)
	for {
		result, err := client.GetGroups(admin.Limit(defaultLimit), admin.Offset(offset))
		rerr := resultToError(result.StatResult, err)
		if rerr != nil {
			plugin.Logger(ctx).Error("duo_group.listGroup", "query_error", rerr)
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
			plugin.Logger(ctx).Error("duo_group.listGroup", "offset_error", err)
			return nil, err
		}
		offset = uint64(offsetInt)
	}
	return nil, nil
}

func getGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	groupID := d.EqualsQuals["group_id"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_group.getGroup", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	result, err := client.GetGroup(groupID)
	rerr := resultToError(result.StatResult, err)
	if rerr != nil {
		plugin.Logger(ctx).Error("duo_group.getGroup", "query_error", rerr, "group_id", groupID)
		return nil, rerr
	}
	return result.Response, nil
}
