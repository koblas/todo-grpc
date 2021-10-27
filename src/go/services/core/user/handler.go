package user

import (
	"log"

	"github.com/google/uuid"
	genpb "github.com/koblas/grpc-todo/genpb/core"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password []byte
}

// Server represents the gRPC server
type UserServer struct {
	genpb.UnimplementedUserServiceServer

	users []User
}

func NewUserServer() *UserServer {
	return &UserServer{
		users: []User{},
	}
}

func (s *UserServer) getById(id string) *User {
	for _, u := range s.users {
		if u.ID == id {
			return &u
		}
	}

	return nil
}

func (s *UserServer) getByEmail(email string) *User {
	for _, u := range s.users {
		log.Printf("CHECKING ", email, u.Email)
		if u.Email == email {
			return &u
		}
	}

	return nil
}

func (s *UserServer) FindBy(ctx context.Context, params *genpb.FindParam) (*genpb.UserEither, error) {
	log.Printf("Received find %s", params.Email)

	user := s.getByEmail(params.Email)

	if user == nil {
		return &genpb.UserEither{Error: &genpb.Error{Message: "Not Found"}}, nil
	}

	log.Printf("found id=%s", user.ID)

	return &genpb.UserEither{
		User: &genpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (s *UserServer) Create(ctx context.Context, params *genpb.CreateParam) (*genpb.UserEither, error) {
	log.Printf("Received create %s", params.Email)

	if s.getByEmail(params.Email) != nil {
		return &genpb.UserEither{Error: &genpb.Error{Message: "Already exists"}}, nil
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		ID:       uuid.New().String(),
		Name:     params.Name,
		Email:    params.Email,
		Password: pass,
	}

	s.users = append(s.users, user)

	return &genpb.UserEither{
		User: &genpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (s *UserServer) Update(ctx context.Context, params *genpb.UpdateParam) (*genpb.UserEither, error) {
	log.Printf("Received find %s", params.UserId)

	user := s.getById(params.UserId)

	if user == nil {
		return &genpb.UserEither{Error: &genpb.Error{Message: "Not Found"}}, nil
	}

	return &genpb.UserEither{
		User: &genpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (s *UserServer) ComparePassword(ctx context.Context, params *genpb.AuthenticateParam) (*genpb.UserEither, error) {
	log.Printf("Received comparePassowrd %s", params.UserId)

	user := s.getById(params.UserId)

	if user == nil {
		return &genpb.UserEither{Error: &genpb.Error{Message: "Not Found"}}, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return &genpb.UserEither{Error: &genpb.Error{Message: "No match"}}, nil
	}

	return &genpb.UserEither{
		User: &genpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
