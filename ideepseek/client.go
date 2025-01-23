package ideepseek

const (
	ModelV3 = "deepseek-chat"
	ModelR1 = "deepseek-reasoner"

	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
)

type Client struct {
	Url    string
	Apikey string
	Model  string
	Stream bool
}

func NewClient(url, apikey string) *Client {
	return &Client{
		Url:    url,
		Apikey: apikey,
		Model:  ModelV3,
		Stream: false,
	}
}

func (c *Client) WithModel(model string) *Client {
	c.Model = model
	return c
}

func (c *Client) WithStream(stream bool) *Client {
	c.Stream = stream
	return c
}
