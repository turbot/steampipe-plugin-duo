---
title: "Steampipe Table: duo_token - Query Duo Security Tokens using SQL"
description: "Allows users to query Duo Security Tokens, providing insights into the two-factor authentication (2FA) tokens in use within an organization."
---

# Table: duo_token - Query Duo Security Tokens using SQL

Duo Security is a cloud-based trusted access provider protecting the world's fastest-growing and largest companies and thousands of organizations worldwide, including Dresser-Rand, Etsy, Facebook, K-Swiss, Random House, Yelp, Zillow, Paramount Pictures, and more. Duo Security's innovative and easy-to-use technology can be quickly deployed to protect users, data, and applications from breaches, credential theft, and account takeover. Token is a hardware device that generates passcodes for login to Duo-protected services and applications.

## Table Usage Guide

The `duo_token` table provides insights into Duo Security Tokens, which are integral to the two-factor authentication process. As a security engineer, explore token-specific details through this table, including token type, serial number, and associated users. Utilize it to manage and track the usage of tokens within your organization, ensuring appropriate distribution and usage.

## Examples

### List all tokens

```sql+postgres
select
  serial,
  token_id,
  type
from
  duo_token
order by
  serial;
```

```sql+sqlite
select
  serial,
  token_id,
  type
from
  duo_token
order by
  serial;
```

### Tokens and their users

```sql+postgres
select
  t.serial,
  u->>'username' as username
from
  duo_token as t,
  jsonb_array_elements(t.users) as u
order by
  serial,
  username;
```

```sql+sqlite
select
  t.serial,
  json_extract(u.value, '$.username') as username
from
  duo_token as t,
  json_each(t.users) as u
order by
  t.serial,
  username;
```

### Tokens by platform

```sql+postgres
select
  type,
  count(*)
from
  duo_token
group by
  type
order by
  type;
```

```sql+sqlite
select
  type,
  count(*)
from
  duo_token
group by
  type
order by
  type;
```
