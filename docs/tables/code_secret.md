---
title: "Steampipe Table: code_secret - Query Code Secrets using SQL"
description: "Allows users to query Code Secrets, specifically the secret's name, description, created and updated timestamps, and the secret's content. This provides insights into the secrets stored in the Code service."
---

# Table: code_secret - Query Code Secrets using SQL

A Code Secret is a resource within Oracle Cloud Infrastructure (OCI) that allows you to manage and store secrets such as database credentials, API keys, and other sensitive information. It provides a centralized way to manage secrets for various OCI resources, including databases, applications, and more. Code Secrets help you keep your sensitive information secure and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `code_secret` table provides insights into secrets within Oracle Cloud Infrastructure (OCI) Code Secrets. As a DevOps engineer, explore secret-specific details through this table, including secret names, descriptions, and associated metadata. Utilize it to uncover information about secrets, such as their creation and update timestamps, and the content of the secrets.

## Examples

### Basic auth is detected
Identify instances where sensitive information like passwords are exposed in your code. This can help in enhancing the security by preventing potential data breaches.

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
Discover the segments that include multiple matched secrets, such as AWS access keys, to enhance security measures and prevent potential data breaches. This query is useful in identifying and managing sensitive information within your codebase.

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
Determine the types of secrets embedded in your code to enhance security measures. This query helps in identifying potential vulnerabilities by highlighting different types of secrets such as AWS keys, Github tokens, etc., that might be accidentally exposed in the code.

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
Determine the areas in which sensitive information might be inadvertently exposed in AWS EC2 instance user data. This is crucial for maintaining security standards and preventing unauthorized access.

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
Analyze the settings to understand potential security risks in your AWS CloudFormation stack templates. This query is useful for identifying hidden secrets within your stack templates, helping you maintain secure and authenticated configurations.

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