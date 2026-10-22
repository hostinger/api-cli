## hostinger reach forms update-template

Update form template

### Synopsis

Save `content` as the latest design of a form and return the stored version.

The HTML is sanitized and, unless the plan allows removing it, the Reach signature is added. Visitors
submit the form with an `email` field (required) and optional `name`, `surname` and `phone`. After
sanitizing, the design must contain a form container (`data-reach-form`), an email input
(`data-reach-field="email"`) and a submit button (`data-reach-submit`); otherwise the request fails
with `422` and `errors.content` lists `missing_form_container`, `missing_email_field` and/or
`missing_submit_button`.

A `draft` or `paused` form stays unreachable until its status is set to `active`, which publishes this
design. On a form that is already `active` the saved design is not public until the form is paused and
set to `active` again. Saving a design does not consume AI credits.

```
hostinger reach forms update-template <profile-uuid> <form-uuid> [flags]
```

### Examples

```
  hostinger reach forms update-template 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000 --content '<form data-reach-form><input type="email" name="email" data-reach-field="email" /><button type="submit" data-reach-submit>Subscribe</button></form>'
```

### Options

```
      --content data-reach-form   HTML of the form design. After sanitizing it must still contain a form container (data-reach-form), an email input (`data-reach-field="email"`) and a submit button (`data-reach-submit`).
  -h, --help                      help for update-template
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

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

