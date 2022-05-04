# Table: duo_administrator

Administrators in the Duo account.

## Examples

### List all administrators

```sql
select
  name,
  admin_id,
  email,
  role
from
  duo_administrator
order by
  name
```

### Most recent 10 administrators to login

```sql
select
  admin_id,
  name,
  email,
  last_login
from
  duo_administrator
order by
  last_login desc
limit 10
```

### Administrators who have never logged in

```sql
select
  admin_id,
  name,
  email
from
  duo_administrator
where
  last_login is null
order by
  name
```
