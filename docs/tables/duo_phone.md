# Table: duo_phone

Phones in the Duo account.

## Examples

### List all phones

```sql
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
  extension
```

### Phones and their users

```sql
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
  username
```

### Phones that are not yet activated

```sql
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
  extension
```

### Phones by platform

```sql
select
  platform,
  count(*)
from
  duo_phone
group by
  platform
order by
  platform
```

### Users of phones that have been tampered with

```sql
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
  extension
```

### Users of phones without encryption

```sql
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
  extension
```
