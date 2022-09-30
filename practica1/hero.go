package main

type Hero struct {

}

func (a *Hero) GetBike() Bike {
	return &HeroBike{
		Bike: Bike{

		}
	}
}

func (a *Honda) GetScooter() Scooter {
	return &HeroScooter{
		Scooter: Scooter{
			
		}
	}
}