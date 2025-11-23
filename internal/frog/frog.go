// Package frog provides business logic for managing frogs and candidates.
package frog

import (
	"errors"
	"fmt"

	"github.com/ryanmalonzo/eat-that-frog/internal/db"
)

// GetTodayFrog returns today's picked frog with validation.
// Returns an error if no frog has been picked for today.
func GetTodayFrog() (string, error) {
	frog, err := db.GetTodayFrog()
	if err != nil {
		return "", err
	}

	if frog == "" {
		return "", errors.New("no frog has been picked for today")
	}

	return frog, nil
}

// GetAllCandidates returns all candidate frogs.
func GetAllCandidates() ([]string, error) {
	return db.GetAllCandidates()
}

// ValidateCandidatesExist checks if there are any candidates available.
// Returns the count of candidates and an error if the check fails.
func ValidateCandidatesExist() (int, error) {
	count, err := db.CountCandidates()
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, fmt.Errorf("no candidate frogs available to pick")
	}

	return count, nil
}

// PickCandidate picks a candidate frog at the given index.
// The index should be 0-based. Returns an error if the index is invalid
// or if the pick operation fails.
func PickCandidate(index int) error {
	count, err := db.CountCandidates()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("no candidate frogs available to pick")
	}

	if index < 0 || index >= count {
		return fmt.Errorf("invalid choice: index must be between 1 and %d", count)
	}

	return db.PickCandidate(index)
}

