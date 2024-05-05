package main

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/kelseyhightower/envconfig"
	"github.com/mateusrlopez/url-shortener-server/clients"
	"github.com/mateusrlopez/url-shortener-server/http"
	"github.com/mateusrlopez/url-shortener-server/http/handlers"
	"github.com/mateusrlopez/url-shortener-server/repositories"
	"github.com/mateusrlopez/url-shortener-server/services"
	"github.com/mateusrlopez/url-shortener-server/settings"
	"github.com/mateusrlopez/url-shortener-server/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

var s settings.Settings

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := envconfig.Process("", &s); err != nil {
		log.Fatal().Err(err).Msg("failed to load the environment variables")
	}
}

func main() {
	postgresClient, err := clients.NewPostgresClient(s.Postgres.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to the Postgres instance")
	}

	defer func() {
		if err = clients.DisconnectPostgresClient(postgresClient); err != nil {
			log.Fatal().Err(err).Msg("failed to disconnect to the Postgres instance")
		}
	}()

	snowflakeNode, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize snowflake node")
	}

	snowflakeIdGenerator := utils.NewSnowflakeIdGenerator(snowflakeNode)
	base62Encryptor := utils.NewBase62Encryptor()

	recordsRepository := repositories.NewRecords(postgresClient)
	recordsService := services.NewRecords(recordsRepository, snowflakeIdGenerator, base62Encryptor)
	recordsHandler := handlers.NewRecords(recordsService)

	router := http.NewRouter(s.Router.Context, s.CORS.AllowedOrigins, recordsHandler)

	server := http.NewServer(router, s.Server.Port, s.Server.ReadTimeout, s.Server.WriteTimeout)

	go func() {
		log.Info().Uint("port", s.Server.Port).Msg("successfully started te server")
		if err = server.ListenAndServe(); err != nil {
			log.Fatal().Err(err).Msg("failed to start the server")
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	<-signalChannel

	if err = server.Shutdown(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed to properly shutdown the server")
	}
}
