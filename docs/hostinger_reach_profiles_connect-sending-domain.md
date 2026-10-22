## hostinger reach profiles connect-sending-domain

Connect sending domain

### Synopsis

Connect a sending domain to the profile so emails can be sent from addresses on it.

Connecting does not make the domain ready to send: check the DNS status endpoint next for
records that are still missing. A domain the account owns or has already verified connects
immediately. Any other domain starts a TXT ownership verification and returns `409`;
complete the verification in Reach, then retry. Returns `422` when the profile or another
profile already has the domain, the domain is blocked, or its ownership cannot be verified.
An already connected domain is never replaced.

On success the sender address of draft and scheduled campaigns, and of published automation
and double opt-in emails, moves onto the new domain in the background, keeping the part
before the `@`.

```
hostinger reach profiles connect-sending-domain <profile-uuid> [flags]
```

### Examples

```
  hostinger reach profiles connect-sending-domain 550e8400-e09b-41d4-a716-400055000000 --domain example.com
```

### Options

```
      --domain string   
  -h, --help            help for connect-sending-domain
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

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

