package gomock

import (
//"strconv"
//"time"
)

// Name will generate a random First and Last Name
func Name() string {
	return "Test"
	//return LastName() + FirstName(Gender())
}

func Age() int {
	return Number(0, 100)
}
