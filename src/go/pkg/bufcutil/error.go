package bufcutil

import (
	"github.com/bufbuild/connect-go"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func InvalidArgumentError(field, msg string) *connect.Error {
	e := connect.NewError(connect.CodeInvalidArgument, errors.New(msg))

	data := &errdetails.BadRequest{
		FieldViolations: []*errdetails.BadRequest_FieldViolation{
			// &errdetails.BadRequest_FieldViolation{
			{
				Field:       field,
				Description: msg,
			},
		},
	}
	if detail, err := connect.NewErrorDetail(data); err == nil {
		e.AddDetail(detail)
	}

	return e
}

func InternalError(err error, msg ...string) *connect.Error {
	if len(msg) == 0 {
		return connect.NewError(connect.CodeInternal, err)
	}
	if err == nil {
		return connect.NewError(connect.CodeInternal, errors.New(msg[0]))
	}

	return connect.NewError(connect.CodeInternal, errors.Wrap(err, msg[0]))
}

func NotFoundError(msg ...string) *connect.Error {
	if len(msg) == 0 {
		return connect.NewError(connect.CodeNotFound, nil)
	}

	return connect.NewError(connect.CodeNotFound, errors.New(msg[0]))
}

func FailedPreconditionError(msg ...string) *connect.Error {
	if len(msg) == 0 {
		return connect.NewError(connect.CodeFailedPrecondition, nil)
	}

	return connect.NewError(connect.CodeFailedPrecondition, errors.New(msg[0]))
}

func PermissionDeniedError(msg ...string) *connect.Error {
	if len(msg) == 0 {
		return connect.NewError(connect.CodeFailedPrecondition, nil)
	}

	return connect.NewError(connect.CodePermissionDenied, errors.New(msg[0]))
}
