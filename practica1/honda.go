package main

type Honda struct {

}

func (a *Honda) GetBike() Bike {
	return &HondaBike{
		Bike: Bike{

		}
	}
}

func (a *Honda) GetScooter() Scooter {
	return &HondaScooter{
		Scooter: Scooter{
			
		}
	}
}