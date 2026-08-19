# LinkResan for developers

This page is the public-safe developer entry point for LinkResan. The commercial Production implementation remains private.

## Public API

- API documentation: https://linkresan-api.onrender.com/docs
- Product website: https://linkresan.ir
- Authentication for supported developer integrations is managed through the product dashboard.

## Supported developer-facing capabilities

- Link creation and management
- API keys
- Webhooks
- Analytics-facing product endpoints
- Custom-domain, bio and team product surfaces where exposed by the public API contract

## API keys

API keys are created and managed inside the authenticated LinkResan dashboard. Never commit a real API key to GitHub, examples, issue comments or screenshots.

Use placeholders in documentation:

```http
Authorization: Bearer <YOUR_TOKEN>
X-API-Key: <YOUR_API_KEY>
```

## Webhooks

Webhook support is a shipped product capability. Public documentation may describe integration behavior, but private delivery evidence, customer endpoints, credentials and production support data are not published here.

## Examples policy

Examples in this repository must:

1. use placeholders only;
2. avoid production account identifiers;
3. avoid customer/user data;
4. avoid admin/CRM internals;
5. avoid billing reconciliation internals;
6. be reviewable as standalone public material.

## Source of truth

Feature claims and entitlement behavior originate from the private `AmirMotefaker/LinkResan-Production` repository. This public repository is a curated developer showcase, not a deployable Production mirror.
