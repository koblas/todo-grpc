package user

import (
	"log"

	"github.com/google/uuid"
	genpb "github.com/koblas/grpc-todo/genpb/core"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		log.Print("CHECKING ", email, u.Email)
		if u.Email == email {
			return &u
		}
	}

	return nil
}

func (s *UserServer) FindBy(ctx context.Context, params *genpb.FindParam) (*genpb.User, error) {
	log.Printf("Received find %s", params.Email)

	user := s.getByEmail(params.Email)

	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Email address not found")
	}

	log.Printf("found id=%s", user.ID)

	return &genpb.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServer) Create(ctx context.Context, params *genpb.CreateParam) (*genpb.User, error) {
	log.Printf("Received create %s", params.Email)

	if s.getByEmail(params.Email) != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Email address not found")
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

	return &genpb.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServer) Update(ctx context.Context, params *genpb.UpdateParam) (*genpb.User, error) {
	log.Printf("Received find %s", params.UserId)

	user := s.getById(params.UserId)

	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, "ID address not found")
	}

	return &genpb.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServer) ComparePassword(ctx context.Context, params *genpb.AuthenticateParam) (*genpb.User, error) {
	log.Printf("Received comparePassowrd %s", params.UserId)

	user := s.getById(params.UserId)

	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, "ID address not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Password mismatch")
	}

	return &genpb.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
