## hostinger agency-hosting files import-website-from-archive

Import website from archive

### Synopsis

Imports an Agency Plan website from an already-uploaded archive.

Upload the archive to the website's root directory via file browser first, then provide its
filename in this request. Website contents are overwritten by the archive contents. Supported
archive types: .zip, .tar, .tar.gz, .tgz.

```
hostinger agency-hosting files import-website-from-archive <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting files import-website-from-archive zpwlGlp19 --archive-name archive.zip
```

### Options

```
      --archive-name string   Archive filename (e.g., archive.zip). The file must already be uploaded to the website's .h5g/ directory.
  -h, --help                  help for import-website-from-archive
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

* [hostinger agency-hosting files](hostinger_agency-hosting_files.md)	 - Files commands

