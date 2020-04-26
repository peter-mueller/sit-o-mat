package sitomat

import (
	"math/rand"
	"time"

	"context"

	"github.com/peter-mueller/sit-o-mat/httperror"
	"github.com/peter-mueller/sit-o-mat/user"
	"github.com/peter-mueller/sit-o-mat/workplace"
)

type UserService interface {
	FindAllUsers(ctx context.Context) ([]user.User, error)
}
type WorkplaceService interface {
	FindAllWorkplacesSortByRating(ctx context.Context) ([]workplace.Workplace, error)
	UpdateWorkplace(ctx context.Context, w workplace.Workplace) (workplace.Workplace, error)
}

type Service struct {
	UserService      UserService
	WorkplaceService WorkplaceService
}

func (s Service) AssignWorkplaces(ctx context.Context) error {
	users, err := s.UserService.FindAllUsers(ctx)
	if err != nil {
		return httperror.Wrap("failed to load all users", err)
	}
	workplaces, err := s.WorkplaceService.FindAllWorkplacesSortByRating(ctx)
	if err != nil {
		return httperror.Wrap("failed to load all workplaces", err)
	}

	for index, _ := range workplaces {
		workplaces[index].CurrentOwner = ""
	}

	for _, u := range users {
		if u.WeeklyRequests.RequestForToday() && u.Fix != "" {
			for index, wp := range workplaces {
				if wp.Name == u.Fix {
					workplaces[index].CurrentOwner = u.Name
				}
			}
		}
	}

	shuffledUser := make([]user.User, 0)
	for _, u := range users {
		if u.WeeklyRequests.RequestForToday() && u.Fix == "" {
			shuffledUser = append(shuffledUser, u)
		}
	}

	rand.Seed(int64(time.Now().YearDay() * time.Now().Year()))
	rand.Shuffle(len(shuffledUser), func(i, j int) { shuffledUser[i], shuffledUser[j] = shuffledUser[j], shuffledUser[i] })

	var index = 0
	for _, wp := range workplaces {
		if wp.CurrentOwner == "" && index < len(shuffledUser) {
			wp.CurrentOwner = shuffledUser[index].Name
			index++
		}

		_, err := s.WorkplaceService.UpdateWorkplace(ctx, wp)
		if err != nil {
			return httperror.Wrap("failed to assign workplace", err)
		}
	}
	return nil
}
