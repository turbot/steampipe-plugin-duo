---
title: "Steampipe Table: duo_admin_log_record - Query Duo Security Admin Log Records using SQL"
description: "Allows users to query Admin Log Records in Duo Security, providing detailed information about administrator actions and system changes."
---

# Table: duo_admin_log_record - Query Duo Security Admin Log Records using SQL

Duo Security is a cloud-based access security provider protecting the world's fastest-growing and largest companies and thousands of organizations worldwide, including Dresser-Rand, Etsy, Facebook, K-Swiss, Random House, Yelp, Zillow, Paramount Pictures, and more. The platform provides a suite of security products including Duo's Trusted Access platform, one of the most secure access platforms in the world. It offers numerous services like secure single sign-on (SSO), Duo Push, Universal Prompt, and more.

## Table Usage Guide

The `duo_admin_log_record` table provides insights into Admin Log Records within Duo Security. As a Security Analyst, explore record-specific details through this table, including action types, object details, and associated metadata. Utilize it to monitor administrator actions, track system changes, and ensure adherence to security policies.

## Examples

### Admin log records for the last 30 days (default time range)

```sql+postgres
select
  *
from
  duo_admin_log_record
order by
  timestamp desc;
```

```sql+sqlite
select
  *
from
  duo_admin_log_record
order by
  timestamp desc;
```

### Admin log records for the last 24 hours

```sql+postgres
select
  *
from
  duo_admin_log_record
where
  timestamp > current_timestamp - interval '24 hours'
order by
  timestamp desc;
```

```sql+sqlite
select
  *
from
  duo_admin_log_record
where
  timestamp > datetime('now', '-24 hours')
order by
  timestamp desc;
```

### Admin log records for a specific time range

```sql+postgres
select
  *
from
  duo_admin_log_record
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
  duo_admin_log_record
where
  timestamp >= '2022-04-17T07:00:00-04:00'
  and timestamp < '2022-04-17T08:00:00-04:00'
order by
  timestamp desc;
```

### Failed login attempts in the last 7 days

```sql+postgres
select
  *
from
  duo_admin_log_record
where
  action = 'admin_login_error'
  and timestamp > current_timestamp - interval '7 days'
order by
  timestamp desc;
```

```sql+sqlite
select
  *
from
  duo_admin_log_record
where
  action = 'admin_login_error'
  and timestamp > datetime('now', '-7 days')
order by
  timestamp desc;
```
