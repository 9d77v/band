package base

type PageQuery struct {
	Page  int `form:"page"`  //页数
	Limit int `form:"limit"` //每页个数
}

func (pq *PageQuery) InjectDefault() {
	if pq.Page < 1 {
		pq.Page = 1
	}
	if pq.Limit == -1 {
	} else if pq.Limit < 1 {
		pq.Limit = 10
	} else if pq.Limit > 100 {
		pq.Limit = 100
	}
}

func (pq *PageQuery) Offset() int {
	return pq.Limit * (pq.Page - 1)
}
