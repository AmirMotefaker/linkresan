# Public and Production release policy

LinkResan uses two repositories:

- `AmirMotefaker/LinkResan` is the public Community and documentation repository.
- `AmirMotefaker/LinkResan-Production` is the private Production source repository.

Public releases may mirror verified Production milestones without publishing proprietary Production source code, infrastructure secrets, credentials or private operational evidence.

A public metadata release must:

1. Clearly identify itself as a public documentation mirror.
2. Describe externally meaningful capabilities without exposing private implementation details.
3. Link the milestone to public-safe GitHub records.
4. Never imply that private Production source exists in the public tag.
5. Use the same release version as the corresponding Production milestone when the public release is a synchronized metadata record.
6. Identify the public tag target separately from the verified private Production application commit.
7. Never imply that review of public metadata is a retroactive code review of historical private implementation commits.
8. Disclose legacy direct-to-main history when the original change did not use the current Issue → Branch → Commit → Pull Request → Review → Merge process.
