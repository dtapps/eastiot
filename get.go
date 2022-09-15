package eastiot

import "go.dtapp.net/golog"

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetApiKey() string {
	return c.config.apiKey
}

func (c *Client) GetLogGorm() *golog.ApiClient {
	return c.log.logGormClient
}

func (c *Client) GetLogMongo() *golog.ApiClient {
	return c.log.logMongoClient
}
