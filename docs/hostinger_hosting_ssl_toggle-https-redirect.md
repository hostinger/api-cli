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

### Options

```
  -h, --help         help for toggle-https-redirect
      --is-enabled   Turn the HTTP to HTTPS redirect on (true) or off (false) for the website.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting ssl](hostinger_hosting_ssl.md)	 - SSL commands

