package main

type IBike interface {
	Name() string
}

type Bike struct {
	name string
}

func (s *Bike) Name() string {
	return s.name
}