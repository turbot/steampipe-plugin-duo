---
title: "Steampipe Table: duo_account_summary - Query Duo Security Account Summaries using SQL"
description: "Allows users to query Duo Security Account Summaries, providing access to details such as account name, account ID, and API hostname."
---

# Table: duo_account_summary - Query Duo Security Account Summaries using SQL

Duo Security is a cloud-based access security provider protecting the world's fastest-growing and largest companies and thousands of organizations worldwide. It provides a comprehensive security solution that verifies the identity of users and the health of their devices before granting them access to applications. Duo Security offers a range of authentication methods and flexible policies to ensure secure access to all applications.

## Table Usage Guide

The `duo_account_summary` table provides insights into Duo Security Account Summaries. As a security administrator, explore account-specific details through this table, including account name, account ID, and API hostname. Utilize it to uncover information about your Duo Security accounts, such as the identification of accounts for specific applications and the verification of API hostnames.

## Examples

### Get account summary

```sql+postgres
select
  *
from
  duo_account_summary;
```

```sql+sqlite
select
  *
from
  duo_account_summary;
```
