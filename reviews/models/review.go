package models

import (
	"errors"
	"fmt"
	"time"
)

const maxLengthInComments = 400

// Review represent a review from some website or App
type Review struct {
	Id      int64
	Stars   int       // 1 - 5
	Comment string    // max len is maxLengthInComments
	Date    time.Time // created at
}

// CreateReviewCMD command to create a new review
type CreateReviewCMD struct {
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("stars must be between 1 - 5")
	}
	if len(cmd.Comment) > maxLengthInComments {
		var errorString = fmt.Sprintf("comment must be less than %d chars", maxLengthInComments)
		return errors.New(errorString)
	}
	return nil
}
