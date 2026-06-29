## hostinger hosting nodejs create-build-from-archive

Create NodeJS build from archive

### Synopsis

Upload a project archive, auto-detect build settings, and immediately start a Node.js build.

This is the recommended single-step approach for deploying a Node.js application.
The archive is uploaded to the website's file storage, build settings are auto-detected
from the package.json inside the archive, and the build process starts automatically.
Optional override fields take precedence over auto-detected values.
Maximum archive size is 50MB.

Before archiving, exclude `node_modules/` and any build output directories
(e.g. `dist/`, `.next/`, `build/`) — they are not needed because the build
process runs the install step automatically, and including them unnecessarily
increases the archive size. This also helps keep the archive well under the 50MB limit.

Example (zip):
```
zip -r archive.zip . --exclude "node_modules/*" --exclude "dist/*"
```

The returned build `uuid` can be used to poll progress and retrieve logs via
the `Get Node.js Build Logs` endpoint.

```
hostinger hosting nodejs create-build-from-archive <username> <domain> [flags]
```

### Options

```
      --app-type string           Node.js application type override (one of: create-react-app, vite, angular, react, vue, parcel, express, fastify, nest)
      --archive string            Project archive file (.zip, .tar.gz, or .tgz), maximum 50MB
      --build-script string       Build script override
      --entry-file string         Main entry point file override
  -h, --help                      help for create-build-from-archive
      --node-version int          Node.js version override (auto-detected from package.json if omitted) (one of: 18, 20, 22, 24)
      --output-directory string   Build output directory override relative to the root directory
      --package-manager string    Package manager override (one of: npm, yarn, pnpm)
      --root-directory string     Application root directory override (where package.json is located) relative to public_html
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

