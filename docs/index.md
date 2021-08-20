---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/code.svg"
brand_color: "#00b050"
display_name: "Code"
short_name: "code"
description: "Steampipe plugin to query secrets and more from Code."
og_description: "Query source code with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/code-social-graphic.png"
---

# Code + Steampipe

Source code can be any string or data for querying.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List instances in your Code account:

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

```
+--------------------+----------------------+------+-----+
| secret_type        | secret               | line | col |
+--------------------+----------------------+------+-----+
| aws_iam_access_key | AKIA4YFAKEKEYXTDS252 | 1    | 23  |
+--------------------+----------------------+------+-----+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/code/tables)**

## Get started

### Install

Download and install the latest Code plugin:

```bash
steampipe plugin install code
```

### Credentials

No credentials are required.

### Configuration

Installing the latest code plugin will create a config file (`~/.steampipe/config/code.spc`) with a single connection named `code`:

```hcl
connection "code" {
  plugin = "code"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-code
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
