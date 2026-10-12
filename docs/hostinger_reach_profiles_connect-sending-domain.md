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

### Options

```
      --domain string   
  -h, --help            help for connect-sending-domain
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

