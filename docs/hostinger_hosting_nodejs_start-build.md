## hostinger hosting nodejs start-build

Start Node.js build

### Synopsis

Start a Node.js build process using files already present on the website's file storage.

WARNING: on success this overwrites the website's existing contents and cannot be
undone — verify this is intended before calling this endpoint.

The `source_type` must be `archive` and `source_options.archive_path` must point to an
existing archive file on the server (relative to the website document root).
Use the `Generate Upload URL` endpoint to obtain credentials and upload the archive first.

To auto-detect build settings from an archive before starting, first call the
`Get Node.js Build Settings from Archive` endpoint.

The returned build `uuid` can be used to poll progress and retrieve logs via
the `Get Node.js Build Logs` endpoint.

```
hostinger hosting nodejs start-build <username> <domain> [flags]
```

### Options

```
      --app-type string           Node.js application type (one of: create-react-app, gatsby, vite, angular, react, vue, parcel, next, nuxt, nest, express, fastify, astro, svelte, svelte-kit, hono, react-router, nitro, other)
      --build-script string       Build script that will be ran to build the application
      --entry-file string         The main entry point file for the application
  -h, --help                      help for start-build
      --node-version int          Node.js version (one of: 18, 20, 22, 24)
      --output-directory string   Build output directory relative to the root directory
      --package-manager string    Package manager (one of: npm, yarn, pnpm)
      --root-directory string     Application root directory (where package.json is located) relative to public_html
      --source-options string     Source-specific options (JSON)
      --source-type string        The source type of the files (one of: archive)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

