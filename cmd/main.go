package cmd

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/divyaanshjha/Eth_Indexer/config"
	"github.com/divyaanshjha/Eth_Indexer/internal/indexer"
	"github.com/divyaanshjha/Eth_Indexer/internal/repository"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)
func main(){

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	log.Info().Msgf("config: %+v", cfg)

	indexer := indexer.New(&cfg, repository.New())

	log.Info().Msg("starting indexer")
	if err := indexer.Start(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to start indexer")
	}
	log.Info().Msg("indexer stopped!")


}