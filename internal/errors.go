package internal

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"go.uber.org/cadence"
)

// encodedError is a serialized version of the AWS SDK error.
type encodedError struct {
	Code    string
	Message string
	OrigErr string
}

// Error implements the error interface.
func (e encodedError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// EncodeError encodes an error returned by the AWS SDK as a Cadence error.
func EncodeError(err error) error {
	if err == nil {
		return nil
	}

	if aerr, ok := err.(awserr.Error); ok {
		newErr := encodedError{
			Code:    aerr.Code(),
			Message: aerr.Message(),
		}

		// TODO: try to encode the original error as well?
		if aerr.OrigErr() != nil {
			newErr.OrigErr = aerr.OrigErr().Error()
		}

		return cadence.NewCustomError("AWS_SDK_"+aerr.Code(), newErr)
	}

	return err
}
