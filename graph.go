package util

// The graph is a simple struct that manages connected nodes. There is no
// distinction made here on directed vs undirected (it's directed implicitly,
// though), link weights, or anything like that. This is as simple as it gets,
// and you can add on to it if you need/want to.

type GraphNode struct {
  Value interface{}
  Links [](*GraphNode)

  currentChild int
}

func (this *GraphNode) LinkTo(other *GraphNode) {
  this.Links = append(this.Links, other)
}

func (this *GraphNode) DoublyLinkTo(other *GraphNode) {
  this.LinkTo(other)
  other.LinkTo(this)
}

func (this *GraphNode) First() interface{} {
  this.currentChild = 0
  return this.Value
}

func (this *GraphNode) Next() interface{} {
  if this.currentChild >= len(this.Links) {
    return nil
  }

  thing := this.Links[this.currentChild]
  this.currentChild += 1
  return thing
}

// TODO: build a general BFS system
