## hostinger horizons websites publish

Publish website

### Synopsis

Publish a Hostinger Horizons website so its latest changes go live.\n
Use this tool when the user asks to publish, deploy or make their website live.\n
This tool starts the publish process and returns the URL the website will be live on.
Publishing happens asynchronously and takes a few minutes.\n
Set `is_template` only when the user explicitly asks to share the website as a template:
true adds a "Use template" banner to its published pages that copies the website into the
visitor's own account, and false removes it. Leave it out to keep the current setting.\n
After invoking this tool, your chat reply must be EXACTLY 1 sentence summarizing
that the website is being published and you should provide the published URL to the user immediately.

```
hostinger horizons websites publish <website-id> [flags]
```

### Examples

```
  hostinger horizons websites publish 123e4567-e89b-12d3-a456-426614174000
```

### Options

```
  -h, --help          help for publish
      --is-template   Set to true to publish the website as a template: its published pages show a "Use template" banner that copies the website into the visitor's own account. Set to false to remove the banner. Leave it out to keep the current setting.
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

* [hostinger horizons websites](hostinger_horizons_websites.md)	 - Websites commands

