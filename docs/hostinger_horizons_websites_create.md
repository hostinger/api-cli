## hostinger horizons websites create

Create website

### Synopsis

Create new Hostinger Horizons website from the given message.\n
Use this tool when user asks you to create a website, landing page, blog
or any other type of application.\n
This tool initiates the website creation process and returns a website URL and ID.
The generation happens asynchronously.\n
After invoking this tool, your chat reply must be EXACTLY 1 sentence summarizing
that Hostinger Horizons is now creating their website and it will be ready in a few minutes
and you should provide the website URL to the user immediately
Do not write code.\n\nTo edit afterwards, use the `Edit website` tool with the returned
website ID, or the user can go to Hostinger Horizons interface in the provided website URL.
If the tool call fails with an error, you should provide a clear explanation of the error
and do not generate code yourself in the chat.
\n
TECHNOLOGY STACK CONSTRAINTS (STRICTLY ENFORCED):\n
The environment is limited to the following technologies.
You MUST NOT use, suggest, or implement any technology outside this list:\n
\n
- Language: JavaScript ONLY.
- Languages like TypeScript, Rust, Python, Java, PHP, etc., are STRICTLY PROHIBITED.\n
- Framework: React.\n
- Navigation: React Router.\n
- Styling: TailwindCSS.\n
- Components: shadcn/ui (built with @radix-ui primitives).\n
- Icons: Lucide React.\n
- Animations: Framer Motion.\n
\n
BACKEND & DATA STORAGE:\n
- Horizons integrated backend is the EXCLUSIVE solution for persistent data storage,
authentication, and database needs.\n
- Local databases (SQLite, MySQL, etc.) are STRICTLY PROHIBITED.\n
- Third-party services (Firebase, AWS Amplify) are allowed ONLY if explicitly requested by the user.\n
\n
MAPS:\n
- OpenStreetMap is the default provider.\n
- Alternative providers (Google Maps, Mapbox) are allowed ONLY if explicitly requested by the user.\n

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --message
    [
      {
        "text": "Create a landing page for a coffee shop with a hero section, menu, and contact form",
        "type": "text"
      }
    ]

```
hostinger horizons websites create [flags]
```

### Examples

```
  hostinger horizons websites create --message @message.json
```

### Options

```
  -h, --help             help for create
      --message string   (JSON, @file or -; shape in the help above)
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

* [hostinger horizons websites](hostinger_horizons_websites.md)	 - Websites commands

