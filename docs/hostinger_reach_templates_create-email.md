## hostinger reach templates create-email

Create an email template

### Synopsis

Create an email template in a profile.

The template holds the HTML body a campaign reuses, so it can be created before any
campaign exists. Only the template metadata comes back - keep the returned `uuid` to
reference it as the `template_uuid` of a campaign.

```
hostinger reach templates create-email <profile-uuid> [flags]
```

### Examples

```
  hostinger reach templates create-email 550e8400-e09b-41d4-a716-400055000000 --template-content '<html><body><h1>Email Template</h1><p>Hello, traveler!</p></body></html>'
```

### Options

```
  -h, --help                      help for create-email
      --template-content string   The email body as HTML. It is sanitised before it is stored, so the saved template can differ from what was sent - inline any styles the email clients need and keep the markup self-contained.
      --title string              Name the template is listed under. Not shown to the recipients.
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

