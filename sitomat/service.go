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

	filteredUsers := make([]user.User, 0)
	for _, user := range users {
		if user.WeeklyRequests.RequestForToday() {
			filteredUsers = append(filteredUsers, user)
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(filteredUsers), func(i, j int) { filteredUsers[i], filteredUsers[j] = filteredUsers[j], filteredUsers[i] })

	for index, workplace := range workplaces {
		if index >= len(filteredUsers) {
			workplace.CurrentOwner = ""
		} else {
			workplace.CurrentOwner = filteredUsers[index].Name
		}
		_, err := s.WorkplaceService.UpdateWorkplace(ctx, workplace)
		if err != nil {
			return httperror.Wrap("failed to assign workplace", err)
		}
	}
	return nil
}
