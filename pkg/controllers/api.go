package controllers

import (
	"context"

	"github.com/katallaxie/fiber-goth/pkg/apis"
)

var _ apis.StrictServerInterface = (*APIController)(nil)

type APIController struct{}

func NewAPIController() *APIController {
	return &APIController{}
}

// (GET /account-info)
func (c *APIController) GetAccountInfo(ctx context.Context, request apis.GetAccountInfoRequestObject) (apis.GetAccountInfoResponseObject, error) {
	return apis.GetAccountInfo200JSONResponse{}, nil
}

// (POST /change-email)
func (c *APIController) ChangeEmail(ctx context.Context, request apis.ChangeEmailRequestObject) (apis.ChangeEmailResponseObject, error) {
	return apis.ChangeEmail200JSONResponse{}, nil
}

// (POST /change-password)
func (c *APIController) ChangePassword(ctx context.Context, request apis.ChangePasswordRequestObject) (apis.ChangePasswordResponseObject, error) {
	return apis.ChangePassword200JSONResponse{}, nil
}

// (POST /delete-user)
func (c *APIController) DeleteUser(ctx context.Context, request apis.DeleteUserRequestObject) (apis.DeleteUserResponseObject, error) {
	return apis.DeleteUser200JSONResponse{}, nil
}

// (GET /delete-user/callback)
func (c *APIController) GetDeleteUserCallback(ctx context.Context, request apis.GetDeleteUserCallbackRequestObject) (apis.GetDeleteUserCallbackResponseObject, error) {
	return apis.GetDeleteUserCallback200JSONResponse{}, nil
}

// (GET /error)
func (c *APIController) GetError(ctx context.Context, request apis.GetErrorRequestObject) (apis.GetErrorResponseObject, error) {
	return apis.GetError200TexthtmlResponse{}, nil
}

// (POST /get-access-token)
func (c *APIController) PostGetAccessToken(ctx context.Context, request apis.PostGetAccessTokenRequestObject) (apis.PostGetAccessTokenResponseObject, error) {
	return apis.PostGetAccessToken200JSONResponse{}, nil
}

// (GET /get-session)
func (c *APIController) GetSession(ctx context.Context, request apis.GetSessionRequestObject) (apis.GetSessionResponseObject, error) {
	return apis.GetSession200JSONResponse{}, nil
}

// (POST /link-social)
func (c *APIController) LinkSocialAccount(ctx context.Context, request apis.LinkSocialAccountRequestObject) (apis.LinkSocialAccountResponseObject, error) {
	return apis.LinkSocialAccount200JSONResponse{}, nil
}

// (GET /list-accounts)
func (c *APIController) ListUserAccounts(ctx context.Context, request apis.ListUserAccountsRequestObject) (apis.ListUserAccountsResponseObject, error) {
	return apis.ListUserAccounts200JSONResponse{}, nil
}

// (GET /list-sessions)
func (c *APIController) ListUserSessions(ctx context.Context, request apis.ListUserSessionsRequestObject) (apis.ListUserSessionsResponseObject, error) {
	return apis.ListUserSessions200JSONResponse{}, nil
}

// (GET /ok)
func (c *APIController) GetOk(ctx context.Context, request apis.GetOkRequestObject) (apis.GetOkResponseObject, error) {
	return apis.GetOk200JSONResponse{}, nil
}

// (POST /organization/accept-invitation)
func (c *APIController) PostOrganizationAcceptInvitation(ctx context.Context, request apis.PostOrganizationAcceptInvitationRequestObject) (apis.PostOrganizationAcceptInvitationResponseObject, error) {
	return apis.PostOrganizationAcceptInvitation200JSONResponse{}, nil
}

// (POST /organization/add-team-member)
func (c *APIController) PostOrganizationAddTeamMember(ctx context.Context, request apis.PostOrganizationAddTeamMemberRequestObject) (apis.PostOrganizationAddTeamMemberResponseObject, error) {
	return apis.PostOrganizationAddTeamMember200JSONResponse{}, nil
}

