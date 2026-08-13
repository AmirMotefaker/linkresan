# Zarinpal Production Readiness

Tracking Issue: #3

This document records the public-safe production-readiness and production-closure status for LinkResan billing. Proprietary implementation details, credentials, environment values, private operational evidence, and deployment data are intentionally excluded.

## Production status

The production-grade Zarinpal billing and reconciliation initiative is now complete in the private Production repository and has passed production acceptance.

Validated behavior includes:

- server-authoritative plan, billing period, price, and gateway amount;
- rejection of free-plan checkout;
- idempotent payment request reuse;
- ambiguous gateway request handling through reconciliation rather than automatic duplicate requests;
- non-authoritative browser callback handling;
- verification replay safety, including already-verified recovery;
- subscription activation and renewal only after verified payment;
- Billing V2 payment history;
- read-only visibility for legacy billing records;
- reconciliation dry-run and apply modes;
- exact Authority and stored-amount matching before reconciliation verification;
- fail-closed amount mismatch behavior;
- no mutation for gateway-only unmatched entries;
- durable manual-review state for unresolved local candidates;
- typed Zarinpal request, verify, and unverified-payment contracts;
- admin CRM visibility for successful historical payments and payment history.

## Browser and API security

Cross-origin checkout requires `Idempotency-Key` to be accepted by the API CORS policy. The validated frontend sends this header without sending a client-controlled payment amount.

Billing observability uses structured events with an explicit safe-field allowlist. Authority values, Merchant ID, tokens, credentials, and other secrets must not be emitted to billing logs or public evidence.

## Verification gates completed

The initiative completed the following gates:

- targeted billing and reconciliation behavior tests;
- complete backend test suite;
- Go vet;
- backend build;
- frontend lint;
- frontend TypeScript checks;
- frontend production build;
- disposable PostgreSQL migration rehearsal for the billing migration;
- production database migration and deployment under private operational controls;
- production payment smoke validation;
- production monetization closure audit;
- production CRM billing visibility acceptance;
- exact-SHA tags and GitHub Releases for private Production milestones.

## Public-safe production lineage

The public record intentionally references only milestone identifiers and sanitized outcomes.

- Private Production monetization closure: Issue #19 / PR #22 in `AmirMotefaker/LinkResan-Production`.
- Private Production CRM billing visibility: Issue #23 / PR #24.
- Private legacy billing visibility fix: PR #26.
- Private production CRM acceptance evidence: Issue #25 / PR #27.
- Final private production acceptance release tag: `issue25-production-crm-acceptance-v1-2026-08-13`.

The accepted production CRM result confirms one historical Basic payment of 490,000 IRR is visible exactly once with successful status, gateway identity, durable payment timing, latest successful payment, and payment history while preserving an explicit no-payment state for users without successful billing.

## Public/private boundary

This public repository is a documentation and community record. Proprietary Production source, infrastructure details, account identifiers, credentials, Merchant ID, database URL, access/session tokens, Authority, RefID, gateway payloads, private evidence bundles, and environment-specific operational values remain private.

Future public monetization changes continue to use Issue → Branch → Commit → Pull Request → Review → Merge, followed by an explicit tag/Release gate for completed milestones.
