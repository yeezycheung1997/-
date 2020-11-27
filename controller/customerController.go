package controller

import (
	"fmt"
	"客户信息管理软件02/model"
)

type CustomersController struct {
	Customers []model.Customer
}

func (this *CustomersController) Add(name string, age int, email string) {
	newCustomer := model.NewCustomer(name, age, email)
	newCustomer.Id = len(this.Customers)
	this.Customers = append(this.Customers, newCustomer)
	fmt.Println("添加成功")
}

func (this *CustomersController) ShowList() []model.Customer {
	return this.Customers
}

func (this *CustomersController) Delete(id int) {
	if id >= len(this.Customers) {
		this.Customers = append(this.Customers[:id])
	} else {
		this.Customers = append(this.Customers[:id], this.Customers[id+1:]...)
	}
	fmt.Println("删除用户成功! ")
}

func (this *CustomersController) Change(id int) {
	if id < 0 || id >= len(this.Customers) {
		fmt.Println("用户Id无效! ")
		return
	}
	customer := &this.Customers[id]
	var age int
	var name, email string
	fmt.Printf("姓名(%v): ", customer.Name)
	fmt.Scanln(&name)
	if name != "\r" {
		customer.Name = name
	}
	fmt.Printf("年龄(%v): ", customer.Age)
	fmt.Scanln(&age)
	if string(age) != "\r" {
		customer.Age = age
	}
	fmt.Printf("邮箱(%v): ", customer.Email)
	fmt.Scanln(&email)
	if email != "\r" {
		customer.Email = email
	}
}
