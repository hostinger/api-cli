## hostinger agency-hosting files import-plan-website-from-archive

Import Agency Plan website from archive

### Synopsis

Imports an Agency Plan website from an already-uploaded archive.

Upload the archive to the website's root directory via file browser first, then provide its
filename in this request. Website contents are overwritten by the archive contents. Supported
archive types: .zip, .tar, .tar.gz, .tgz.

```
hostinger agency-hosting files import-plan-website-from-archive <website_uid> [flags]
```

### Options

```
      --archive-name string   Archive filename (e.g., archive.zip). The file must already be uploaded to the website's .h5g/ directory.
  -h, --help                  help for import-plan-website-from-archive
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting files](hostinger_agency-hosting_files.md)	 - Files commands

