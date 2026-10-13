package endpoints

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicEndpoints(t *testing.T) {
	tests := []struct {
		name     string
		actual   Endpoint
		expected Endpoint
	}{
		{
			name:     "Project",
			actual:   Project(123),
			expected: "/projects/0.1/projects/123",
		},
		{
			name:     "ProjectInvite",
			actual:   ProjectInvite(123),
			expected: "/projects/0.1/projects/123/invite",
		},
		{
			name:     "ProjectBids",
			actual:   ProjectBids(123),
			expected: "/projects/0.1/projects/123/bids",
		},
		{
			name:     "ProjectBidsInfo",
			actual:   ProjectBidsInfo(123),
			expected: "/projects/0.1/projects/123/bids_info",
		},
		{
			name:     "ProjectMilestones",
			actual:   ProjectMilestones(123),
			expected: "/projects/0.1/projects/123/milestones",
		},
		{
			name:     "ProjectMilestoneRequests",
			actual:   ProjectMilestoneRequests(123),
			expected: "/projects/0.1/projects/123/milestone_requests",
		},
		{
			name:     "ProjectIPContractInfo",
			actual:   ProjectIPContractInfo(123),
			expected: "/projects/0.1/projects/123/ip_contract_info",
		},
		{
			name:     "ProjectCollaborations",
			actual:   ProjectCollaborations(123),
			expected: "/projects/0.1/projects/123/collaborations",
		},
		{
			name:     "ProjectCollaborationsActions",
			actual:   ProjectCollaborationsActions(123, 456),
			expected: "/projects/0.1/projects/123/collaborations/456/actions",
		},
		{
			name:     "ServicesOrder",
			actual:   ServicesOrder("design", 123),
			expected: "/projects/0.1/services/design/123/order",
		},
		{
			name:     "Bid",
			actual:   Bid(123),
			expected: "/projects/0.1/bids/123",
		},
		{
			name:     "BidTimeTracking",
			actual:   BidTimeTracking(123),
			expected: "/projects/0.1/bids/123/time_tracking",
		},
		{
			name:     "BidEditRequests",
			actual:   BidEditRequests(123),
			expected: "/projects/0.1/bids/123/edit_requests",
		},
		{
			name:     "BidEditRequest",
			actual:   BidEditRequest(123, 456),
			expected: "/projects/0.1/bids/123/edit_requests/456",
		},
		{
			name:     "BidsRatings",
			actual:   BidsRatings(123),
			expected: "/projects/0.1/bids/123/bid_ratings",
		},
		{
			name:     "BidRating",
			actual:   BidRating(123, 456),
			expected: "/projects/0.1/bids/123/bid_ratings/456",
		},
		{
			name:     "Milestone",
			actual:   Milestone(123),
			expected: "/projects/0.1/milestones/123",
		},
		{
			name:     "MilestoneRequest",
			actual:   MilestoneRequest(123),
			expected: "/projects/0.1/milestone_requests/123",
		},
		{
			name:     "Review",
			actual:   Review(123),
			expected: "/projects/0.1/reviews/123",
		},
		{
			name:     "ExpertGuarantee",
			actual:   ExpertGuarantee(123),
			expected: "/projects/0.1/expert_guarantees/123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Endpoint(tt.expected), tt.actual)
		})
	}
}

func TestDynamicEndpoints_LargeID(t *testing.T) {
	const id int64 = 9223372036854775807

	assert.Equal(
		t,
		Endpoint("/projects/0.1/projects/9223372036854775807"),
		Project(id),
	)
}
