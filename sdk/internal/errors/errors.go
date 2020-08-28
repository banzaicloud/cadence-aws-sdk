package errors

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"go.uber.org/cadence"
)

// reason identifies a custom Cadence error as an AWS SDK error.
const reason = "AWS_SDK_ERROR"

// Error is a serialized version of the AWS SDK error.
type Error struct {
	Code    string
	Message string
	OrigErr error
}

// Error implements the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// DecodeError decodes an error returned by a Cadence activity.
// If the error is an AWS SDK error, it gets converted to an AWS SDK compatible error interface.
// Otherwise the original error is returned.
func DecodeError(err error) error {
	if customError, ok := err.(*cadence.CustomError); ok && customError.Reason() == reason {
		var aerr Error

		_ = customError.Details(&aerr)

		return awserr.New(aerr.Code, aerr.Message, aerr.OrigErr)
	}

	return err
}
