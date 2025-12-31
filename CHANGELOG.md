# Changelog

## [1.3.1] - 2025-12-31

### Bug Fixes

Project Service: Fixed a critical issue where project URLs were not being constructed correctly. The SDK now ensures the full URL is generated for project-specific requests, preventing 404 or malformed request errors.

## [1.3.0] - 2025-12-31

### Refactoring & Improvements

- New Internal Query Package: Introduced a new internal/query package to handle repetitive logic, significantly reducing boilerplate code across all services.

- Service Architecture Consolidation: Refactored service structures from multiple fragmented files into consolidated single-file modules for better maintainability.

- Response Handling: Updated API responses to return pointer value objects, allowing for better nil-checking and memory efficiency in the SDK consumer layer.

### Testing & Development

- Full Coverage for Query Package: Achieved 100% test coverage for the new internal query package.

- Global Test Coverage: Improved overall SDK test coverage to 32%.

- Real-world Testing Examples: Added golang.org/x/net/proxy within the examples/ package. This allows for testing the client with real SOCKS proxies without adding the dependency to the SDK core.

## [1.2.2] - 2025-12-30

### Important Fix: Slice Query Parameters

The primary focus of this release is a critical fix for how the SDK handles array/slice fields in query parameters across all services.

- Fixed: Corrected the parsing logic for slice parameters (e.g., []int64, []string). Previously, these were not being correctly serialized into URL query strings.

- Global Application: This fix has been applied universally across all service endpoints that support filtering by multiple IDs or statuses.

### Architectural Changes

- Encapsulation: Moved the endpoints package into the internal/ directory. This hides raw API URL constants from the public API, preventing users from accidentally depending on internal routing logic.

- Project Helpers: Added a new GetFullUrl method to the Project struct to easily retrieve the web link for a project.
