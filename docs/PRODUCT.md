# LinkResan product overview

LinkResan is a Persian-first SaaS platform for professional link creation, management, analytics and distribution. This document describes the public-safe product surface only; it does not document private commercial implementation details.

## Product promise

LinkResan helps users turn links into manageable product assets: create them quickly, control how they behave, understand usage, connect them to a brand or campaign, and integrate them into developer workflows.

## Primary audiences

### Individuals and creators
- Short links and custom aliases
- Expiration and click limits
- Password-protected links
- Link-in-bio
- Persian-first RTL product experience

### Growth and marketing teams
- Core click analytics
- Device / OS and weekly analytics in the current partial state
- Campaign / UTM tooling in the current partial state
- Custom domains
- Team management

### Developers and product teams
- API key management
- Webhook management
- Publicly documented integration entry points
- Server-authoritative plan and entitlement behavior

## Capability map

### Shipped
- Core link shortening and management
- Custom aliases
- Expiration and click limits
- Password-protected links
- Bulk CSV creation
- AI-assisted naming
- Core click analytics
- Custom domains
- Link-in-bio
- API keys
- Webhooks
- Team management
- Persian / RTL experience
- Rial billing

### Partial / preview
- Extended device and operating-system analytics
- Weekly analytics views
- Analytics retention behavior by plan
- Campaign / UTM tooling
- Native mobile client
- Selected public/open-source materials

### Not shipped
- Cryptocurrency billing

`Partial` and `Preview` are deliberate truth states. Public material must not present them as fully shipped.

## Product tiers

LinkResan exposes four tiers: Free, Basic, Pro and Enterprise. Public price is owned by the product pricing page and may change independently of this repository.

Current entitlement examples:

| Tier | Link allowance | API keys | Webhooks | Team size |
|---|---:|---:|---:|---:|
| Free | 50/month | 1 | 1 | 2 |
| Basic | Unlimited | 2 | 2 | 5 |
| Pro | — | 5 | 5 | 10 |
| Enterprise | — | 20 | 20 | 25 |

Enterprise currently includes entitlement for up to five custom domains.

## Product principles

1. **Persian-first, not translated-later.** RTL and Persian usability are product-level concerns.
2. **Server-authoritative enforcement.** Plans and entitlements are not treated as client-only presentation rules.
3. **Truthful capability status.** Partial and preview states remain visible as such.
4. **Developer access without private implementation disclosure.** Public API and integration concepts can be documented without exposing commercial source or operational topology.
5. **Public-safe by design.** Public GitHub material is curated from the private canonical Production source through an allowlisted review flow.

## What this repository does not promise

- It is not a deployable mirror of Production.
- It does not expose admin/CRM implementation details.
- It does not publish billing reconciliation evidence or gateway credentials.
- It does not publish private deployment identifiers or database configuration.
- It does not imply future roadmap dates or availability commitments.

## Canonical source

`AmirMotefaker/LinkResan-Production` remains the sole canonical Production repository. `AmirMotefaker/LinkResan` is a sanitized public product and developer showcase.
