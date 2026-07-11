---
name: valgo
description: Use when writing, reviewing, debugging, or migrating Go validation code with github.com/cohesivestack/valgo. Covers Valgo validator selection, Is vs Check behavior, nested namespaces, error output, localization/factory options, custom validators, migration notes, and tests. Prefer the current Markdown docs in docs/src/content/docs, and fetch upstream docs when local docs are unavailable.
---

# Valgo

Use this skill for Go projects that use, add, or migrate validations with `github.com/cohesivestack/valgo`.

## Documentation Source

Do not duplicate Valgo docs into the skill context. Load only the docs needed for the user request.

1. First look for local docs in the current workspace:
   - `docs/src/content/docs`
   - For older versions: `docs/src/content/docs/<version>`
2. If local docs are unavailable, fetch the matching Markdown file from:
   - `https://raw.githubusercontent.com/cohesivestack/valgo/master/docs/src/content/docs/<path>`
3. If the user names a version, prefer the versioned docs folder when it exists.
4. If docs and local source code disagree, trust the checked-out code and tests, then mention the mismatch.

Use this routing table to choose the smallest relevant docs set:

- Setup and first validation: `getting-started.md`
- Breaking changes and pre-v1 behavior: `migration.md`
- Sessions, `Is`, `Check`, `New`: `using-valgo/validation-sessions.md`
- Nested structs, slices, indexed errors: `using-valgo/namespaces.md`
- Checking validation state: `using-valgo/querying-results.md`
- Conditional validation flows: `using-valgo/conditional-flows.md`
- Errors, `ToError`, `ToValgoError`, `Merge`, custom output: `using-valgo/errors.md`
- Localization and factories: `using-valgo/localization.md`
- Validator model, names, titles, `Not`: `validators/overview.md`
- Rules by validator family: `validators/rule-index.md`, plus the relevant validator page under `validators/`
- OR validation behavior: `validators/or-operators.md`
- Custom validators: `extending/custom-validators.md`
- Applied examples: `cookbook/*.md` and `cookbook/index.mdx`

## Working Pattern

When implementing or reviewing Valgo code:

1. Inspect the existing Go code for import alias, validation style, error handling, and tests.
2. Read the relevant docs files from the source map above before changing code.
3. Prefer `import v "github.com/cohesivestack/valgo"` for new examples unless the project already uses another alias.
4. Use field names consistently. Add human-friendly titles only when messages need different display text.
5. Return `val.ToError()` for normal Go error flows. Use `ToValgoError()` when structured field/message access is needed.
6. Add or update focused tests for validation behavior and error shape.
7. Run `go test ./...` when practical.

## Valgo Conventions

- Use `v.Is(...)` for normal validation. It short-circuits rules per value after the first failure.
- Use `v.Check(...)` only when all failing messages for a value should be collected.
- Use `v.New()` when validations are built conditionally over multiple statements.
- Use `In(...)`, `InRow(...)`, and `InCell(...)` for nested structs and slices instead of hand-building field paths.
- Use `PathValid`, `AllValid`, and `AnyValid` to inspect validation results without re-validating.
- Use `Merge(...)` to compose reusable validation functions.
- For strings, distinguish rune-length rules (`MaxLength`, `OfLengthBetween`) from byte-length rules (`MaxBytes`, `OfByteLengthBetween`).
- For optional fields, prefer pointer validators and documented optional-field patterns instead of sentinel values.
- For domain-specific rules, create custom validators around `*valgo.ValidatorContext` and provide locale fallbacks for custom error keys.

## Example Skeleton

```go
func ValidateSignup(input Signup) error {
  val := v.Is(
    v.String(input.Email, "email", "Email").Not().Blank(),
    v.String(input.Name, "name", "Name").Not().Blank().OfLengthBetween(2, 80),
  )

  return val.ToError()
}
```

Treat this only as a shape example. Read the relevant docs and existing project tests before choosing exact validators, names, titles, and output handling.
