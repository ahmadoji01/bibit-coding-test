package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	_logRepo "go_bibit_test/log/repository/mysql"
	_movieGRPCHandler "go_bibit_test/movie/delivery/grpc"
	_movieHttpDelivery "go_bibit_test/movie/delivery/http"
	_movieHttpDeliveryMiddleware "go_bibit_test/movie/delivery/http/middleware"
	_movieUcase "go_bibit_test/movie/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _movieHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	logRepo := _logRepo.NewMysqlLogRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _movieUcase.NewMovieUsecase(logRepo, timeoutContext)
	_movieHttpDelivery.NewMovieHandler(e, au)

	//log.Fatal(e.Start(viper.GetString("server.address")))

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf(err.Error())
	}

	grpcServer := grpc.NewServer()
	_movieGRPCHandler.NewMovieServerGRPC(grpcServer, au)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to open gRPC Server")
	}
}
