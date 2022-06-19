package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"wiiki_server/common/config"
	"wiiki_server/common/wiikierr"
	"wiiki_server/domain/service"
	"wiiki_server/domain/usecase"
	"wiiki_server/infra/graph"
	"wiiki_server/infra/graph/generated"
	"wiiki_server/infra/graph/middleware"
	"wiiki_server/infra/logger"
	"wiiki_server/infra/postgres"
	"wiiki_server/infra/postgres/psglrepository"

	"github.com/go-chi/chi/v5"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

var (
	configPath = flag.String("f", "config/local.toml", "config file path")
)

func main() {

	flag.Parse()

	// setup
	log.Println("read config file")
	conf, err := config.New(*configPath)
	wiikierr.MustNil(err)

	// logger
	logger, err := logger.New(conf)
	wiikierr.MustNil(err)

	// postgres
	log.Println("new postgres")
	postgresEngine, err := postgres.New(conf.Postgres)
	wiikierr.MustNil(err)

	// repository
	todoRepository := psglrepository.NewTodo()
	userRepository := psglrepository.NewUser()

	// hash
	hashService := service.NewHash(10)

	// usecase
	todoUsecase := usecase.NewTodo(todoRepository)
	userUsecase := usecase.NewUser(userRepository, hashService)

	// resolver
	resolver := &graph.Resolver{
		PostgresEngine: postgresEngine,
		TodoUsecase:    todoUsecase,
		UserUsecase:    userUsecase,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	// middleware
	errHandlingMiddleware := middleware.NewErrorHandling(logger, conf)
	authMiddleware := middleware.NewAuth()
	// transactionMiddleware := middleware.NewTransactionMiddleware(postgresEngine)

	r := chi.NewRouter()

	r.Get("/", playground.Handler("GraphQL playground", "/query"))
	r.With(errHandlingMiddleware.ErrorHandling(), authMiddleware.Auth()).Post("/query", srv.ServeHTTP)

	log.Println("======== start wiiki server ==========")
	log.Printf("listen : %s\n", conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), r))

}
