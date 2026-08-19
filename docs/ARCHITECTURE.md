# LinkResan Public Architecture

This document describes LinkResan at product level only.

It intentionally excludes:

- deployment topology
- environment values
- credentials
- database dumps
- customer data
- admin/CRM internals
- private operational evidence

## High-level flow

```mermaid
flowchart LR
 User --> Web[LinkResan Web]
 Developer --> API[Public API]
 Web --> API
 API --> Links[Link Management]
 API --> Analytics[Analytics]
 API --> Product[Domains / Bio / Teams]
 API --> Billing[Entitlements]
 API --> Data[(Application Data)]
```

## Repository boundary

`AmirMotefaker/LinkResan-Production` is the canonical production source.

`AmirMotefaker/LinkResan` is a public-safe product showcase and developer resource.

Production implementation is not mirrored automatically.
