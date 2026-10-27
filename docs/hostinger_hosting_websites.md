## hostinger hosting websites

Websites commands

```
hostinger hosting websites [flags]
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

* [hostinger hosting](hostinger_hosting.md)	 - Hosting commands
* [hostinger hosting websites create](hostinger_hosting_websites_create.md)	 - Create website
* [hostinger hosting websites delete](hostinger_hosting_websites_delete.md)	 - Delete website
* [hostinger hosting websites deploy-static-site-archive](hostinger_hosting_websites_deploy-static-site-archive.md)	 - Deploy static site archive
* [hostinger hosting websites list](hostinger_hosting_websites_list.md)	 - List websites
* [hostinger hosting websites list-setups](hostinger_hosting_websites_list-setups.md)	 - List website setups
* [hostinger hosting websites start-setup](hostinger_hosting_websites_start-setup.md)	 - Start website setup

