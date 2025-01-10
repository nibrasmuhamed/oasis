package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Balance struct {
	VersionId       int64             `json:"versionId"`
	BalanceID       string            `json:"balanceId"`
	CustomerID      string            `json:"customerId"`
	CreatedAt       string            `json:"createdAt"`
	UpdatedAt       string            `json:"updatedAt"`
	CreatedBy       string            `json:"createdBy"`
	UpdatedBy       string            `json:"UpdatedBy"`
	BalanceItemSpec []BalanceItemSpec `json:"balanceItemSpec,omitempty"`
	OperatorId      string            `json:"operatorId"`
}

type BalanceItemSpec struct {
	SubscriptionID         string        `json:"subscriptionId"`
	SubscriptionInternalId string        `json:"subscriptionInternalId"`
	OfferCode              string        `json:"offerCode,omitempty"`
	Id                     string        `json:"id,omitempty"`
	ResourceID             int           `json:"resourceID,omitempty"`
	ResourceName           string        `json:"resourceName,omitempty"`
	CreditFloor            float64       `json:"creditFloor"`
	CreditLimit            float64       `json:"creditLimit"`
	BalanceItem            []BalanceItem `json:"balanceItem,omitempty"`
}

type BalanceItem struct {
	OriginalBalance    float64              `json:"originalBalance,omitempty"`
	ConsumedBalance    float64              `json:"consumedBalance,omitempty"`
	ValidFrom          time.Time            `json:"validFrom,omitempty"`
	ValidTo            time.Time            `json:"validTo,omitempty"`
	BalanceReservation []BalanceReservation `json:"balanceReservation,omitempty"`
}

type BalanceReservation struct {
	ReservationID   string    `json:"reservationId,omitempty"`
	ReservedBalance float64   `json:"reservedBalance,omitempty"`
	ExpirationTime  time.Time `json:"expirationTime,omitempty"`
}

func fetch() {
	// MongoDB connection settings
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Database and collection
	db := client.Database("balance")
	collection := db.Collection("Balances")

	// Customer IDs to fetch
	customerIDs := []string{}
	num := 10000
	for i := 1; i <= num; i++ {
		k := strconv.Itoa(i)
		customerIDs = append(customerIDs, "cust"+k)
	}
	var balances []Balance

	for _, customerID := range customerIDs {
		// Filter by customer ID
		filter := bson.M{"customerid": customerID}
		var balance Balance
		// Fetch documents
		err = collection.FindOne(context.Background(), filter).Decode(&balance)
		if err != nil {
			log.Fatal(err)
		}

		balances = append(balances, balance)
	}

	// Write to JSON file
	data, err := json.MarshalIndent(balances, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling data:", err)
	}

	err = os.WriteFile("x.json", data, 0644)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Println("Data written to balances.json")
}
