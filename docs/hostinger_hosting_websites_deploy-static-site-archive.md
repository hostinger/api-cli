## hostinger hosting websites deploy-static-site-archive

Deploy static site archive

### Synopsis

Deploy a static application from an archive file.

WARNING: this overwrites the website's existing contents and cannot be undone —
verify this is intended before calling this endpoint.

This endpoint allows you to deploy a static application from an archive
file that has been uploaded to the website's directory.

This only works for static sites (pre-built HTML/CSS/JS with no build step). For
Node.js applications, use `Create NodeJS build from archive` instead, or
`Start Node.js build` if the archive is already uploaded. For WordPress sites,
use `Import WordPress website`.

```
hostinger hosting websites deploy-static-site-archive <username> <domain> [flags]
```

### Options

```
      --archive-path string   Relative path to the archive file from website root directory
  -h, --help                  help for deploy-static-site-archive
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

