---
title: "Steampipe Table: duo_auth_log_record - Query Duo Security Authentication Log Records using SQL"
description: "Allows users to query Duo Security Authentication Log Records, providing insights into user authentication activities and potential security breaches."
---

# Table: duo_auth_log_record - Query Duo Security Authentication Log Records using SQL

Duo Security is a cloud-based security platform that provides two-factor authentication, endpoint security, remote access solutions and more. This platform helps prevent breaches, reduce risk, and ensure regulatory compliance. The Authentication Log Records in Duo Security provide detailed information about user authentication attempts, including the user, device, location, and result.

## Table Usage Guide

The `duo_auth_log_record` table provides insights into user authentication activities within Duo Security. As a security analyst, explore detailed information about user authentication attempts through this table, including the user, device, location, and result. Utilize it to uncover information about user activities, detect potential security breaches, and ensure regulatory compliance.

## Examples

### Authentication log records for the last 30 days (default time range)

```sql+postgres
select
  *
from
  duo_auth_log_record
order by
  timestamp desc;
```

```sql+sqlite
select
  *
from
  duo_auth_log_record
order by
  timestamp desc;
```

### Authentication log records for the last 24 hours

```sql+postgres
select
  *
from
  duo_auth_log_record
where
  timestamp > current_timestamp - interval '24 hours'
order by
  timestamp desc;
```

```sql+sqlite
select
  *
from
  duo_auth_log_record
where
  timestamp > datetime('now', '-24 hours')
order by
  timestamp desc;
```

### Authentication log records for a specific time range

```sql+postgres
select
  *
from
  duo_auth_log_record
where
  timestamp >= '2022-04-17T07:00:00-04:00'
  and timestamp < '2022-04-17T08:00:00-04:00'
order by
  timestamp desc;
```

```sql+sqlite
select
  *
from
  duo_auth_log_record
where
  timestamp >= '2022-04-17T07:00:00-04:00'
  and timestamp < '2022-04-17T08:00:00-04:00'
order by
  timestamp desc;
```
