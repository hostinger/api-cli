## hostinger agency-hosting php update-website-version

Update website PHP version

### Synopsis

Switches an Agency Plan website to a different PHP version. Call the available versions endpoint first to see which versions can be selected. The website restarts on the new version, so requests served during the switch may fail and code that is incompatible with the target version will break.

```
hostinger agency-hosting php update-website-version <website_uid> [flags]
```

### Options

```
  -h, --help             help for update-website-version
      --version string   PHP version to switch the website to, as major.minor. Must be one of the versions returned by the available versions endpoint.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting php](hostinger_agency-hosting_php.md)	 - PHP commands

