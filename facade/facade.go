package facade

type Facade struct{}

// Facade method for executing systems with even number
func (f *Facade) ExecuteEvenSubSystems() []int {
	s2 := &SubSystem2{}
	s4 := &SubSystem4{}

	return []int{
		s2.DoSomething(),
		s4.DoSomething(),
	}
}

// Facade method for executing systems with odd number
func (f *Facade) ExecuteOddSubSystems() []int {
	s1 := &SubSystem1{}
	s3 := &SubSystem3{}
	s5 := &SubSystem5{}

	return []int{
		s1.DoSomething(),
		s3.DoSomething(),
		s5.DoSomething(),
	}
}
