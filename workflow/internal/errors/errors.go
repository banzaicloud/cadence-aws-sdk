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
	OrigErr string
}

// Error implements the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// EncodeError encodes an error returned by the AWS SDK as a Cadence error.
func EncodeError(err error) error {
	if aerr, ok := err.(awserr.Error); ok {
		newErr := Error{
			Code:    aerr.Code(),
			Message: aerr.Message(),
		}

		// TODO: try to encode the original error as well?
		if aerr.OrigErr() != nil {
			newErr.OrigErr = aerr.OrigErr().Error()
		}

		return cadence.NewCustomError(reason, newErr)
	}

	return err
}
