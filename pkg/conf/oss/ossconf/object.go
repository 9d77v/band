package ossconf

type ObjectInfo struct {
	Size int64 `json:"size"`
}

type GetObjectOption struct {
	bytesRange *Range
}

type Range struct {
	Start int64
	End   int64
}

func (opt *GetObjectOption) SetRange(start, end int64) {
	opt.bytesRange = &Range{Start: start, End: end}
}

func (opt *GetObjectOption) GetRange() *Range {
	return opt.bytesRange
}
