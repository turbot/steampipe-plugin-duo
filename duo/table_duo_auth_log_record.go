package duo

import (
	"context"
	"time"

	"github.com/duosecurity/duo_api_golang/admin"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableDuoAuthLogRecord(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "duo_auth_log_record",
		Description: "Auth log records in the Duo account.",
		List: &plugin.ListConfig{
			Hydrate: listAuthLogRecord,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "timestamp", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional, CacheMatch: "exact"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("isotimestamp"), Description: "Time when the event occurred."},
			{Name: "action", Type: proto.ColumnType_STRING, Transform: transform.FromField("action"), Description: "The type of change that was performed, e.g. auth_login, group_create, user_update."},
			{Name: "username", Type: proto.ColumnType_STRING, Transform: transform.FromField("username"), Description: "The full name of the authistrator who performed the action in the Duo Auth Panel. If the action was performed with the API this will be 'API'. Automatic actions like deletion of inactive users have 'System' for the username. Changes synchronized from Directory Sync will have a username of the form (example) 'AD Sync: name of directory'."},
			{Name: "object", Type: proto.ColumnType_STRING, Transform: transform.FromField("object"), Description: "The object that was acted on. For example: 'jsmith' (for users), '(555) 713-6275 x456' (for phones), or 'HOTP 8-digit 123456' (for tokens)."},
			{Name: "description", Type: proto.ColumnType_JSON, Transform: transform.FromField("description"), Description: "Details of what changed, format varies based on the action."},
		},
	}
}

func listAuthLogRecord(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// Get the client
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("duo_auth_log_record.listAuthLogRecord", "connection_error", err)
		return nil, err
	}
	client := admin.New(*conn)

	// Find the time range from optional quals.
	maxtime := time.Now()
	mintime := time.Time{}
	if d.Quals["timestamp"] != nil {
		for _, q := range d.Quals["timestamp"].Quals {
			ts := q.Value.GetTimestampValue().AsTime()
			switch q.Operator {
			case ">=", ">":
				if ts.After(mintime) {
					mintime = ts
				}
			case "=":
				if ts.After(mintime) {
					mintime = ts
				}
				if ts.Before(maxtime) {
					maxtime = ts
				}
			case "<", "<=":
				if ts.Before(maxtime) {
					maxtime = ts
				}
			}
		}
	}
	// If mintime has not been set via a timestamp qual, then default to last 30 days.
	if mintime.IsZero() {
		mintime = time.Now().Add(-30 * 24 * time.Hour)
	}

	logsOffset := admin.Offset(0)
	period := maxtime.Sub(mintime)
	for {
		result, err := client.GetAuthLogs(mintime, period, logsOffset)
		rerr := resultToError(result.StatResult, err)
		if rerr != nil {
			plugin.Logger(ctx).Error("duo_auth_log_record.listAuthLogRecord", "query_error", rerr)
			return nil, rerr
		}
		for _, i := range result.Response.Logs {
			d.StreamListItem(ctx, i)
		}
		logsOffset = result.Response.Metadata.GetNextOffset()
		if logsOffset == nil {
			break
		}
	}
	return nil, nil
}
