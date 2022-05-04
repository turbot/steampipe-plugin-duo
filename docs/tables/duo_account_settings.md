# Table: duo_account_settings

Settings for the Duo account.

## Examples

### Get account settings

```sql
select
  *
from
  duo_account_settings
```

### Check password settings are secure

```sql
select
  name,
  case when password_requires_lower_alpha then '✅' else '❌' end as requires_lower_alpha,
  case when password_requires_upper_alpha then '✅' else '❌' end as requires_upper_alpha,
  case when password_requires_numeric then '✅' else '❌' end as requires_numeric,
  case when password_requires_special then '✅' else '❌' end as requires_special,
  case when minimum_password_length >= 12 then '✅' else '❌' end as min_length_gte_12
from
  duo_account_settings
```
