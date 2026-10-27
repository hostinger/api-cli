## hostinger reach contact-fields create

Create a contact field

### Synopsis

Define a new custom contact field in a profile.

The `slug` is derived from the label and, like the field type, cannot be changed later.
Use the returned uuid to set values on contacts.

```
hostinger reach contact-fields create <profile-uuid> [flags]
```

### Examples

```
  hostinger reach contact-fields create 550e8400-e09b-41d4-a716-400055000000 --label 'Job title' --type text
```

### Options

```
  -h, --help              help for create
      --label string      
      --options strings   Required for single_choice and multi_choice, ignored for the scalar types. Labels must be unique regardless of casing.
      --type string       Immutable once the field exists (one of: text, number, date, single_choice, multi_choice)
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger reach contact-fields](hostinger_reach_contact-fields.md)	 - Contact Fields commands

