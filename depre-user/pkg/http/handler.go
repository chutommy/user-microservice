package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"

	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/multierr"

	"user/pkg/endpoint"
	"user/pkg/service"
)

// ErrJSONDecode is returned jf invalid JSON is provided.
var ErrJSONDecode = errors.New("could not decode JSON body")

// makeAddGenderHandler creates the handler logic
func makeAddGenderHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodPost).Path("/add-gender").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodPost}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.AddGenderEndpoint,
		decodeAddGenderRequest,
		encodeAddGenderResponse,
		options...,
	)))
}

// decodeAddGenderRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddGenderRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.AddGenderRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeAddGenderResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddGenderResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetGenderHandler creates the handler logic
func makeGetGenderHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodGet).Path("/get-gender").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodGet}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.GetGenderEndpoint,
		decodeGetGenderRequest,
		encodeGetGenderResponse,
		options...,
	)))
}

// decodeGetGenderRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetGenderRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetGenderRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeGetGenderResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetGenderResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListGendersHandler creates the handler logic
func makeListGendersHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodGet).Path("/list-genders").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodGet}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.ListGendersEndpoint,
		decodeListGendersRequest,
		encodeListGendersResponse,
		options...,
	)))
}

// decodeListGendersRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListGendersRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ListGendersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeListGendersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListGendersResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeRemoveGenderHandler creates the handler logic
func makeRemoveGenderHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodDelete).Path("/remove-gender").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodDelete}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.RemoveGenderEndpoint,
		decodeRemoveGenderRequest,
		encodeRemoveGenderResponse,
		options...,
	)))
}

// decodeRemoveGenderRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRemoveGenderRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.RemoveGenderRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeRemoveGenderResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRemoveGenderResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateUserHandler creates the handler logic
func makeCreateUserHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodPost).Path("/create-user").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodPost}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.CreateUserEndpoint,
		decodeCreateUserRequest,
		encodeCreateUserResponse,
		options...,
	)))
}

// decodeCreateUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateUserRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeCreateUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateUserResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserByIDHandler creates the handler logic
func makeGetUserByIDHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodGet).Path("/get-user-by-id").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodGet}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.GetUserByIDEndpoint,
		decodeGetUserByIDRequest,
		encodeGetUserByIDResponse,
		options...,
	)))
}

// decodeGetUserByIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserByIDRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetUserByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeGetUserByIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserByIDResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserByEmailHandler creates the handler logic
func makeGetUserByEmailHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodGet).Path("/get-user-by-email").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodGet}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.GetUserByEmailEndpoint,
		decodeGetUserByEmailRequest,
		encodeGetUserByEmailResponse,
		options...,
	)))
}

// decodeGetUserByEmailRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserByEmailRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetUserByEmailRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeGetUserByEmailResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserByEmailResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateUserEmailHandler creates the handler logic
func makeUpdateUserEmailHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodPut).Path("/update-user-email").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodPut}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.UpdateUserEmailEndpoint,
		decodeUpdateUserEmailRequest,
		encodeUpdateUserEmailResponse,
		options...,
	)))
}

// decodeUpdateUserEmailRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateUserEmailRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateUserEmailRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeUpdateUserEmailResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateUserEmailResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateUserPasswordHandler creates the handler logic
func makeUpdateUserPasswordHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodPut).Path("/update-user-password").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodPut}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.UpdateUserPasswordEndpoint,
		decodeUpdateUserPasswordRequest,
		encodeUpdateUserPasswordResponse,
		options...,
	)))
}

// decodeUpdateUserPasswordRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateUserPasswordRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateUserPasswordRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeUpdateUserPasswordResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateUserPasswordResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateUserInfoHandler creates the handler logic
func makeUpdateUserInfoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodPut).Path("/update-user-info").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodPut}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.UpdateUserInfoEndpoint,
		decodeUpdateUserInfoRequest,
		encodeUpdateUserInfoResponse,
		options...,
	)))
}

// decodeUpdateUserInfoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateUserInfoRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateUserInfoRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeUpdateUserInfoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateUserInfoResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteUserSoftHandler creates the handler logic
func makeDeleteUserSoftHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodDelete).Path("/delete-user-soft").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodDelete}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.DeleteUserSoftEndpoint,
		decodeDeleteUserSoftRequest,
		encodeDeleteUserSoftResponse,
		options...,
	)))
}

// decodeDeleteUserSoftRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteUserSoftRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.DeleteUserSoftRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeDeleteUserSoftResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteUserSoftResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeRecoverUserHandler creates the handler logic
func makeRecoverUserHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodPost).Path("/recover-user").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodPost}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.RecoverUserEndpoint,
		decodeRecoverUserRequest,
		encodeRecoverUserResponse,
		options...,
	)))
}

// decodeRecoverUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRecoverUserRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.RecoverUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeRecoverUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRecoverUserResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteUserPermanentHandler creates the handler logic
func makeDeleteUserPermanentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodDelete).Path("/delete-user-permanent").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodDelete}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.DeleteUserPermanentEndpoint,
		decodeDeleteUserPermanentRequest,
		encodeDeleteUserPermanentResponse,
		options...,
	)))
}

// decodeDeleteUserPermanentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteUserPermanentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.DeleteUserPermanentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeDeleteUserPermanentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteUserPermanentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeVerifyPasswordHandler creates the handler logic
func makeVerifyPasswordHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods(http1.MethodPost).Path("/verify-password").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{http1.MethodPost}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(
		endpoints.VerifyPasswordEndpoint,
		decodeVerifyPasswordRequest,
		encodeVerifyPasswordResponse,
		options...,
	)))
}

// decodeVerifyPasswordRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeVerifyPasswordRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.VerifyPasswordRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = multierr.Append(ErrJSONDecode, err)
	}
	return req, err
}

// encodeVerifyPasswordResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeVerifyPasswordResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	_ = json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status.
func err2code(err error) int {
	switch {
	case errors.Is(err, service.ErrMissingRequestField), errors.Is(err, ErrJSONDecode):
		return http1.StatusBadRequest
	case errors.Is(err, service.ErrDuplicatedValue):
		return http1.StatusConflict
	case errors.Is(err, service.ErrNotFound):
		return http1.StatusNotFound
	case errors.Is(err, service.ErrWrongPassword):
		return http1.StatusUnauthorized
	}

	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
