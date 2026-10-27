## hostinger reach templates list-email

List email templates

### Synopsis

Get a list of the email templates in a profile, most recently updated first.

Templates are the reusable email bodies a campaign is built from. The list is not paginated
and only the metadata is returned, without the HTML content. Use the `uuid` of a template
as the `template_uuid` when creating a campaign. To read one template's HTML use the get
email template endpoint, and only when the content itself is needed.

```
hostinger reach templates list-email <profile-uuid> [flags]
```

### Examples

```
  hostinger reach templates list-email 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for list-email
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger reach templates](hostinger_reach_templates.md)	 - Templates commands

