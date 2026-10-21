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

### Options

```
      --content data-reach-form   HTML of the form design. After sanitizing it must still contain a form container
                                  (data-reach-form), an email input (`data-reach-field="email"`) and a submit button
                                  (`data-reach-submit`).
  -h, --help                      help for update-template
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

