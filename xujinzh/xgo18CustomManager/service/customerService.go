package service

import (
	"xgo18/model"
)

// 对 Customer 的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
	// 当前切片含有多少个客户
	customerNum int
}

func FactoryCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.FactoryCustomer(1, "张三", "男", 20, "112", "zs@gmail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

// 添加客户到 customers 切片
func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.customerNum += 1
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}
