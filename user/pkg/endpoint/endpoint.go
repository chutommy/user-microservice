package endpoint

import (
	"context"

	"user/pkg/repo"
	"user/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

// AddGenderRequest collects the request parameters for the AddGender method.
type AddGenderRequest struct {
	Title string `json:"title"`
}

// AddGenderResponse collects the response parameters for the AddGender method.
type AddGenderResponse struct {
	Gender repo.Gender `json:"gender"`
	Err    error       `json:"err"`
}

// MakeAddGenderEndpoint returns an endpoint that invokes AddGender on the service.
func MakeAddGenderEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddGenderRequest)
		gender, err := s.AddGender(ctx, req.Title)
		return AddGenderResponse{
			Err:    err,
			Gender: gender,
		}, nil
	}
}

// Failed implements Failer.
func (r AddGenderResponse) Failed() error {
	return r.Err
}

// GetGenderRequest collects the request parameters for the GetGender method.
type GetGenderRequest struct {
	Id int16 `json:"id"`
}

// GetGenderResponse collects the response parameters for the GetGender method.
type GetGenderResponse struct {
	Gender repo.Gender `json:"gender"`
	Err    error       `json:"err"`
}

// MakeGetGenderEndpoint returns an endpoint that invokes GetGender on the service.
func MakeGetGenderEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGenderRequest)
		gender, err := s.GetGender(ctx, req.Id)
		return GetGenderResponse{
			Err:    err,
			Gender: gender,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGenderResponse) Failed() error {
	return r.Err
}

// ListGendersRequest collects the request parameters for the ListGenders method.
type ListGendersRequest struct{}

// ListGendersResponse collects the response parameters for the ListGenders method.
type ListGendersResponse struct {
	Genders []repo.Gender `json:"genders"`
	Err     error         `json:"err"`
}

// MakeListGendersEndpoint returns an endpoint that invokes ListGenders on the service.
func MakeListGendersEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		genders, err := s.ListGenders(ctx)
		return ListGendersResponse{
			Err:     err,
			Genders: genders,
		}, nil
	}
}

// Failed implements Failer.
func (r ListGendersResponse) Failed() error {
	return r.Err
}

// RemoveGenderRequest collects the request parameters for the RemoveGender method.
type RemoveGenderRequest struct {
	Id int16 `json:"id"`
}

// RemoveGenderResponse collects the response parameters for the RemoveGender method.
type RemoveGenderResponse struct {
	Err error `json:"err"`
}

// MakeRemoveGenderEndpoint returns an endpoint that invokes RemoveGender on the service.
func MakeRemoveGenderEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveGenderRequest)
		err := s.RemoveGender(ctx, req.Id)
		return RemoveGenderResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r RemoveGenderResponse) Failed() error {
	return r.Err
}

// CreateUserRequest collects the request parameters for the CreateUser method.
type CreateUserRequest struct {
	User repo.User `json:"user"`
}

// CreateUserResponse collects the response parameters for the CreateUser method.
type CreateUserResponse struct {
	User repo.User `json:"user"`
	Err  error     `json:"err"`
}

// MakeCreateUserEndpoint returns an endpoint that invokes CreateUser on the service.
func MakeCreateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		user, err := s.CreateUser(ctx, req.User)
		return CreateUserResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateUserResponse) Failed() error {
	return r.Err
}

// GetUserByIDRequest collects the request parameters for the GetUserByID method.
type GetUserByIDRequest struct {
	Id int64 `json:"id"`
}

// GetUserByIDResponse collects the response parameters for the GetUserByID method.
type GetUserByIDResponse struct {
	User repo.User `json:"user"`
	Err  error     `json:"err"`
}

// MakeGetUserByIDEndpoint returns an endpoint that invokes GetUserByID on the service.
func MakeGetUserByIDEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByIDRequest)
		user, err := s.GetUserByID(ctx, req.Id)
		return GetUserByIDResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserByIDResponse) Failed() error {
	return r.Err
}

// GetUserByEmailRequest collects the request parameters for the GetUserByEmail method.
type GetUserByEmailRequest struct {
	Email string `json:"email"`
}

// GetUserByEmailResponse collects the response parameters for the GetUserByEmail method.
type GetUserByEmailResponse struct {
	User repo.User `json:"user"`
	Err  error     `json:"err"`
}

// MakeGetUserByEmailEndpoint returns an endpoint that invokes GetUserByEmail on the service.
func MakeGetUserByEmailEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByEmailRequest)
		user, err := s.GetUserByEmail(ctx, req.Email)
		return GetUserByEmailResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserByEmailResponse) Failed() error {
	return r.Err
}

// UpdateUserEmailRequest collects the request parameters for the UpdateUserEmail method.
type UpdateUserEmailRequest struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

