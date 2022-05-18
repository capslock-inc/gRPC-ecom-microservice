package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	postgresclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/postgresclient"
	core "github.com/capslock-inc/gprc-demo/Servers/postgresclient/Core"
	model "github.com/capslock-inc/gprc-demo/Servers/postgresclient/Model"
	_ "github.com/lib/pq"
)

// althaf

const (
	host     = "localhost"
	user     = "admin"
	password = "mypassword"
	dbname   = "ecom"
)

func PostgresINIT(port int) (*sql.DB, *gorm.DB) {
	var connstring string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psql, err := sql.Open("postgres", connstring)
	if err != nil {
		log.Fatalf("error connecting postgres: %v", err)
	} else {
		log.Printf("connection established")
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: psql,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm :: error establishing connection : %v", err)
	} else {
		log.Println("gorm :: connection established")
	}

	return psql, db
}

func main() {
	p := flag.Int("Postgresport", 5434, "postgres port")
	flag.Parse()

	dbconn, ormdb := PostgresINIT(*p)

	// create table if not exist
	ormdb.AutoMigrate(&model.Product{})
	ormdb.AutoMigrate(&model.User{})

	serverport := flag.Int("port", 8401, "server port")
	port := fmt.Sprintf(":%d", *serverport)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening to the port:%s :: %v", port, err)
	}
	grpcserver := grpc.NewServer()
	postgresclientmodel.RegisterPostgresClientServiceServer(grpcserver, &core.PostgresClientServer{
		Db: ormdb,
	})
	log.Printf("listen to %s", port)
	if err := grpcserver.Serve(listen); err != nil {
		log.Fatalf("error serving grpc server :: %v", err)
	}

	defer dbconn.Close()

}
