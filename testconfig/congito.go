package testconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoMock struct {
	CreateGroupFunc               func(ctx context.Context, params *cognitoidentityprovider.CreateGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateGroupOutput, error)
	AdminCreateUserFunc           func(ctx context.Context, params *cognitoidentityprovider.AdminCreateUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminCreateUserOutput, error)
	AdminAddUserToGroupFunc       func(ctx context.Context, params *cognitoidentityprovider.AdminAddUserToGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error)
	AdminGetUserFunc              func(ctx context.Context, params *cognitoidentityprovider.AdminGetUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminGetUserOutput, error)
	AdminListGroupsForUserFunc    func(ctx context.Context, params *cognitoidentityprovider.AdminListGroupsForUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error)
	ListUsersInGroupFunc          func(ctx context.Context, params *cognitoidentityprovider.ListUsersInGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
	GetGroupFunc                  func(ctx context.Context, params *cognitoidentityprovider.GetGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetGroupOutput, error)
	AdminRemoveUserFromGroupFunc  func(ctx context.Context, params *cognitoidentityprovider.AdminRemoveUserFromGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error)
	AdminUpdateUserAttributesFunc func(ctx context.Context, params *cognitoidentityprovider.AdminUpdateUserAttributesInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
	AdminSetUserPasswordFunc      func(ctx context.Context, params *cognitoidentityprovider.AdminSetUserPasswordInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error)
}

func (c *CognitoMock) CreateGroup(ctx context.Context, params *cognitoidentityprovider.CreateGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateGroupOutput, error) {
	if c.CreateGroupFunc != nil {
		return c.CreateGroupFunc(ctx, params)
	}
	return &cognitoidentityprovider.CreateGroupOutput{}, nil
}

func (c *CognitoMock) AdminCreateUser(ctx context.Context, params *cognitoidentityprovider.AdminCreateUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
	if c.AdminCreateUserFunc != nil {
		return c.AdminCreateUserFunc(ctx, params)
	}
	return &cognitoidentityprovider.AdminCreateUserOutput{}, nil
}

func (c *CognitoMock) AdminAddUserToGroup(ctx context.Context, params *cognitoidentityprovider.AdminAddUserToGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error) {
	if c.AdminAddUserToGroupFunc != nil {
		return c.AdminAddUserToGroupFunc(ctx, params)
	}
	return &cognitoidentityprovider.AdminAddUserToGroupOutput{}, nil
}

func (c *CognitoMock) AdminGetUser(ctx context.Context, params *cognitoidentityprovider.AdminGetUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	if c.AdminGetUserFunc != nil {
		return c.AdminGetUserFunc(ctx, params, optFns...)
	}
	return &cognitoidentityprovider.AdminGetUserOutput{}, nil
}

func (c *CognitoMock) AdminListGroupsForUser(ctx context.Context, params *cognitoidentityprovider.AdminListGroupsForUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error) {
	if c.AdminListGroupsForUserFunc != nil {
		return c.AdminListGroupsForUserFunc(ctx, params, optFns...)
	}
	return &cognitoidentityprovider.AdminListGroupsForUserOutput{}, nil
}

func (c *CognitoMock) ListUsersInGroup(ctx context.Context, params *cognitoidentityprovider.ListUsersInGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
	if c.ListUsersInGroupFunc != nil {
		return c.ListUsersInGroupFunc(ctx, params, optFns...)
	}
	return &cognitoidentityprovider.ListUsersInGroupOutput{}, nil
}

func (c *CognitoMock) GetGroup(ctx context.Context, params *cognitoidentityprovider.GetGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetGroupOutput, error) {
	if c.GetGroupFunc != nil {
		return c.GetGroupFunc(ctx, params, optFns...)
	}
	return &cognitoidentityprovider.GetGroupOutput{}, nil
}

func (c *CognitoMock) AdminRemoveUserFromGroup(ctx context.Context, params *cognitoidentityprovider.AdminRemoveUserFromGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error) {
	if c.AdminRemoveUserFromGroupFunc != nil {
		return c.AdminRemoveUserFromGroupFunc(ctx, params, optFns...)
	}
	return &cognitoidentityprovider.AdminRemoveUserFromGroupOutput{}, nil
}

func (c *CognitoMock) AdminUpdateUserAttributes(ctx context.Context, params *cognitoidentityprovider.AdminUpdateUserAttributesInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {
	if c.AdminUpdateUserAttributesFunc != nil {
		return c.AdminUpdateUserAttributesFunc(ctx, params, optFns...)
	}
	return &cognitoidentityprovider.AdminUpdateUserAttributesOutput{}, nil
}

func (c *CognitoMock) AdminSetUserPassword(ctx context.Context, params *cognitoidentityprovider.AdminSetUserPasswordInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error) {
	if c.AdminSetUserPasswordFunc != nil {
		return c.AdminSetUserPasswordFunc(ctx, params, optFns...)
	}
	return &cognitoidentityprovider.AdminSetUserPasswordOutput{}, nil
}
