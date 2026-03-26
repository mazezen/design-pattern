package iterator

import "testing"

// go test -v ./...
func TestUserIterator(t *testing.T) {
	users := []User{
		{"A", 18},
		{"B", 25},
		{"C", 30},
	}

	ut := NewUserIterator(users)

	count := 0
	for ut.HasNext() {
		u := ut.Next().(User)
		if u.Name == "" {
			t.Error("invalid user")
		}
		count++
	}

	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func TestNumberIterator(t *testing.T) {
	nums := []int{1, 2, 3}

	nt := NewNumberIterator(nums)

	sum := 0

	for nt.HasNext() {
		sum += nt.Next().(int)
	}

	if sum != 6 {
		t.Errorf("expected 6, got %d", sum)
	}
}

func TestFilterIterator(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	base := NewNumberIterator(nums)

	// 过滤偶数
	filter := NewFilterIterator(base, func(v interface{}) bool {
		return v.(int) % 2 == 0
	})

	var result []int
	for filter.HasNext() {
		result = append(result, filter.Next().(int))
	}

	if len(result) != 2 || result[0] != 2 || result[1] != 4 {
		t.Errorf("unexpected result: %v", result)
	}
 
}