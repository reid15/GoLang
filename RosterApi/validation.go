// Validate data inputs

package main

import (
	"errors"
	"strconv"
)

func isValidJerseyNumberString(jerseyNumber string) error {
	if len(jerseyNumber) == 0 {
		return errors.New("jerseyNumber Is Required")
	}
	jerseyNumberInt, err := strconv.ParseInt(jerseyNumber, 10, 32)
	if err != nil {
		return errors.New("Jersey Number Should Be An Integer")
	}
	err = isValidJerseyNumber(int(jerseyNumberInt))
	if err != nil {
		return err
	}

	return nil
}

func isValidJerseyNumber(jerseyNumber int) error {
	if jerseyNumber < 0 {
		return errors.New("Jersey Number Should Not Be A Negative Number")
	}
	if jerseyNumber > 99 {
		return errors.New("Jersey Number Should Be Less Than 100")
	}

	return nil
}

func isValidPlayer(player Player) error {
	err := isValidJerseyNumber(player.JerseyNumber)
	if err != nil {
		return err
	}
	if len(player.FirstName) == 0 {
		return errors.New("FirstName Is Required")
	}
	if len(player.LastName) == 0 {
		return errors.New("LastName Is Required")
	}
	if len(player.Position) == 0 {
		return errors.New("Position Is Required")
	}
	if len(player.Position) > 2 {
		return errors.New("Position Should be 1 or 2 characters")
	}

	return nil
}

func isValidPlayerForUpdate(player Player) error {
	err := isValidJerseyNumber(player.JerseyNumber)
	if err != nil {
		return err
	}
	if len(player.Position) == 0 {
		return errors.New("Position Is Required")
	}
	if len(player.Position) > 2 {
		return errors.New("Position Should be 1 or 2 characters")
	}

	return nil
}
