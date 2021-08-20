![image](https://hub.steampipe.io/images/plugins/turbot/code-social-graphic.png)

# Code Plugin for Steampipe

Use SQL to query secrets and more from source code.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/code)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/code/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-code/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install code
```

Run a query:

```sql
select
  secret_type,
  secret,
  line,
  col
from
  code_secret
where
  src = 'Detect AWS access key AKIA4YFAKEKEYXTDS252!'
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-code.git
cd steampipe-plugin-code
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/code.spc
```

Try it!

```
steampipe query
> .inspect code
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)


## Credits

* The `code_secret` table is based on [Yelp's detect secrets](https://github.com/Yelp/detect-secrets)
  project. The general matching approach and regular expressions are copied and
  based on their amazing work.


## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-code/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Code Plugin](https://github.com/turbot/steampipe-plugin-code/labels/help%20wanted)
