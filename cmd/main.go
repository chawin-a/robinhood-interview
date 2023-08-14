package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	healthcheck "github.com/aschenmaker/fiber-health-check"
	"github.com/chawin-a/robinhood-interview/configs"
	middleware_logger "github.com/chawin-a/robinhood-interview/internal/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

func main() {
	conf := configs.InitConfig()
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	errWg, errCtx := errgroup.WithContext(ctx)

	interviewHandler := func(ctx *fiber.Ctx) error {
		return nil
	}

	commentHandler := func(ctx *fiber.Ctx) error {
		return nil
	}

	app := fiber.New()

	app.
		Use(favicon.New()).
		Use(healthcheck.New()).
		Use(middleware_logger.New())

	interviewAPI := app.Group("/interview")
	interviewAPI.Post("/", interviewHandler)
	interviewAPI.Post("/:id", interviewHandler)
	interviewAPI.Post("/:id/archive", interviewHandler)
	interviewAPI.Post("/:id/update", interviewHandler)

	commentAPI := interviewAPI.Group("/:id/comment")
	commentAPI.Post("/", commentHandler)
	commentAPI.Post("/add", commentHandler)

	errWg.Go(func() error {
		if err := app.Listen(fmt.Sprintf(":%s", conf.App.Port)); err != nil {
			return err
		}
		return nil
	})

	errWg.Go(func() error {
		<-errCtx.Done()
		return app.Shutdown()
	})

	if err := errWg.Wait(); err != nil {
		log.Panic(err)
	}
	log.Println("gracefully shutdown")
}
