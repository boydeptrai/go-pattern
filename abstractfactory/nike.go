package abstractfactory

type Nike struct {

}

func (n *Nike) MakeShoe() IShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 12,
		},
	}
}

func (n *Nike) MakeShort() IShort {
	return &nikeShort{
		short: short{
			logo: "Nike",
			size: 12,
		},
	}
}