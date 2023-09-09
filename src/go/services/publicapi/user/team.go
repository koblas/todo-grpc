package user

import (
	"errors"

	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/user/v1"
	userv1 "github.com/koblas/grpc-todo/gen/core/user/v1"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"golang.org/x/net/context"
)

var apiTeamStatus = map[userv1.TeamStatus]string{
	userv1.TeamStatus_TEAM_STATUS_ACTIVE:      "active",
	userv1.TeamStatus_TEAM_STATUS_INVITED:     "invited",
	userv1.TeamStatus_TEAM_STATUS_UNSPECIFIED: "*unknown*",
}

func coreTeamMemberToApi(input *userv1.TeamMember) *apiv1.TeamMember {
	output := apiv1.TeamMember{
		Id:       input.Id,
		UserId:   input.UserId,
		TeamName: input.TeamName,
		Role:     input.Role,
		Status:   apiTeamStatus[input.Status],
	}

	return &output
}

func (svc *UserServer) TeamCreate(ctx context.Context, request *connect.Request[apiv1.TeamCreateRequest]) (*connect.Response[apiv1.TeamCreateResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("TeamCreate BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	resp, err := svc.user.TeamCreate(ctx, connect.NewRequest(&userv1.TeamCreateRequest{
		UserId: userId,
		Name:   request.Msg.Name,
	}))

	if err != nil {
		log.With("error", err).Info("TeamCreate failed")
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&apiv1.TeamCreateResponse{
		Team: coreTeamMemberToApi(resp.Msg.Team),
	}), nil
}

func (svc *UserServer) TeamDelete(ctx context.Context, request *connect.Request[apiv1.TeamDeleteRequest]) (*connect.Response[apiv1.TeamDeleteResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("TeamDelete BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	_, err = svc.user.TeamDelete(ctx, connect.NewRequest(&userv1.TeamDeleteRequest{
		UserId: userId,
		TeamId: request.Msg.TeamId,
	}))

	if err != nil {
		log.With("error", err).Info("TeamDelete failed")
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&apiv1.TeamDeleteResponse{}), nil
}

func (svc *UserServer) TeamList(ctx context.Context, _ *connect.Request[apiv1.TeamListRequest]) (*connect.Response[apiv1.TeamListResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("TeamList BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	resp, err := svc.user.TeamList(ctx, connect.NewRequest(&userv1.TeamListRequest{
		UserId: userId,
	}))

	if err != nil {
		log.With("error", err).Info("TeamList failed")
		return nil, bufcutil.InternalError(err)
	}

	teams := []*apiv1.TeamMember{}
	for _, team := range resp.Msg.Teams {
		teams = append(teams, coreTeamMemberToApi(team))
	}

	return connect.NewResponse(&apiv1.TeamListResponse{
		Teams: teams,
	}), nil
}

func (svc *UserServer) TeamInvite(ctx context.Context, request *connect.Request[apiv1.TeamInviteRequest]) (*connect.Response[apiv1.TeamInviteResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("TeamInvite BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	var findBy *connect.Response[userv1.FindByResponse]

	// You can invite by email address or UserId (if they're already somebody you know)
	if request.Msg.Email != nil {
		findBy, err = svc.user.FindBy(ctx, connect.NewRequest(&userv1.FindByRequest{
			FindBy: &userv1.FindBy{
				Email: *request.Msg.Email,
			},
		}))
	} else if request.Msg.UserId != nil {
		findBy, err = svc.user.FindBy(ctx, connect.NewRequest(&userv1.FindByRequest{
			FindBy: &userv1.FindBy{
				UserId: *request.Msg.UserId,
			},
		}))
	} else {
		return nil, bufcutil.InvalidArgumentError("email", "email is not provided")
	}

	if err != nil && connect.CodeOf(err) != connect.CodeNotFound {
		log.With("error", err).Info("FindBy failed")
		if connect.CodeOf(err) != connect.CodeUnknown {
			return nil, bufcutil.InternalError(err)
		}
		return nil, err
	}

	var user *userv1.User
	if findBy != nil {
		user = findBy.Msg.User
	}
	if user == nil && request.Msg.Email != nil {
		create, err := svc.user.Create(ctx, connect.NewRequest(&userv1.CreateRequest{
			Email:  *request.Msg.Email,
			Status: userv1.UserStatus_USER_STATUS_INVITED,
		}))
		if err != nil {
			log.With("error", err).Info("Create failed")
			return nil, bufcutil.InternalError(err)
		}
		user = create.Msg.User
	}
	if user == nil {
		return nil, bufcutil.InvalidArgumentError("email", "unable to invite user")
	}
	if user.ClosedStatus != userv1.ClosedStatus_CLOSED_STATUS_UNSPECIFIED {
		return nil, bufcutil.InvalidArgumentError("user", "user not able to be added")
	}

	// Add this user to the team in INVITED state
	_, err = svc.user.TeamAddMembers(ctx, connect.NewRequest(&userv1.TeamAddMembersRequest{
		UserId: userId,
		TeamId: request.Msg.TeamId,
		Members: []*userv1.TeamMember{
			{
				UserId:    user.Id,
				TeamId:    request.Msg.TeamId,
				Status:    userv1.TeamStatus_TEAM_STATUS_INVITED,
				InvitedBy: &userId,
				Role:      "member",
			},
		},
	}))
	if err != nil {
		return nil, bufcutil.InternalError(err, "unable to invite user")
	}

	return connect.NewResponse(&apiv1.TeamInviteResponse{}), nil
}
