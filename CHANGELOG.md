# Changelog

All notable changes to this provider are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Initial provider surface for the Kinde Management API.
  - Resources: `kinde_api`, `kinde_application`, `kinde_application_connection`,
    `kinde_connection`, `kinde_organization`, `kinde_organization_user`,
    `kinde_permission`, `kinde_role`, `kinde_user`, `kinde_user_role`.
  - Data sources: `kinde_api`, `kinde_application`, `kinde_connections`.

### Changed

- Upgraded to Go toolchain 1.27.3 (minimum `go` 1.26.0) and to
  `terraform-plugin-framework` v1.19.0, `terraform-plugin-go` v0.31.0,
  `terraform-plugin-log` v0.11.0 and `terraform-plugin-testing` v1.16.0.

[Unreleased]: https://github.com/kinde-oss/terraform-provider-kinde/commits/main
