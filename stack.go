// Copyright 2014, Ali Najafizadeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// Author: Ali Najafizadeh

//Inspired by: https://gist.github.com/bemasher/1777766

// As the header implies, this is copy-pasta'd for simplicity. In the
// real-world, do thing this as a linked list is silly for speed reasons
// (arrays and slices in go would be SUPER fast by comparison), I just didn't
// want to write Yet Another Goddamn Stack (tm)

package util

type item struct {
  value interface{}
  next  *item
}

//Stack the implementation of stack
//this stack is not thread safe!
type Stack struct {
  top  *item
  size int
}

//Len returns the size of stack
func (s *Stack) Len() int {
  return s.size
}

//Push pushs a value to the top of stack
func (s *Stack) Push(value interface{}) {
  s.top = &item{
    value: value,
    next:  s.top,
  }
  s.size++
}

//Pop returns a top value. make sure to check exists
//it is possible to push nil value. So again check the exists value
func (s *Stack) Pop() (value interface{}, exists bool) {
  exists = false
  if s.size > 0 {
    value, s.top = s.top.value, s.top.next
    s.size--
    exists = true
  }

  return
}

//Peek returns a top without removing it from list
func (s *Stack) Peek() (value interface{}, exists bool) {
  exists = false
  if s.size > 0 {
    value = s.top.value
    exists = true
  }

  return
}
