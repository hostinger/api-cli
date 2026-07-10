## hostinger agency-hosting websites build-plan-nodejs-assets

Build Agency Plan website NodeJS assets

### Synopsis

Builds and deploys a Node.js application for an Agency Plan website from an already-uploaded archive.

Upload the archive to file browser first, then provide its relative path from document root in this request.
Website contents are overwritten by the build result, which is deployed to public_html.

```
hostinger agency-hosting websites build-plan-nodejs-assets <website_uid> [flags]
```

### Options

```
      --archive-path public_html   Directory, relative to the website document root, where the uploaded site archive currently lives. Most commonly this is simply public_html.
  -h, --help                       help for build-plan-nodejs-assets
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting websites](hostinger_agency-hosting_websites.md)	 - Websites commands

