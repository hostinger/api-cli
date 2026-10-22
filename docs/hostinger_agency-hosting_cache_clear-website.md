## hostinger agency-hosting cache clear-website

Clear website cache

### Synopsis

Clears cache for all domains associated with an Agency Plan website, including its preview domain.

This operation clears all cache types for the website.

```
hostinger agency-hosting cache clear-website <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting cache clear-website zpwlGlp19
```

### Options

```
  -h, --help   help for clear-website
  -y, --yes    Skip the confirmation prompt
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

* [hostinger agency-hosting cache](hostinger_agency-hosting_cache.md)	 - Cache commands

