## hapi hosting wordpress install

Install WordPress

### Synopsis

Install WordPress on an existing website.

The website must already exist before calling this endpoint. To create a new
website first, use POST /api/hosting/v1/websites and poll
GET /api/hosting/v1/websites until it appears.

Call GET /api/hosting/v1/wordpress/installations filtered by username and
domain before proceeding to check whether WordPress is already installed on
the target domain/path. If WordPress already exists and `overwrite` is false
(the default), the async job will fail.

This operation is asynchronous: a successful response only means the install
job has been queued, not that WordPress is ready. Installation typically
takes 1-2 minutes. Poll GET /api/hosting/v1/wordpress/installations filtered
by username and domain to track progress. When the installation appears in
that list, WordPress is ready.

```
hapi hosting wordpress install <username> [flags]
```

### Options

```
      --auto-updates string   WordPress core auto-update policy (one of: all, none, minor)
      --credentials string    WordPress admin credentials (JSON)
      --database string       Optional. If the named database already exists, it will be used for this WordPress install. Otherwise a new database is created with a generated name and random credentials. (JSON)
      --directory string      Relative directory to install WordPress into. Defaults to the website root when omitted.
      --domain string         Domain of the existing website where WordPress will be installed
  -h, --help                  help for install
      --language string       WordPress locale. Defaults to en_US when omitted.
      --overwrite             When false (default), does not replace an existing installation. If WordPress is already installed on the domain/path, the async install job fails unless true.
      --site-title string     Title of the WordPress site
      --version string        WordPress core version to install. If omitted, the latest core version compatible with the account vhost PHP version is selected.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi hosting wordpress](hapi_hosting_wordpress.md)	 - Wordpress commands

