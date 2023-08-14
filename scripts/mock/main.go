package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/chawin-a/robinhood-interview/configs"
	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/chawin-a/robinhood-interview/internal/postgres"
	repository_comment "github.com/chawin-a/robinhood-interview/internal/repository/comment"
	repository_interview "github.com/chawin-a/robinhood-interview/internal/repository/interview"
	repository_user "github.com/chawin-a/robinhood-interview/internal/repository/user"
	"github.com/google/uuid"
)

func main() {
	conf := configs.InitConfig()
	db := postgres.New(conf.Postgres)
	ctx := context.Background()
	userRepo := repository_user.NewUserRepository(db)
	interviewRepo := repository_interview.NewInterviewRepository(db)
	commentRepo := repository_comment.NewCommentRepository(db)
	userIds := []uuid.UUID{}
	for i := 1; i <= 5; i++ {
		user, err := userRepo.Create(ctx, &entity.User{
			Username: fmt.Sprintf("user%d", i),
			Name:     fmt.Sprintf("user%d", i),
			Email:    fmt.Sprintf("user%d@robinhood.co.th", i),
		})
		if err != nil {
			panic(err)
		}
		userIds = append(userIds, user.Id)
	}
	for i := 1; i <= 50; i++ {
		interview, err := interviewRepo.Create(ctx, &entity.Interview{
			UserId:      userIds[rand.Intn(5)],
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas ligula purus, pulvinar vel nisi in, feugiat gravida lorem. Phasellus elit nunc, posuere ac ante sit amet, bibendum iaculis mi.",
			Status:      entity.TO_DO,
			IsArchived:  false,
		})
		if err != nil {
			panic(err)
		}
		nComment := rand.Intn(3) + 3
		for j := 0; j < nComment; j++ {
			_, err := commentRepo.Create(ctx, &entity.Comment{
				UserId:      userIds[rand.Intn(5)],
				InterviewId: interview.Id,
				Comment:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			})
			if err != nil {
				panic(err)
			}
		}
	}
}
