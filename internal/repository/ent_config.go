package repository

import (
	"context"
	"leg3nd-agora/ent"
	"log"
)

func ProvideClient() (*ent.Client, func(), error) {
	client, err := ent.Open("mysql", "root:test123@tcp(localhost:3306)/agora?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	cleanup := func() {
		if err := client.Close(); err != nil {
			log.Fatalln(err)
		}
	}

	return client, cleanup, nil
}
