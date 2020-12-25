package main

import "fmt"

var (
	start   = `318946572`
	min     = -1
	max     = -1
	cups    []*cup
	cupsMap map[int]*cup
)

type cup struct {
	val  int
	prev *cup
	next *cup
}

func main() {
	cupsMap = make(map[int]*cup)
	cups = getCups(start)
	for i := range cups {
		cupsMap[cups[i].val] = cups[i]
	}

	min, max = cups[0].val, cups[0].val
	for i := range cups[1:] {
		if cups[i].val < min {
			min = cups[i].val
		}
		if cups[i].val > max {
			max = cups[i].val
		}
	}

	ptr := cups[0]
	i := 0
	for i < 100 {
		ptr = move(ptr)
		i++
	}

	print(cupsMap[1])

}

func print(ptr *cup) {
	vals := make([]byte, 0, len(cups))
	for p := ptr; p.next != ptr; p = p.next {
		vals = append(vals, byte(p.next.val+48))
	}

	fmt.Println(string(vals))
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

func getCups(s string) []*cup {
	vals := []byte(s)
	cups := make([]*cup, len(vals), len(vals))
	for i := range vals {
		vals[i] = vals[i] - 48
		cups[i] = &cup{val: int(vals[i])}
		if i > 0 {
			cups[i].prev = cups[i-1]
			cups[i-1].next = cups[i]
		}
	}
	cups[0].prev = cups[len(cups)-1]
	cups[len(cups)-1].next = cups[0]
	return cups

}
