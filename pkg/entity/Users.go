package entity

type Users struct {
	Name    string
	Age     int
	Address string
}

func (u Users) GetName() string {
	return u.Name
}

func (u Users) GetAge() int {
	return u.Age
}

func (u Users) GetAddress() string {
	return u.Address
}

// method multiple return
// GetUsers implements
func (u Users) GetUsers() (string, int, string) {
	name := u.GetName()
	age := u.GetAge()
	addreess := u.GetAddress()
	return name, age, addreess
}
