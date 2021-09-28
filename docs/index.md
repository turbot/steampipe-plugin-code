---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/code.svg"
brand_color: "#000000"
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

| Secrets                      | Secret slug                  |
| :--------------------------- | :--------------------------- |
| AWS access key id            | aws_access_key_id            |
| Azure storage account key    | azure_storage_account_key    |
| Basic auth                   | basic_auth                   |
| Facebook access token        | facebook_access_token        |
| Facebook oauth               | facebook_oauth               |
| Facebook secret key          | facebook_secret_key          |
| Github app token             | github_app_token             |
| Github oauth access token    | github_oauth_access_token    |
| Github personal access token | github_personal_access_token |
| Github refresh token         | github_refresh_token         |
| Google api key               | google_api_key               |
| JWT                          | jwt                          |
| Mailchimp access key         | mailchimp_access_key         |
| Okta token                   | okta_token                   |
| Slack api token              | slack_api_token              |
| Stripe api key               | stripe_api_key               |
| Twilio auth token            | twilio_auth_token            |
| Twitter secret key           | twitter_secret_key           |

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-code
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
