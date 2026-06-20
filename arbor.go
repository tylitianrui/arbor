package arbor

type Engine struct {
	// internal fields
}

func Open(path string, opts ...Option) (*Engine, error) {
	return &Engine{}, nil
}
