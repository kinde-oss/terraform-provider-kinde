# Changelog

All notable changes to this provider are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.1] - 2026-10-26

### Changed

- Upgraded to Go toolchain 1.27.3 (minimum `go` 1.26.0) and to
  `terraform-plugin-framework` v1.19.0, `terraform-plugin-go` v0.31.0,
  `terraform-plugin-log` v0.11.0 and `terraform-plugin-testing` v1.16.0.
- Expanded the provider documentation page with setup, required Management API
  scopes, usage and import guidance.

## [0.1.0] - 2026-08-21

### Added

- Initial provider surface for the Kinde Management API.
  - Resources: `kinde_api`, `kinde_application`, `kinde_application_connection`,
    `kinde_connection`, `kinde_organization`, `kinde_organization_user`,
    `kinde_permission`, `kinde_role`, `kinde_user`, `kinde_user_role`.
  - Data sources: `kinde_api`, `kinde_application`, `kinde_connections`.

[Unreleased]: https://github.com/kinde-oss/terraform-provider-kinde/compare/v0.1.1...HEAD
[0.1.1]: https://github.com/kinde-oss/terraform-provider-kinde/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/kinde-oss/terraform-provider-kinde/releases/tag/v0.1.0
