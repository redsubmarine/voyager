package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	mockdb "github.com/yangoneseok/voyager/db/mock"
	db "github.com/yangoneseok/voyager/db/sqlc"
	"github.com/yangoneseok/voyager/util"
	"go.uber.org/mock/gomock"
)

func TestGetUserAPI(t *testing.T) {
	admin := randomAdmin()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	// build stub
	store.EXPECT().
		GetUser(gomock.Any(), gomock.Eq(admin.ID)).
		Times(1).
		Return(admin, nil)

	// start test server and send request
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/users/%d", admin.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchUser(t, recorder.Body, admin)
}

func randomAdmin() db.User {
	return db.User{
		ID:       util.RandomInt(1, 1000),
		Username: util.RandomName(),
		Password: util.RandomString(20),
		Email:    util.RandomEmail(),
		Role:     "admin",
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
}
