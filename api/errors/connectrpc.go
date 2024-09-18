package errors

import (
	"errors"

	commonv1 "buf.build/gen/go/redpandadata/common/protocolbuffers/go/redpanda/api/common/v1"
	"connectrpc.com/connect"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// NewConnectError is a helper function to construct a new connect error. This
// function should always be used over instantiating connect errors directly, as
// we can ensure that certain error details will always be provided.
func NewConnectError(
	code connect.Code,
	innerErr error,
	errInfo *errdetails.ErrorInfo,
	errDetails ...proto.Message,
) *connect.Error {
	connectErr := connect.NewError(code, innerErr)

	if detail, detailErr := connect.NewErrorDetail(errInfo); detailErr == nil {
		connectErr.AddDetail(detail)
	}

	for _, msg := range errDetails {
		// We may sometimes pass in a nil object so that this function is easier
		// to use. In this case we just want to skip it.
		if msg == nil {
			continue
		}
		detail, detailErr := connect.NewErrorDetail(msg)
		if detailErr != nil {
			continue
		}
		connectErr.AddDetail(detail)
	}

	return connectErr
}

// KeyVal is a key/value pair that is used to provide additional metadata labels.
type KeyVal struct {
	Key   string
	Value string
}

// NewErrorInfo is a helper function to create a new ErrorInfo detail.
func NewErrorInfo(domain Domain, reason string, metadata ...KeyVal) *errdetails.ErrorInfo {
	var md map[string]string
	if len(metadata) > 0 {
		md = make(map[string]string, len(metadata))

		for _, keyVal := range metadata {
			md[keyVal.Key] = keyVal.Value
		}
	}

	return &errdetails.ErrorInfo{
		Reason:   reason,
		Domain:   string(domain),
		Metadata: md,
	}
}

// NewBadRequest is a constructor for creating bad request
func NewBadRequest(fieldValidations ...*errdetails.BadRequest_FieldViolation) *errdetails.BadRequest {
	return &errdetails.BadRequest{FieldViolations: fieldValidations}
}

// NewHelp constructs a new errdetails.Help with one or more provided errdetails.Help_Link.
func NewHelp(links ...*errdetails.Help_Link) *errdetails.Help {
	return &errdetails.Help{Links: links}
}

// NewHelpLink constructs a new link that can be put into the errdetails.Help.
func NewHelpLink(description, url string) *errdetails.Help_Link {
	return &errdetails.Help_Link{
		Description: description,
		Url:         url,
	}
}

// NewExternalErrorDetail can be called to transform an internal error into an
// external error that is safe to be returned in gRPC/Connect.
// Error messages are removed.
func NewExternalErrorDetail(message string, details ...proto.Message) *commonv1.ExternalError {
	e := commonv1.ExternalError{
		Message: message,
	}

	for _, detail := range details {
		anyDetail, err := anypb.New(detail)
		if err != nil {
			// Ignore errors
			continue
		}
		e.Details = append(e.Details, anyDetail)
	}

	return &e
}

// NewSafePublicError can be called to transform an internal error into an
// external error that is safe to be returned in gRPC/Connect.
// Error messages are removed.
func NewSafePublicError(internalError *status.Status) *status.Status {
	ext := status.New(internalError.Code(), "")

	for _, detail := range internalError.Details() {
		if d, ok := detail.(*commonv1.ExternalError); ok {
			pb := &spb.Status{
				Code:    int32(internalError.Code()), //nolint:gosec // code conversion
				Message: d.Message,
				Details: d.Details,
			}

			return status.FromProto(pb)
		}
	}

	return ext
}

// NewSafePublicErrorConnect can be called to transform an internal error into an
// external error that is safe to be returned in gRPC/Connect.
// Error messages are removed.
func NewSafePublicErrorConnect(internalError *connect.Error) *connect.Error {
	for _, d := range internalError.Details() {
		detail, err := d.Value()
		if err != nil {
			continue
		}

		if d, ok := detail.(*commonv1.ExternalError); ok {
			result := connect.NewError(internalError.Code(), errors.New(d.Message))

			for _, detail := range d.Details {
				connectSDKDetail, err := connect.NewErrorDetail(detail)
				if err != nil {
					continue
				}
				result.AddDetail(connectSDKDetail)
			}

			return result
		}
	}

	return connect.NewError(internalError.Code(), nil)
}
