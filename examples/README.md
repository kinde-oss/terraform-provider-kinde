# Examples

This directory contains examples that are mostly used for documentation, but can also be run/tested manually via the Terraform CLI.

The document generation tool looks for files in the following locations by default. All other \*.tf files besides the ones mentioned below are ignored by the documentation tool. This is useful for creating examples that can run and/or are testable even if some parts are not relevant for the documentation.

- **provider/provider.tf** example file for the provider index page
- **data-sources/`full data source name`/data-source.tf** example file for the named data source page
- **resources/`full resource name`/resource.tf** example file for the named resource page

## Local smoke testing

When testing the provider locally before publishing it to a registry, Terraform must be pointed at a locally built provider binary through a development override. The repository provides a convenience target for this:

```sh
make example-smoke EXAMPLE=examples/resources/kinde_role
```

A plain `go install .` followed by `terraform init` in an example directory is not enough. Terraform will still try to resolve `kinde-oss/kinde` from `registry.terraform.io`, and that will fail until the provider is published there. Use the development-override workflow instead.

This target will:

- build the local provider into `.tmp/`
- create a temporary Terraform CLI config with a provider development override
- run `init`, `validate`, `plan`, `apply`, a second `plan`, and `destroy`

> [!IMPORTANT] Smoke examples assume the M2M client used by the provider has the required management scopes for the full lifecycle of each resource (`create`, `read`, `update` when applicable, and `delete`). If `delete` is missing, rollback or destroy steps may require manual cleanup in Kinde.

To run that smoke flow across every resource example folder, use:

```sh
make example-smoke-all
```

The current runnable resource example modules are:

- `examples/resources/kinde_api`
- `examples/resources/kinde_application`
- `examples/resources/kinde_organization_user`
- `examples/resources/kinde_permission`
- `examples/resources/kinde_role`
- `examples/resources/kinde_user_role`

## Example authoring rules

To keep examples useful for both docs and manual testing, each example directory should remain a valid Terraform module as a whole.

- Define `terraform.required_providers` only once per directory.
- Define the default `provider "kinde"` block only once per directory.
- Keep `resource.tf` simple for documentation, and put more advanced runnable compositions in additional files such as `complete_example.tf`.
- Any additional `.tf` files in the same directory must complement the module rather than conflict with the top-level provider or Terraform blocks.
