package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resource-Services
// --------------------------------------

// Orders one of the available services.
// It maps to the `POST` `/projects/0.1/services/{service_type}/{service_id}/order` endpoint
func (r *Services) Order(
	ctx context.Context,
	serviceID int64,
	serviceType rr.ServiceType,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.ServicesOrder(string(serviceType), serviceID),
		nil,
		nil,
	)
}

// Returns a list of services.
// it maps to the `GET` `/projects/0.1/services` endpoint
func (r *Services) List(
	ctx context.Context,
	opts *rr.ListServicesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Services,
		opts,
		nil,
	)
}

// Returns active services.
// it maps to the `GET` `/projects/0.1/services/active` endpoint
func (r *Services) SearchActive(
	ctx context.Context,
	opts *rr.SearchActiveServicesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ServicesActive,
		opts,
		nil,
	)
}
