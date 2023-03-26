package repository

import (
	"context"
	"fmt"
	"leg3nd-agora/ent"
	"leg3nd-agora/internal/util"
	"log"
)

func ProvideClient() (*ent.Client, func(), error) {
	mysqlUser, err := util.Config("MYSQL_USER")
	if err != nil {
		return nil, nil, err
	}
	mysqlPassword, err := util.Config("MYSQL_PASSWORD")
	if err != nil {
		return nil, nil, err
	}
	mysqlHost, err := util.Config("MYSQL_HOST")
	if err != nil {
		return nil, nil, err
	}
	mysqlPort, err := util.Config("MYSQL_PORT")
	if err != nil {
		return nil, nil, err
	}
	mysqlDatabase, err := util.Config("MYSQL_DATABASE")
	if err != nil {
		return nil, nil, err
	}

	dataSourceUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase,
	)
	client, err := ent.Open("mysql", dataSourceUrl)
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
