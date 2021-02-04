package grpc

import (
	"context"

	"github.com/lucasres/code-pix/application/grpc/pb"
	"github.com/lucasres/code-pix/application/usecase"
)

type PixGrpcService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixeServiceServer
}

//RegisterPixKey implementa o service de criar uma key
func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreateResult, error) {
	key, err := p.PixUseCase.RegisterKey(in.Key, in.Kind, in.AccountID)
	if err != nil {
		return &pb.PixKeyCreateResult{
			Status: "not created",
			Error:  err.Error(),
			Id:     "",
		}, err
	}

	response := pb.PixKeyCreateResult{
		Error:  "",
		Id:     key.ID,
		Status: key.Status,
	}

	return &response, nil
}

//Find implementa a funcao do grpc para encontra uma pix key
func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	key, err := p.PixUseCase.FindKey(in.Kind, in.Key)

	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id: key.ID,
		Account: &pb.Account{
			AccountID:     key.AccountID,
			OwnerName:     key.Account.OwnerName,
			AccountNumber: key.Account.Number,
			BankName:      key.Account.Bank.Name,
			BankId:        key.Account.BankID,
			CreatedAt:     key.Account.CreatedAt.String(),
		},
		Kind:      key.Kind,
		Key:       key.Key,
		CreatedAt: key.CreatedAt.String(),
	}, nil
}

//NewPixGrpcService cria uma nova instancia de PixServiceGrpc
func NewPixGrpcService(usecase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUseCase: usecase,
	}
}
