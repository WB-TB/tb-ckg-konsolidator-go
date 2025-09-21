package api

import (
	"fhir-sirs/app/config"
	dbMongo "fhir-sirs/app/db/mongo"
	"fhir-sirs/app/server"

	// exampleRoutesV1 "fhir-sirs/pkg/api/v1/rlreport/routes"
	ckgTBRepo "fhir-sirs/pkg/api/v1/ckg_tb/repository/impl"
	ckgTBRoutes "fhir-sirs/pkg/api/v1/ckg_tb/routes"
	ckgTBUC "fhir-sirs/pkg/api/v1/ckg_tb/usecase"
	tunjanganRepo "fhir-sirs/pkg/api/v1/data_tunjangan_khusus/repository/impl"
	tunjanganRoutes "fhir-sirs/pkg/api/v1/data_tunjangan_khusus/routes"
	tunjanganUC "fhir-sirs/pkg/api/v1/data_tunjangan_khusus/usecase"
	rlRoutes "fhir-sirs/pkg/api/v1/rlreport/routes"
	rlUC "fhir-sirs/pkg/api/v1/rlreport/usecase"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/time/rate"

	apikey "fhir-sirs/app/middleware/apikey"
)

// function to establish mongodb connection
func initMongoDatabase() *mongo.Client {
	dbConn, err := dbMongo.NewConnectionDB()
	if err != nil {
		log.Fatal(err)
	}
	return dbConn
}

// function as the maintenance switcher
func maintenanceMode(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if config.GetConfig().UnderMaintenance {
			return server.ResponseStatusServiceUnavailable(c, "The service is under maintenance", nil, nil, nil)
		}
		return next(c)
	}
}

// The entrypoint function
func Start() {

	// Please comment this part when deploying into dev, staging or prod environment
	//err := godotenv.Load()
	//if err != nil {
	//	fmt.Print("unable to load .env file: ", err)
	//}
	//

	// NOT USED FOR NOW
	var (
		mongoConn *mongo.Client
	)
	// Init mongodb connection, just call one time to prevent connection flood
	mongoConn = initMongoDatabase()

	// Init the Echo Framework
	e := server.InitEcho()

	// Init for routing group
	exampleV1 := e.Group("/v1")
	exampleV1.Use(maintenanceMode)
	exampleV1.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(config.GetConfig().RateLimit))))
	// exampleRoutesV1.NewHTTP(exampleUCV1.Initialize(mongoConn), exampleV1)

	rlV1 := e.Group("/v1/rlreport")
	rlV1.Use(apikey.APIKeyAuthMiddleware())
	rlRoutes.NewHTTP(rlUC.NewRLReport(), rlV1)

	tunjanganV1 := e.Group("/v1/tunjangankhusus")
	tunjanganV1.Use(apikey.APIKeyAuthMiddleware())

	tunjanganUC := tunjanganUC.NewDataTunjanganKhusus(
		tunjanganRepo.NewTunjanganKhususRepo(), // ✅ sudah sesuai
		mongoConn,
	)
	tunjanganRoutes.NewHTTP(tunjanganUC, tunjanganV1)

	ckgTB := e.Group("/v1/ckg/tb")
	ckgTB.Use(apikey.APIKeyAuthMiddleware())
	ckgTBUC := ckgTBUC.NewDataCKGTB(
		ckgTBRepo.NewCKGTBRepo(), // ✅ sudah sesuai
		mongoConn,
	)
	ckgTBRoutes.NewHTTP(ckgTBUC, ckgTB)

	// Start the server
	server.Start(e)
}