// (POST /organization/cancel-invitation)
func (c *APIController) PostOrganizationCancelInvitation(ctx context.Context, request apis.PostOrganizationCancelInvitationRequestObject) (apis.PostOrganizationCancelInvitationResponseObject, error) {
	return nil, nil
}

// (POST /organization/check-slug)
func (c *APIController) PostOrganizationCheckSlug(ctx context.Context, request apis.PostOrganizationCheckSlugRequestObject) (apis.PostOrganizationCheckSlugResponseObject, error) {
	return nil, nil
}

// (POST /organization/create)
func (c *APIController) PostOrganizationCreate(ctx context.Context, request apis.PostOrganizationCreateRequestObject) (apis.PostOrganizationCreateResponseObject, error) {
	return nil, nil
}

// (POST /organization/create-team)
func (c *APIController) PostOrganizationCreateTeam(ctx context.Context, request apis.PostOrganizationCreateTeamRequestObject) (apis.PostOrganizationCreateTeamResponseObject, error) {
	return nil, nil
}

// (POST /organization/delete)
func (c *APIController) PostOrganizationDelete(ctx context.Context, request apis.PostOrganizationDeleteRequestObject) (apis.PostOrganizationDeleteResponseObject, error) {
	return nil, nil
}

// (GET /organization/get-active-member)
func (c *APIController) GetOrganizationGetActiveMember(ctx context.Context, request apis.GetOrganizationGetActiveMemberRequestObject) (apis.GetOrganizationGetActiveMemberResponseObject, error) {
	return apis.GetOrganizationGetActiveMember200JSONResponse{}, nil
}

// (GET /organization/get-active-member-role)
func (c *APIController) GetOrganizationGetActiveMemberRole(ctx context.Context, request apis.GetOrganizationGetActiveMemberRoleRequestObject) (apis.GetOrganizationGetActiveMemberRoleResponseObject, error) {
	return nil, nil
}

// (GET /organization/get-full-organization)
func (c *APIController) GetOrganization(ctx context.Context, request apis.GetOrganizationRequestObject) (apis.GetOrganizationResponseObject, error) {
	return apis.GetOrganization200JSONResponse{}, nil
}

// (GET /organization/get-invitation)
func (c *APIController) GetOrganizationGetInvitation(ctx context.Context, request apis.GetOrganizationGetInvitationRequestObject) (apis.GetOrganizationGetInvitationResponseObject, error) {
	return apis.GetOrganizationGetInvitation200JSONResponse{}, nil
}

// (POST /organization/has-permission)
func (c *APIController) PostOrganizationHasPermission(ctx context.Context, request apis.PostOrganizationHasPermissionRequestObject) (apis.PostOrganizationHasPermissionResponseObject, error) {
	return nil, nil
}

// (POST /organization/invite-member)
func (c *APIController) CreateOrganizationInvitation(ctx context.Context, request apis.CreateOrganizationInvitationRequestObject) (apis.CreateOrganizationInvitationResponseObject, error) {
	return nil, nil
}

// (POST /organization/leave)
func (c *APIController) PostOrganizationLeave(ctx context.Context, request apis.PostOrganizationLeaveRequestObject) (apis.PostOrganizationLeaveResponseObject, error) {
	return nil, nil
}

// (GET /organization/list)
func (c *APIController) GetOrganizationList(ctx context.Context, request apis.GetOrganizationListRequestObject) (apis.GetOrganizationListResponseObject, error) {
	return apis.GetOrganizationList200JSONResponse{}, nil
}

// (GET /organization/list-invitations)
func (c *APIController) GetOrganizationListInvitations(ctx context.Context, request apis.GetOrganizationListInvitationsRequestObject) (apis.GetOrganizationListInvitationsResponseObject, error) {
	return nil, nil
}

