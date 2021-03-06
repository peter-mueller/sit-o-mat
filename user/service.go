package user

import (
	"context"
	"io"

	"github.com/peter-mueller/sit-o-mat/httperror"
	"gocloud.dev/docstore"
)

// Service contains all actions on users
type Service struct {
}

// RegisterUser as a new user in the sit-o-mat.
// A password is auto-generated.
func (s *Service) RegisterUser(ctx context.Context, Name string) (User, error) {

	pwd := GeneratePassword()
	hash := hashAndSalt(pwd)

	user := User{
		Name:     Name,
		Password: hash,
	}
	coll := userCollection()
	defer coll.Close()

	err := coll.Create(ctx, &user)

	if err != nil {
		return User{}, httperror.Wrap("failed to register new user", err)
	}

	return User{
		Name:     Name,
		Password: pwd,
	}, nil
}

// FindUserByName by the name key
func (s *Service) FindUserByName(ctx context.Context, Name string) (User, error) {
	user := User{Name: Name}

	coll := userCollection()
	defer coll.Close()

	err := coll.Get(ctx, &user)

	if err != nil {
		return user, httperror.Wrap("failed to find user by name", err)
	}
	return user, nil
}

// UpdateUser to the new values
func (s *Service) UpdateUser(ctx context.Context, user User) (User, error) {
	coll := userCollection()
	defer coll.Close()

	err := coll.Put(ctx, &user)

	if err != nil {
		return user, httperror.Wrap("failed to put user", err)
	}
	return user, nil
}

// LoginUser attempt, returns true if ok
func (s *Service) LoginUser(ctx context.Context, user User) (bool, error) {
	givenPassword := user.Password
	coll := userCollection()
	defer coll.Close()

	err := coll.Get(ctx, &user)

	if err != nil {
		return false, httperror.Wrap("failed to validate user", err)
	}

	ok := comparePasswords(user.Password, givenPassword)
	return ok, nil
}

// DeleteUser from the application
func (s *Service) DeleteUser(ctx context.Context, user User) error {
	coll := userCollection()
	defer coll.Close()

	err := coll.Delete(ctx, &user)
	if err != nil {
		return httperror.Wrap("failed to delete user", err)
	}
	return nil
}

func (s *Service) FindAllUsers(ctx context.Context) ([]User, error) {
	coll := userCollection()
	defer coll.Close()

	iter := coll.Query().OrderBy("Name", docstore.Ascending).Get(ctx)
	defer iter.Stop()

	workplaces := make([]User, 0)

	for {
		var u User
		err := iter.Next(ctx, &u)
		if err == io.EOF {
			break
		}
		if err != nil {
			return workplaces, httperror.Wrap("failed to find all users", err)
		}

		workplaces = append(workplaces, u)

	}
	return workplaces, nil
}
