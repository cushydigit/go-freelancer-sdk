package reqres

import (
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
)

func TestGetFullUrl(t *testing.T) {
	tests := []struct {
		name     string
		base     string
		seoURL   string
		expected string
	}{
		{
			seoURL:   "my-project",
			expected: endpoints.Base + "/projects/my-project",
		},
		{
			seoURL:   "my-project2",
			expected: endpoints.Base + "/projects/my-project2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Project{SeoURL: tt.seoURL}
			if actual := p.GetFullUrl(); actual != tt.expected {
				t.Errorf("GetFullUrl() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}
