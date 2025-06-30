package testdata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_CreateUser(t *testing.T) {
	u := NewMockUserIFace(t)
	u.On("CreateUser", mock.Anything, User{}).Return(nil) // want `mock\.On\(\"CreateUser\", \.\.\.\) could be replaced with mock\.EXPECT\(\)\.CreateUser\(\.\.\.\)`

	err := u.CreateUser(context.Background(), User{})
	if err != nil {
		t.Error(err)
	}
}

func Test_GetUser(t *testing.T) {
	userMock := &MockUserIFace{}
	userMock. // want `mock\.On\(\"GetUser\", \.\.\.\) could be replaced with mock\.EXPECT\(\)\.GetUser\(\.\.\.\)`
			On(
			"GetUser",
			context.Background(),
			"test",
		).Return(User{}, nil)
}

func Test_Expecter(t *testing.T) {
	u := NewMockUserIFace(t)
	u.EXPECT().GetUser(context.Background(), "Bob").Return(User{}, nil) // OK
}

func Test_EmptyMethod(t *testing.T) {
	m := NewMockUserIFace(t)
	m.On("", mock.Anything, User{}).Return(nil) // ignore empty method name
}

func Test_InvalidMethod(t *testing.T) {
	i := NewMockUserIFace(t)
	// no function i.MOCK().DoesNotExist(...)
	i.On("DoesNotExist", mock.Anything, User{}, 123).Return(nil)
}

func Test_Void(t *testing.T) {
	u := NewMockUserIFace(t)
	u.On("Void")        // want `mock\.On\(\"Void\", \.\.\.\) could be replaced with mock\.EXPECT\(\)\.Void\(\.\.\.\)`
	u.On("Void").Once() // want `mock\.On\(\"Void\", \.\.\.\) could be replaced with mock\.EXPECT\(\)\.Void\(\.\.\.\)`
}

func Test_Count(t *testing.T) {
	u := NewMockUserIFace(t)
	u.On("CountUsers").Return(123) // want `mock\.On\(\"CountUsers\", \.\.\.\) could be replaced with mock\.EXPECT\(\)\.CountUsers\(\.\.\.\)`
}
