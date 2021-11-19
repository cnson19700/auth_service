package auth

import (
	"context"
	"log"

	"github.com/cnson19700/auth_service/config"
	"github.com/cnson19700/auth_service/package/auth"
	checkform "github.com/cnson19700/auth_service/package/checkForm"
	"github.com/cnson19700/auth_service/util"
	"github.com/cnson19700/auth_service/util/myerror"
	"github.com/cnson19700/user_service/model"
)

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Token    string `json:"token"`
	IP       string `json:"ip"`
}

func (u *Usecase) Register(ctx context.Context, req RegisterRequest) (*model.User, error) {
	//Email Format Error
	isMail, email := checkform.CheckFormatValue("email", req.Email)
	if !isMail {
		return &model.User{}, myerror.ErrEmailFormat(nil)
	}

	// if u.userRepo.CheckEmailExist(ctx, email) {
	// 	return nil, myerror.ErrEmailExist(nil)
	// }

	//password format error
	isPass, password := checkform.CheckFormatValue("password", req.Password)
	if !isPass {
		return &model.User{}, myerror.ErrEmailFormat(nil)
	}

	//password format error
	isFullName, fullname := checkform.CheckFormatValue("full_name", req.FullName)
	if !isFullName {
		return &model.User{}, myerror.ErrEmailFormat(nil)
	}

	// All data validated
	responseCaptcha := req.Token
	cfg := config.GetConfig()

	husCaptcha := util.HusCaptcha{PrivateKey: cfg.Captcha.SecretKey}
	result, _ := husCaptcha.Verify(req.IP, "register", responseCaptcha, 0.7)

	if !result {
		return nil, myerror.ErrInvalidCaptcha(nil)
	}

	passwordHash, err := auth.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	var user = &model.User{
		FullName: fullname,
		Password: passwordHash,
		Age:      req.Age,
		Email:    email,
	}

	res, err := u.userRepo.Create(ctx, user)

	if err != nil {
		return &model.User{}, nil
	}
	return res, nil
}
