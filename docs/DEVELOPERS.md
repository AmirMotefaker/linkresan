# LinkResan for developers

This page is the public-safe developer entry point for LinkResan. The commercial Production implementation remains private.

## Public developer surface

Current public-facing concepts:

- API keys
- Webhooks
- Integration-oriented documentation
- Product-level API concepts

## Credential safety

Never publish:

- API tokens
- session cookies
- private keys
- environment files
- customer identifiers
- production endpoints or internal service URLs

## Production boundary

The public repository is not a deployable SDK mirror or backend source tree.

For production implementation, operational access, and commercial development workflows, the private canonical repository is:

`AmirMotefaker/LinkResan-Production`
