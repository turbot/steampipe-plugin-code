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
