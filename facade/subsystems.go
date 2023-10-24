package facade

type SubSystem1 struct{}

func (s *SubSystem1) DoSomething() int {
	return 1
}

type SubSystem2 struct{}

func (s *SubSystem2) DoSomething() int {
	return 2
}

type SubSystem3 struct{}

func (s *SubSystem3) DoSomething() int {
	return 3
}

type SubSystem4 struct{}

func (s *SubSystem4) DoSomething() int {
	return 4
}

type SubSystem5 struct{}

func (s *SubSystem5) DoSomething() int {
	return 5
}
