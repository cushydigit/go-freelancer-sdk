package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
	"github.com/stretchr/testify/assert"
)

func TestReviews_List(t *testing.T) {
	opts := rr.ListReviewsOptions{
		Projects:    []int64{1, 2},
		ReviewTypes: []rr.ReviewType{rr.ReviewTypeProject},
		UserStatus:  rr.Bool(true),
		Limit:       rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Reviews), r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_status"))
		assert.Equal(t, "10", q.Get("limit"))
		assert.ElementsMatch(t, []string{string(rr.ReviewTypeProject)}, q["review_types[]"])
		assert.ElementsMatch(t, []string{"1", "2"}, q["projects[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Reviews.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestReviews_CreateForEmployer(t *testing.T) {
	body := rr.CreateReviewForEmployerBody{
		ReviewBody: rr.ReviewBody{
			ProjectID:  100,
			ToUserID:   101, // employer user id
			FromUserID: 102, // freelancer user id
			ReviewType: rr.ReviewTypeProject,
			Comment:    "test",
			// regardless of explicitly wrong role it should be corrected
			Role: rr.RoleFreelancer,
		},
		ReputationData: rr.EmployerReputationData{
			Category: rr.EmployerCategoryRatings{
				Communication:   5,
				Professionalism: 4,
				Clarity:         3,
				Payment:         2,
				WorkForAgain:    1,
			},
		},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, string(endpoints.Reviews), r.URL.Path)
		// body
		var res rr.CreateReviewForEmployerBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.ProjectID, res.ProjectID)
		assert.Equal(t, body.FromUserID, res.FromUserID)
		assert.Equal(t, body.ToUserID, res.ToUserID)
		// regardless of explicitly wrong role it should be corrected
		assert.Equal(t, rr.RoleEmployer, res.Role)
		assert.Equal(t, body.Comment, body.Comment)
		assert.Equal(t, body.ReviewType, res.ReviewType)
		assert.Equal(t, body.ReputationData.Category.Clarity, res.ReputationData.Category.Clarity)
		assert.Equal(t, body.ReputationData.Category.Communication, res.ReputationData.Category.Communication)
		assert.Equal(t, body.ReputationData.Category.Payment, res.ReputationData.Category.Payment)
		assert.Equal(t, body.ReputationData.Category.WorkForAgain, res.ReputationData.Category.WorkForAgain)
		assert.Equal(t, body.ReputationData.Category.Professionalism, res.ReputationData.Category.Professionalism)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Reviews.CreateForEmployer(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestReviews_CreateForFreelancer(t *testing.T) {
	body := rr.CreateReviewForFreelancerBody{
		ReviewBody: rr.ReviewBody{
			ProjectID:  100,
			ToUserID:   101, // freelancer user id
			FromUserID: 102, // employer user id
			ReviewType: rr.ReviewTypeProject,
			Comment:    "test",
			// regardless of explicitly wrong role it should be corrected
			Role: rr.RoleEmployer,
		},
		ReputationData: rr.FreelancerReputationData{
			OnBudget: 5,
			OnTime:   3,
			Category: rr.FreelancerCategoryRatings{
				Communication:   5,
				Professionalism: 4,
				Expertise:       3,
				Quality:         2,
				HireAgain:       1,
			},
		},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, string(endpoints.Reviews), r.URL.Path)
		// body
		var res rr.CreateReviewForFreelancerBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.ProjectID, res.ProjectID)
		assert.Equal(t, body.FromUserID, res.FromUserID)
		assert.Equal(t, body.ToUserID, res.ToUserID)
		// regardless of explicitly wrong role it should be corrected
		assert.Equal(t, rr.RoleFreelancer, res.Role)
		assert.Equal(t, body.Comment, body.Comment)
		assert.Equal(t, body.ReviewType, res.ReviewType)
		assert.Equal(t, body.ReputationData.Category.Communication, res.ReputationData.Category.Communication)
		assert.Equal(t, body.ReputationData.Category.Professionalism, res.ReputationData.Category.Professionalism)
		assert.Equal(t, body.ReputationData.Category.Expertise, res.ReputationData.Category.Expertise)
		assert.Equal(t, body.ReputationData.Category.Quality, res.ReputationData.Category.Quality)
		assert.Equal(t, body.ReputationData.Category.HireAgain, res.ReputationData.Category.HireAgain)
		assert.Equal(t, body.ReputationData.OnBudget, res.ReputationData.OnBudget)
		assert.Equal(t, body.ReputationData.OnTime, res.ReputationData.OnTime)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Reviews.CreateForFreelancer(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestReviews_FeatureProject(t *testing.T) {
	reviewID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		assert.Equal(t, string(endpoints.Review(reviewID)), r.URL.Path)
		// body
		var res rr.ActionReviewBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, rr.ReviewActionFeature, res.Action)
		assert.Equal(t, rr.ReviewTypeProject, res.ReviewType)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Reviews.FeatureProject(context.Background(), reviewID)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestReviews_UnfeatureProject(t *testing.T) {
	reviewID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		assert.Equal(t, string(endpoints.Review(reviewID)), r.URL.Path)
		// body
		var res rr.ActionReviewBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, rr.ReviewActionUnfeature, res.Action)
		assert.Equal(t, rr.ReviewTypeProject, res.ReviewType)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Reviews.UnfeatureProject(context.Background(), reviewID)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestReviews_FeatureContest(t *testing.T) {
	reviewID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		assert.Equal(t, string(endpoints.Review(reviewID)), r.URL.Path)
		// body
		var res rr.ActionReviewBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, rr.ReviewActionFeature, res.Action)
		assert.Equal(t, rr.ReviewTypeContest, res.ReviewType)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Reviews.FeatureContest(context.Background(), reviewID)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestReviews_UnfeatureContest(t *testing.T) {
	reviewID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		assert.Equal(t, string(endpoints.Review(reviewID)), r.URL.Path)
		// body
		var res rr.ActionReviewBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, rr.ReviewActionUnfeature, res.Action)
		assert.Equal(t, rr.ReviewTypeContest, res.ReviewType)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Reviews.UnfeatureContest(context.Background(), reviewID)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}
