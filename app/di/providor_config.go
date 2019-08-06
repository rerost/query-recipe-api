package di

type Config interface {
	AccessToken() string
	Owner() string
	Repository() string
}

type config struct {
	accessToken string
	owner       string
	repository  string
}

func ProvidorConfig(
	AccessToken string,
	owner string,
	repository string,
) Config {
	return &config{
		accessToken: AccessToken,
		owner:       owner,
		repository:  repository,
	}
}

func (c *config) AccessToken() string {
	return c.accessToken
}

func (c *config) Owner() string {
	return c.owner
}

func (c *config) Repository() string {
	return c.repository
}
