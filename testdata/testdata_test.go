package testdata

import (
	"log"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_CreateUser(t *testing.T) {
	u := NewMockUserIFace(t)
	u.On("CreateUser", mock.Anything, User{}).Return(nil) // want `mock\.On\(\"CreateUser\", \.\.\.\) could be replaced with mock\.EXPECT\(\)\.CreateUser\(\.\.\.\)`

	err := u.CreateUser(t.Context(), User{})
	if err != nil {
		log.Fatal(err)
	}
}

func Test_GetUser(t *testing.T) {
	userMock := &MockUserIFace{}
	userMock. // want `mock\.On\(\"GetUser\", \.\.\.\) could be replaced with mock\.EXPECT\(\)\.GetUser\(\.\.\.\)`
			On(
			"GetUser",
			t.Context(),
			"test",
		).Return(User{}, nil)
}

func Test_Expecter(t *testing.T) {
	u := NewMockUserIFace(t)
	u.EXPECT().GetUser(t.Context(), "Bob").Return(User{}, nil) // OK
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
