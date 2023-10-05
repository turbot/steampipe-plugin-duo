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
