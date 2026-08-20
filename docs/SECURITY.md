# LinkResan public security guidance

This document defines the public disclosure boundary for the LinkResan showcase repository. It is not a production security runbook.

## Never publish

The public synchronization process must reject or stop on material containing:

- real API keys, access tokens, session tokens or private keys;
- database credentials, database URLs or dumps;
- payment gateway credentials, reconciliation records or private payment evidence;
- customer/user data or private support data;
- admin/CRM implementation details or private operational fields;
- deployment credentials, environment values or non-public service identifiers;
- private audit evidence or internal operational topology.

## Documentation examples

Use placeholders only, for example:

```text
<YOUR_API_KEY>
<YOUR_TOKEN>
<PUBLIC_EXAMPLE_ID>
```

Never replace placeholders with live values in public files, issues, pull requests or screenshots.

## Public claims

Security statements in public material must describe only behavior that has been verified in Production. Avoid absolute claims such as "unbreakable", "fully secure" or guarantees that cannot be demonstrated.

## Vulnerability reporting

Do not place sensitive vulnerability details, credentials, customer data or exploit evidence in a public GitHub issue. Use the contact/support path published by LinkResan instead.

## Repository boundary

`AmirMotefaker/LinkResan-Production` remains the private canonical Production source. `AmirMotefaker/LinkResan` is a sanitized public showcase and developer record.
