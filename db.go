package arbor

type Options struct {
}

func (o *Options) Clone() *Options {
	if o == nil {
		return &Options{}
	}
	newOpt := *o
	return &newOpt
}

func (o *Options) Validate() error {
	return nil
}

func (o *Options) SetDefaultOptions() {

}

type DB struct {
	dirname string
	opts    *Options
}

func Open(dirname string, opts *Options) (*DB, error) {
	opts = opts.Clone()
	opts.SetDefaultOptions()

	if err := opts.Validate(); err != nil {
		return nil, err
	}
	db := &DB{
		dirname: dirname,
		opts:    opts,
	}
	return db, nil
}

func (db *DB) Close() error {
	return nil
}
