package data_structure

type Set struct {
	m map[interface{}]struct{}
}

var exist = struct {}{}

func (s *Set) Add(items ...interface{}) bool {
	for _, val := range items {
		s.m[val] = exist
	}
	return true
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) IsEqual(other *Set) bool {
	if s.Size()  != other.Size() {
		return false
	}

	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}




