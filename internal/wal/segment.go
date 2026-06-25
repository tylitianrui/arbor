package wal

// Segment represents a segment file.
type Segment struct {
	dir string
	idx int
}

// Index returns the index of the segment.
func (s *Segment) Index() int {
	return s.idx
}

// Directory returns the directory of the segment.
func (s *Segment) Directory() string {
	return s.dir
}
