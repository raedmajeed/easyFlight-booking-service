package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/utils"
	"gorm.io/gorm"
	"log"
	"time"
)

func (svc *BookingServiceStruct) UserLogin(p *pb.LoginRequest) (string, error) {
	user, err := svc.repo.FindUserByEmail(p.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return "", err
		} else {
			log.Printf("unable to login %v, err: %v", p.Email, err.Error())
			return "", err
		}
	}

	check := utils.CheckPasswordMatch([]byte(user.Password), p.Password)
	if !check {
		log.Printf("password mismatch for user %v", p.Email)
		return "", fmt.Errorf("password mismatch for user %v", p.Email)
	}

	token, err := utils.GenerateToken(p.Email, p.Role, svc.cfg)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", p.Email, err.Error())
		return "", err
	}
	return token, err
}

func (svc *BookingServiceStruct) RegisterUserSvc(p *pb.UserRequest) (*dom.UserData, error) {
	hashPassword, err := utils.HashPassword(p.Password)
	if err != nil {
		log.Printf("unable to hash password in RegisterAirlineSvc() - service, err: %v", err.Error())
		return nil, err
	}
	user := &dom.UserData{
		Phone:    p.PhoneNumber,
		Email:    p.Email,
		Password: string(hashPassword),
		Name:     p.UserName,
	}
	// send otp to phone number
	resp, err := svc.twilio.SentTwilioOTP(p.PhoneNumber)
	if err != nil {
		return nil, err
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}
	userData, err := json.Marshal(&user)
	if err != nil {
		log.Printf("error parsing JSON, err: %v", err.Error())
		return nil, err
	}

	registerUser := fmt.Sprintf("register_user_%v", p.Email)
	svc.redis.Set(context.Background(), registerUser, userData, time.Minute*2)
	return user, nil
}

func (svc *BookingServiceStruct) VerifyUserRequest(p *pb.OTPRequest) (*dom.UserData, error) {
	registerUser := fmt.Sprintf("register_user_%v", p.Email)
	redisVal := svc.redis.Get(context.Background(), registerUser)

	if redisVal.Err() != nil {
		log.Printf("unable to get value from redis err: %v", redisVal.Err().Error())
		return nil, redisVal.Err()
	}

	var userData dom.UserData
	err := json.Unmarshal([]byte(redisVal.Val()), &userData)
	if err != nil {
		log.Println("unable to unmarshal json")
		return nil, err
	}

	code := fmt.Sprintf("%v", p.Otp)
	resp, err := svc.twilio.VerifyTwilioOTP(userData.Phone, code)
	if err != nil {
		return nil, err
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}

	_, err = svc.repo.FindUserByEmail(userData.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Existing record found  of airline %v", p.Email)
		return nil, errors.New("user already exists")
	}

	user, err := svc.repo.CreateUser(&userData)
	if err != nil {
		return nil, err
	}
	return user, nil
}
