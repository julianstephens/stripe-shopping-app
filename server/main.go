package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	stripe.Key = os.Getenv("STRIPE_SECRECT_KEY")

	e := echo.New()

	// Routes
	e.GET("/checkout/success", func(c echo.Context) error {
		return c.String(http.StatusOK, "Payment successful!")
	})
	e.GET("/checkout/cancel", func(c echo.Context) error {
		return c.String(http.StatusOK, "Payment cancelled.")
	})
	e.GET("/create-session", createSession)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

type createSessionResponse struct {
	SessionID string `json:"id"`
}

func createSession(c echo.Context) error {
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("/checkout/success"),
		CancelURL:  stripe.String("/checkout/cancel"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(string(stripe.CurrencyUSD)),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("T-shirt"),
					},
					UnitAmount: stripe.Int64(2000),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	}

	session, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	data := createSessionResponse{
		SessionID: session.ID,
	}
	resp, _ := json.Marshal(data)

	return c.JSON(http.StatusOK, resp)
}
