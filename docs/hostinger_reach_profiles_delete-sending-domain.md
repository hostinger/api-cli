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

### Options

```
      --domain string   
  -h, --help            help for delete-sending-domain
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

