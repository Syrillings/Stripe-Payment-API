# Stripe-Payment-API
This project is a simple backend server implemented in Go for handling Stripe payment intents. It provides endpoints to create a payment intent and check the server's health status

Prerequisites
Go (version 1.16 or later)
Stripe account and API keys
Insomnia (or any other API testing tool)

Endpoints
1. Create Payment Intent
URL: /create-payment-intent
Method: POST
Description: Creates a new payment intent and returns the client secret.
2. Health Check
URL: /health
Method: GET
Description: Checks if the server is active.
Testing with Insomnia
Install Insomnia: Download and install Insomnia from here.
https://insomnia.rest/download

Create a new request:

Open Insomnia and click on Create > Request.
Name your request (e.g., Create Payment Intent).
Configure the request:

Method: POST
URL: http://localhost:8080/create-payment-intent
Body: Select JSON and add content in the format below:

{
    "product_id": "example_product_id",
    "first_name": "John",
    "last_name": "Doe",
    "address_1": "123 Main St",
    "address_2": "Apt 4B",
    "city": "Anytown",
    "state": "CA",
    "zip": 12345,
    "country": "USA"
}

Send the request:

Click on Send.
You should receive a response with the clientSecret.
Health Check:

Create another request named Health Check.
Method: GET
URL: http://localhost:8080/health
Click on Send.
You should receive a response with the message "Server is active".

Head to your stripe dashboard and check payments. You should see the requested items 
Conclusion
This project demonstrates how to set up a basic Stripe payment intent server in Go and test it using Insomnia. Feel free to extend the functionality as needed.

