# Zarinpal Billing & Reconciliation

Tracking Issue: #3

## Objective

LinkResan already enforces server-side plans, entitlements, and quotas. This initiative connects verified payments to subscriptions in a durable, auditable, and idempotent way so monetization can operate safely in production.

## Payment lifecycle

1. Create an internal payment record before redirecting the user to Zarinpal.
2. Snapshot the server-selected plan, billing period, amount, and currency on that record.
3. Treat browser callback parameters as untrusted input.
4. Perform verification server-to-server with Zarinpal.
5. Activate or renew a subscription only after successful verification.
6. Make callback and verification handling idempotent so replay cannot duplicate subscription changes.
7. Persist gateway reference data and status transitions for audit and support.

## Internal payment states

- `created`
- `redirected`
- `verification_pending`
- `paid`
- `failed`
- `cancelled`
- `reconciliation_required`

Transitions must be explicit and invalid transitions must fail closed.

## Billing invariants

- Price, plan, user, and billing period are server-authoritative.
- A client-supplied amount must never determine the charged or activated plan.
- Merchant credentials remain server-side and are never logged or published.
- Authority and successful RefID values must be unique at the persistence layer.
- Duplicate callback or verify requests must be safe no-ops after a payment is finalized.
- Subscription activation must be atomic with the finalized payment transition.
- Active-subscription renewal must preserve remaining paid time where applicable.

## Reconciliation

Reconciliation must detect and classify at least:

- verified gateway payment with no local subscription activation;
- locally finalized payment with inconsistent subscription state;
- stale verification-pending payments;
- duplicate or conflicting gateway reference data.

A reconciliation run should report:

- scanned;
- matched;
- repaired;
- unresolved.

Repairs must be idempotent and must not guess when gateway evidence is insufficient.

## Observability

Each payment lifecycle should carry a correlation identifier. Structured logs must contain operational status and failure reason without credentials, Merchant ID, access tokens, or other secrets.

## Verification matrix

- request success;
- duplicate request/idempotency;
- invalid callback status;
- missing or invalid authority;
- successful verification;
- verification replay;
- amount/plan mismatch fail-closed;
- failed verification;
- subscription activation;
- subscription renewal;
- concurrent verification race;
- persistence uniqueness and atomicity;
- reconciliation no-op;
- reconciliation repair of verified-but-not-activated payment.

## Public/private synchronization rule

This public repository is the continuously maintained public product record for LinkResan. Public-safe code and documentation should be updated for each milestone through Issue → Branch → Commit → Pull Request → Review → Merge.

Credentials, secrets, private operational evidence, and environment-specific values must never be copied into the public repository.

Git tags and GitHub Releases are handled by a separate release gate.