## hostinger reach profiles delete-sending-domain

Delete sending domain

### Synopsis

Disconnect the sending domain from the profile.

The domain record is removed, so the profile has no sending domain afterwards and campaigns
cannot be sent until a domain is connected again. Sender addresses already stored on
campaigns, automations and settings are left untouched and keep pointing at the removed
domain. A domain that is not connected to the profile also succeeds, so success is not proof
that anything was deleted. DNS records at the DNS provider are not removed. To switch
domains, delete this one and connect the new one. Confirm the domain with the customer
before calling.

```
hostinger reach profiles delete-sending-domain <profile-uuid> [flags]
```

### Examples

```
  hostinger reach profiles delete-sending-domain 550e8400-e09b-41d4-a716-400055000000 --domain example.com
```

### Options

```
      --domain string   
  -h, --help            help for delete-sending-domain
  -y, --yes             Skip the confirmation prompt
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

