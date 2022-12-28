package entity

type Users struct {
	Name    string
	Age     int
	Address string
}

func (u *Users) getName() string {
	return u.Name
}

func (u *Users) getAge() int {
	return u.Age
}

func (u *Users) getAddress() string {
	return u.Address
}

// method multiple return
// GetUsers implements
func (u Users) GetUsers() (string, int, string) {
	name := u.getName()
	age := u.getAge()
	addreess := u.getAddress()
	return name, age, addreess
}
