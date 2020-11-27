package view

import (
	"fmt"
	"客户信息管理软件02/controller"
)

type CustomerView struct {
	CustomersController controller.CustomersController
	loop                bool
}

func NewCustomerView() CustomerView {
	return CustomerView{
		loop: true,
	}
}

func (this CustomerView) MainView() {
	for {
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println("          1. 查看用户列表")
		fmt.Println("          2. 添加用户")
		fmt.Println("          3. 修改用户")
		fmt.Println("          4. 删除用户")
		fmt.Println("          5. 退出软件")
		fmt.Print("请输入(1-5): ")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			this.showList()
		case "2":
			this.add()
		case "3":
			this.change()
		case "4":
			this.delete()
		case "5":
			this.exit()
		}
		if !this.loop{
			break
		}
	}
	fmt.Println("再见！")
}

func (this *CustomerView) add() {
	var age int
	var name, email string
	fmt.Print("请输入用户姓名: ")
	fmt.Scanln(&name)
	fmt.Print("请输入用户年龄: ")
	fmt.Scanln(&age)
	fmt.Print("请输入用户邮箱: ")
	fmt.Scanln(&email)
	this.CustomersController.Add(name, age, email)
}

func (this *CustomerView) showList() {
	customers := this.CustomersController.ShowList()
	if len(customers) == 0 {
		fmt.Println("用户列表为空！")
	} else {
		for _, v := range customers {
			fmt.Println(v)
		}
	}
}

func (this *CustomerView) delete() {
	fmt.Print("请输入要删除的用户Id: ")
	var id int
	fmt.Scanln(&id)
	this.CustomersController.Delete(id)
}

func (this *CustomerView) change() {
	fmt.Print("请输入要修改的用户Id: ")
	var id int
	fmt.Scanln(&id)
	this.CustomersController.Change(id)
}

func (this *CustomerView) exit() {
	fmt.Print("请输入(y/n): ")
	var flag string
	for {
		fmt.Scanln(&flag)
		if flag == "y" || flag == "n" {
			break
		} else {
			fmt.Print("输入有误，请重新输入(y/n): ")
		}
	}
	if flag == "y" {
		this.loop = false
	}
}
