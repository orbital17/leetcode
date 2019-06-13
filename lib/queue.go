package lib

type Queue struct {
	start *ListNodeAbstract
	end   *ListNodeAbstract
}

func (q *Queue) Add(v interface{}) {
	if q.start == nil {
		q.start = &ListNodeAbstract{Val: v}
		q.end = q.start
	} else {
		q.end.Next = &ListNodeAbstract{Val: v}
		q.end = q.end.Next
	}
}

func (q *Queue) Poll() (r interface{}) {
	r = q.start.Val
	if q.start == q.end {
		q.start = nil
		q.end = nil
	} else {
		q.start = q.start.Next
	}
	return
}

func (q *Queue) IsEmpty() bool {
	return q.start == nil
}
