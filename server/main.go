package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/components/appContext"
	"server/middleware"
	"server/routes"
	"server/smtp"

	"server/storage/servicestore"
	"server/storage/userstore"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func main() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file" + err.Error())
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	appCtx := appContext.InitAppContext(uri)
	defer appCtx.CloseDB()

	userStore := userstore.NewMongoStore(appCtx.DB)
	serviceStore := servicestore.NewMongoStore(appCtx.DB)
	smtpStore := smtp.NewMongoStore(appCtx.DB)

	v1 := http.NewServeMux()
	v1.Handle("/v1/user/", http.StripPrefix("/v1", routes.UserRoutes(userStore)))
	v1.Handle("/v1/service/", http.StripPrefix("/v1", routes.ServiceRoutes(serviceStore, userStore)))
	v1.Handle("/v1/bill/", http.StripPrefix("/v1", routes.BillRoutes(userStore)))

	//middle ware stack
	stack := middleware.CreateStack(

		middleware.Logging,
		// middleware.VerifyJWT,
	)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Change "*" to your allowed origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(stack(v1))

	server := http.Server{
		Addr:    ":8080",
		Handler: corsHandler,
	}
	fmt.Println("Server is running on port 8080")

	//Calculate the bill price of each customer every month
	c := cron.New()
	c.AddFunc("@every monthly", func() {
		smtp.CalculateUserBill(smtpStore)
		fmt.Println("User bill calculated")
	})
	c.Start()

	_ = server.ListenAndServe()

}
