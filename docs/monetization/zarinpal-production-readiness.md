# Zarinpal Production Readiness

Tracking Issue: #3

This document records the public-safe production-readiness contract for LinkResan billing. Proprietary implementation details, credentials, environment values, private operational evidence, and deployment data are intentionally excluded.

## Readiness status

The local billing implementation has completed its exact-commit Product Test gate with a clean worktree and unchanged tested commit.

Validated behavior includes:

- server-authoritative plan, billing period, price, and gateway amount;
- rejection of free-plan checkout;
- idempotent payment request reuse;
- ambiguous gateway request handling through reconciliation rather than automatic duplicate requests;
- non-authoritative browser callback handling;
- verification replay safety, including already-verified recovery;
- subscription activation and renewal only after verified payment;
- Billing V2 payment history;
- reconciliation dry-run and apply modes;
- exact Authority and stored-amount matching before reconciliation verification;
- fail-closed amount mismatch behavior;
- no mutation for gateway-only unmatched entries;
- durable manual-review state for unresolved local candidates;
- typed Zarinpal request, verify, and unverified-payment contracts.

## Browser and API security

Cross-origin checkout requires `Idempotency-Key` to be accepted by the API CORS policy. The readiness gate verifies this contract together with the frontend requirement to send the header while never sending a client-controlled payment amount.

Billing observability uses structured events with an explicit safe-field allowlist. Authority values, Merchant ID, tokens, credentials, and other secrets must not be emitted to billing logs.

## Verification gates

The Product Test gate covers:

- targeted billing and reconciliation behavior tests;
- complete backend test suite;
- Go vet;
- backend build;
- frontend lint;
- frontend production build;
- disposable PostgreSQL migration rehearsal for the billing migration;
- readiness security contracts;
- final clean-worktree and unchanged-commit verification.

## Deployment boundary

Passing the Product Test gate does not itself authorize Production deployment.

The following remain separate explicit gates:

1. synchronize reviewed public-safe documentation;
2. synchronize proprietary source to the private repository when authorized;
3. apply the reviewed Production database migration through the migration runbook;
4. deploy the exact reviewed application commit;
5. run Production smoke tests;
6. publish tag and GitHub Release only through the release gate.

No credential, Merchant ID, database URL, access token, session token, private evidence bundle, or environment-specific operational value belongs in the public repository.
