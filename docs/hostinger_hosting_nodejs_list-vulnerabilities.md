## hostinger hosting nodejs list-vulnerabilities

List Node.js vulnerabilities

### Synopsis

Lists known npm package vulnerabilities detected on a Node.js website, enriched with
advisory metadata (severity, CVSS score, CVE, advisory URL). Results are sorted from
the most severe to the least severe, then by publish date (newest first). Use the
`severities` query parameter to filter.

Vulnerabilities with `is_patchable` set to `true` can be auto-fixed via the
`Patch Node.js Vulnerabilities` endpoint, which opens a GitHub pull request with
updated package versions. Auto-fix is only available for websites deployed from a
connected GitHub repository. Vulnerabilities with `is_patching_in_progress` set to
`true` are already included in an open patch pull request; while any patch pull
request is open, new patch requests for this website are rejected until it is merged
or closed.

Data comes from periodic dependency scans, so it may lag behind the latest deployment.
An empty list means the most recent scan found no vulnerabilities; it does not
guarantee the current deployment is vulnerability-free. Available on Business and
Cloud Hosting plans.

```
hostinger hosting nodejs list-vulnerabilities <username> <domain> [flags]
```

### Examples

```
  hostinger hosting nodejs list-vulnerabilities u123456789 mydomain.tld
```

### Options

```
  -h, --help                 help for list-vulnerabilities
      --severities strings   Severities to filter by (one of: low, moderate, high, critical, unknown)
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

