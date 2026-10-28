## hostinger up

Deploy a directory

### Synopsis

Deploy a directory (default: the current one) to a Web or Cloud hosting website.

The website comes from --domain, the project file written by an earlier run or by
`hostinger link`, or a prompt. A website that does not exist yet is created on a
hosting plan (--order-id, or a prompt), with a free subdomain when no domain is given.
The directory is deployed as a static site (HTML, CSS, JavaScript, PHP) or a Node.js
app, which is built on the server, from an uploaded copy of the directory or from its
pushed Git branch (--source git, asked once and remembered, which also redeploys on
every push). Deploying overwrites the website's current files.

A headless website (--headless creates one: empty, wordpress with WordPress in /cms,
ecommerce or pocketbase) serves a Node.js frontend, so a folder with a package.json
is deployed there as a Node.js app and its backend is left alone. In a terminal, a new
website is confirmed as a regular one, or as the headless website its package.json
points to: a PocketBase or WordPress client, else an empty one for a web frontend
(Next.js, Nuxt, Astro, Vite and similar); --headless skips that.

The directory is linked to the website in .hostinger/site.json, written in the
current directory with the deployed folder recorded, so `hostinger up` from the
project root redeploys the same folder. --domain on a linked directory deploys there
once and leaves the link alone; `hostinger link` changes it.

Without a terminal (CI, pipes, --format json) nothing is asked: pass --domain for an
existing website, or --create for a new one (--order-id when the account has more than
one hosting plan; the only plan is used by itself), and --yes to replace a website that
already has content. up never deploys to default_domain from the config file: only to the
linked website, --domain or HOSTINGER_DOMAIN.

```
hostinger up [path] [flags]
```

### Examples

```
  hostinger up
  hostinger up ./dist --domain example.com
  hostinger up --headless wordpress
  hostinger up --create --order-id 123 --type nodejs --format json
```

### Options

```
      --app-type string           Node.js framework, e.g. express, next, vite, nest, other
      --build-script string       npm script that builds the app
      --create                    Create a new website instead of deploying to an existing one
      --datacenter string         Where the first website on a plan is hosted, a code from hostinger hosting datacenters list (default: asked, or the API's best match)
      --dry-run                   Show what would be deployed where, as JSON, without creating, uploading or building anything
      --entry-file string         Entry file of the server (compiled apps)
      --headless string           Create a headless website for a Node.js frontend: empty, wordpress (installed in /cms), ecommerce or pocketbase. A new website unless the linked website or --domain is already headless
  -h, --help                      help for up
      --no-wait                   Return as soon as a Node.js build is started instead of waiting for it
      --node-version int          Node.js major version (18, 20, 22, 24)
      --order-id int              Hosting plan (order ID) a new website is created on
      --output-directory string   Build output directory, relative to the root directory
      --package-manager string    npm, yarn or pnpm
      --root-directory string     Directory holding package.json, relative to the archive root
      --settings                  Review the Node.js build settings before building (a first deploy in a terminal reviews them anyway; without a terminal the detected settings and flags are used)
      --source string             Where the files come from: upload (pack this directory) or git (the pushed branch, through a connected Git account); default: remembered, else asked
      --type string               What to deploy as: static or nodejs (default: detected, or remembered)
  -y, --yes                       Accept the detected type, and replace a website's content or deploy static files onto a Node.js site without asking (required for those without a terminal)
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger](hostinger.md)	 - Hostinger API Command Line Interface

