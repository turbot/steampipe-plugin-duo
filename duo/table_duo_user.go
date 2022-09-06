package duo

import (
	"context"
	"net/url"

	"github.com/duosecurity/duo_api_golang/admin"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableDuoUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "duo_user",
		Description: "Users in the Duo account.",
		List: &plugin.ListConfig{
			Hydrate: listUser,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "username", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getUser,
			KeyColumns: plugin.SingleColumn("user_id"),
		},
		Columns: []*plugin.Column{

			// Top columns
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "The user's ID."},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "The user's username."},
			{Name: "realname", Type: proto.ColumnType_STRING, Transform: transform.FromField("RealName"), Description: "The user's real name (or full name)."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The user's email address."},

			// Other columns
			{Name: "aliases", Type: proto.ColumnType_JSON, Description: "Map of the user's username alias(es). Up to eight aliases may exist."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "The user's creation date."},
			{Name: "firstname", Type: proto.ColumnType_STRING, Transform: transform.FromField("FirstName"), Description: "The user's given name."},
			{Name: "groups", Type: proto.ColumnType_JSON, Description: "List of groups to which this user belongs."},
			{Name: "is_enrolled", Type: proto.ColumnType_BOOL, Description: "Is true if the user has a phone, hardware token, U2F token, WebAuthn security key, or other WebAuthn method available for authentication. Otherwise, false."},
			{Name: "last_directory_sync", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("LastDirectorySync").Transform(transform.UnixToTimestamp), Description: "The last update to the user via directory sync, or null if the user has never synced with an external directory or if the directory that originally created the user has been deleted from Duo."},
			{Name: "last_login", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("LastLogin").Transform(transform.UnixToTimestamp), Description: "The last time this user logged in, or null if the user has not logged in."},
			{Name: "lastname", Type: proto.ColumnType_STRING, Transform: transform.FromField("LastName"), Description: "The user's surname."},
			{Name: "notes", Type: proto.ColumnType_STRING, Description: "Notes about this user."},
			{Name: "phones", Type: proto.ColumnType_JSON, Description: "A list of phones that this user can use."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The user's status: active, bypass, disabled, locked out, pending deletion."},
			{Name: "tokens", Type: proto.ColumnType_JSON, Description: "A list of tokens that this user can use."},
		},
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_user.listUser", "connection_error", err)
		return nil, err
	}

	// URL parameters for all queries
	baseParams := []func(*url.Values){admin.Limit(defaultLimit)}
	keyQuals := d.KeyColumnQuals
	if keyQuals["username"] != nil {
		baseParams = append(baseParams, admin.GetUsersUsername(keyQuals["username"].GetStringValue()))
	}

	client := admin.New(*conn)
	offset := uint64(0)
	for {
		result, err := client.GetUsers(append(baseParams, admin.Offset(offset))...)
		rerr := resultToError(result.StatResult, err)
		if rerr != nil {
			plugin.Logger(ctx).Error("duo_user.listUser", "query_error", rerr)
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
			plugin.Logger(ctx).Error("duo_user.listUser", "offset_error", err)
			return nil, err
		}
		offset = uint64(offsetInt)
	}
	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	userID := d.KeyColumnQuals["user_id"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_user.getUser", "connection_error", err, "user_id", userID)
		return nil, err
	}
	client := admin.New(*conn)
	result, err := client.GetUser(userID)
	rerr := resultToError(result.StatResult, err)
	if rerr != nil {
		plugin.Logger(ctx).Error("duo_user.getUser", "query_error", rerr, "user_id", userID)
		return nil, rerr
	}
	return result.Response, nil
}
