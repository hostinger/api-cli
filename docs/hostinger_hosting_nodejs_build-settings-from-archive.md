## hostinger hosting nodejs build-settings-from-archive

Get Node.js build settings from archive

### Synopsis

Auto-detect Node.js build settings from a package.json inside an archive already on the server.

Use this before calling `Start Node.js Build` to preview what settings will be used,
or to let the user review and override values (framework, node version, root directory,
output directory, build script) before committing to a build.

The archive must already be present on the website's file storage. Use the
`Generate Upload URL` endpoint to obtain credentials and upload the archive first.

```
hostinger hosting nodejs build-settings-from-archive <username> <domain> [flags]
```

### Options

```
      --archive-path string   The path to the archive file relative to the document root of the vhost
  -h, --help                  help for build-settings-from-archive
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

