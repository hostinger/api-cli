## hostinger hosting ssl toggle-https-redirect

Toggle HTTPS redirect

### Synopsis

Turns the HTTP to HTTPS redirect of the website on or off, based on `is_enabled`. Does
nothing when the redirect is already in the requested state. Turning it on requires an
installed certificate (`status` `active` or `expired` on `Get SSL status`) and returns 422
when there is none; turning it off is always accepted.

```
hostinger hosting ssl toggle-https-redirect <username> <domain> [flags]
```

### Examples

```
  hostinger hosting ssl toggle-https-redirect u123456789 mydomain.tld --is-enabled=true
```

### Options

```
  -h, --help         help for toggle-https-redirect
      --is-enabled   Turn the HTTP to HTTPS redirect on (true) or off (false) for the website.
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

* [hostinger hosting ssl](hostinger_hosting_ssl.md)	 - SSL commands

