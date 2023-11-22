package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"onlineshop/pkg/models"

	razorpay "github.com/razorpay/razorpay-go"
)

func CreateTransactionRazorPay(amount int, recieptNot string) {
	var client = razorpay.NewClient("rzp_test_pDqm2g3OXuOBnj", "7tzCkIsDyl2bgnBZ8gvlh955")

	// data := map[string]interface{}{
	// 	"name":          "Gaurav Kumar",
	// 	"contact":       9123456780,
	// 	"email":         "gaurav.kumar@example.com",
	// 	"fail_existing": 0,
	// 	"notes": map[string]interface{}{
	// 		"notes_key_1": "Tea, Earl Grey, Hot",
	// 		"notes_key_2": "Tea, Earl Grey… decaf.",
	// 	},
	// }

	order := map[string]interface{}{
		"amount":   50000,
		"currency": "INR",
		// "receipt":         reciepy,
		"partial_payment": false,
	}
	body, err := client.Order.Create(order, nil)
	if err != nil {
		fmt.Println("error is ", err)
	}
	fmt.Println("order  ==== ", body)
}

func ProcessPaymentHelper(w http.ResponseWriter, r *http.Request, order models.Order) (map[string]interface{}, error) {

	client := razorpay.NewClient("rzp_test_pDqm2g3OXuOBnj", "7tzCkIsDyl2bgnBZ8gvlh955")
	reciept := fmt.Sprintf("%d", order.ID)

	fmt.Println("order.total is ", order.Total)
	razorpaydata := map[string]interface{}{

		"amount":   order.Total * 100,
		"currency": "INR",
		"receipt":  reciept,
	}
	body, err := client.Order.Create(razorpaydata, nil)
	if err != nil {
		// return models.Payment{}, err
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return body, nil
}

func RazorPaymentVerification(sign, orderId, paymentId string) error {
	signature := sign
	secret := "7tzCkIsDyl2bgnBZ8gvlh955"
	data := orderId + "|" + paymentId

	h := hmac.New(sha256.New, []byte(secret))

	_, err := h.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
		return errors.New("Payment failed")
	} else {
		return nil
	}
}
