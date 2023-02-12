package repository_test

import (
	"errors"
	"testing"

	"github.com/go-git/go-git/v5/config"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

type MockRepositoryUser struct {
	name  string
	email string
	err   error
}

func (m MockRepositoryUser) Config() (*config.Config, error) {
	var cfg config.Config

	cfg.User.Name = m.name
	cfg.User.Email = m.email

	return &cfg, m.err
}

func MockGlobalConfig(name, email string, err error) func(scope config.Scope) (*config.Config, error) {
	return func(scope config.Scope) (*config.Config, error) {
		var cfg config.Config

		cfg.User.Name = name
		cfg.User.Email = email

		return &cfg, err
	}
}

var errMockUser = errors.New("error")

func TestUser(t *testing.T) {
	t.Parallel()

	type userSet struct {
		user repository.User
		err  error
	}

	type args struct {
		repositoryUser userSet
		globalUser     userSet
		ignoreGlobal   bool
	}

	type want struct {
		users []repository.User
		err   error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "repository_both_name_email",
			args: args{
				repositoryUser: userSet{
					user: repository.User{Name: "John Doe", Email: "john.doe@example.com"},
					err:  nil,
				},
			},
			want: want{
				users: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
				},
				err: nil,
			},
		},
		{
			name: "repository_only_name",
			args: args{
				repositoryUser: userSet{
					user: repository.User{Name: "John Doe"},
					err:  nil,
				},
			},
			want: want{
				users: []repository.User{
					{Name: "John Doe"},
				},
				err: nil,
			},
		},
		{
			name: "repository_only_email",
			args: args{
				repositoryUser: userSet{
					user: repository.User{Email: "john.doe@example.com"},
					err:  nil,
				},
			},
			want: want{
				users: []repository.User{
					{Email: "john.doe@example.com"},
				},
				err: nil,
			},
		},
		{
			name: "global",
			args: args{
				globalUser: userSet{
					user: repository.User{Name: "Jim Bob", Email: "jim.bob@example.org"},
					err:  nil,
				},
			},
			want: want{
				users: []repository.User{
					{Name: "Jim Bob", Email: "jim.bob@example.org"},
				},
				err: nil,
			},
		},
		{
			name: "both_repository_global",
			args: args{
				repositoryUser: userSet{
					user: repository.User{Name: "John Doe", Email: "john.doe@example.com"},
					err:  nil,
				},
				globalUser: userSet{
					user: repository.User{Name: "Jim Bob", Email: "jim.bob@example.org"},
					err:  nil,
				},
			},
			want: want{
				users: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
					{Name: "Jim Bob", Email: "jim.bob@example.org"},
				},
				err: nil,
			},
		},
		{
			name: "both_repository_global_ignore",
			args: args{
				repositoryUser: userSet{
					user: repository.User{Name: "John Doe", Email: "john.doe@example.com"},
					err:  nil,
				},
				globalUser: userSet{
					user: repository.User{Name: "Jim Bob", Email: "jim.bob@example.org"},
					err:  nil,
				},
				ignoreGlobal: true,
			},
			want: want{
				users: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
				},
				err: nil,
			},
		},
		{
			name: "global_ignore",
			args: args{
				globalUser: userSet{
					user: repository.User{Name: "Jim Bob", Email: "jim.bob@example.org"},
					err:  nil,
				},
				ignoreGlobal: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "repository_error",
			args: args{
				repositoryUser: userSet{
					err: errMockUser,
				},
			},
			want: want{
				err: errMockUser,
			},
		},
		{
			name: "global_error",
			args: args{
				globalUser: userSet{
					err: errMockUser,
				},
			},
			want: want{
				err: errMockUser,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var r repository.Repository

			r.Configer = MockRepositoryUser{
				name:  tt.args.repositoryUser.user.Name,
				email: tt.args.repositoryUser.user.Email,
				err:   tt.args.repositoryUser.err,
			}

			r.GlobalConfig = MockGlobalConfig(tt.args.globalUser.user.Name, tt.args.globalUser.user.Email, tt.args.globalUser.err)

			if tt.args.ignoreGlobal {
				r.IgnoreGlobalConfig()
			}

			users, err := r.Users()
			if tt.want.err != nil {
				assert.ErrorContains(t, err, tt.want.err.Error())
				return
			}
			assert.Nil(t, err)

			if !assert.Len(t, tt.want.users, len(users)) {
				t.FailNow()
			}

			for i := range users {
				assert.Equal(t, tt.want.users[i].Name, users[i].Name)
				assert.Equal(t, tt.want.users[i].Email, users[i].Email)
			}
		})
	}
}

func TestIgnoreGlobalConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input func(config.Scope) (*config.Config, error)
	}{
		{
			name: "default",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var r repository.Repository

			r.IgnoreGlobalConfig()

			got, err := r.GlobalConfig(config.GlobalScope)
			assert.NoError(t, err)
			assert.Equal(t, &config.Config{}, got)
		})
	}
}
