# Changelog

## [1.2.2] - 2025-12-30

### Important Fix: Slice Query Parameters

The primary focus of this release is a critical fix for how the SDK handles array/slice fields in query parameters across all services.

- Fixed: Corrected the parsing logic for slice parameters (e.g., []int64, []string). Previously, these were not being correctly serialized into URL query strings.

- Global Application: This fix has been applied universally across all service endpoints that support filtering by multiple IDs or statuses.

### Architectural Changes

- Encapsulation: Moved the endpoints package into the internal/ directory. This hides raw API URL constants from the public API, preventing users from accidentally depending on internal routing logic.

- Project Helpers: Added a new GetFullUrl method to the Project struct to easily retrieve the web link for a project.
