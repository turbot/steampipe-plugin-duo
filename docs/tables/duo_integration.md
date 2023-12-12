---
title: "Steampipe Table: duo_integration - Query Duo Security Integrations using SQL"
description: "Allows users to query Duo Security Integrations, particularly the integration key, secret key, and API hostname, providing insights into integration configurations and potential security issues."
---

# Table: duo_integration - Query Duo Security Integrations using SQL

Duo Security is a cloud-based security platform that protects access to sensitive data across applications and devices. It provides a range of security features, including two-factor authentication, secure single sign-on, and adaptive authentication. Duo Security Integration is a feature that allows the platform to work seamlessly with various applications and systems, enhancing the overall security posture.

## Table Usage Guide

The `duo_integration` table provides insights into Duo Security Integrations. As a security engineer, explore integration-specific details through this table, including integration keys, secret keys, and API hostnames. Utilize it to uncover information about integrations, such as the security configuration of each integration, the associated applications, and potential security issues.

## Examples

### List all integrations

```sql+postgres
select
  name,
  integration_key,
  type
from
  duo_integration
order by
  name;
```

```sql+sqlite
select
  name,
  integration_key,
  type
from
  duo_integration
order by
  name;
```

### Integrations granted permission to administrator methods

```sql+postgres
select
  name,
  integration_key,
  type
from
  duo_integration
where
  adminapi_admins = 1
order by
  name;
```

```sql+sqlite
select
  name,
  integration_key,
  type
from
  duo_integration
where
  adminapi_admins = 1
order by
  name;
```
