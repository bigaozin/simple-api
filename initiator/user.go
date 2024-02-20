package initiator

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"errors"
)

// const (
// 	postgresURL = "postgresql://%s:%s@%s/iDevoid-db?sslmode=disable"

// 	domain = "user"
// )

// User initializes the domain user
func User(testInit bool) {
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// dbURL := fmt.Sprintf(postgresURL, dbUser, dbPass, dbHost)
	// psql := cptx.Initialize(dbURL, dbURL, domain)
	// postgresDB, postgresTX := psql.Open()

	// databaseUser := persistence.UserInit(postgresDB)
	// databaseProfile := persistence.ProfileInit(postgresDB)

	// repo := repository.UserInit(encryptKey)

	// usecase := user.Initialize(postgresTX, repo, databaseUser, databaseProfile)

	// handler := rest.UserInit(usecase)
	// router := routing.UserRouting(handler)

	port := os.Getenv("HOST_PORT")
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	fmt.Printf("Server starting on %s in PORT: %s", allowedOrigins, port)
	// server := routers.Initialize(port, allowedOrigins, router, domain)

	if testInit {
		logrus.Info("Initialize test mode Finished!")
		os.Exit(0)
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		errors.New("Error on ListenAndServer\n" + err.Error())
	}
}
