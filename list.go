package util

type item struct {
  value interface{}
  next *item
}

type List struct {
  head *item
  size int
  next *item
}

//Len returns the size of stack
func (l *List) Len() int {
  return l.size
}

func (l *List) First() (value interface{}) {
  if l.size > 0 {
    value = l.head.value
    l.next = l.head.next
  }
  return
}

func (l *List) Next() (value interface{}) {
  if l.next != nil {
    value = l.next.value
    l.next = l.next.next
  }
  return
}

func (l *List) Add(value interface{}) {
  if l.head == nil {
    l.head = &item {
      value: value,
      next: nil,
    }
    l.size = 1
    return
  }

  var temp *item
  for temp = l.head; temp.next != nil; temp = temp.next {}
  temp.next = &item {
    value: value,
    next: nil,
  }
  l.size += 1
}
