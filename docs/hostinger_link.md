## hostinger link

Link this directory to a website

### Synopsis

Remember which website the commands run from this directory act on, in
.hostinger/site.json. Commands that take [domain] then use it when the domain is
left out, and `hostinger up` redeploys to it.

With --global the website is written as default_domain in the config file instead,
for commands run outside any linked directory.

With --show nothing is changed: it prints what commands run from this directory
would act on and where each value comes from (flags, environment, the link, the
config file).

```
hostinger link [domain] [flags]
```

### Options

```
      --global   Set default_domain in the config file instead of linking this directory
  -h, --help     help for link
      --show     Print what commands run from this directory act on, and where each value comes from
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

* [hostinger](hostinger.md)	 - Hostinger API Command Line Interface

