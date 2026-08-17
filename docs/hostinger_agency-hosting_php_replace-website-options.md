## hostinger agency-hosting php replace-website-options

Replace website PHP options

### Synopsis

Replaces the custom php.ini values on an Agency Plan website with the ones provided. Any option not in the request is reset to its default, so call the options endpoint first and send the full desired set. Sending an empty array resets every option to its default.

```
hostinger agency-hosting php replace-website-options <website_uid> [flags]
```

### Options

```
  -h, --help             help for replace-website-options
      --options string   Option names and values. Each name must be one of the options returned by the options endpoint, and each value must satisfy that option's allowed_values when it declares them. (JSON)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting php](hostinger_agency-hosting_php.md)	 - PHP commands

