# LinkResan for developers

This is the public-safe developer entry point for LinkResan. The commercial Production implementation remains private.

## Public developer surface

- Product website: https://linkresan.ir
- API keys are created and managed inside the authenticated LinkResan dashboard.
- Webhooks are a shipped integration capability.
- Public API/reference material is published only through intentionally public LinkResan surfaces.

## Supported developer-facing capabilities

- Link creation and management
- API keys
- Webhooks
- Analytics-facing product endpoints
- Publicly documented product integrations

## Credential hygiene

Never publish a real API key, token, session credential, webhook secret, production identifier or customer endpoint in GitHub files, issues, pull requests or screenshots.

Documentation examples must use placeholders only, for example:

```http
Authorization: Bearer <YOUR_TOKEN>
X-API-Key: <YOUR_API_KEY>
```

## Public example policy

Examples synchronized to the public repository must:

1. use placeholders only;
2. contain no customer/user data;
3. contain no admin/CRM implementation details;
4. contain no billing reconciliation evidence;
5. contain no deployment/service identifiers that are not intentionally public;
6. be explicitly allowlisted by the Production public-sync contract.

## Source of truth

Feature claims and entitlement behavior originate from the private `AmirMotefaker/LinkResan-Production` repository. The public `AmirMotefaker/LinkResan` repository is a curated developer showcase, not a deployable Production mirror.
