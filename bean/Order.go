package bean

import "fmt"

type Order struct {
	acctId   int
	acctUser string
	currPair string
	dealQty  float64
	ctrQty   float64
}

func (order Order) New(acctId int, acctUser string, currPair string, dealQty float64, ctrQty float64) Order {
	order.acctId = acctId
	order.acctUser = acctUser
	order.currPair = currPair
	order.dealQty = dealQty
	order.ctrQty = ctrQty
	return order
}
func (order Order) GetAcctId() int {
	return order.acctId
}
func (order Order) GetAcctUser() string {
	return order.acctUser
}
func (order Order) GetCurrPair() string {
	return order.currPair
}
func (order Order) GetDealQty() float64 {
	return order.dealQty
}
func (order Order) GetCtrQty() float64 {
	return order.ctrQty
}
func (order Order) PrintOrderDetail() {
	fmt.Println("acctId:", order.acctId, "acctUser: ", order.acctUser, "dealQty:", order.dealQty, "ctrQty:", order.ctrQty)
}
