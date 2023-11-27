package utils

import pb "github.com/raedmajeed/booking-service/pkg/pb"

// dom "github.com/raedmajeed/booking-service/pkg/DOM"
// pb "github.com/raedmajeed/booking-service/pkg/pb"

func ConvertLoginRequestToResponse(token string, p *pb.LoginRequest) *pb.LoginResponse {
	return &pb.LoginResponse{
		Email: p.Email,
		Token: token,
	}
}
