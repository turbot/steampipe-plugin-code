# Table: code_secret

Detect, and verify if possible, secrets in a given source string.

Note: All queries to this table must provide the `src` column.

## Examples

### Basic auth is detected

```sql
select
  secret_type,
  secret,
  line,
  col
from
  code_secret
where
  src = 'Text with a secret postgresql://user:secret@localhost:5432/mydb.'
```

### Multiple secrets are matched, including AWS access keys

```sql
select
  secret_type,
  secret,
  line,
  col
from
  code_secret
where
  src =
    E'AWS access keys should be detected:\n'
    '* AKIA4YFAKEKEYXTDS252\n* AKIA9YFBKFGZYZTW387K'
```

### Secrets of multiple types

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
    '* Replace and try <AWS_ACCESS_KEY>:<AWS_SECRET_KEY> \n'
    '* Basic auth: https://joe:passwd123@example.com/secret'
    '* Github Personal Access Token: 45ab6f911111f9f376a5b52c25d22113f2b45fa1'
    '* Okta Token: 00Am7B2M_U-63q_Ppd6tDzAbBOkvcCht-kDG-baM7t'
    '* Stripe Api Key: sk_live_tR3PYbcVNZZ796tH88S4VQ2u'
    '* Azure Storage Account Key: mllhBNrG467B7Q5iT+ePFr6eLCE24ij9vT/fCeckOunfqzoGm8k5X9vKCphDaO81gmuzr89ldN+gKB0vlEHahg=='
```

### Detect secrets in AWS EC2 instance user data (requires AWS plugin)

```sql
select
  instance_id,
  region as instance_region,
  secret_type,
  secret,
  authenticated,
  line,
  col
from
  code_secret,
  aws_ec2_instance
where
  src = user_data;
```

### Detect secrets in AWS CloudFormation stack template body (requires AWS plugin)

```sql
select
  id as stack_id,
  name as stack_name,
  region as stack_region,
  secret_type,
  secret,
  authenticated,
  line,
  col
from
  code_secret,
  aws_cloudformation_stack
where
  src = template_body;
```
