---
title: "Steampipe Table: duo_phone - Query Duo Security Phones using SQL"
description: "Allows users to query Duo Security Phones, specifically the phone number, platform, and type, providing insights into the phone's information and its status."
---

# Table: duo_phone - Query Duo Security Phones using SQL

Duo Security is a cloud-based security solution that verifies the identity of users and the health of their devices before granting them access to applications. It provides a two-factor authentication service to protect against account takeover and data theft. Duo Security helps you secure access to all applications, for any user and device, from anywhere.

## Table Usage Guide

The `duo_phone` table provides insights into Phones within Duo Security. As a security engineer, explore phone-specific details through this table, including phone number, platform, and type. Utilize it to uncover information about phones, such as their status, the type of phone, and the platform it's running on.

## Examples

### List all phones

```sql+postgres
select
  number,
  extension,
  type,
  platform,
  model,
  phone_id
from
  duo_phone
order by
  number,
  extension;
```

```sql+sqlite
select
  number,
  extension,
  type,
  platform,
  model,
  phone_id
from
  duo_phone
order by
  number,
  extension;
```

### Phones and their users

```sql+postgres
select
  p.number,
  p.extension,
  u->>'username' as username
from
  duo_phone as p,
  jsonb_array_elements(p.users) as u
order by
  number,
  extension,
  username;
```

```sql+sqlite
select
  p.number,
  p.extension,
  json_extract(u.value, '$.username') as username
from
  duo_phone as p,
  json_each(p.users) as u
order by
  number,
  extension,
  username;
```

### Phones that are not yet activated

```sql+postgres
select
  number,
  extension,
  phone_id
from
  duo_phone
where
  not activated
order by
  number,
  extension;
```

```sql+sqlite
select
  number,
  extension,
  phone_id
from
  duo_phone
where
  not activated
order by
  number,
  extension;
```

### Phones by platform

```sql+postgres
select
  platform,
  count(*)
from
  duo_phone
group by
  platform
order by
  platform;
```

```sql+sqlite
select
  platform,
  count(*)
from
  duo_phone
group by
  platform
order by
  platform;
```

### Users of phones that have been tampered with

```sql+postgres
select
  u->>'username' as username,
  p.number,
  p.extension
from
  duo_phone as p,
  jsonb_array_elements(p.users) as u
where
  p.tampered = 'Tampered'
order by
  username,
  number,
  extension;
```

```sql+sqlite
select
  json_extract(u.value, '$.username') as username,
  p.number,
  p.extension
from
  duo_phone as p,
  json_each(p.users) as u
where
  p.tampered = 'Tampered'
order by
  username,
  number,
  extension;
```

### Users of phones without encryption

```sql+postgres
select
  u->>'username' as username,
  p.number,
  p.extension,
  p.encrypted
from
  duo_phone as p,
  jsonb_array_elements(p.users) as u
where
  p.encrypted is null
  or p.encrypted != 'Encrypted'
order by
  username,
  number,
  extension;
```

```sql+sqlite
select
  json_extract(u.value, '$.username') as username,
  p.number,
  p.extension,
  p.encrypted
from
  duo_phone as p,
  json_each(p.users) as u
where
  p.encrypted is null
  or p.encrypted != 'Encrypted'
order by
  username,
  number,
  extension;
```
