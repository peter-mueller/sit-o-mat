package workplace

import (
	"context"
	"io"

	"github.com/peter-mueller/sit-o-mat/httperror"
	"gocloud.dev/docstore"
)

// Service contains all actions workplaces
type Service struct {
	Collection *docstore.Collection
}

func (s *Service) DeleteWorkplaceByName(ctx context.Context, name string) error {
	workplace := Workplace{Name: name}
	err := s.Collection.Delete(ctx, &workplace)
	if err != nil {
		return httperror.Wrap("failed to delete workplace", err)
	}
	return nil
}

func (s *Service) UpdateWorkplace(ctx context.Context, w Workplace) (Workplace, error) {
	err := s.Collection.Put(ctx, &w)
	if err != nil {
		return w, httperror.Wrap("failed to update workplace", err)
	}
	return w, nil
}

func (s *Service) CreateWorkplace(ctx context.Context, w Workplace) (Workplace, error) {
	err := s.Collection.Create(ctx, &w)
	if err != nil {
		return w, httperror.Wrap("failed to create workplace", err)
	}
	return w, nil
}

func (s *Service) FindAllWorkplacesSortByRating(ctx context.Context) ([]Workplace, error) {
	iter := s.Collection.Query().OrderBy("Ranking", docstore.Ascending).Get(ctx)
	defer iter.Stop()

	workplaces := make([]Workplace, 0)

	for {
		var w Workplace
		err := iter.Next(ctx, &w)
		if err == io.EOF {
			break
		}
		if err != nil {
			return workplaces, httperror.Wrap("failed to find all workplaces", err)
		}

		workplaces = append(workplaces, w)

	}
	return workplaces, nil
}
