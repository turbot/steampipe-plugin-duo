---
title: "Steampipe Table: duo_administrator - Query Duo Security Administrators using SQL"
description: "Allows users to query Duo Security Administrators, specifically their roles, names, and other associated details, providing insights into administrative access and roles."
---

# Table: duo_administrator - Query Duo Security Administrators using SQL

Duo Security Administrators are the users who have access to manage and configure settings within the Duo Security system. These administrators have different roles, such as Owner, Billing, User Manager, and more, each with specific permissions and access rights. Understanding these roles and their associated details can provide valuable insights into the management and security of the Duo system.

## Table Usage Guide

The `duo_administrator` table provides insights into Duo Security Administrators. As a security professional, you can explore administrator-specific details through this table, including roles, names, and other associated details. Utilize it to uncover information about administrators, such as their roles in the system, their access rights, and the security implications of these roles.

## Examples

### List all administrators

```sql+postgres
select
  name,
  admin_id,
  email,
  role
from
  duo_administrator
order by
  name;
```

```sql+sqlite
select
  name,
  admin_id,
  email,
  role
from
  duo_administrator
order by
  name;
```

### Most recent 10 administrators to login

```sql+postgres
select
  admin_id,
  name,
  email,
  last_login
from
  duo_administrator
order by
  last_login desc
limit 10;
```

```sql+sqlite
select
  admin_id,
  name,
  email,
  last_login
from
  duo_administrator
order by
  last_login desc
limit 10;
```

### Administrators who have never logged in

```sql+postgres
select
  admin_id,
  name,
  email
from
  duo_administrator
where
  last_login is null
order by
  name;
```

```sql+sqlite
select
  admin_id,
  name,
  email
from
  duo_administrator
where
  last_login is null
order by
  name;
```
