package gram

type Pipeline struct {
}

func Cmd(name, desc string) *Pipeline {
	return &Pipeline{}
}

func Any() *Pipeline {
	return &Pipeline{}
}

func Step(func(*Context)) *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Step(func(*Context)) *Pipeline {
	return &Pipeline{}
}
