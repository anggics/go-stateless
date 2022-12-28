package entity

type Customer struct {
	Id        int
	FirstName string
	LastName  string
}

func (c *Customer) GetId() int {
	return c.Id
}

func (c *Customer) GetFirstName() string {
	return c.FirstName
}

func (c *Customer) GetLastName() string {
	return c.LastName
}
