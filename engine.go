package arbor

type Engine struct {
	// todo: add internal fields
}

func Open(path string, opts ...Option) (*Engine, error) {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	return &Engine{}, nil
}
