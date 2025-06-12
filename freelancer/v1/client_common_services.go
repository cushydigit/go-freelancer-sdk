package freelancer

// Returns a list of countries
func (c *Client) NewCountriesListService() *countriesListService {
	return &countriesListService{client: c}
}

// Returns a list of time zones.
func (c *Client) NewTimezonesListService() *timezonesListService {
	return &timezonesListService{client: c}
}
