package user

import (
	"context"
	"fmt"
	"time"

	"github.com/bufbuild/connect-go"
	userv1 "github.com/koblas/grpc-todo/gen/core/user/v1"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
)

const MAX_NAME_LEN = 200
const ADMIN_ROLE = "admin"

var teamStatusToPbStatus = map[TeamStatus]userv1.TeamStatus{
	TeamStatus_UNSET:   userv1.TeamStatus_TEAM_STATUS_UNSPECIFIED,
	TeamStatus_ACTIVE:  userv1.TeamStatus_TEAM_STATUS_ACTIVE,
	TeamStatus_INVITED: userv1.TeamStatus_TEAM_STATUS_INVITED,
}

func marshalMember(member *TeamMember, team *Team) *userv1.TeamMember {
	if member == nil {
		return &userv1.TeamMember{
			TeamId:   team.TeamId,
			TeamName: team.Name,
			Role:     ADMIN_ROLE,
			Status:   userv1.TeamStatus_TEAM_STATUS_ACTIVE,
		}
	}

	return &userv1.TeamMember{
		Id:       member.MemberId,
		UserId:   member.UserId,
		TeamId:   team.TeamId,
		TeamName: team.Name,
		Status:   teamStatusToPbStatus[member.Status],
		Role:     member.Role,
	}
}

func (s *UserServer) isTeamAdmin(ctx context.Context, teamId, userId string) error {
	member, err := s.store.TeamGetMember(ctx, teamId, userId)
	if member == nil {
		return bufcutil.PermissionDeniedError("user is not member of team")
	}
	if err != nil {
		return bufcutil.InternalError(err)
	}
	if member.Role != ADMIN_ROLE || member.Status != TeamStatus_ACTIVE {
		return bufcutil.FailedPreconditionError("user is not admin of team")
	}

	return nil
}

func (s *UserServer) isTeamMember(ctx context.Context, teamId, userId string) error {
	member, err := s.store.TeamGetMember(ctx, teamId, userId)
	if member == nil {
		return bufcutil.PermissionDeniedError("user is not member of team")
	}
	if err != nil {
		return bufcutil.InternalError(err)
	}
	if member.Status != TeamStatus_ACTIVE {
		return bufcutil.FailedPreconditionError("user is not active member of team")
	}

	return nil
}

func (s *UserServer) TeamCreate(ctx context.Context, request *connect.Request[userv1.TeamCreateRequest]) (*connect.Response[userv1.TeamCreateResponse], error) {
	log := logger.FromContext(ctx)
	name := request.Msg.Name

	if len(name) == 0 || len(name) > MAX_NAME_LEN {
		return nil, bufcutil.InvalidArgumentError("name", fmt.Sprintf("name must be between 1 and %d characters", MAX_NAME_LEN))
	}

	team, err := s.store.TeamCreate(ctx, name, TeamMember{
		UserId: request.Msg.UserId,
		Role:   ADMIN_ROLE,
		Status: TeamStatus_ACTIVE,
	})
	if err != nil {
		log.With(zap.Error(err)).Error("TeamCreate failed")
		return nil, bufcutil.InternalError(err, "failed to create team")
	}

	return connect.NewResponse(&userv1.TeamCreateResponse{
		Team: marshalMember(nil, team),
	}), nil
}

func (s *UserServer) TeamDelete(ctx context.Context, request *connect.Request[userv1.TeamDeleteRequest]) (*connect.Response[userv1.TeamDeleteResponse], error) {
	log := logger.FromContext(ctx)
	userId := request.Msg.UserId
	teamId := request.Msg.TeamId

	if err := s.isTeamAdmin(ctx, teamId, userId); err != nil {
		return nil, err
	}

	if err := s.store.TeamDelete(ctx, teamId); err != nil {
		log.With(zap.Error(err)).Error("TeamDelete failed")
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&userv1.TeamDeleteResponse{}), nil
}

func (s *UserServer) TeamList(ctx context.Context, request *connect.Request[userv1.TeamListRequest]) (*connect.Response[userv1.TeamListResponse], error) {
	log := logger.FromContext(ctx)
	userId := request.Msg.UserId

	memberships, err := s.store.TeamList(ctx, userId)
	if err != nil {
		log.With(zap.Error(err)).Error("TeamList failed")
		return nil, bufcutil.InternalError(err)
	}
	tlist := []*userv1.TeamMember{}
	for _, member := range memberships {
		team, err := s.store.TeamGet(ctx, member.TeamId)
		if err != nil {
			log.With("teamId", member.TeamId).Error("failed to fetch team")
			return nil, bufcutil.InternalError(err)
		}

		tlist = append(tlist, &userv1.TeamMember{
			Id:       member.MemberId,
			UserId:   member.UserId,
			TeamId:   team.TeamId,
			TeamName: team.Name,
			Status:   teamStatusToPbStatus[member.Status],
			Role:     member.Role,
		})
	}

	return connect.NewResponse(&userv1.TeamListResponse{Teams: tlist}), nil
}

