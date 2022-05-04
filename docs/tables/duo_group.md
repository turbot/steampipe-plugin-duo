# Table: duo_group

Groups in the Duo account.

## Examples

### List all groups

```sql
select
  name,
  group_id,
  status,
  description
from
  duo_group
order by
  name
```

### Group statistics by status

```sql
select
  status,
  count(*)
from
  duo_group
group by
  status
order by
  status
```

### List disabled groups

```sql
select
  name
from
  duo_group
where
  status = 'disabled'
order by
  name
```
