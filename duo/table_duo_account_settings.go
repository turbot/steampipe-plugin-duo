package duo

import (
	"context"

	"github.com/duosecurity/duo_api_golang/admin"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDuoAccountSettings(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "duo_account_settings",
		Description: "Get settings for the Duo account.",
		List: &plugin.ListConfig{
			Hydrate: listAccountSettings,
		},
		// Counts should return zero not null, so change the default transform
		DefaultTransform: transform.FromJSONTag(),
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The customer name."},
			// Other columns
			{Name: "caller_id", Type: proto.ColumnType_STRING, Description: "Automated calls will appear to come from this number. This does not apply to text messages."},
			{Name: "fraud_email", Type: proto.ColumnType_STRING, Description: "The email address to be notified when a user reports a fraudulent authentication attempt or is locked out due to failed authentication attempts. All administrators will be notified if this is not set."},
			{Name: "fraud_email_enabled", Type: proto.ColumnType_BOOL, Description: "If true, emailed notifications of user-reported fraudulent authentication attempts and user lockouts due to failed authentication are sent to the email address defined for fraud_email, or to all administrators if fraud_email is not defined. If set to false, no fraud alert emails are sent."},
			{Name: "helpdesk_bypass", Type: proto.ColumnType_STRING, Description: "Grants permission for administrators with the Help Desk role to generate bypass codes for users. One of allow (default value), limit, or deny."},
			{Name: "helpdesk_bypass_expiration", Type: proto.ColumnType_INT, Description: "Integer specifying a default expiration for bypass codes generated by Help Desk admins, in minutes. If not set, Help Desk admins may change bypass code expiration from the default 60 minutes after creation if helpdesk_bypass is set to allow."},
			{Name: "helpdesk_can_send_enroll_email", Type: proto.ColumnType_BOOL, Description: "Permits Help Desk administrators to send or resend enrollment emails to users. One of true or false (default)."},
			{Name: "helpdesk_message", Type: proto.ColumnType_STRING, Description: "Custom help message shown to end-users in the traditional Duo Prompt and Device Health application. Up to 200 characters; HTML formatting or hyperlinks are not allowed."},
			{Name: "inactive_user_expiration", Type: proto.ColumnType_INT, Description: "Users will be automatically deleted if they are inactive (no successful logins) for a this amount of days."},
			{Name: "keypress_confirm", Type: proto.ColumnType_STRING, Description: "The key for users to press to authenticate, or empty if any key should be pressed to authenticate."},
			{Name: "keypress_fraud", Type: proto.ColumnType_STRING, Description: "The key for users to press to report fraud, or empty if any key should be pressed to authenticate."},
			{Name: "language", Type: proto.ColumnType_STRING, Description: "The language used in the traditional Duo browser-based user authentication prompt. One of: EN, DE, FR. Default: EN"},
			{Name: "lockout_expire_duration", Type: proto.ColumnType_INT, Description: "If non-zero, an integer indicating the time in minutes until a locked-out user's status reverts to Active. If null or 0, a user remains locked out until their status is manually changed (By an admin or API call). Minimum: 5 minutes. Maximum: 30000 minutes."},
			{Name: "lockout_threshold", Type: proto.ColumnType_INT, Description: "The number of consecutive failed authentication attempts before the user's status is set to 'Locked Out' and the user is denied access."},
			{Name: "minimum_password_length", Type: proto.ColumnType_INT, Description: "An integer indicating the minimum number of characters that an administrator's Duo Admin Panel password must contain. This is only enforced on password creation and reset; existing passwords will not be invalidated. Default: 12."},
			{Name: "password_requires_lower_alpha", Type: proto.ColumnType_BOOL, Description: "If true, administrator passwords will be required to contain a lower case alphabetic character. If false, administrator passwords will not be required to contain a lower case alphabetic character. This is only enforced on password creation and reset; existing passwords will not be invalidated. Default: false."},
			{Name: "password_requires_numeric", Type: proto.ColumnType_BOOL, Description: "If true, administrator passwords will be required to contain a numeric character. If false, administrator passwords will not be required to contain a numeric character. This is only enforced on password creation and reset; existing passwords will not be invalidated. Default: false."},
			{Name: "password_requires_special", Type: proto.ColumnType_BOOL, Description: "If true, administrator passwords will be required to contain a special (non-alphanumeric) character. If false, administrator passwords will not be required to contain a special (non-alphanumeric) character. This is only enforced on password creation and reset; existing passwords will not be invalidated. Default: false."},
			{Name: "password_requires_upper_alpha", Type: proto.ColumnType_BOOL, Description: "If true, administrator passwords will be required to contain an upper case alphabetic character. If false, administrator passwords will not be required to contain an upper case alphabetic character. This is only enforced on password creation and reset; existing passwords will not be invalidated. Default: false."},
			{Name: "sms_batch", Type: proto.ColumnType_INT, Description: "An integer that indicates how many passcodes to send at one time, up to 10."},
			{Name: "sms_expiration", Type: proto.ColumnType_INT, Description: "The time in minutes to expire and invalidate SMS passcodes."},
			{Name: "sms_message", Type: proto.ColumnType_STRING, Description: "Description sent with every batch of SMS passcodes."},
			{Name: "sms_refresh", Type: proto.ColumnType_INT, Description: "If 1, a new set of SMS passcodes will automatically be sent after the last one is used. If 0, a new set will not be sent."},
			{Name: "telephony_warning_min", Type: proto.ColumnType_INT, Description: "An integer indicating the number of telephony credits at which an alert will be sent for low credits."},
			{Name: "timezone", Type: proto.ColumnType_STRING, Description: "This is the timezone used when displaying timestamps in the Duo Admin Panel."},
			{Name: "user_telephony_cost_max", Type: proto.ColumnType_DOUBLE, Description: "An integer indicating the maximum number of telephony credits a user may consume in a single authentication event. This excludes Duo administrators authenticating to the Duo administration panel. Default: 20."},
		},
	}
}

func listAccountSettings(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_account_settings.getAccountSettings", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)
	result, err := client.GetAccountSettings()
	rerr := resultToError(result.StatResult, err)
	if rerr != nil {
		plugin.Logger(ctx).Error("duo_account_settings.listAccountSettings", "query_error", rerr)
		return nil, rerr
	}
	d.StreamListItem(ctx, result.Response)
	return nil, nil
}