// (GET /organization/list-members)
func (c *APIController) GetOrganizationListMembers(ctx context.Context, request apis.GetOrganizationListMembersRequestObject) (apis.GetOrganizationListMembersResponseObject, error) {
	return nil, nil
}

// (GET /organization/list-team-members)
func (c *APIController) GetOrganizationListTeamMembers(ctx context.Context, request apis.GetOrganizationListTeamMembersRequestObject) (apis.GetOrganizationListTeamMembersResponseObject, error) {
	return apis.GetOrganizationListTeamMembers200JSONResponse{}, nil
}

// (GET /organization/list-teams)
func (c *APIController) GetOrganizationListTeams(ctx context.Context, request apis.GetOrganizationListTeamsRequestObject) (apis.GetOrganizationListTeamsResponseObject, error) {
	return apis.GetOrganizationListTeams200JSONResponse{}, nil
}

// (GET /organization/list-user-invitations)
func (c *APIController) GetOrganizationListUserInvitations(ctx context.Context, request apis.GetOrganizationListUserInvitationsRequestObject) (apis.GetOrganizationListUserInvitationsResponseObject, error) {
	return apis.GetOrganizationListUserInvitations200JSONResponse{}, nil
}

// (GET /organization/list-user-teams)
func (c *APIController) GetOrganizationListUserTeams(ctx context.Context, request apis.GetOrganizationListUserTeamsRequestObject) (apis.GetOrganizationListUserTeamsResponseObject, error) {
	return apis.GetOrganizationListUserTeams200JSONResponse{}, nil
}

// (POST /organization/reject-invitation)
func (c *APIController) PostOrganizationRejectInvitation(ctx context.Context, request apis.PostOrganizationRejectInvitationRequestObject) (apis.PostOrganizationRejectInvitationResponseObject, error) {
	return nil, nil
}

// (POST /organization/remove-member)
func (c *APIController) PostOrganizationRemoveMember(ctx context.Context, request apis.PostOrganizationRemoveMemberRequestObject) (apis.PostOrganizationRemoveMemberResponseObject, error) {
	return nil, nil
}

// (POST /organization/remove-team)
func (c *APIController) PostOrganizationRemoveTeam(ctx context.Context, request apis.PostOrganizationRemoveTeamRequestObject) (apis.PostOrganizationRemoveTeamResponseObject, error) {
	return nil, nil
}

// (POST /organization/remove-team-member)
func (c *APIController) PostOrganizationRemoveTeamMember(ctx context.Context, request apis.PostOrganizationRemoveTeamMemberRequestObject) (apis.PostOrganizationRemoveTeamMemberResponseObject, error) {
	return nil, nil
}

// (POST /organization/set-active)
func (c *APIController) SetActiveOrganization(ctx context.Context, request apis.SetActiveOrganizationRequestObject) (apis.SetActiveOrganizationResponseObject, error) {
	return nil, nil
}

// (POST /organization/set-active-team)
func (c *APIController) PostOrganizationSetActiveTeam(ctx context.Context, request apis.PostOrganizationSetActiveTeamRequestObject) (apis.PostOrganizationSetActiveTeamResponseObject, error) {
	return apis.PostOrganizationSetActiveTeam200JSONResponse{}, nil
}

// (POST /organization/update)
func (c *APIController) PostOrganizationUpdate(ctx context.Context, request apis.PostOrganizationUpdateRequestObject) (apis.PostOrganizationUpdateResponseObject, error) {
	return nil, nil
}

// (POST /organization/update-member-role)
func (c *APIController) UpdateOrganizationMemberRole(ctx context.Context, request apis.UpdateOrganizationMemberRoleRequestObject) (apis.UpdateOrganizationMemberRoleResponseObject, error) {
	return apis.UpdateOrganizationMemberRole200JSONResponse{}, nil
}

// (POST /organization/update-team)
func (c *APIController) PostOrganizationUpdateTeam(ctx context.Context, request apis.PostOrganizationUpdateTeamRequestObject) (apis.PostOrganizationUpdateTeamResponseObject, error) {
	return nil, nil
}

