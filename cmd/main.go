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
	handler_comment "github.com/chawin-a/robinhood-interview/internal/handler/comment"
	handler_interview "github.com/chawin-a/robinhood-interview/internal/handler/interview"
	handler_user "github.com/chawin-a/robinhood-interview/internal/handler/user"
	middleware_logger "github.com/chawin-a/robinhood-interview/internal/middleware/logger"
	"github.com/chawin-a/robinhood-interview/internal/postgres"
	repository_comment "github.com/chawin-a/robinhood-interview/internal/repository/comment"
	repository_interview "github.com/chawin-a/robinhood-interview/internal/repository/interview"
	repository_user "github.com/chawin-a/robinhood-interview/internal/repository/user"
	usecase_comment "github.com/chawin-a/robinhood-interview/internal/usecase/comment"
	usecase_interview "github.com/chawin-a/robinhood-interview/internal/usecase/interview"
	usecase_user "github.com/chawin-a/robinhood-interview/internal/usecase/user"
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

	db := postgres.New(conf.Postgres)
	userRepo := repository_user.NewUserRepository(db)
	interviewRepo := repository_interview.NewInterviewRepository(db)
	commentRepo := repository_comment.NewCommentRepository(db)

	// usecases
	interviewUsecase := usecase_interview.NewUsecase(interviewRepo)
	commentUsecase := usecase_comment.NewUsecase(commentRepo)
	userUsecase := usecase_user.NewUsecase(userRepo)

	// handler
	interviewHandler := handler_interview.NewHandler(interviewUsecase)
	commentHandler := handler_comment.NewHandler(commentUsecase)
	userHandler := handler_user.NewHandler(userUsecase)

	errWg, errCtx := errgroup.WithContext(ctx)

	app := fiber.New()

	app.
		Use(favicon.New()).
		Use(healthcheck.New()).
		Use(middleware_logger.New())

	userAPI := app.Group("/user")
	userAPI.Post("/:id", userHandler.Get)

	interviewAPI := app.Group("/interview")
	interviewAPI.Post("/", interviewHandler.List)
	interviewAPI.Post("/:id", interviewHandler.Get)
	interviewAPI.Post("/:id/archive", interviewHandler.Archive)
	interviewAPI.Post("/:id/update_status", interviewHandler.UpdateStatus)

	commentAPI := interviewAPI.Group("/:interviewId/comment")
	commentAPI.Post("/", commentHandler.List)
	commentAPI.Post("/create", commentHandler.Create)

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
