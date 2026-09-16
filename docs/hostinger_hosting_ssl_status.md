## hostinger hosting ssl status

Get SSL status

### Synopsis

Returns the SSL state of the website: the certificate `status` and `provider`, whether the
certificate is a lifetime one managed by the platform, whether HTTP requests are redirected to
HTTPS, when the certificate stops being valid and the last installation error.

`installing` and `waiting_for_retry` mean an installation is in progress. `failed` means the
last installation gave up, or the website was not updated for 60 minutes while `installing`;
`last_error` holds the reason when it is a known message, otherwise it is null. `expired`
means the assigned certificate's validity has ended. `not_installed` means no certificate is
assigned. Free subdomains use a platform-managed certificate: with no installation recorded
they report `active` with `provider` and `expires_at` null.

```
hostinger hosting ssl status <username> <domain> [flags]
```

### Options

```
  -h, --help   help for status
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting ssl](hostinger_hosting_ssl.md)	 - SSL commands

