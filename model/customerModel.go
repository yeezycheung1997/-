package model

type Customer struct {
	Id    int
	Name  string
	Age   int
	Email string
}

func NewCustomer(name string, age int, email string) Customer {
	return Customer{
		Name:  name,
		Age:   age,
		Email: email,
	}
}
