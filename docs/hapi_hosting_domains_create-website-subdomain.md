## hapi hosting domains create-website-subdomain

Create website subdomain

### Synopsis

Create a new subdomain for the selected website.

Provide a subdomain prefix and, optionally, a custom directory or the
website public directory to use as the subdomain root.

```
hapi hosting domains create-website-subdomain <username> <domain> [flags]
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
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi hosting domains](hapi_hosting_domains.md)	 - Domains commands

