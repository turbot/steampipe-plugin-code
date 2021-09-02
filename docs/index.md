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
  verified,
  line,
  col
from
  code_secret
where
  src =
    'Mixed secrets are matched:\n'
    '* Slack: xoxp-5228148520-5228148525-1323104836872-10674849628c43b9d4b4660f7f9a7b65\n'
    '* AWS: AKIA4YFAKFKFYXTDS353\n'
    '* Basic auth: https://joe:passwd123@example.com/secret'
    '* Stripe: sk_live_tR3PYbcVNZZ796tH88S4VQ2u'
order by
  secret_type;
```

```
+-------------------+---------------------------------------------------------------------------+----------------+------+-----+
| secret_type       | secret                                                                    | verified       | line | col |
+-------------------+---------------------------------------------------------------------------+----------------+------+-----+
| aws_access_key_id | AKIA4YFAKFKFYXTDS353                                                      | unverified     | 1    | 120 |
| basic_auth        | https://joe:passwd123                                                     | unverified     | 1    | 156 |
| slack_api_token   | xoxp-5228148520-5228148525-1323104836872-10674849628c43b9d4b4660f7f9a7b65 | verified false | 1    | 38  |
| stripe_api_key    | sk_live_tR3PYbcVNZZ796tH88S4VQ2u                                          | verified false | 1    | 206 |
+-------------------+---------------------------------------------------------------------------+----------------+------+-----+
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

Installing the latest code plugin will create a config file (`~/.steampipe/config/code.spc`) with a single connection named `code`.

```hcl
connection "code" {
  plugin = "code"
}
```

### Code plugin currently scans for the following secrets:



| Secret                    | Regex                  |
| :------------------------ | :--------------------- |
| AWS access key id|`(?m)AKIA[0-9A-Z]{16}`|
| Azure storage account key|`[a-zA-Z0-9+/=]{88}`|
| Basic auth|`(?m)([a-zA-Z0-9+-\.]+://[^:/\?#\[\]@!\$&\\'\(\)\*\+,;=\s]+:[^:/\?#\[\]@!\$&\\'\(\)\*\+,;=\s]+)@`|
| Facebook access token|`(?m)EAACEdEose0cBA[0-9A-Za-z]+`|
| Facebook oauth|`(?im)[f|F][a|A][c|C][e|E][b|B][o|O][o|O][k|K].*['|\"][0-9a-f]{32}['|\"]`|
| Facebook secret key|`(?im)(facebook|fb)(.{0,20})?(?-i)['\"][0-9a-f]{32}`|
| Github app token|`(?m)(ghu|ghs)_[0-9a-zA-Z]{36}`|
| Github oauth access token|`(?m)gho_[0-9a-zA-Z]{36}`|
| Github personal access token|`(?m)(ghp|gho|ghu|ghs|ghr)_[A-Za-z0-9_]{36}` or `(?m)[0-9a-f]{40}`|
| Github refresh token|`(?m)ghr_[0-9a-zA-Z]{76}`|
| Google api key|`(?m)AIza[0-9A-Za-z\\-_]{35}`|
| JWT|`(?m)eyJ[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*`|
| Mailchimp access key|`(?m)[0-9a-z]{32}-us[0-9]{1,2}`|
| Okta token|`(?m)00[a-zA-Z0-9\-\_]{40}`|
| Slack api token|`(?m)xox(?:a|b|p|o|s|r)-(?:\d+-)+[a-z0-9]+`|
| Stripe api key|`(?m)(?:r|s)k_live_[0-9a-zA-Z]{24}`|
| Twilio auth token|`(?m)AC[a-z0-9]{32}` or `(?m)SK[a-z0-9]{32}`|
| Twitter secret key|`(?im)twitter(.{0,20})?['\"][0-9a-z]{35,44}`|


## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-code
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
