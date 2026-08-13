# Zarinpal Reconciliation Persistence Taxonomy

Tracking Issue: #3

LinkResan keeps two levels of reconciliation classification:

1. detailed runtime outcomes used by the reconciliation engine and operator-facing action report;
2. a stable canonical persistence taxonomy stored in the reconciliation audit table.

The canonical database taxonomy remains intentionally small so schema constraints and long-lived audit records do not depend on every implementation-level outcome name.

## Canonical persistent classifications

- `gateway_paid_local_not_activated`
- `local_paid_subscription_inconsistent`
- `stale_verification_pending`
- `duplicate_gateway_reference`
- `gateway_evidence_insufficient`

## Canonical persistent statuses

- `matched`
- `repairable`
- `repaired`
- `unresolved`
- `ignored`

Detailed action classification and status are retained in structured reconciliation evidence rather than being written directly into constrained database columns.

The production-readiness gate includes a disposable PostgreSQL integration test that exercises both repaired and unresolved apply-mode persistence through the real service and repository path.

No credentials, Merchant ID, database URL, gateway authority, access token, session token, or private evidence belongs in this public document.
