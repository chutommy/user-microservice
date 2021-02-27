package grpc

import (
	"context"
	"errors"
	endpoint "user/pkg/endpoint"
	pb "user/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeAddGenderHandler creates the handler logic
func makeAddGenderHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddGenderEndpoint, decodeAddGenderRequest, encodeAddGenderResponse, options...)
}

// decodeAddGenderResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain AddGender request.
// TODO implement the decoder
func decodeAddGenderRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeAddGenderResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddGenderResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) AddGender(ctx context1.Context, req *pb.AddGenderRequest) (*pb.AddGenderReply, error) {
	_, rep, err := g.addGender.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddGenderReply), nil
}

// makeGetGenderHandler creates the handler logic
func makeGetGenderHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGenderEndpoint, decodeGetGenderRequest, encodeGetGenderResponse, options...)
}

// decodeGetGenderResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetGender request.
// TODO implement the decoder
func decodeGetGenderRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeGetGenderResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetGenderResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) GetGender(ctx context1.Context, req *pb.GetGenderRequest) (*pb.GetGenderReply, error) {
	_, rep, err := g.getGender.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGenderReply), nil
}

// makeListGendersHandler creates the handler logic
func makeListGendersHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ListGendersEndpoint, decodeListGendersRequest, encodeListGendersResponse, options...)
}

// decodeListGendersResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ListGenders request.
// TODO implement the decoder
func decodeListGendersRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeListGendersResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeListGendersResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) ListGenders(ctx context1.Context, req *pb.ListGendersRequest) (*pb.ListGendersReply, error) {
	_, rep, err := g.listGenders.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListGendersReply), nil
}

// makeRemoveGenderHandler creates the handler logic
func makeRemoveGenderHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RemoveGenderEndpoint, decodeRemoveGenderRequest, encodeRemoveGenderResponse, options...)
}

// decodeRemoveGenderResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain RemoveGender request.
// TODO implement the decoder
func decodeRemoveGenderRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeRemoveGenderResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeRemoveGenderResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) RemoveGender(ctx context1.Context, req *pb.RemoveGenderRequest) (*pb.RemoveGenderReply, error) {
	_, rep, err := g.removeGender.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RemoveGenderReply), nil
}

// makeCreateUserHandler creates the handler logic
func makeCreateUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateUserEndpoint, decodeCreateUserRequest, encodeCreateUserResponse, options...)
}

// decodeCreateUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateUser request.
// TODO implement the decoder
func decodeCreateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeCreateUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) CreateUser(ctx context1.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	_, rep, err := g.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateUserReply), nil
}

// makeGetUserByIDHandler creates the handler logic
func makeGetUserByIDHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserByIDEndpoint, decodeGetUserByIDRequest, encodeGetUserByIDResponse, options...)
}

// decodeGetUserByIDResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserByID request.
// TODO implement the decoder
func decodeGetUserByIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeGetUserByIDResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserByIDResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) GetUserByID(ctx context1.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDReply, error) {
	_, rep, err := g.getUserByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserByIDReply), nil
}

// makeGetUserByEmailHandler creates the handler logic
func makeGetUserByEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserByEmailEndpoint, decodeGetUserByEmailRequest, encodeGetUserByEmailResponse, options...)
}

// decodeGetUserByEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserByEmail request.
// TODO implement the decoder
func decodeGetUserByEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeGetUserByEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserByEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) GetUserByEmail(ctx context1.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailReply, error) {
	_, rep, err := g.getUserByEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserByEmailReply), nil
}

// makeUpdateUserEmailHandler creates the handler logic
func makeUpdateUserEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateUserEmailEndpoint, decodeUpdateUserEmailRequest, encodeUpdateUserEmailResponse, options...)
}

// decodeUpdateUserEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateUserEmail request.
// TODO implement the decoder
func decodeUpdateUserEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeUpdateUserEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateUserEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) UpdateUserEmail(ctx context1.Context, req *pb.UpdateUserEmailRequest) (*pb.UpdateUserEmailReply, error) {
	_, rep, err := g.updateUserEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserEmailReply), nil
}

// makeUpdateUserPasswordHandler creates the handler logic
func makeUpdateUserPasswordHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateUserPasswordEndpoint, decodeUpdateUserPasswordRequest, encodeUpdateUserPasswordResponse, options...)
}

// decodeUpdateUserPasswordResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateUserPassword request.
// TODO implement the decoder
func decodeUpdateUserPasswordRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeUpdateUserPasswordResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateUserPasswordResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) UpdateUserPassword(ctx context1.Context, req *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordReply, error) {
	_, rep, err := g.updateUserPassword.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserPasswordReply), nil
}

// makeUpdateUserInfoHandler creates the handler logic
func makeUpdateUserInfoHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateUserInfoEndpoint, decodeUpdateUserInfoRequest, encodeUpdateUserInfoResponse, options...)
}

// decodeUpdateUserInfoResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateUserInfo request.
// TODO implement the decoder
func decodeUpdateUserInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeUpdateUserInfoResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateUserInfoResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) UpdateUserInfo(ctx context1.Context, req *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoReply, error) {
	_, rep, err := g.updateUserInfo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserInfoReply), nil
}

// makeDeleteUserSoftHandler creates the handler logic
func makeDeleteUserSoftHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteUserSoftEndpoint, decodeDeleteUserSoftRequest, encodeDeleteUserSoftResponse, options...)
}

// decodeDeleteUserSoftResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteUserSoft request.
// TODO implement the decoder
func decodeDeleteUserSoftRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeDeleteUserSoftResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteUserSoftResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) DeleteUserSoft(ctx context1.Context, req *pb.DeleteUserSoftRequest) (*pb.DeleteUserSoftReply, error) {
	_, rep, err := g.deleteUserSoft.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteUserSoftReply), nil
}

// makeRecoverUserHandler creates the handler logic
func makeRecoverUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RecoverUserEndpoint, decodeRecoverUserRequest, encodeRecoverUserResponse, options...)
}

// decodeRecoverUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain RecoverUser request.
// TODO implement the decoder
func decodeRecoverUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeRecoverUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeRecoverUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) RecoverUser(ctx context1.Context, req *pb.RecoverUserRequest) (*pb.RecoverUserReply, error) {
	_, rep, err := g.recoverUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RecoverUserReply), nil
}

// makeDeleteUserPermanentHandler creates the handler logic
func makeDeleteUserPermanentHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteUserPermanentEndpoint, decodeDeleteUserPermanentRequest, encodeDeleteUserPermanentResponse, options...)
}

// decodeDeleteUserPermanentResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteUserPermanent request.
// TODO implement the decoder
func decodeDeleteUserPermanentRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeDeleteUserPermanentResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteUserPermanentResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) DeleteUserPermanent(ctx context1.Context, req *pb.DeleteUserPermanentRequest) (*pb.DeleteUserPermanentReply, error) {
	_, rep, err := g.deleteUserPermanent.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteUserPermanentReply), nil
}

// makeVerifyPasswordHandler creates the handler logic
func makeVerifyPasswordHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.VerifyPasswordEndpoint, decodeVerifyPasswordRequest, encodeVerifyPasswordResponse, options...)
}

// decodeVerifyPasswordResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain VerifyPassword request.
// TODO implement the decoder
func decodeVerifyPasswordRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeVerifyPasswordResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeVerifyPasswordResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) VerifyPassword(ctx context1.Context, req *pb.VerifyPasswordRequest) (*pb.VerifyPasswordReply, error) {
	_, rep, err := g.verifyPassword.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.VerifyPasswordReply), nil
}
