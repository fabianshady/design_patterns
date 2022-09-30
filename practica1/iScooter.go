package main

type IScooter interface {
	Name() string
}

type Scooter struct {
	name string
}

func (s *Scooter) Name() string {
	return s.name
}