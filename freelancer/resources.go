package freelancer

// resource contains the shared client used by API resources.
type resource struct {
	client *Client
}

type Projects struct{ resource }
type Collaborations struct{ resource }
type Services struct{ resource }
type Reviews struct{ resource }
type Bids struct{ resource }
type Milestones struct{ resource }
type ExpertGuarantees struct{ resource }
type Users struct{ resource }
type Self struct{ resource }
type Profiles struct{ resource }
type Common struct{ resource }

// Resources provides access to the Freelancer API resources.
//
// Each resource is exposed as a top-level field and shares the client
// configured on the parent Client.
//
// Example:
//
//	res, _, err := client.Resources.Projects.Get(ctx, projectID)
//	if err != nil {
//		// handle error
//	}
//
//	res, _, err := client.Resources.Bids.Accept(ctx, bidID)
type Resources struct {
	Projects
	Collaborations
	Reviews
	Bids
	Milestones
	ExpertGuarantees
	Services
	Users
	Self
	Profiles
	Common
}

func newResources(c *Client) *Resources {
	r := resource{client: c}
	return &Resources{
		Projects:         Projects{r},
		Collaborations:   Collaborations{r},
		Reviews:          Reviews{r},
		Bids:             Bids{r},
		Milestones:       Milestones{r},
		ExpertGuarantees: ExpertGuarantees{r},
		Services:         Services{r},
		Users:            Users{r},
		Self:             Self{r},
		Profiles:         Profiles{r},
		Common:           Common{r},
	}
}
