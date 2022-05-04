# Table: duo_integration

Integrations in the Duo account.

## Examples

### List all integrations

```sql
select
  name,
  integration_key,
  type
from
  duo_integration
order by
  name
```

### Integrations granted permission to administrator methods

```sql
select
  name,
  integration_key,
  type
from
  duo_integration
where
  adminapi_admins = 1
order by
  name
```
