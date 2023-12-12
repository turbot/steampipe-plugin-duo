---
title: "Steampipe Table: duo_account_settings - Query Duo Security Account Settings using SQL"
description: "Allows users to query Duo Security Account Settings, providing detailed information about the security settings and configurations applied to the Duo account."
---

# Table: duo_account_settings - Query Duo Security Account Settings using SQL

Duo Security Account Settings is a feature within Duo Security that allows administrators to manage and configure the security settings for their Duo account. It provides a centralized way to set up and manage settings for various Duo resources, including user policies, integrations, and more. Duo Security Account Settings helps you stay informed about the security and configuration status of your Duo resources and take appropriate actions when required.

## Table Usage Guide

The `duo_account_settings` table provides insights into Account Settings within Duo Security. As a Security Administrator, explore specific details through this table, including security restrictions, policy settings, and associated metadata. Utilize it to uncover information about settings, such as those related to user policies, integrations, and to verify the current configuration status.

## Examples

### Get account settings

```sql+postgres
select
  *
from
  duo_account_settings;
```

```sql+sqlite
select
  *
from
  duo_account_settings;
```

### Check password settings are secure

```sql+postgres
select
  name,
  case when password_requires_lower_alpha then '✅' else '❌' end as requires_lower_alpha,
  case when password_requires_upper_alpha then '✅' else '❌' end as requires_upper_alpha,
  case when password_requires_numeric then '✅' else '❌' end as requires_numeric,
  case when password_requires_special then '✅' else '❌' end as requires_special,
  case when minimum_password_length >= 12 then '✅' else '❌' end as min_length_gte_12
from
  duo_account_settings;
```

```sql+sqlite
select
  name,
  case when password_requires_lower_alpha then '✅' else '❌' end as requires_lower_alpha,
  case when password_requires_upper_alpha then '✅' else '❌' end as requires_upper_alpha,
  case when password_requires_numeric then '✅' else '❌' end as requires_numeric,
  case when password_requires_special then '✅' else '❌' end as requires_special,
  case when minimum_password_length >= 12 then '✅' else '❌' end as min_length_gte_12
from
  duo_account_settings;
```
