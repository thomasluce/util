package util

type Queue []interface{}

func (q *Queue) Push(value interface{}) {
  *q = append(*q, value)
}

func (q *Queue) Pop() interface{} {
  if len(*q) == 0 {
    return nil
  }

  n := (*q)[0]
  *q = (*q)[1:]
  return n
}
