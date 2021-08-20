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
    'AWS access keys should be detected:\n'
    '* AKIA4YFAKEKEYXTDS252\n* AKIA9YFBKFGZYZTW387K'
```

### Secrets of multiple types

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
    'Mixed secrets are matched:\n'
    '* Slack: xoxp-5228148520-5228148525-1323104836872-10674849628c43b9d4b4660f7f9a7b65\n'
    '* AWS: AKIA4YFAKFKFYXTDS353\n'
    '* Basic auth: https://joe:passwd123@example.com/secret'
```
