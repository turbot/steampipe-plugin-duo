package duo

import (
	"context"

	"github.com/duosecurity/duo_api_golang/admin"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableDuoToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "duo_token",
		Description: "Tokens in the Duo account.",
		List: &plugin.ListConfig{
			Hydrate: listToken,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("token_id"),
			Hydrate:    getToken,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "token_id", Type: proto.ColumnType_STRING, Description: "The token's ID."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Specify a type and serial number to look up a single hardware token. One of: h6 (HOTP-6 hardware token), h8 (HOTP-8 hardware token), yk (YubiKey AES hardware token), d1 (Duo-D100 hardware token)."},
			{Name: "serial", Type: proto.ColumnType_STRING, Description: "The serial number of the hardware token; used to uniquely identify the hardware token when paired with type."},
			{Name: "totp_step", Type: proto.ColumnType_STRING, Description: "Value is null for all supported token types."},
			{Name: "users", Type: proto.ColumnType_JSON, Description: "A list of end users associated with this hardware token."},
		},
	}
}

func listToken(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_token.listToken", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	offset := uint64(0)
	for {
		result, err := client.GetTokens(admin.Limit(defaultLimit), admin.Offset(offset))
		rerr := resultToError(result.StatResult, err)
		if rerr != nil {
			plugin.Logger(ctx).Error("duo_token.listToken", "query_error", rerr)
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
			plugin.Logger(ctx).Error("duo_token.listToken", "offset_error", err)
			return nil, err
		}
		offset = uint64(offsetInt)
	}
	return nil, nil
}

func getToken(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	tokenID := d.KeyColumnQuals["token_id"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_token.getToken", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	result, err := client.GetToken(tokenID)
	rerr := resultToError(result.StatResult, err)
	if rerr != nil {
		plugin.Logger(ctx).Error("duo_token.getToken", "query_error", rerr, "token_id", tokenID)
		return nil, rerr
	}
	return result.Response, nil
}
