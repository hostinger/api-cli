## hostinger hosting nodejs patch-vulnerabilities

Patch Node.js vulnerabilities

### Synopsis

Patches the selected Node.js vulnerabilities by updating the affected package versions
in `package.json` and opening a GitHub pull request in the connected repository. The
customer reviews and merges the pull request; merging triggers the automatic deployment.

Auto-fix is only available for websites deployed from a connected GitHub repository.
Websites deployed from an archive have no auto-fix path and return a 404. The Hostinger
GitHub App needs write access to the repository; without it the request fails with a
403 explaining the missing permission.

Only vulnerabilities with `is_patchable` set to `true` can be patched. Non-patchable
IDs in the selection are skipped; the pull request covers the patchable subset, listed
in `patched_vulnerability_ids`. Selections without any patchable vulnerability are
rejected with a 422. Only one patch pull request can be open at a time per website;
close or merge it before patching again. Available on Business and Cloud Hosting plans.

```
hostinger hosting nodejs patch-vulnerabilities <username> <domain> [flags]
```

### Examples

```
  hostinger hosting nodejs patch-vulnerabilities u123456789 mydomain.tld --vulnerability-ids value
```

### Options

```
  -h, --help                        help for patch-vulnerabilities
      --vulnerability-ids strings   List of vulnerability IDs to patch, as returned by the list vulnerabilities endpoint.
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

