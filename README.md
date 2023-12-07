![image](https://hub.steampipe.io/images/plugins/turbot/duo-social-graphic.png)

# Duo Security Plugin for Steampipe

Use SQL to query users, logs and more from [Duo Security](https://duosecurity.com).

* **[Get started →](https://hub.steampipe.io/plugins/turbot/duo)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/duo/tables)
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-duo/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install duo
```

Run steampipe:

```shell
steampipe query
```

Run a query:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-duo.git
cd steampipe-plugin-duo
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/duo.spc
```

Try it!

```
steampipe query
> .inspect duo
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-duo/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-duo/blob/main/docs/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Duo Plugin](https://github.com/turbot/steampipe-plugin-duo/labels/help%20wanted)
