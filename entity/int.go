package entity

import "strconv"

type MyInt int

func (myInt MyInt) String() string {
	return strconv.Itoa(int(myInt))
}

