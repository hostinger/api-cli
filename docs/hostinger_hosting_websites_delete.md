## hostinger hosting websites delete

Delete website

### Synopsis

This endpoint permanently removes a website and all of its data. This action
cannot be undone. Before calling it, make sure the user understands the
consequences and explicitly confirms that they want to proceed.

All website files, databases and related configuration will be removed.
The hosting plan itself is kept, so a new website can be created on it afterwards.

Supported websites: main and addon domain websites on web hosting plans, and
Website Builder websites. Parked domains and subdomains cannot be deleted with
this endpoint. The domain must be the exact website domain, not a preview
domain or an alias.

Returns 404 when the domain does not exist or does not belong to the
authenticated client.

Website removal is processed asynchronously and can take a few minutes to
complete. The response returns before the removal finishes.

The website can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting websites delete [domain] [flags]
```

### Examples

```
  hostinger hosting websites delete mydomain.tld
```

### Options

```
  -h, --help   help for delete
  -y, --yes    Skip the confirmation prompt
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

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

