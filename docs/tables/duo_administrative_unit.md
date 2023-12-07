---
title: "Steampipe Table: duo_administrative_unit - Query Duo Security Administrative Units using SQL"
description: "Allows users to query Administrative Units in Duo Security, specifically the details of each unit such as ID, name, and description, providing insights into the organization's structure and user management."
---

# Table: duo_administrative_unit - Query Duo Security Administrative Units using SQL

Duo Security's Administrative Units is a feature that allows you to organize users and devices into different groups based on their roles, departments, or any other criteria. These units can be used to apply different policies and controls, ensuring appropriate access and security measures are in place. It is a crucial component in managing and securing access in an organization.

## Table Usage Guide

The `duo_administrative_unit` table provides insights into Administrative Units within Duo Security. As a security administrator, explore unit-specific details through this table, including ID, name, and description. Utilize it to uncover information about each unit, such as its role within the organization, the users and devices it contains, and the policies applied to it.

## Examples

### List all administrative units

```sql+postgres
select
  name,
  admin_unit_id,
  description
from
  duo_administrative_unit
order by
  name;
```

```sql+sqlite
select
  name,
  admin_unit_id,
  description
from
  duo_administrative_unit
order by
  name;
```

### Admin units with their specific groups

```sql+postgres
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
  group_name;
```

```sql+sqlite
select
  au.name,
  g.name as group_name
from
  duo_administrative_unit as au,
  json_each(au.groups) as aug,
  duo_group as g
where
  aug.value = g.group_id
order by
  au.name,
  group_name;
```

### Admin units with their specific integrations

```sql+postgres
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
  integration_name;
```

```sql+sqlite
select
  au.name,
  i.name as integration_name
from
  duo_administrative_unit as au,
  json_each(au.integrations) as aui,
  duo_integration as i
where
  aui.value = i.integration_key
order by
  au.name,
  integration_name;
```
