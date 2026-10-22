## hostinger hosting domains create-website-subdomain

Create website subdomain

### Synopsis

Create a new subdomain for the selected website.

Provide a subdomain prefix and, optionally, a custom directory or the
website public directory to use as the subdomain root.

```
hostinger hosting domains create-website-subdomain <username> <domain> [flags]
```

### Examples

```
  hostinger hosting domains create-website-subdomain u123456789 mydomain.tld --subdomain blog
```

### Options

```
      --directory string            Directory name for the subdomain relative to the website root
  -h, --help                        help for create-website-subdomain
      --is-using-public-directory   Use the website public directory as the subdomain root directory
      --subdomain string            Subdomain prefix to create under the selected website
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

* [hostinger hosting domains](hostinger_hosting_domains.md)	 - Domains commands

