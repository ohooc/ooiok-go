package options_partten

type HttpClient struct {
	Timeout     int
	MaxIdle     int
	ErrCallback func(err error)
}

type ClientOption func(client *HttpClient)

type ClientOptions []ClientOption

func (opts ClientOptions) apply(c *HttpClient) {
	for _, opt := range opts {
		opt(c)
	}
}

func NewHttpClient(opts ...ClientOption) *HttpClient {
	cli := &HttpClient{}
	ClientOptions(opts).apply(cli)
	return cli
}

func WithTimeout(time int) ClientOption {
	return func(client *HttpClient) {
		client.Timeout = time
	}
}

func WithMaxIdle(max int) ClientOption {
	return func(client *HttpClient) {
		client.MaxIdle = max
	}
}

func WithCallback(f func(err error)) ClientOption {
	return func(client *HttpClient) {
		client.ErrCallback = f
	}
}
