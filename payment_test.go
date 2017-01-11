package paypal

import (
	"fmt"
	"testing"
	//"github.com/smartwalle/go-uuid/uuid"
)

var paypal *PayPal

func getPayPal() *PayPal {
	if paypal == nil {
		paypal = New("AS8XSa9JrOJ3rf0kxVqCgRLIlMpgaKhLTShpYxISysR1VpnN6AMLfrvj-upOMuNkXdb9bTIzsFH4umB5", "ECA3_usif2DUgGxgcBTddOKgg2rbjUT7J3B3-Ud9z9y54AK9mYTDDFyadmMLSo1QOiO2rci99FSq1PbZ", false)
	}
	return paypal
}

//func TestPayPal_CreatePayment(t *testing.T) {
//	var p = &Payment{}
//	p.Intent = K_PAYPAL_PAYMENT_INTENT_SALE
//	p.Payer = &Payer{}
//	p.Payer.PaymentMethod = "paypal"
//	p.RedirectURLs = &RedirectURLs{}
//	p.RedirectURLs.CancelURL = "http://www.baidu.com"
//	p.RedirectURLs.ReturnURL = "http://127.0.0.1:9001/paypal"
//	var ti = &Transaction{}
//	p.Transactions = []*Transaction{ti}
//
//	ti.Amount = &Amount{}
//	ti.Amount.Total = "30.11"
//	ti.Amount.Currency = "USD"
//	ti.Amount.Details = &AmountDetails{}
//	ti.Amount.Details.Subtotal = "30.00"
//	ti.Amount.Details.Tax = "0.07"
//	ti.Amount.Details.Shipping = "0.03"
//	ti.Amount.Details.HandlingFee = "1.00"
//	ti.Amount.Details.ShippingDiscount = "-1.00"
//	ti.Amount.Details.Insurance = "0.01"
//
//	ti.Description = "This is the payment transaction description."
//	ti.Custom = "EBAY_EMS_90048630024435"
//	ti.InvoiceNumber = uuid.New()
//
//	ti.PaymentOptions = &PaymentOptions{}
//	ti.PaymentOptions.AllowedPaymentMethod = "INSTANT_FUNDING_SOURCE"
//	ti.SoftDescriptor = "ECHI5786786"
//
//	ti.ItemList = &ItemList{}
//	ti.ItemList.ShippingAddress = &ShippingAddress{}
//	ti.ItemList.ShippingAddress.RecipientName = "Hello World"
//	ti.ItemList.ShippingAddress.Line1 = "4thFloor"
//	ti.ItemList.ShippingAddress.Line2 = "unit#34"
//	ti.ItemList.ShippingAddress.City = "SAn Jose"
//	ti.ItemList.ShippingAddress.CountryCode = "US"
//	ti.ItemList.ShippingAddress.PostalCode = "95131"
//	ti.ItemList.ShippingAddress.Phone = "011862212345678"
//	ti.ItemList.ShippingAddress.State = "CA"
//
//	var i1, i2 = &Item{}, &Item{}
//	ti.ItemList.Items = []*Item{i1, i2}
//
//	i1.Name = "hat"
//	i1.Description = "Brown color hat"
//	i1.Quantity = "5"
//	i1.Price = "3"
//	i1.Tax = "0.01"
//	i1.SKU = "1"
//	i1.Currency = "USD"
//
//	i2.Name = "handbag"
//	i2.Description = "Black color hand bag"
//	i2.Quantity = "1"
//	i2.Price = "15"
//	i2.Tax = "0.02"
//	i2.SKU = "product34"
//	i2.Currency = "USD"
//
//	p.NoteToPayer = "Contact us for any questions on your order."
//
//	fmt.Println(getPayPal().CreatePayment(p))
//}

//func TestPayPal_GetPaymentList(t *testing.T) {
//	var p = &PaymentListParam{}
//	fmt.Println(getPayPal().GetPaymentList(p))
//}
//
//func TestPayPal_GetPaymentDetails(t *testing.T) {
//	fmt.Println(getPayPal().GetPaymentDetails("PAY-1SJ16214TY566804MLB26S5I"))
//}
//
//func TestPayPal_ExecuteApprovedPayment(t *testing.T) {
//	fmt.Println(getPayPal().ExecuteApprovedPayment("PAY-0E809244MY2080201LB26LWI", "XV9HF9K25FB38"))
//}
//
//func TestPayPal_GetSaleDetails(t *testing.T) {
//	fmt.Println(getPayPal().GetSaleDetails("84E33686SW065691F"))
//}

//func TestPayPal_RefundSale(t *testing.T) {
//	fmt.Println(getPayPal().RefundSale("5SW33389HH3038001", "48787589677", "30.11", "USD"))
//}

func TestPayPal_GetRefundDetails(t *testing.T) {
	fmt.Println(getPayPal().GetRefundDetails("7AG891841V2503230"))
}