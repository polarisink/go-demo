package main

import (
	"fmt"
	"strconv"
)

//声明接口
type Good interface {
	settleAccount() int
	orderInfo() string
}

//两个结构体
type Pad struct {
	name     string
	quantity int
	price    int
}
type FreeGift struct {
	name     string
	quantity int
	price    int
}

//实现接口
func (phone Pad) settleAccount() int {
	return phone.quantity * phone.price
}
func (phone Pad) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + ",合计：" + strconv.Itoa(phone.settleAccount()) + "元"
}
func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + ",合计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func calculatePrice(goods[] Good)int{
	var allPrice int
	for _,good :=range goods{
		fmt.Println(good.orderInfo())
		allPrice+=good.settleAccount()
	}
	return allPrice
}

func main() {
	iPad := Pad{
		name:     "iPad Pro2021",
		quantity: 2,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 3,
		price:    200,
	}
	goods := []Good{iPad, earphones}
	allPrice:=calculatePrice(goods);
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}
