package cmd

import (
	"fmt"
	"github.com/jsmvaldivia/shorty/internal/app"
	"github.com/jsmvaldivia/shorty/internal/web"
	"github.com/spf13/cobra"
	"log"
)

func newStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		addr := fmt.Sprintf(":%d", 8888)
		log.Printf("API listening on %s", addr)

		return Setup()
	}

	return cmd
}

func Setup() error {
	shortenerService := app.ShortenerService{
		Shortener:  app.NewRandomShortener(),
		Repository: app.NewInMemoryRepository(),
	}

	shortenerController := app.NewShortenerController(shortenerService)
	return web.SetupWeb(shortenerController)
}
