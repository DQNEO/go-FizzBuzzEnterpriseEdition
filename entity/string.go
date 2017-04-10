package entity

const FIZZ string = "Fizz"
const BUZZ string = "Buzz"
const FIZZ_BUZZ string  = "FizzBuzz"


type MyString struct {
	Data string
}

func (myString MyString) String() string {
	return myString.Data;
}

type Stringable interface {
	String() string
}


var FIZZ_ENTITY MyString = MyString{Data: FIZZ}

var BUZZ_ENTITY MyString = MyString{Data:BUZZ}

var FIZZ_BUZZ_ENTITY MyString = MyString{Data:FIZZ_BUZZ}