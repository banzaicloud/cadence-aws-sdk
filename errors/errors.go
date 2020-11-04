package errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"go.uber.org/cadence"
)

// decodedError is a serialized version of the AWS SDK error.
type decodedError struct {
	Code    string
	Message string
	OrigErr string
}

// decodedError implements the error interface.
func (e decodedError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// DecodeError decodes an error returned by a Cadence activity.
// If the error is an AWS SDK error, it gets converted to an AWS SDK compatible error interface.
// Otherwise the original error is returned.
func DecodeError(err error) error {
	if err == nil {
		return nil
	}

	customError, ok := err.(*cadence.CustomError)
	if !ok {
		return err
	}

	if !strings.HasPrefix(customError.Reason(), "AWS_SDK_") {
		return err
	}

	if !customError.HasDetails() {
		return awserr.New(
			strings.TrimPrefix(customError.Reason(), "AWS_SDK_"),
			customError.Error(),
			nil,
		)
	}

	var decErr decodedError

	if e := customError.Details(&decErr); e != nil {
		return err
	}

	return awserr.New(decErr.Code, decErr.Message, errors.New(decErr.OrigErr))
}
