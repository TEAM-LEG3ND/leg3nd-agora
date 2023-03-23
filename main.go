package main

import (
	"context"
	"fmt"
	"leg3nd-agora/ent"
	"leg3nd-agora/ent/account"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateAccount(ctx context.Context, client *ent.Client) (*ent.Account, error) {
	ac, err := client.Account.
		Create().
		SetEmail("limdongyoung0@gmail.com").
		SetNickname("빠른거북이").
		SetFullName("임동영").
		SetOauthProvider("google").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create ac: %w", err)
	}
	log.Println("account was created: ", ac)
	return ac, nil
}

func QueryAccount(ctx context.Context, client *ent.Client) (*ent.Account, error) {
	ac, err := client.Account.Query().Where(account.Email("limdongyoung0@gmail.com")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying account: %w", err)
	}
	log.Println("account queried: ", ac)
	return ac, nil
}

func main() {
	client, err := ent.Open("mysql", "root:test123@tcp(localhost:3306)/agora?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	QueryAccount(context.Background(), client)
}
