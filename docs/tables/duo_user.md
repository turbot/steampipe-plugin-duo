---
title: "Steampipe Table: duo_user - Query Duo Security Users using SQL"
description: "Allows users to query Duo Security Users, providing details such as user ID, username, real name, status, and last login information."
---

# Table: duo_user - Query Duo Security Users using SQL

Duo Security is a cloud-based access security provider protecting the world's fastest-growing companies and thousands of organizations worldwide. It provides comprehensive security solutions to address the security needs of users, devices, and applications, ensuring that only trusted users and devices can access protected applications. The Duo Security Users are the entities that have been granted access to the Duo-protected applications.

## Table Usage Guide

The `duo_user` table provides detailed insights into Duo Security Users within the Duo Security System. As a security analyst, you can explore user-specific details through this table, including user status, last login information, and associated metadata. Utilize it to uncover information about users, such as those with active or inactive status, the last login details of users, and the verification of user identities.

## Examples

### List all users

```sql+postgres
select
  user_id,
  username,
  email
from
  duo_user
order by
  username;
```

```sql+sqlite
select
  user_id,
  username,
  email
from
  duo_user
order by
  username;
```

### Users who have not yet enrolled

```sql+postgres
select
  user_id,
  username,
  email
from
  duo_user
where
  not is_enrolled
order by
  username;
```

```sql+sqlite
select
  user_id,
  username,
  email
from
  duo_user
where
  not is_enrolled
order by
  username;
```

### Most recent 10 users to login

```sql+postgres
select
  user_id,
  username,
  email,
  last_login
from
  duo_user
order by
  last_login desc
limit 10;
```

```sql+sqlite
select
  user_id,
  username,
  email,
  last_login
from
  duo_user
order by
  last_login desc
limit 10;
```

### Users who have never logged in

```sql+postgres
select
  user_id,
  username,
  email
from
  duo_user
where
  last_login is null
order by
  username;
```

```sql+sqlite
select
  user_id,
  username,
  email
from
  duo_user
where
  last_login is null
order by
  username;
```

### Users who have not logged in for more than 30 days

```sql+postgres
select
  user_id,
  username,
  email,
  last_login,
  age(last_login) as age
from
  duo_user
where
  age(last_login) > interval '30 days'
order by
  age desc;
```

```sql+sqlite
Error: SQLite does not support the AGE function or INTERVAL keyword used in PostgreSQL.
```

### Users who are locked out

```sql+postgres
select
  user_id,
  username,
  email,
  last_login,
  status
from
  duo_user
where
  status = 'locked_out'
order by
  username;
```

```sql+sqlite
select
  user_id,
  username,
  email,
  last_login,
  status
from
  duo_user
where
  status = 'locked_out'
order by
  username;
```

### User statistics by status

```sql+postgres
select
  status,
  count(*)
from
  duo_user
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
  duo_user
group by
  status
order by
  status;
```

### Users and their group memberships

```sql+postgres
select
  u.username,
  g->>'name' as groupname
from
  duo_user as u,
  jsonb_array_elements(u.groups) as g
order by
  username,
  groupname;
```

```sql+sqlite
select
  u.username,
  json_extract(g.value, '$.name') as groupname
from
  duo_user as u,
  json_each(u.groups) as g
order by
  username,
  groupname;
```

### Users and their phones

```sql+postgres
select
  u.username,
  p->>'number' as phone_number,
  p->>'extension' as phone_extension
from
  duo_user as u,
  jsonb_array_elements(u.phones) as p
order by
  username,
  phone_number,
  phone_extension;
```

```sql+sqlite
select
  u.username,
  json_extract(p.value, '$.number') as phone_number,
  json_extract(p.value, '$.extension') as phone_extension
from
  duo_user as u,
  json_each(u.phones) as p
order by
  username,
  phone_number,
  phone_extension;
```

### Users and their hardware tokens

```sql+postgres
select
  u.username,
  t->>'serial' as token_serial
from
  duo_user as u,
  jsonb_array_elements(u.tokens) as t
order by
  username,
  token_serial;
```

```sql+sqlite
select
  u.username,
  json_extract(t.value, '$.serial') as token_serial
from
  duo_user as u,
  json_each(u.tokens) as t
order by
  username,
  token_serial;
```
