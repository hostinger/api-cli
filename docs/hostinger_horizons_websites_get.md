## hostinger horizons websites get

Get website

### Synopsis

Get the link for the user to open their website in Hostinger Horizons interface.\n
Use this tool when the user wants the link to an existing website, or when you need its
website URL before or after editing it.\n
`is_in_progress` is true while changes are being generated or the website is being published;
wait until it is false before publishing. `published_at` is when the website was last published,
`is_template` is whether its published pages show the "Use template" banner, and
`has_ecommerce_store` is whether it has an online store.\n
Websites can be edited with the `Edit website` tool, or by the user in Hostinger Horizons
interface in the provided website URL.

```
hostinger horizons websites get <website-id> [flags]
```

### Examples

```
  hostinger horizons websites get 123e4567-e89b-12d3-a456-426614174000
```

### Options

```
  -h, --help   help for get
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

