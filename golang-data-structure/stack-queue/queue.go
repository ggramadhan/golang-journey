package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue : menambahakan int pada bagian belakang
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Denqueue : menghapus int pada bagian depan queue
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(100)
	fmt.Println(myQueue)
	myQueue.Dequeue()

}