// (POST /refresh-token)
func (c *APIController) PostRefreshToken(ctx context.Context, request apis.PostRefreshTokenRequestObject) (apis.PostRefreshTokenResponseObject, error) {
	return nil, nil
}

// (POST /request-password-reset)
func (c *APIController) RequestPasswordReset(ctx context.Context, request apis.RequestPasswordResetRequestObject) (apis.RequestPasswordResetResponseObject, error) {
	return apis.RequestPasswordReset200JSONResponse{}, nil
}

// (POST /reset-password)
func (c *APIController) ResetPassword(ctx context.Context, request apis.ResetPasswordRequestObject) (apis.ResetPasswordResponseObject, error) {
	return nil, nil
}

// (GET /reset-password/{token})
func (c *APIController) ResetPasswordCallback(ctx context.Context, request apis.ResetPasswordCallbackRequestObject) (apis.ResetPasswordCallbackResponseObject, error) {
	return apis.ResetPasswordCallback200JSONResponse{}, nil
}

// (POST /revoke-other-sessions)
func (c *APIController) PostRevokeOtherSessions(ctx context.Context, request apis.PostRevokeOtherSessionsRequestObject) (apis.PostRevokeOtherSessionsResponseObject, error) {
	return apis.PostRevokeOtherSessions200JSONResponse{}, nil
}

// (POST /revoke-session)
func (c *APIController) PostRevokeSession(ctx context.Context, request apis.PostRevokeSessionRequestObject) (apis.PostRevokeSessionResponseObject, error) {
	return nil, nil
}

// (POST /revoke-sessions)
func (c *APIController) PostRevokeSessions(ctx context.Context, request apis.PostRevokeSessionsRequestObject) (apis.PostRevokeSessionsResponseObject, error) {
	return nil, nil
}

// (POST /send-verification-email)
func (c *APIController) SendVerificationEmail(ctx context.Context, request apis.SendVerificationEmailRequestObject) (apis.SendVerificationEmailResponseObject, error) {
	return apis.SendVerificationEmail200JSONResponse{}, nil
}

// (POST /sign-in/email)
func (c *APIController) SignInEmail(ctx context.Context, request apis.SignInEmailRequestObject) (apis.SignInEmailResponseObject, error) {
	return apis.SignInEmail200JSONResponse{}, nil
}

// (POST /sign-in/social)
func (c *APIController) SocialSignIn(ctx context.Context, request apis.SocialSignInRequestObject) (apis.SocialSignInResponseObject, error) {
	return apis.SocialSignIn200JSONResponse{}, nil
}

// (POST /sign-out)
func (c *APIController) SignOut(ctx context.Context, request apis.SignOutRequestObject) (apis.SignOutResponseObject, error) {
	return apis.SignOut200JSONResponse{}, nil
}

// (POST /sign-up/email)
func (c *APIController) SignUpWithEmailAndPassword(ctx context.Context, request apis.SignUpWithEmailAndPasswordRequestObject) (apis.SignUpWithEmailAndPasswordResponseObject, error) {
	return apis.SignUpWithEmailAndPassword200JSONResponse{}, nil
}

// (POST /unlink-account)
func (c *APIController) PostUnlinkAccount(ctx context.Context, request apis.PostUnlinkAccountRequestObject) (apis.PostUnlinkAccountResponseObject, error) {
	return nil, nil
}

// (POST /update-user)
func (c *APIController) UpdateUser(ctx context.Context, request apis.UpdateUserRequestObject) (apis.UpdateUserResponseObject, error) {
	return apis.UpdateUser200JSONResponse{}, nil
}

// (GET /verify-email)
func (c *APIController) GetVerifyEmail(ctx context.Context, request apis.GetVerifyEmailRequestObject) (apis.GetVerifyEmailResponseObject, error) {
	return apis.GetVerifyEmail200JSONResponse{}, nil
}
