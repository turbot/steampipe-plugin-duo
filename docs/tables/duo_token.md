# Table: duo_token

Tokens in the Duo account.

## Examples

### List all tokens

```sql
select
  serial,
  token_id,
  type
from
  duo_token
order by
  serial
```

### Tokens and their users

```sql
select
  t.serial,
  u->>'username' as username
from
  duo_token as t,
  jsonb_array_elements(t.users) as u
order by
  serial,
  username
```

### Tokens by platform

```sql
select
  type,
  count(*)
from
  duo_token
group by
  type
order by
  type
```
