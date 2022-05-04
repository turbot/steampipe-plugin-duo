# Table: duo_administrative_unit

Administrators in the Duo account.

## Examples

### List all administrative units

```sql
select
  name,
  admin_unit_id,
  description
from
  duo_administrative_unit
order by
  name
```

### Admin units with their specific groups

For admin units restricted to specific groups, list the group associations.
Unrestricted admin units are not included in results.

```sql
select
  au.name,
  g.name as group_name
from
  duo_administrative_unit as au,
  jsonb_array_elements_text(au.groups) as aug,
  duo_group as g
where
  aug = g.group_id
order by
  au.name,
  group_name
```

### Admin units with their specific integrations

For admin units restricted to specific integrations, list the integration
associations. Unrestricted admin units are not included in results.

```sql
select
  au.name,
  i.name as integration_name
from
  duo_administrative_unit as au,
  jsonb_array_elements_text(au.integrations) as aui,
  duo_integration as i
where
  aui = i.integration_key
order by
  au.name,
  integration_name
```
