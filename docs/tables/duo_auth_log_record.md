# Table: duo_auth_log_record

Query records in the authentication log for your Duo account.

Notes:
* Unless you specify a time range with `timestamp`, defaults to the last 30 days.
* Duo has [documentation for the action types](https://duo.com/docs/authapi#authistrator-logs).

## Examples

### Authentication log records for the last 30 days (default time range)

```sql
select
  *
from
  duo_auth_log_record
order by
  timestamp desc
```

### Authentication log records for the last 24 hours

```sql
select
  *
from
  duo_auth_log_record
where
  timestamp > current_timestamp - interval '24 hours'
order by
  timestamp desc
```

### Authentication log records for a specific time range

```sql
select
  *
from
  duo_auth_log_record
where
  timestamp >= '2022-04-17T07:00:00-04:00'
  and timestamp < '2022-04-17T08:00:00-04:00'
order by
  timestamp desc
```
