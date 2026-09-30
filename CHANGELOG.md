# Changelog

## [1.4.0] - 2025-12-31

### Refactoring & Improvements

- Refactored request-related packages, including `actions`, `body`, and `request`, into the `reqres` package.
- Renamed several enums and types to improve naming consistency and API clarity.
- Replaced the previous debug logging mechanism with Go's structured `slog` logging.
- Added `IsAPIError` helper for detecting and extracting `*APIError` from wrapped errors.
- Improved query parameter encoding with explicit error handling.
- Added support for `time.Time`
- Project search and other service options now use `*time.Time` as query parameters.
- Added timestamp helper methods for converting Freelancer API Unix timestamps to `time.Time`.
- Query encoding errors are now returned instead of silently logged.
- Query parameter encoding is now centralized in the request execution layer `execute`.
- Improved internal endpoint handling with typed endpoints and centralized
  dynamic endpoint construction.

### Breaking Changes

- Removed explicit client-side rate limiting from the SDK.
- Added response metadata support for exposing API rate-limit information, including rate-limit limits and windows.
- Refactor `JobBundleService` and `JobBundleCategories` to main `JobsService`.
- Refactor `BidEditRequestsService` and `BidRatingsService` to main `BidsService`.
- Refactor and merge `MilestoneRequestsService` to `MilestonesService`.
- Refactor and move `CurrenciesService`, `BudgetsService` and `CategoriesService` to `CommonService`.
- Refactor and merge `SelfJobsService` and `ProfilesService` to `SelfService`.
- Refactor and merge `ReputationsService`, `EnterprisesService`, `ViolationsService`, `PortfoliosService` and `PoolsService` to main `UsersService`.
- The top-level `Client.Services` field has been renamed to `Client.Resources`
to avoid ambiguity with the Freelancer `/services` API resource.
- Services (Resources) are now exposed as top-level resources instead of being nested
under `Projects` and `Users`.

This aligns the SDK service structure with the Freelancer API resources.



### Bug Fixes

- fixes the response type of jon in list category response `ListCategoriesResponse``
- fixes the URL of sandbox 

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
