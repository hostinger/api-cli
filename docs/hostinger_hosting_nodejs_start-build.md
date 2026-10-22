## hostinger hosting nodejs start-build

Start Node.js build

### Synopsis

Start a Node.js build process using files already present on the website's file storage.

WARNING: on success this overwrites the website's existing contents and cannot be
undone — verify this is intended before calling this endpoint.

With `source_type` `archive`, `source_options.archive_path` must point to an existing
archive file on the server (relative to the website document root). Use the
`Generate Upload URL` endpoint to obtain credentials and upload the archive first. To
auto-detect build settings from an archive before starting, first call the
`Get Node.js Build Settings from Archive` endpoint.

With `source_type` `git`, `source_options` carries `owner`, `repository`, `branch` and
`installation_uuid`. Take the installation from `List Git installations` and the owner and
repository from `List Git installation repositories`; the branch is cloned at its current
head. The installation must belong to the same customer as the website.

The returned build `uuid` can be used to poll progress and retrieve logs via
the `Get Node.js Build Logs` endpoint.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --source-options
    {
      "archive_path": "example.zip",
      "branch": "main",
      "installation_uuid": "018f5e2a-1234-7890-abcd-1234567890ab",
      "owner": "octocat",
      "repository": "my-site"
    }

```
hostinger hosting nodejs start-build <username> <domain> [flags]
```

### Examples

```
  hostinger hosting nodejs start-build u123456789 mydomain.tld --app-type vite --build-script build --node-version 20 --output-directory dist --root-directory webapp --source-options @source-options.json --source-type archive
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
      --source-options archive    Source-specific options. For archive send `archive_path`. For `git` send `owner`, `repository`, `branch` and `installation_uuid`, taken from `List Git installations` and `List Git installation repositories`. (JSON, @file or -; shape in the help above)
      --source-type archive       Where the files come from: archive (an uploaded archive on the website) or `git` (a branch of a repository reachable through a Git installation). (one of: archive, git)
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

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

