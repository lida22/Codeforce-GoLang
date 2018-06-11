package main

import (
	"fmt"
)

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	q := new(Queue)
	q.elements = make([]interface{}, 0)
	return q
}

func (q *Queue) Enqueue(e interface{}) *Queue {
	q.elements = append(q.elements, e)
	return q
}

func (q *Queue) Dequeue() (e interface{}) {
	e, q.elements = q.elements[0], q.elements[1:len(q.elements)]
	return
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func main() {
	var f, s, g, u, d, next, next1, dis, ret int
	fmt.Scan(&f, &s, &g, &u, &d)
	seen := []int{}
	q := NewQueue()
	q.Enqueue(s)
	seen = append(seen, s)
	dis = 0
	ret = 1
	var result, result1 int
	for q.Size() > 0 {

		vertex := q.Dequeue().(int)
		next = vertex + u
		//fmt.Println(vertex)
		next1 = vertex - d
		if next == g || next1 == g {
			ret = 0
		}
		for i := 0; i < len(seen); i++ {
			if seen[i] == next {
				result = 1
				break
			} else {
				result = 0
			}
		}
		for i := 0; i < len(seen); i++ {
			if seen[i] == next1 {
				result1 = 1
				break
			} else {
				result1 = 0
			}
		}
		//fmt.Println(result,result1,next1)
		if next > f && result1 == 0 {
			dis = dis + 1
		}
		if next <= f && result == 0 {
			seen = append(seen, next)
			q.Enqueue(next)

			dis = dis + 1

			if next == g {
				fmt.Println(dis)
			}

		} 

		if next1 >= 1 && result1 == 0 {
			seen = append(seen, next1)
			q.Enqueue(next1)
			if next1 == g {
				fmt.Println(dis)
			}
		} 
	}
	if ret == 1 {
		fmt.Println("use")
	}
}

