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

| Secret                       | Slug                         | Verification |
| :--------------------------- | :--------------------------- | :----------- |
| AWS access key ID            | aws_access_key_id            | Available    |
| Azure storage account key    | azure_storage_account_key    | N/A          |
| Basic auth                   | basic_auth                   | N/A          |
| Facebook OAuth               | facebook_oauth               | N/A          |
| Facebook access token        | facebook_access_token        | N/A          |
| Facebook secret key          | facebook_secret_key          | N/A          |
| GitHub OAuth access token    | github_oauth_access_token    | N/A          |
| GitHub app token             | github_app_token             | N/A          |
| GitHub personal access token | github_personal_access_token | N/A          |
| GitHub refresh token         | github_refresh_token         | N/A          |
| Google API key               | google_api_key               | N/A          |
| JWT                          | jwt                          | N/A          |
| Mailchimp access key         | mailchimp_access_key         | Available    |
| Okta token                   | okta_token                   | N/A          |
| Slack API token              | slack_api_token              | Available    |
| Stripe API key               | stripe_api_key               | Available    |
| Twilio auth token            | twilio_auth_token            | N/A          |
| Twitter secret key           | twitter_secret_key           | N/A          |

## Verification Status

Verification status of the secret. Valid values are:

- `verified true` secret is still active can we used for leakage.
- `verified false` means secret is inactive (i.e it may have expired or does not exit anymore, etc... ) and is not working currently.
- `unverified` means status of the key could not be verified.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-code
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
