## hostinger wordpress installations install

Install WordPress

### Synopsis

Install WordPress on an existing website.

The website must already exist before calling this endpoint. To create a new
website first, use POST /api/hosting/v1/orders/{order_id}/onboardings
(`hosting_websites_start-setup`) and poll GET /api/hosting/v1/onboardings
(`hosting_websites_list-setups`) with the `domain` filter until its `status` is
`completed`.

Send `type: headless` to turn a headless website (for example one created by
`hosting_websites_start-setup` with `type: headless_empty`) into a headless WordPress
website: WordPress is installed into the `cms` directory with the Hostinger headless
MU plugin and the website is marked as a headless WordPress website. `directory` is
not allowed with it. The target website must already be headless: the headless
install rewrites the root `.htaccess`, so on any other website the request is
rejected with 422.

Call GET /api/hosting/v1/wordpress/installations (`wordpress_installations_list`)
filtered by username and domain before proceeding to check whether WordPress is
already installed on the target domain/path. If WordPress already exists and
`overwrite` is false (the default), the async job will fail.

This operation is asynchronous: a successful response only means the install
job has been queued, not that WordPress is ready. Installation typically
takes 1-2 minutes. Poll GET /api/hosting/v1/wordpress/installations
(`wordpress_installations_list`) filtered by username and domain to track
progress. When the installation appears in that list, WordPress is ready.

Returns 422 when the account already uses all the databases its plan allows,
unless `database.name` is an existing database on the account.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --credentials
    {
      "email": "owner@example.com",
      "login": "admin",
      "password": "********"
    }

  --database
    {
      "name": "mydb",
      "password": "********"
    }

```
hostinger wordpress installations install <username> [flags]
```

### Examples

```
  hostinger wordpress installations install u123456789 --credentials @credentials.json --domain example.com --site-title 'My site'
```

### Options

```
      --auto-updates string        WordPress core auto-update policy (one of: all, none, minor)
      --credentials string         WordPress admin credentials (JSON, @file or -; shape in the help above)
      --database string            Optional. If the named database already exists on the account, it is used for this WordPress install. Otherwise a new database is created with this name, or with a generated name when database is omitted or null. A new database gets a random database user and counts toward the plan's database limit. (JSON, @file or -; shape in the help above)
      --directory type: headless   Installation subdirectory relative to the document root; omit for the root. Not allowed with type: headless, which always installs into `cms`.
      --domain string              Domain of the existing website where WordPress will be installed
  -h, --help                       help for install
      --language string            WordPress locale. Defaults to en_US when omitted.
      --overwrite                  When false (default), does not replace an existing installation. If WordPress is already installed on the domain/path, the async install job fails unless true.
      --site-title string          Title of the WordPress site
      --type default               Installation kind; default when omitted. `headless` installs WordPress into the `cms` directory with the Hostinger headless MU plugin and marks the website as a headless WordPress website. Accepted only for a website that is already headless, e.g. one created by `POST /api/hosting/v1/orders/{order_id}/onboardings` (`hosting_websites_start-setup`) with `type: headless_empty`; any other website is rejected with 422. (one of: default, headless)
      --version string             WordPress core version to install. If omitted, the latest core version compatible with the account vhost PHP version is selected.
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

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

