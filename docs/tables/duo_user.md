# Table: duo_user

Users in the Duo account.

## Examples

### List all users

```sql
select
  user_id,
  username,
  email
from
  duo_user
order by
  username
```

### Users who have not yet enrolled

```sql
select
  user_id,
  username,
  email
from
  duo_user
where
  not is_enrolled
order by
  username
```

### Most recent 10 users to login

```sql
select
  user_id,
  username,
  email,
  last_login
from
  duo_user
order by
  last_login desc
limit 10
```

### Users who have never logged in

```sql
select
  user_id,
  username,
  email
from
  duo_user
where
  last_login is null
order by
  username
```

### Users who have not logged in for more than 30 days

```sql
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
  age desc
```

### Users who are locked out

```sql
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
  username
```

### User statistics by status

```sql
select
  status,
  count(*)
from
  duo_user
group by
  status
order by
  status
```

### Users and their group memberships

```sql
select
  u.username,
  g->>'name' as groupname
from
  duo_user as u,
  jsonb_array_elements(u.groups) as g
order by
  username,
  groupname
```

### Users and their phones

```sql
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
  phone_extension
```

### Users and their hardware tokens

```sql
select
  u.username,
  t->>'serial' as token_serial
from
  duo_user as u,
  jsonb_array_elements(u.tokens) as t
order by
  username,
  token_serial
```
