## hostinger agency-hosting ssl website-status

Get website SSL status

### Synopsis

Returns the SSL state of one domain of an Agency Plan website: the certificate `status`,
whether the certificate was uploaded by the customer, and when it stops being valid.

`installing` means a certificate setup is running or retrying; the `ssl_setup` entry of
`List website processes` shows the same progress. `active` means a valid certificate is in
place: uploaded by the customer, issued by the platform, or a lifetime certificate bought for
the domain. `failed` means the last setup gave up and no valid certificate is in place.
`expired` means the certificate has run out. `not_installed` means the domain has no
certificate and no setup process. Returns 404 when the website or the domain does not exist.

```
hostinger agency-hosting ssl website-status <website_uid> <domain> [flags]
```

### Options

```
  -h, --help   help for website-status
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting ssl](hostinger_agency-hosting_ssl.md)	 - SSL commands

