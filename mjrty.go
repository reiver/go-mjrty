package main


// T is a type passed to the stores the state of an instance of mjrty.
type T struct {
	count uint64
	datum string
}


// New return a new instance mjrty.
func New() (*T) {

	me := T{
		count: 0,
		datum: "",
	}

	return &me
}


// Insert inserts a new element into the instance of mjrty.
func (me *T) Insert(datum string) {

	switch {
		case 0 == me.count:
			me.datum = datum
			me.count    = 1
		case me.datum == datum:
			me.count += 1
		default:
			me.count -= 1
	}
}


// Majority returns what the instance of mjrty currently believes is the majority.
// Note that if the sequennce does not actually have a majority, what this returns
// is nonsense.
func (me *T) Majority() string {
	return me.datum
}
