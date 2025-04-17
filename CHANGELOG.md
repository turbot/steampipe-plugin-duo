## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#35](https://github.com/turbot/steampipe-plugin-duo/pull/35))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#35](https://github.com/turbot/steampipe-plugin-duo/pull/35))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#32](https://github.com/turbot/steampipe-plugin-duo/pull/32))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#32](https://github.com/turbot/steampipe-plugin-duo/pull/32))

## v0.4.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#21](https://github.com/turbot/steampipe-plugin-duo/pull/21))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#21](https://github.com/turbot/steampipe-plugin-duo/pull/21))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-duo/blob/main/docs/LICENSE). ([#21](https://github.com/turbot/steampipe-plugin-duo/pull/21))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#20](https://github.com/turbot/steampipe-plugin-duo/pull/20))

## v0.3.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#13](https://github.com/turbot/steampipe-plugin-duo/pull/13))

## v0.3.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#11](https://github.com/turbot/steampipe-plugin-duo/pull/11))
- Recompiled plugin with Go version `1.21`. ([#11](https://github.com/turbot/steampipe-plugin-duo/pull/11))

## v0.2.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#7](https://github.com/turbot/steampipe-plugin-duo/pull/7))

## v0.1.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#4](https://github.com/turbot/steampipe-plugin-duo/pull/4))
- Recompiled plugin with Go version `1.19`. ([#4](https://github.com/turbot/steampipe-plugin-duo/pull/4))

## v0.0.1 [2022-05-10]

_What's new?_

- New tables added
  - [duo_account_settings](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_account_settings)
  - [duo_account_summary](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_account_summary)
  - [duo_admin_log_record](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_admin_log_record)
  - [duo_administrative_unit](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_administrative_unit)
  - [duo_administrator](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_administrator)
  - [duo_auth_log_record](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_auth_log_record)
  - [duo_group](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_group)
  - [duo_integration](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_integration)
  - [duo_phone](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_phone)
  - [duo_token](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_token)
  - [duo_user](https://hub.steampipe.io/plugins/turbot/duo/tables/duo_user)
