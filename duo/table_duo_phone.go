package duo

import (
	"context"
	"net/url"

	"github.com/duosecurity/duo_api_golang/admin"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableDuoPhone(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "duo_phone",
		Description: "Phones in the Duo account.",
		List: &plugin.ListConfig{
			Hydrate: listPhone,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "number", Require: plugin.Optional},
				{Name: "extension", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPhone,
			KeyColumns: plugin.SingleColumn("phone_id"),
		},
		Columns: []*plugin.Column{

			// Top columns
			{Name: "phone_id", Type: proto.ColumnType_STRING, Description: "The phone's ID."},
			{Name: "number", Type: proto.ColumnType_STRING, Description: "The phone number. A phone with a smartphone platform but no number is a tablet."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Free-form label for the phone."},

			// Other columns
			{Name: "activated", Type: proto.ColumnType_BOOL, Description: "Has this phone been activated for Duo Mobile yet?"},
			{Name: "capabilities", Type: proto.ColumnType_JSON, Description: "List of strings, each a factor that can be used with the device: push, phone, sms, mobile_otp."},
			{Name: "encrypted", Type: proto.ColumnType_STRING, Description: "The encryption status of an Android or iOS device file system. One of: Encrypted, Unencrypted, or Unknown. Blank for other platforms."},
			{Name: "extension", Type: proto.ColumnType_STRING, Description: "An extension, if necessary."},
			{Name: "fingerprint", Type: proto.ColumnType_STRING, Description: "Whether an Android or iOS phone is configured for biometric verification. One of: Configured, Disabled, or Unknown. Blank for other platforms."},
			// TODO - string parsing {Name: "last_seen", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("LastSeen").Transform(transform.UnixToTimestamp), Description: "The timestamp of the last contact between Duo's service and the activated Duo Mobile app installed on the phone. Blank if the device has never activated Duo Mobile or if the platform does not support it."},
			{Name: "model", Type: proto.ColumnType_STRING, Description: "The phone's model."},
			{Name: "platform", Type: proto.ColumnType_STRING, Description: "The phone platform. One of: 'unknown', 'google android', 'apple ios', 'windows phone 7', 'rim blackberry', 'java j2me', 'palm webos', 'symbian os', 'windows mobile', or 'generic smartphone'."},
			{Name: "postdelay", Type: proto.ColumnType_INT, Description: "The time (in seconds) to wait after the extension is dialed and before the speaking the prompt."},
			{Name: "predelay", Type: proto.ColumnType_INT, Description: "The time (in seconds) to wait after the number picks up and before dialing the extension."},
			{Name: "screenlock", Type: proto.ColumnType_STRING, Description: "Whether screen lock is enabled on an Android or iOS phone. One of: Locked, Unlocked, or Unknown. Blank for other platforms."},
			{Name: "sms_passcodes_sent", Type: proto.ColumnType_BOOL, Description: "Have SMS passcodes been sent to this phone?"},
			{Name: "tampered", Type: proto.ColumnType_STRING, Description: "Whether an iOS or Android device is jailbroken or rooted. One of: Not Tampered, Tampered, or Unknown. Blank for other platforms."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of phone. One of: unknown, mobile, or landline."},
			{Name: "users", Type: proto.ColumnType_JSON, Description: "List of users to which this phone belongs."},
		},
	}
}

func listPhone(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_phone.listPhone", "connection_error", err)
		return nil, err
	}

	// URL parameters for all queries
	baseParams := []func(*url.Values){admin.Limit(defaultLimit)}
	keyQuals := d.KeyColumnQuals
	if keyQuals["number"] != nil {
		baseParams = append(baseParams, admin.GetPhonesNumber(keyQuals["number"].GetStringValue()))
	}
	if keyQuals["extension"] != nil {
		baseParams = append(baseParams, admin.GetPhonesExtension(keyQuals["extension"].GetStringValue()))
	}

	client := admin.New(*conn)
	offset := uint64(0)
	for {
		result, err := client.GetPhones(append(baseParams, admin.Offset(offset))...)
		rerr := resultToError(result.StatResult, err)
		if rerr != nil {
			plugin.Logger(ctx).Error("duo_phone.listPhone", "query_error", rerr)
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
			plugin.Logger(ctx).Error("duo_phone.listPhone", "offset_error", err)
			return nil, err
		}
		offset = uint64(offsetInt)
	}
	return nil, nil
}

func getPhone(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	phoneID := d.KeyColumnQuals["phone_id"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_phone.getPhone", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	result, err := client.GetPhone(phoneID)
	rerr := resultToError(result.StatResult, err)
	if rerr != nil {
		plugin.Logger(ctx).Error("duo_phone.getPhone", "query_error", rerr, "phone_id", phoneID)
		return nil, rerr
	}
	return result.Response, nil
}
