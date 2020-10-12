package computer

type Spec struct{ //exported struct start with capital letter
	Maker string //exported field
	Price int    //exported field
	model string //unexported field start with small letter
}
