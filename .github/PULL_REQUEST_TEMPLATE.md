## Description

<!-- Briefly describe what model or tooling changes this PR introduces and why. -->
<!-- For model changes, reference the Jira ticket (e.g., OCM-XXXXX). -->

## Type of Change

- [ ] Model change (new or modified types, resources, methods, or parameters)
- [ ] Generated code update (`make update` after model change)
- [ ] Breaking change (removes or renames existing model elements)
- [ ] Documentation update
- [ ] CI/CD or tooling change

## Verification

- [ ] `make check` passes (model syntax validation)
- [ ] `make verify` passes (generated code matches model)
- [ ] `make lint` passes (indentation enforcement)

## Checklist

- [ ] Model files follow the project's DSL conventions (see [README](../README.md#concepts))
- [ ] Documentation comments use Markdown format
- [ ] Generated `clientapi/` and `openapi/` are updated if model changed (`make update`)
