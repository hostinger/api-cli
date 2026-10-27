## hostinger agency-hosting websites

Websites commands

```
hostinger agency-hosting websites [flags]
```

### Options

```
  -h, --help   help for websites
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

* [hostinger agency-hosting](hostinger_agency-hosting.md)	 - Agency Hosting commands
* [hostinger agency-hosting websites build-nodejs-assets](hostinger_agency-hosting_websites_build-nodejs-assets.md)	 - Build website NodeJS assets
* [hostinger agency-hosting websites delete](hostinger_agency-hosting_websites_delete.md)	 - Delete website
* [hostinger agency-hosting websites get](hostinger_agency-hosting_websites_get.md)	 - Get website details
* [hostinger agency-hosting websites list-plan](hostinger_agency-hosting_websites_list-plan.md)	 - List Agency Plan websites
* [hostinger agency-hosting websites list-processes](hostinger_agency-hosting_websites_list-processes.md)	 - List website processes

