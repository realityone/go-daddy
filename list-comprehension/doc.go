/*
	List Comprehension in Golang:

		```go
		type Student struct {
			ID   int
			Name string
			Age  int
		}

		// assume there are a lot of students
		stus := []Student{.........}

		// collect the students ID who is older than 20
		sids := make([]int, 0, len(stus))
		for _, s := range stus {
			if s.Age <= 20 {
				continue
			}
			sids = append(sids, s.ID)
		}

		// using list comprehension is much simpler
		sids := []int{s.ID for _, s := range stus if s.Age <= 20}
		```
*/

package comprehension
