package main

import (
	"fmt"
	"strings"
)

var (
	start   = `318946572`
	min     = 1
	max     = -1
	cupsMap map[int]*cup
)

type cup struct {
	val  int
	prev *cup
	next *cup
}

func main() {

	first := setupCups(start)

	//results := make(map[string]bool)

	ptr := first
	i := 0
	for i < 10000000 {
		ptr = move(ptr)
		//val := getN(ptr,20)
		//if _, ok := results[val]; ok {
		//	fmt.Println(i)
		//	results = make(map[string]bool)
		//}
		//results[val] = true
		i++
	}

	val := getN(cupsMap[1], 20)
	fmt.Println(val)

	result := int64(cupsMap[1].next.val) * int64(cupsMap[1].next.next.val)
	fmt.Println(result)

}

func getN(ptr *cup, n int) string {
	vals := make([]string, 0, n)
	for i, p := 0, ptr; i < n; i++ {
		vals = append(vals, fmt.Sprintf("%d", p.val))
		p = p.next
	}
	return strings.Join(vals, ",")
}

func move(ptr *cup) *cup {
	links, end, unlinked := unlink(ptr)
	d := destination(ptr, unlinked)
	end.next, d.next, links.prev, d.next.prev = d.next, links, d, end
	return ptr.next
}

func destination(ptr *cup, unlinked map[int]*cup) *cup {
	val := ptr.val
	for {
		val = val - 1
		if val < min {
			val = max
		}
		if _, ok := unlinked[val]; !ok {
			return cupsMap[val]
		}
	}
}

func unlink(ptr *cup) (links *cup, end *cup, unlinked map[int]*cup) {
	unlinked = make(map[int]*cup)
	links = ptr.next
	end = ptr
	for i := 0; i < 3; i++ {
		end = end.next
	}
	end.next.prev, ptr.next = ptr, end.next
	end.next, links.prev = nil, nil

	unlinked = make(map[int]*cup)

	for p := links; p != nil; p = p.next {
		unlinked[p.val] = p
	}
	return
}

func setupCups(s string) *cup {
	cupsMap = make(map[int]*cup)
	vals := []byte(s)

	first := &cup{val: int(vals[0] - 48)}
	last := first

	cupsMap[first.val] = first

	min = first.val
	max = first.val

	for i := range vals[1:] {
		tmp := &cup{val: int(vals[i+1] - 48)}
		tmp.prev = last
		last.next = tmp
		last = tmp

		if last.val < min {
			min = last.val
		}
		if last.val > max {
			max = last.val
		}
		cupsMap[last.val] = last
	}

	for i := max + 1; i <= 1000000; i++ {
		tmp := &cup{val: i}
		tmp.prev = last
		last.next = tmp
		last = tmp
		max = i
		cupsMap[last.val] = last
	}

	last.next, first.prev = first, last
	return first

}