// UpdateUserEmailResponse collects the response parameters for the UpdateUserEmail method.
type UpdateUserEmailResponse struct {
	User repo.User `json:"user"`
	Err  error     `json:"err"`
}

// MakeUpdateUserEmailEndpoint returns an endpoint that invokes UpdateUserEmail on the service.
func MakeUpdateUserEmailEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserEmailRequest)
		user, err := s.UpdateUserEmail(ctx, req.Id, req.Email)
		return UpdateUserEmailResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserEmailResponse) Failed() error {
	return r.Err
}

// UpdateUserPasswordRequest collects the request parameters for the UpdateUserPassword method.
type UpdateUserPasswordRequest struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}

// UpdateUserPasswordResponse collects the response parameters for the UpdateUserPassword method.
type UpdateUserPasswordResponse struct {
	User repo.User `json:"user"`
	Err  error     `json:"err"`
}

// MakeUpdateUserPasswordEndpoint returns an endpoint that invokes UpdateUserPassword on the service.
func MakeUpdateUserPasswordEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserPasswordRequest)
		user, err := s.UpdateUserPassword(ctx, req.Id, req.Password)
		return UpdateUserPasswordResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserPasswordResponse) Failed() error {
	return r.Err
}

// UpdateUserInfoRequest collects the request parameters for the UpdateUserInfo method.
type UpdateUserInfoRequest struct {
	Id   int64     `json:"id"`
	User repo.User `json:"user"`
}

// UpdateUserInfoResponse collects the response parameters for the UpdateUserInfo method.
type UpdateUserInfoResponse struct {
	User repo.User `json:"user"`
	Err  error     `json:"err"`
}

// MakeUpdateUserInfoEndpoint returns an endpoint that invokes UpdateUserInfo on the service.
func MakeUpdateUserInfoEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserInfoRequest)
		user, err := s.UpdateUserInfo(ctx, req.Id, req.User)
		return UpdateUserInfoResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserInfoResponse) Failed() error {
	return r.Err
}

// DeleteUserSoftRequest collects the request parameters for the DeleteUserSoft method.
type DeleteUserSoftRequest struct {
	Id int64 `json:"id"`
}

// DeleteUserSoftResponse collects the response parameters for the DeleteUserSoft method.
type DeleteUserSoftResponse struct {
	Err error `json:"err"`
}

// MakeDeleteUserSoftEndpoint returns an endpoint that invokes DeleteUserSoft on the service.
func MakeDeleteUserSoftEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserSoftRequest)
		err := s.DeleteUserSoft(ctx, req.Id)
		return DeleteUserSoftResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteUserSoftResponse) Failed() error {
	return r.Err
}

// RecoverUserRequest collects the request parameters for the RecoverUser method.
type RecoverUserRequest struct {
	Id int64 `json:"id"`
}

// RecoverUserResponse collects the response parameters for the RecoverUser method.
type RecoverUserResponse struct {
	User repo.User `json:"user"`
	Err  error     `json:"err"`
}

// MakeRecoverUserEndpoint returns an endpoint that invokes RecoverUser on the service.
func MakeRecoverUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RecoverUserRequest)
		user, err := s.RecoverUser(ctx, req.Id)
		return RecoverUserResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r RecoverUserResponse) Failed() error {
	return r.Err
}

// DeleteUserPermanentRequest collects the request parameters for the DeleteUserPermanent method.
type DeleteUserPermanentRequest struct {
	Id int64 `json:"id"`
}

// DeleteUserPermanentResponse collects the response parameters for the DeleteUserPermanent method.
type DeleteUserPermanentResponse struct {
	Err error `json:"err"`
}

// MakeDeleteUserPermanentEndpoint returns an endpoint that invokes DeleteUserPermanent on the service.
func MakeDeleteUserPermanentEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserPermanentRequest)
		err := s.DeleteUserPermanent(ctx, req.Id)
		return DeleteUserPermanentResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteUserPermanentResponse) Failed() error {
	return r.Err
}

// VerifyPasswordRequest collects the request parameters for the VerifyPassword method.
type VerifyPasswordRequest struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}

// VerifyPasswordResponse collects the response parameters for the VerifyPassword method.
type VerifyPasswordResponse struct {
	Err error `json:"err"`
}

// MakeVerifyPasswordEndpoint returns an endpoint that invokes VerifyPassword on the service.
func MakeVerifyPasswordEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VerifyPasswordRequest)
		err := s.VerifyPassword(ctx, req.Id, req.Password)
		return VerifyPasswordResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r VerifyPasswordResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// AddGender implements Service. Primarily useful in a client.
func (e Endpoints) AddGender(ctx context.Context, title string) (repo.Gender, error) {
	request := AddGenderRequest{Title: title}
	response, err := e.AddGenderEndpoint(ctx, request)
	if err != nil {
		return repo.Gender{}, err
	}
	return response.(AddGenderResponse).Gender, response.(AddGenderResponse).Err
}

