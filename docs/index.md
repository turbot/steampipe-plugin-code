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

[Steampipe](https://steampipe.io) is an open source CLI to instantly query local data or APIs (cloud or not) using SQL.

List secrets in your source code:

```sql
select
  secret_type,
  secret,
  authenticated,
  line,
  col
from
  code_secret
where
  src =
    E'Mixed secrets are matched:\n'
    '* Slack: xoxp-5228148520-5228148525-1323104836872-10674849628c43b9d4b4660f7f9a7b65\n'
    '* AWS: AKIA4YFAKFKFYXTDS353\n'
    '* Basic auth: https://joe:passwd123@example.com/secret'
    '* Stripe: sk_live_tR3PYbcVNZZ796tH88S4VQ2u'
order by
  secret_type;
```

```
+-------------------+---------------------------------------------------------------------------+-----------------+------+-----+
| secret_type       | secret                                                                    | authenticated   | line | col |
+-------------------+---------------------------------------------------------------------------+-----------------+------+-----+
| aws_access_key_id | AKIA4YFAKFKFYXTDS353                                                      | not_implemented | 3    | 8   |
| basic_auth        | https://joe:passwd123                                                     | not_implemented | 4    | 15  |
| slack_api_token   | xoxp-5228148520-5228148525-1323104836872-10674849628c43b9d4b4660f7f9a7b65 | unauthenticated | 2    | 10  |
| stripe_api_key    | sk_live_tR3PYbcVNZZ796tH88S4VQ2u                                          | unauthenticated | 5    | 11  |
+-------------------+---------------------------------------------------------------------------+-----------------+------+-----+
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

## Supported secret types

| Secret                       | Slug                         | Authentication |
| :--------------------------- | :--------------------------- | :-----------   |
| AWS Access Key ID            | aws_access_key_id            | Available      |
| Azure Storage Account Key    | azure_storage_account_key    | N/A            |
| Basic Auth                   | basic_auth                   | N/A            |
| Facebook Access Token        | facebook_access_token        | N/A            |
| Facebook OAuth               | facebook_oauth               | N/A            |
| Facebook Secret Key          | facebook_secret_key          | N/A            |
| GitHub App Token             | github_app_token             | N/A            |
| GitHub OAuth Access Token    | github_oauth_access_token    | N/A            |
| GitHub Personal Access Token | github_personal_access_token | N/A            |
| GitHub Refresh Token         | github_refresh_token         | N/A            |
| Google API Key               | google_api_key               | N/A            |
| JWT                          | jwt                          | N/A            |
| Mailchimp Access Key         | mailchimp_access_key         | Available      |
| Okta Token                   | okta_token                   | N/A            |
| Slack API Token              | slack_api_token              | Available      |
| Stripe API Key               | stripe_api_key               | Available      |
| Twilio Auth Token            | twilio_auth_token            | N/A            |
| Twitter Secret Key           | twitter_secret_key           | N/A            |

### Authentication Status

For secret types that support authentication, the results are returned in the `authenticated` column with one of the following values:

- `authenticated`: Secret is active
- `unauthenticated`: Secret is inactive
- `not_implemented`: Secret was not tested due to lack of authentication function
- `unknown`: Secret was tested but results were inconclusive

## Credits

- The `code_secret` table is based on [Yelp's detect secrets](https://github.com/Yelp/detect-secrets)
  project. The general matching approach and regular expressions are copied and
  based on their amazing work.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-code
- Community: [Slack Channel](https://steampipe.io/community/join)
