//Copyright (C) Subhajit DasGupta 2021

// Package main is an example REST server for persisted structs.
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"sync"

	"github.com/lf8r/example/main/generated/person"
	"github.com/lf8r/example/main/generated/persondao"
	"github.com/lf8r/example/main/generated/resthandler"

	"github.com/lf8r/example/main/generated/book"
	"github.com/lf8r/example/main/generated/bookdao"

	"github.com/lf8r/dbgen/pkg/dbs"

	// nolint
	_ "github.com/lf8r/dbgen/pkg/encoding" // to register the JSON codec for gRPC.
	log "github.com/lf8r/slog/pkg/slog"
	grpc "google.golang.org/grpc"
)

const (
	// defaultGRPCPort is the default REST endpoint.
	defaultRESTPort = 8080

	// defaultGRPCPort is the default gRPC endpoint.
	defaultGRPCPort = 9090
)

// main runs the main program exposing listeners for generated services. It
// expects two optional parameters, viz. "port" and "grpcport", specifying the
// listening ports for REST and gRPC endpoints. Unless given "port" defaults to
// 8080 and "grpcport" defaults to 9090.
func main() {
	log.NewSloggerWithWriter(os.Args[0], log.ErrorLevel, os.Stdout)

	// Profiling - Only one type of profiling must be enabled at any time.
	// Options are profile.CPUProfile, profile.MemProfile and
	// profile.GoroutineProfile.
	f, err := os.Create("cpu.pprof")
	if err != nil {
		log.Errorf("Could not create CPU profile file")

		return
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Errorf("Could not start CPU profile")

		return
	}

	defer pprof.StopCPUProfile()

	// Listening ports.
	restPortPtr := flag.Int("port", defaultRESTPort, "REST listening port")
	grpcPortPtr := flag.Int("grpcport", defaultGRPCPort, "gRPC listening port")

	flag.Parse()

	restPort := *restPortPtr
	grpcPort := *grpcPortPtr

	// Initialize the DB client factory from DB connection information. For a
	// Posgres DB, we need host, port, user, password, dbname and sslenable
	// parameters.
	db, err := getDB()
	if err != nil {
		log.Infof("Could not initialize DB: %v", err)

		return
	}

	persondao.Db = db

	bookdao.Db = db

	// abortListenAndServer asks the listen and serve function (below) to quit.
	abortListenAndServer := make(chan struct{})

	// listenAndServeIsDone informs us that listenAndServer is done after it's
	// been aborted.
	listenAndServeIsDone := make(chan struct{}, 1)

	// Handle Ctrl-C (interrupts) and kill signals properly by closing the
	// listen and serve function (below).
	ctrlC := make(chan os.Signal, 1)
	signal.Notify(ctrlC, os.Interrupt, os.Kill)

	go func() {
		for sig := range ctrlC {
			switch sig {
			case os.Interrupt, os.Kill:
				log.Infof("%v", sig)
				close(abortListenAndServer)
				<-listenAndServeIsDone

				return
			}
		}
	}()

	listenAndServe(restPort, grpcPort, abortListenAndServer, listenAndServeIsDone, nil)
}

// listenAndServe starts gRPC and REST listeners and serves services for the
// data types. If the abort channel is not nil and it becomes readable, closes
// the done channel and returns.
//
// The abort channel is typically closed by the caller when it's shutting down,
// such as when due to an interrupt or kill signal.
func listenAndServe(restPort, grpcPort int, abort <-chan struct{}, done chan struct{}, started *sync.Mutex) {
	// Start the GRPC listener.
	startedCustomGRPCPort := sync.Mutex{}
	startedCustomGRPCPort.Lock()
	go listenAndServeGRPCEndpoint(grpcPort+1, &startedCustomGRPCPort)
	log.Infof("Listen and serve gRPC/JSON on port: %d", (grpcPort + 1))

	// Start the REST listener.
	startedHTTPPort := sync.Mutex{}
	startedHTTPPort.Lock()
	go listenAndServeRESTEndpoint(restPort, &startedHTTPPort)
	log.Infof("Listen and serve REST/HTTP on port: %d", restPort)

	// Wait for both listeners to start.
	startedCustomGRPCPort.Lock()
	startedHTTPPort.Lock()

	if started != nil {
		started.Unlock()
	}

	for {
		select {
		case <-abort:
			if done != nil {
				log.Infof("Closing closed channel.")
				close(done)

				return
			}
		}
	}
}

// getDB returns an sql.DB to connect to a DB instance used to persist books and
// persons, by constructing one from environment variables:
//  DB_HOST
//  DB_PORT
//  DB_DATABASE
//  DB_USER
//  DB_PASSWORD
//  DB_SSL_MODE
func getDB() (*sql.DB, error) {
	return dbs.OpenDB()
}

// listenAndServeGRPCEndpoint serves a gRPC endpoint over the given port. The
// endpoint serves an API to act upon the bound types for which code has been generated.
func listenAndServeGRPCEndpoint(port int, started *sync.Mutex) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	person.RegisterServiceServer(grpcServer, person.Server{})

	book.RegisterServiceServer(grpcServer, book.Server{})

	if started != nil {
		started.Unlock()
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// listenAndServeRESTEndpoint serves a REST endpoint over the given port. The endpoint
// serves an API to act upon the bound types for which code has been generated.
func listenAndServeRESTEndpoint(port int, started *sync.Mutex) {
	address := fmt.Sprintf(":%d", port)

	http.HandleFunc("/", resthandler.Handler)
	log.Infof("Listening on port: %d", port)

	if started != nil {
		started.Unlock()
	}

	log.Fatalf("%v", http.ListenAndServe(address, nil))
}
