## hostinger reach templates list-email

List email templates

### Synopsis

Get a list of the email templates in a profile, most recently updated first.

Templates are the reusable email bodies a campaign is built from. The list is not paginated
and only the metadata is returned - the template content itself is not exposed. Use the
`uuid` of a template as the `template_uuid` when creating a campaign.

```
hostinger reach templates list-email <profile-uuid> [flags]
```

### Options

```
  -h, --help   help for list-email
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach templates](hostinger_reach_templates.md)	 - Templates commands

