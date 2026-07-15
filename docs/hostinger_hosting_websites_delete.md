## hostinger hosting websites delete

Delete website

### Synopsis

Permanently deletes a website and all of its data. This action is destructive
and cannot be undone. Always ask the user for explicit confirmation before
calling this endpoint.

All website files, databases and related configuration will be removed.
The hosting plan itself is kept, so a new website can be created on it afterwards.

The confirm field must be boolean true, otherwise the request is rejected.

Supported websites: main and addon domain websites on web hosting plans, and
Website Builder websites. Parked domains and subdomains cannot be deleted with
this endpoint. The domain must be the exact website domain, not a preview
domain or an alias.

Returns 404 when the domain does not exist or does not belong to the
authenticated client.

Website removal is processed asynchronously and can take a few minutes to
complete. The response returns before the removal finishes.

```
hostinger hosting websites delete <domain> [flags]
```

### Options

```
      --confirm   Must be boolean true to confirm the permanent deletion of the website. (one of: true)
  -h, --help      help for delete
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

