package main

import "fmt"

type PaymnetProvider interface {
	pay(provider string, amount int) string
}

type Strpie struct {
	name     string
	stripeId int
}

type RazorPay struct {
	name       string
	razorPayid int
}

func (s *Strpie) pay(provider string, amount int) string {
	fmt.Println("Paying using ", provider)
	s.stripeId = amount
	return "Done"
}

func (r *RazorPay) pay(provider string, amount int) string {
	fmt.Println("Paying using ", provider)
	r.razorPayid = amount
	return "Done form razor pay"
}

func wrapper(anyy PaymnetProvider, provider string, amount int) {
	anyy.pay(provider, amount)
}

func getProviderName(name string, id int) PaymnetProvider {
	switch name {
	case "stripe":
		val := Strpie{name: name, stripeId: id}
		return &val
	case "razorpay":
		vl := RazorPay{name: name, razorPayid: id}
		return &vl
	default:
		panic("No provider selected")
	}
}

func main() {
	name := getProviderName("stripe", 1000)
	wrapper(name, "stripe", 2300)

}
