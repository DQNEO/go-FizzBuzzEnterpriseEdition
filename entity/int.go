package entity

import "strconv"

type MyInt struct {
	Data int
}

func (myInt MyInt) String() string {
	return strconv.Itoa(myInt.Data)
}


