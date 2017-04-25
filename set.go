package set

//C++ SET https://groups.google.com/forum/#!topic/golang-nuts/xqqaR2N2sAI
type Set map[string]struct{}

func New() *Set {
	s := make(Set)
	return &s
}

func (s *Set) Add(values ...string) {
	for _, v := range values {
		if v != "" {
			(*s)[v] = struct{}{}
		}
	}
}

func (s *Set) GetAll() []string {
	var keys []string
	for k, _ := range *s {
		keys = append(keys, k)
	}
	return keys
}
