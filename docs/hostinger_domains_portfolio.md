## hostinger domains portfolio

Portfolio commands

```
hostinger domains portfolio [flags]
```

### Options

```
  -h, --help   help for portfolio
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

* [hostinger domains](hostinger_domains.md)	 - Domains commands
* [hostinger domains portfolio authorization-code](hostinger_domains_portfolio_authorization-code.md)	 - Get domain authorization code
* [hostinger domains portfolio claim-free](hostinger_domains_portfolio_claim-free.md)	 - Claim free domain
* [hostinger domains portfolio complete-setup](hostinger_domains_portfolio_complete-setup.md)	 - Complete domain setup
* [hostinger domains portfolio disable-lock](hostinger_domains_portfolio_disable-lock.md)	 - Disable domain lock
* [hostinger domains portfolio disable-privacy-protection](hostinger_domains_portfolio_disable-privacy-protection.md)	 - Disable privacy protection
* [hostinger domains portfolio enable-lock](hostinger_domains_portfolio_enable-lock.md)	 - Enable domain lock
* [hostinger domains portfolio enable-privacy-protection](hostinger_domains_portfolio_enable-privacy-protection.md)	 - Enable privacy protection
* [hostinger domains portfolio get](hostinger_domains_portfolio_get.md)	 - Get domain details
* [hostinger domains portfolio list](hostinger_domains_portfolio_list.md)	 - Get domain list
* [hostinger domains portfolio purchase](hostinger_domains_portfolio_purchase.md)	 - Purchase new domain
* [hostinger domains portfolio renewal-information](hostinger_domains_portfolio_renewal-information.md)	 - Get domain renewal information
* [hostinger domains portfolio update-nameservers](hostinger_domains_portfolio_update-nameservers.md)	 - Update domain nameservers

