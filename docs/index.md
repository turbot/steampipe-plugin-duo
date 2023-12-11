---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/duo.svg"
brand_color: "#6BBF4E"
display_name: Duo Security
name: duo
description: Steampipe plugin for querying Duo Security users, logs and more.
og_description: Query Duo Security with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/duo-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Duo Security + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Duo Security](https://duo.com) provides cloud-based two-factor authentication services.

Example query:

```sql
select
  username,
  is_enrolled,
  last_login
from
  duo_user
order by
  username;
```

```
+----------+-------------+---------------------------+
| username | is_enrolled | last_login                |
+----------+-------------+---------------------------+
| dwight   | true        | 2022-04-17T07:36:34-04:00 |
| jim      | true        | 2022-04-17T09:36:34-04:00 |
| michael  | false       | <null>                    |
| pam      | true        | 2022-04-17T08:55:34-04:00 |
+----------+-------------+---------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/duo/tables)**

## Get started

### Install

Download and install the latest Duo plugin:

```bash
steampipe plugin install duo
```

### Configuration

Installing the latest duo plugin will create a config file (`~/.steampipe/config/duo.spc`) with a single connection named `duo`:

```hcl
connection "duo" {
  plugin          = "duo"
  api_hostname    = "api-28bcd3da.duosecurity.com"
  integration_key = "DINXR28B7BSL5NB362QR"
  secret_key      = "Xo8ZbvGLOLkw8iFowK34Mp2LOqQEh7cxeMmDoHSO"
}
```

- `api_hostname` - Unique API endpoint for your account, [learn more](https://duo.com/docs/adminapi#first-steps).
- `integration_key` - Integration key for your account, [learn more](https://duo.com/docs/adminapi#first-steps).
- `secret_key` - Secret key, [learn more](https://duo.com/docs/adminapi#first-steps).

Environment variables are also available as an alternate configuration method:
- `DUO_API_HOSTNAME`
- `DUO_INTEGRATION_KEY`
- `DUO_SECRET_KEY`

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-duo
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
