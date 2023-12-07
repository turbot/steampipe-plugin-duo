---
title: "Steampipe Table: duo_group - Query Duo Security Groups using SQL"
description: "Allows users to query Duo Security Groups, providing detailed information about each group including its ID, name, description, and status."
---

# Table: duo_group - Query Duo Security Groups using SQL

Duo Security Groups are collections of users within Duo's two-factor authentication system, allowing for more efficient management of user access and permissions. Groups can be used to assign different policies or applications to different sets of users. This is a critical aspect of managing security and access control in an organization.

## Table Usage Guide

The `duo_group` table provides insights into Duo Security Groups within Duo's two-factor authentication system. As a security or IT professional, explore group-specific details through this table, including the group's ID, name, description, and status. Utilize it to manage and monitor user access and permissions, enabling better security and control within your organization.

## Examples

### List all groups

```sql+postgres
select
  name,
  group_id,
  status,
  description
from
  duo_group
order by
  name;
```

```sql+sqlite
select
  name,
  group_id,
  status,
  description
from
  duo_group
order by
  name;
```

### Group statistics by status

```sql+postgres
select
  status,
  count(*)
from
  duo_group
group by
  status
order by
  status;
```

```sql+sqlite
select
  status,
  count(*)
from
  duo_group
group by
  status
order by
  status;
```

### List disabled groups

```sql+postgres
select
  name
from
  duo_group
where
  status = 'disabled'
order by
  name;
```

```sql+sqlite
select
  name
from
  duo_group
where
  status = 'disabled'
order by
  name;
```