func (s *UserServer) TeamAddMembers(ctx context.Context, request *connect.Request[userv1.TeamAddMembersRequest]) (*connect.Response[userv1.TeamAddMembersResponse], error) {
	log := logger.FromContext(ctx)
	userId := request.Msg.UserId
	teamId := request.Msg.TeamId

	if len(request.Msg.Members) == 0 {
		return connect.NewResponse(&userv1.TeamAddMembersResponse{}), nil
	}
	if err := s.isTeamAdmin(ctx, teamId, userId); err != nil {
		return nil, err
	}

	tuser := []TeamMember{}
	tnow := time.Now()
	for _, item := range request.Msg.Members {
		member := TeamMember{
			TeamId: teamId,
			UserId: item.UserId,
			Role:   item.Role,
		}

		member.Status = TeamStatus_ACTIVE
		switch item.Status {
		case userv1.TeamStatus_TEAM_STATUS_INVITED:
			member.Status = TeamStatus_INVITED
			member.InvitedBy = &userId
			member.InvitedOn = &tnow
		case userv1.TeamStatus_TEAM_STATUS_ACTIVE:
			member.Status = TeamStatus_ACTIVE
		default:
			return nil, bufcutil.InvalidArgumentError("status", "unknowns status provided")
		}

		tuser = append(tuser, member)
	}

	if err := s.store.TeamAddMembers(ctx, tuser...); err != nil {
		log.With(zap.Error(err)).Error("TeamAddMembers failed")
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&userv1.TeamAddMembersResponse{}), nil
}

func (s *UserServer) TeamAcceptInvite(ctx context.Context, request *connect.Request[userv1.TeamAcceptInviteRequest]) (*connect.Response[userv1.TeamAcceptInviteResponse], error) {
	log := logger.FromContext(ctx)
	userId := request.Msg.UserId
	teamIds := request.Msg.TeamIds

	for _, teamId := range teamIds {
		err := s.store.TeamAcceptInvite(ctx, teamId, userId)
		if err != nil {
			log.With(
				zap.String("userId", userId),
				zap.String("teamId", teamId),
				zap.Error(err),
			).Error("Unable to accept invite")

			return nil, bufcutil.InternalError(err)
		}
	}

	return connect.NewResponse(&userv1.TeamAcceptInviteResponse{}), nil
}

func (s *UserServer) TeamRemoveMembers(ctx context.Context, request *connect.Request[userv1.TeamRemoveMembersRequest]) (*connect.Response[userv1.TeamRemoveMembersResponse], error) {
	log := logger.FromContext(ctx)
	userId := request.Msg.UserId
	teamId := request.Msg.TeamId

	needAdmin := false
	for _, id := range request.Msg.UserIds {
		needAdmin = needAdmin || (id != userId)
	}
	// If one or more userIds in the list is not yourself, then you must be an admin
	//  to delete the relationship.  You can always remove yourself
	if needAdmin {
		if err := s.isTeamAdmin(ctx, teamId, userId); err == nil {
			return nil, err
		}
	}

	if err := s.store.TeamDeleteMembers(ctx, teamId, request.Msg.UserIds...); err != nil {
		log.With(zap.Error(err)).Error("TeamRemoveMembers failed")
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&userv1.TeamRemoveMembersResponse{}), nil
}

func (s *UserServer) TeamListMembers(ctx context.Context, request *connect.Request[userv1.TeamListMembersRequest]) (*connect.Response[userv1.TeamListMembersResponse], error) {
	log := logger.FromContext(ctx)
	userId := request.Msg.UserId
	teamId := request.Msg.TeamId

	if err := s.isTeamMember(ctx, teamId, userId); err != nil {
		return nil, err
	}

	members, err := s.store.TeamListMembers(ctx, teamId)
	if err != nil {
		log.With(zap.Error(err)).Error("TeamListMembers failed")
		return nil, bufcutil.InternalError(err)
	}

	mlist := []*userv1.TeamMember{}

	for _, item := range members {
		mlist = append(mlist, marshalMember(&item, &Team{TeamId: teamId}))
	}

	return connect.NewResponse(&userv1.TeamListMembersResponse{Members: mlist}), nil
}
