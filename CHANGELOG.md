## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#124](https://github.com/turbot/steampipe-plugin-code/pull/124))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#124](https://github.com/turbot/steampipe-plugin-code/pull/124))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#108](https://github.com/turbot/steampipe-plugin-code/pull/108))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#108](https://github.com/turbot/steampipe-plugin-code/pull/108))

## v0.7.0 [2024-03-04]

_Enhancements_

- Updated the regex pattern of `slack_api_token` to also detect the Slack bot tokens. ([#73](https://github.com/turbot/steampipe-plugin-code/pull/73))
- Updated the regex pattern of AWS `access_key_id` to include key resources like AWS SSO credentials. ([#74](https://github.com/turbot/steampipe-plugin-code/pull/74))

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#66](https://github.com/turbot/steampipe-plugin-code/pull/66))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#66](https://github.com/turbot/steampipe-plugin-code/pull/66))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-code/blob/main/docs/LICENSE). ([#66](https://github.com/turbot/steampipe-plugin-code/pull/66))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#65](https://github.com/turbot/steampipe-plugin-code/pull/65))

## v0.5.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#43](https://github.com/turbot/steampipe-plugin-code/pull/43))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#38](https://github.com/turbot/steampipe-plugin-code/pull/38))
- Recompiled plugin with Go version `1.21`. ([#38](https://github.com/turbot/steampipe-plugin-code/pull/38))

## v0.4.0 [2023-04-06]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#25](https://github.com/turbot/steampipe-plugin-code/pull/25))

## v0.3.0 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#21](https://github.com/turbot/steampipe-plugin-code/pull/21))
- Recompiled plugin with Go version `1.19`. ([#21](https://github.com/turbot/steampipe-plugin-code/pull/21))

## v0.2.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#16](https://github.com/turbot/steampipe-plugin-code/pull/16))

## v0.2.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#13](https://github.com/turbot/steampipe-plugin-code/pull/13))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#14](https://github.com/turbot/steampipe-plugin-code/pull/14))

## v0.1.0 [2021-12-08]

_Enhancements_

- Recompiled plugin with Go version 1.17 ([#9](https://github.com/turbot/steampipe-plugin-code/pull/9))
- Updated: Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#8](https://github.com/turbot/steampipe-plugin-code/pull/8))
- Added the Credits section to the `docs/index.md` file

_Bug fixes_

- Fixed the multi-line strings in the `code_secret` table document examples

## v0.0.1 [2021-09-28]

_What's new?_

- New tables:
  - [code_secret](https://hub.steampipe.io/plugins/turbot/code/tables/code_secret)
