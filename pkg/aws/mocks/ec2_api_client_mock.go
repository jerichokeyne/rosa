// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/aws/api_interface/ec2_api_client.go
//
// Generated by this command:
//
//	mockgen-v0.4.0 -source=pkg/aws/api_interface/ec2_api_client.go -package=mocks -destination=pkg/aws/mocks/ec2_api_client_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	gomock "go.uber.org/mock/gomock"
)

// MockEc2ApiClient is a mock of Ec2ApiClient interface.
type MockEc2ApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockEc2ApiClientMockRecorder
}

// MockEc2ApiClientMockRecorder is the mock recorder for MockEc2ApiClient.
type MockEc2ApiClientMockRecorder struct {
	mock *MockEc2ApiClient
}

// NewMockEc2ApiClient creates a new mock instance.
func NewMockEc2ApiClient(ctrl *gomock.Controller) *MockEc2ApiClient {
	mock := &MockEc2ApiClient{ctrl: ctrl}
	mock.recorder = &MockEc2ApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEc2ApiClient) EXPECT() *MockEc2ApiClientMockRecorder {
	return m.recorder
}

// DescribeAvailabilityZones mocks base method.
func (m *MockEc2ApiClient) DescribeAvailabilityZones(ctx context.Context, params *ec2.DescribeAvailabilityZonesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeAvailabilityZonesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAvailabilityZones", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeAvailabilityZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAvailabilityZones indicates an expected call of DescribeAvailabilityZones.
func (mr *MockEc2ApiClientMockRecorder) DescribeAvailabilityZones(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAvailabilityZones", reflect.TypeOf((*MockEc2ApiClient)(nil).DescribeAvailabilityZones), varargs...)
}

// DescribeInstanceTypeOfferings mocks base method.
func (m *MockEc2ApiClient) DescribeInstanceTypeOfferings(ctx context.Context, params *ec2.DescribeInstanceTypeOfferingsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstanceTypeOfferings", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceTypeOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstanceTypeOfferings indicates an expected call of DescribeInstanceTypeOfferings.
func (mr *MockEc2ApiClientMockRecorder) DescribeInstanceTypeOfferings(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstanceTypeOfferings", reflect.TypeOf((*MockEc2ApiClient)(nil).DescribeInstanceTypeOfferings), varargs...)
}

// DescribeInstances mocks base method.
func (m *MockEc2ApiClient) DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstances", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstances indicates an expected call of DescribeInstances.
func (mr *MockEc2ApiClientMockRecorder) DescribeInstances(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockEc2ApiClient)(nil).DescribeInstances), varargs...)
}

// DescribeRouteTables mocks base method.
func (m *MockEc2ApiClient) DescribeRouteTables(ctx context.Context, params *ec2.DescribeRouteTablesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeRouteTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRouteTables", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeRouteTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRouteTables indicates an expected call of DescribeRouteTables.
func (mr *MockEc2ApiClientMockRecorder) DescribeRouteTables(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRouteTables", reflect.TypeOf((*MockEc2ApiClient)(nil).DescribeRouteTables), varargs...)
}

// DescribeSecurityGroups mocks base method.
func (m *MockEc2ApiClient) DescribeSecurityGroups(ctx context.Context, params *ec2.DescribeSecurityGroupsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSecurityGroups", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecurityGroups indicates an expected call of DescribeSecurityGroups.
func (mr *MockEc2ApiClientMockRecorder) DescribeSecurityGroups(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityGroups", reflect.TypeOf((*MockEc2ApiClient)(nil).DescribeSecurityGroups), varargs...)
}

// DescribeSubnets mocks base method.
func (m *MockEc2ApiClient) DescribeSubnets(ctx context.Context, params *ec2.DescribeSubnetsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSubnets", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubnets indicates an expected call of DescribeSubnets.
func (mr *MockEc2ApiClientMockRecorder) DescribeSubnets(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnets", reflect.TypeOf((*MockEc2ApiClient)(nil).DescribeSubnets), varargs...)
}

// DescribeVpcAttribute mocks base method.
func (m *MockEc2ApiClient) DescribeVpcAttribute(ctx context.Context, params *ec2.DescribeVpcAttributeInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpcAttribute", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcAttribute indicates an expected call of DescribeVpcAttribute.
func (mr *MockEc2ApiClientMockRecorder) DescribeVpcAttribute(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcAttribute", reflect.TypeOf((*MockEc2ApiClient)(nil).DescribeVpcAttribute), varargs...)
}
