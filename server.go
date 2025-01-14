package main

import (
	"fmt"
	"log"
	"bytes"
	"io"
  "encoding/json"
	"net/http"
  "github.com/stripe/stripe-go/v81"
  "github.com/stripe/stripe-go/v81/paymentintent"
)

func main() {
  stripe.Key = "sk_test_51Qg5jXBc1QiyXIlBgTrcSXibGnNvhxnoNc5fylHMDERXOMdhfiUJ9UObDXNVL3dOqxZ48nqwCE2HIiBcJs6kANNm00FcBd4qv1"
	
  fmt.Println("Backend Initialized")

	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	log.Println("Listening on Localhost 4242....")
	var err error = http.ListenAndServe("localhost:4242", nil)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Something bad happened")
	}

}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {

	if request.Method != ("POST") {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Request Method Authorized")

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       int `json:"zip"`
		Country   string `json:"country"`
	}

err := json.NewDecoder(request.Body).Decode(&req)
if err!=nil{
  fmt.Println(err)
  http.Error(writer, err.Error(), http.StatusInternalServerError)
  //log.Println(err)
}

params := &stripe.PaymentIntentParams{
Amount: stripe.Int64(calculateOrderAmmount(req.ProductId)),
Currency: stripe.String(string(stripe.CurrencyNGN)),
AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
 Enabled: stripe.Bool(true),
},
  }

paymentIntent, err :=paymentintent.New(params)
if err != nil{
  log.Println(err)
  http.Error(writer, err.Error(), http.StatusInternalServerError)
}
fmt.Println(paymentIntent.ClientSecret)

var response struct{
  ClientSecret string `json:"clientSecret"`
}

response.ClientSecret = paymentIntent.ClientSecret

var buf bytes.Buffer
err = json.NewEncoder(&buf).Encode(response)
if err !=nil{
  log.Println(err)
  http.Error(writer, err.Error(), http.StatusInternalServerError)
  }

  writer.Header().Set("Content-Type", "application/json")
  _, err = io.Copy(writer, &buf)
  if err !=nil{
    log.Println(err)
  }
}


func handleHealth(writer http.ResponseWriter, request *http.Request) {
	response := []byte("Server is active")

	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)

  
	}
}

func calculateOrderAmmount(ProductId string) int64 { 
  switch ProductId {
  case"Forever Pants":
    return 260000
  case"Forever Shirt":
    return 155000
  case"Forever Shorts":
    return 300000
  }   
  return 0
}