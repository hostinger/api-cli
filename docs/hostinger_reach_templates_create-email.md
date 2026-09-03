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

### Options

```
  -h, --help                      help for create-email
      --template-content string   The email body as HTML. It is sanitised before it is stored, so the saved template
                                  can differ from what was sent - inline any styles the email clients need and keep
                                  the markup self-contained.
      --title string              Name the template is listed under. Not shown to the recipients.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach templates](hostinger_reach_templates.md)	 - Templates commands

