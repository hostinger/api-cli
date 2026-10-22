## hostinger agency-hosting files generate-upload-url

Generate upload URL

### Synopsis

Generate a file browser upload URL with authentication credentials for uploading files
to an Agency Plan website's file storage.

Returns `url`, `auth_key` and `rest_auth_key`. Use these to upload a file to the
website's file storage via the TUS resumable upload protocol (TUS 1.0.0). Send
`X-Auth: {auth_key}` and `X-Auth-Rest: {rest_auth_key}` headers on every request below.

1. Create the upload: `POST` to `{url}/{relative_file_path}?override=true` with headers
   `upload-length: {file size in bytes}` and `upload-offset: 0`. Expect `201 Created`.
2. Upload the file: send the file bytes to the same location (any TUS 1.0.0 client, or
   `PATCH` requests with an `upload-offset` header tracking progress) until complete.

`relative_file_path` is the destination path inside the website's file storage, e.g.
`app.zip`.

Instead of a TUS client, plain `curl` also works:
```
FILE=app.zip
SIZE=$(stat -f%z "$FILE")   # stat -c%s on Linux

curl -i -X POST "{url}/${FILE}?override=true" \
  -H "X-Auth: {auth_key}" \
  -H "X-Auth-Rest: {rest_auth_key}" \
  -H "Tus-Resumable: 1.0.0" \
  -H "Upload-Length: ${SIZE}" \
  -H "Upload-Offset: 0"
-> 201 Created

curl -i -X PATCH "{url}/${FILE}?override=true" \
  -H "X-Auth: {auth_key}" \
  -H "X-Auth-Rest: {rest_auth_key}" \
  -H "Tus-Resumable: 1.0.0" \
  -H "Content-Type: application/offset+octet-stream" \
  -H "Upload-Offset: 0" \
  --data-binary "@${FILE}"
-> 204 No Content, Upload-Offset response header equals SIZE when done
```

```
hostinger agency-hosting files generate-upload-url <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting files generate-upload-url zpwlGlp19
```

### Options

```
  -h, --help   help for generate-upload-url
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

* [hostinger agency-hosting files](hostinger_agency-hosting_files.md)	 - Files commands