// GetGender implements Service. Primarily useful in a client.
func (e Endpoints) GetGender(ctx context.Context, id int16) (repo.Gender, error) {
	request := GetGenderRequest{Id: id}
	response, err := e.GetGenderEndpoint(ctx, request)
	if err != nil {
		return repo.Gender{}, err
	}
	return response.(GetGenderResponse).Gender, response.(GetGenderResponse).Err
}

// ListGenders implements Service. Primarily useful in a client.
func (e Endpoints) ListGenders(ctx context.Context) ([]repo.Gender, error) {
	request := ListGendersRequest{}
	response, err := e.ListGendersEndpoint(ctx, request)
	if err != nil {
		return []repo.Gender{}, err
	}
	return response.(ListGendersResponse).Genders, response.(ListGendersResponse).Err
}

// RemoveGender implements Service. Primarily useful in a client.
func (e Endpoints) RemoveGender(ctx context.Context, id int16) error {
	request := RemoveGenderRequest{Id: id}
	response, err := e.RemoveGenderEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return response.(RemoveGenderResponse).Err
}

// CreateUser implements Service. Primarily useful in a client.
func (e Endpoints) CreateUser(ctx context.Context, user repo.User) (repo.User, error) {
	request := CreateUserRequest{User: user}
	response, err := e.CreateUserEndpoint(ctx, request)
	if err != nil {
		return repo.User{}, err
	}
	return response.(CreateUserResponse).User, response.(CreateUserResponse).Err
}

// GetUserByID implements Service. Primarily useful in a client.
func (e Endpoints) GetUserByID(ctx context.Context, id int64) (repo.User, error) {
	request := GetUserByIDRequest{Id: id}
	response, err := e.GetUserByIDEndpoint(ctx, request)
	if err != nil {
		return repo.User{}, err
	}
	return response.(GetUserByIDResponse).User, response.(GetUserByIDResponse).Err
}

// GetUserByEmail implements Service. Primarily useful in a client.
func (e Endpoints) GetUserByEmail(ctx context.Context, email string) (repo.User, error) {
	request := GetUserByEmailRequest{Email: email}
	response, err := e.GetUserByEmailEndpoint(ctx, request)
	if err != nil {
		return repo.User{}, err
	}
	return response.(GetUserByEmailResponse).User, response.(GetUserByEmailResponse).Err
}

// UpdateUserEmail implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUserEmail(ctx context.Context, id int64, email string) (repo.User, error) {
	request := UpdateUserEmailRequest{
		Email: email,
		Id:    id,
	}
	response, err := e.UpdateUserEmailEndpoint(ctx, request)
	if err != nil {
		return repo.User{}, err
	}
	return response.(UpdateUserEmailResponse).User, response.(UpdateUserEmailResponse).Err
}

// UpdateUserPassword implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUserPassword(ctx context.Context, id int64, password string) (repo.User, error) {
	request := UpdateUserPasswordRequest{
		Id:       id,
		Password: password,
	}
	response, err := e.UpdateUserPasswordEndpoint(ctx, request)
	if err != nil {
		return repo.User{}, err
	}
	return response.(UpdateUserPasswordResponse).User, response.(UpdateUserPasswordResponse).Err
}

// UpdateUserInfo implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUserInfo(ctx context.Context, id int64, user repo.User) (repo.User, error) {
	request := UpdateUserInfoRequest{
		Id:   id,
		User: user,
	}
	response, err := e.UpdateUserInfoEndpoint(ctx, request)
	if err != nil {
		return repo.User{}, err
	}
	return response.(UpdateUserInfoResponse).User, response.(UpdateUserInfoResponse).Err
}

// DeleteUserSoft implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUserSoft(ctx context.Context, id int64) error {
	request := DeleteUserSoftRequest{Id: id}
	response, err := e.DeleteUserSoftEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return response.(DeleteUserSoftResponse).Err
}

// RecoverUser implements Service. Primarily useful in a client.
func (e Endpoints) RecoverUser(ctx context.Context, id int64) (repo.User, error) {
	request := RecoverUserRequest{Id: id}
	response, err := e.RecoverUserEndpoint(ctx, request)
	if err != nil {
		return repo.User{}, err
	}
	return response.(RecoverUserResponse).User, response.(RecoverUserResponse).Err
}

// DeleteUserPermanent implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUserPermanent(ctx context.Context, id int64) error {
	request := DeleteUserPermanentRequest{Id: id}
	response, err := e.DeleteUserPermanentEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return response.(DeleteUserPermanentResponse).Err
}

// VerifyPassword implements Service. Primarily useful in a client.
func (e Endpoints) VerifyPassword(ctx context.Context, id int64, password string) error {
	request := VerifyPasswordRequest{
		Id:       id,
		Password: password,
	}
	response, err := e.VerifyPasswordEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return response.(VerifyPasswordResponse).Err
}
