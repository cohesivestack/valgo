# Valgo

Valgo is a type-safe, expressive, and extensible validator for Go with built-in i18n support.

Unlike validation libraries that rely on struct tags, Valgo defines validation rules as functions. This gives you greater flexibility to validate any value, compose rules programmatically, and decide where validation belongs within your application.

Valgo can be customized to fit your application's needs, from overriding validation messages to localizing them for different languages and contexts.

## Quick example

```go
package main

import (
  "encoding/json"
  "fmt"

  v "github.com/cohesivestack/valgo"
)

func main() {
  val := v.Is(
    v.String("Bob", "full_name").Not().Blank().OfLengthBetween(4, 20),
    v.Number(17, "age").GreaterThan(18),
  )

  if !val.Valid() {
    out, _ := json.MarshalIndent(val.ToError(), "", "  ")
    fmt.Println(string(out))
  }
}
```

Output:

```json
{
  "age": [
    "Age must be greater than \"18\""
  ],
  "full_name": [
    "Full name must have a length between \"4\" and \"20\""
  ]
}
```

## Website and documentation

[valgo.build](https://valgo.build)

## Installing

```bash
go get github.com/cohesivestack/valgo
```

## Agent skill

This repository includes a Valgo Agent Skill installable with [`npx skills`](https://github.com/vercel-labs/skills):

```bash
npx skills add cohesivestack/valgo --skill valgo
```

## Docs

- Start Here
  - [Getting Started](docs/src/content/docs/getting-started.md)
  - [Migration Notes](docs/src/content/docs/migration.md)
- Using Valgo
  - [Validation Sessions](docs/src/content/docs/using-valgo/validation-sessions.md)
  - [Namespaces](docs/src/content/docs/using-valgo/namespaces.md)
  - [Querying Results](docs/src/content/docs/using-valgo/querying-results.md)
  - [Conditional Flows](docs/src/content/docs/using-valgo/conditional-flows.md)
  - [Errors & Output](docs/src/content/docs/using-valgo/errors.md)
  - [Localization & Factory](docs/src/content/docs/using-valgo/localization.md)
- Validators
  - [Overview](docs/src/content/docs/validators/overview.md)
  - [String](docs/src/content/docs/validators/string.md)
  - [Numbers](docs/src/content/docs/validators/numbers.md)
  - [Boolean](docs/src/content/docs/validators/boolean.md)
  - [Time](docs/src/content/docs/validators/time.md)
  - [Comparable](docs/src/content/docs/validators/comparable.md)
  - [Typed & Any](docs/src/content/docs/validators/typed-any.md)
  - [OR Operators (Or / OrElse)](docs/src/content/docs/validators/or-operators.md)
  - [Rule Index](docs/src/content/docs/validators/rule-index.md)
- Extending
  - [Custom Validators](docs/src/content/docs/extending/custom-validators.md)
- Cookbook
  - [Overview](docs/src/content/docs/cookbook/index.mdx)
  - [Sign-up Form](docs/src/content/docs/cookbook/signup-form.md)
  - [Nested Structs](docs/src/content/docs/cookbook/nested-structs.md)
  - [Slices & Indexed Errors](docs/src/content/docs/cookbook/slices.md)
  - [Optional Fields (Pointers)](docs/src/content/docs/cookbook/optional-fields.md)
  - [Conditional Rules](docs/src/content/docs/cookbook/conditional-rules.md)
  - [Custom Messages](docs/src/content/docs/cookbook/custom-messages.md)
  - [Localization](docs/src/content/docs/cookbook/localization.md)
  - [Reusable Validations](docs/src/content/docs/cookbook/reusable-validations.md)
- About
  - [License](docs/src/content/docs/about/license.md)

# Github Code Contribution Guide

We welcome contributions to our project! To make the process smooth and efficient, please follow these guidelines when submitting code:

* **Discuss changes with the community**: We encourage contributors to discuss their proposed changes or improvements with the [community](https://github.com/cohesivestack/valgo/discussions/categories/ideas) before starting to code. This ensures that the changes align with the focus and purpose of the project, and that other contributors are aware of the work being done.

* **Make commits small and cohesive**: It is important to keep your commits focused on a single task or change. This makes it easier to review and understand your changes.

* **Check code formatting with go fmt**: Before submitting your code, please ensure that it is properly formatted using the go fmt command.

* **Make tests to cover your changes**: Please include tests that cover the changes you have made. This ensures that your code is functional and reduces the likelihood of bugs.

* **Update golang docs and README to cover your changes**: If you have made changes that affect documentation or the README file, please update them accordingly.

* **Keep a respectful language with a collaborative tune**: We value a positive and collaborative community. Please use respectful language when communicating with other contributors or maintainers.

* **Go version support:**: Valgo supports the Go versions currently supported by the Go project, plus two previous Go major versions when compatible. The minimum supported Go version is declared in `go.mod`.

# License

Copyright © 2026 Carlos Forero

Valgo is developed and maintained by [Cohesive Stack LLC](https://cohesivestack.com) and released under the [MIT License](LICENSE).
