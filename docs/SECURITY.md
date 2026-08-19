# LinkResan public security guidance

This document defines what security information belongs in the public repository.

## Publicly documented

- Secure development practices
- Public API usage guidance
- Responsible disclosure expectations
- Public/private repository boundary

## Never publish here

- API keys
- Access tokens
- Session cookies
- Database credentials
- Payment gateway secrets
- Customer information
- Internal admin records
- CRM data
- Production audit evidence
- Private deployment configuration

## Reporting

Do not publish sensitive vulnerability details in public issues. Contact the LinkResan team through the official product channels with enough information to reproduce the issue safely.

## Production boundary

Production operations remain in `AmirMotefaker/LinkResan-Production`. This repository is intended for public-safe product information and developer resources.
