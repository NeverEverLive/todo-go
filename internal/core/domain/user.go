package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
)


type User struct {
	ID          string
	Version     int

	FirstName   string
	LastName    string
	PhoneNumber *string
}

func NewUser(
	id string,
	version int,
	firstName string,
	lastName string,
	phoneNumber *string,
) User {
	return User{
		ID: id,
		Version: version,
		FirstName: firstName,
		LastName: lastName,
		PhoneNumber: phoneNumber,
	}
}
	

func NewUserUninitialized(
	firstName string,
	lastName string,
	phoneNumber *string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		firstName,
		lastName,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	firstNameLength := len([]rune(u.FirstName))
	if firstNameLength < 3 || firstNameLength > 100 {
		return fmt.Errorf(
			"Invalid `FirstName` length: %d. %w",
			firstNameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	lastNameLength := len([]rune(u.LastName))
	if lastNameLength < 3 || lastNameLength > 100 {
		return fmt.Errorf(
			"Invalid `LastName` length: %d. %w",
			lastNameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if u.PhoneNumber != nil {
		phoneNumberLength := len([]rune(*u.PhoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf(
				"Invalid `PhoneNumber` length: %d. %w",
				phoneNumberLength,
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"Invalid `PhoneNumber` format: %s. %w",
				*u.PhoneNumber,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}


type UserPatch struct {
	FirstName   Nullable[string]
	LastName    Nullable[string]
	PhoneNumber Nullable[string]
}

func (p *UserPatch) Validate() error {
	if p.FirstName.Set && p.FirstName.Value == nil {
		return fmt.Errorf(
			"'FirstName' can't be patched to NULL: %w", 
			core_errors.ErrInvalidArgument,
		)
	}

	if p.LastName.Set && p.LastName.Value == nil {
		return fmt.Errorf(
			"'LastName' can't be patched to NULL: %w", 
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.FirstName.Set {
		tmp.FirstName = *patch.FirstName.Value
	}

	if patch.LastName.Set {
		tmp.LastName = *patch.LastName.Value
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp
	return nil
}