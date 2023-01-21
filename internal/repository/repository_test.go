package repository_test

import (
	"errors"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

var (
	errMockOpen     = errors.New("error")
	errMockDescribe = errors.New("error")
)

func MockOpen(err error) func(string, *git.PlainOpenOptions) (*git.Repository, error) {
	return func(string, *git.PlainOpenOptions) (*git.Repository, error) {
		return nil, err
	}
}

func TestOpen(t *testing.T) {
	type args struct {
		err error
	}

	type want struct {
		err string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
		},
		{
			name: "open_repository_not_exists",
			args: args{
				err: git.ErrRepositoryNotExists,
			},
			want: want{
				err: "repository does not exist",
			},
		},
		{
			name: "open_error",
			args: args{
				err: errMockOpen,
			},
			want: want{
				err: "unable to open git repository: .: error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.New()
			repo.Opener = MockOpen(tt.args.err)

			err := repo.Open()
			if tt.want.err != "" {
				assert.NotNil(t, err)
				assert.Equal(t, tt.want.err, err.Error())
				return
			}
			assert.Nil(t, err)
		})
	}
}

func TestDescribe(t *testing.T) {
	type args struct {
		localBranch string
		userErr     error
		remoteErr   error
		headErr     error
		branchErr   error
	}

	type want struct {
		users   []repository.User
		remotes []string
		head    repository.Head
		branch  repository.Branch
		err     error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "valid",
			args: args{
				localBranch: "master",
			},
			want: want{
				head: repository.Head{
					Hash: "0000000000000000000000000000000000000000",
				},
				branch: repository.Branch{
					Local: "master",
				},
			},
		},
		{
			name: "error_user",
			args: args{
				userErr: errMockDescribe,
			},
			want: want{
				err: errMockDescribe,
			},
		},
		{
			name: "error_remote",
			args: args{
				remoteErr: errMockDescribe,
			},
			want: want{
				err: errMockDescribe,
			},
		},
		{
			name: "error_head",
			args: args{
				headErr: errMockDescribe,
			},
			want: want{
				err: errMockDescribe,
			},
		},
		{
			name: "error_branch",
			args: args{
				branchErr:   errMockDescribe,
				localBranch: "master",
			},
			want: want{
				head: repository.Head{
					Hash: "0000000000000000000000000000000000000000",
				},
				err: errMockDescribe,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.Repository{
				Configer:     MockRepositoryUser{err: tt.args.userErr},
				GlobalConfig: MockGlobalConfig("", "", nil),
				Remoter:      MockRepositoryRemote{err: tt.args.remoteErr},
				Header:       MockRepositoryHead{headErr: tt.args.headErr},
				Brancher:     MockRepositoryBranch{local: tt.args.localBranch, headErr: tt.args.branchErr},
			}

			d, err := r.Describe()
			if tt.want.err != nil {
				assert.NotNil(t, err)
				assert.ErrorContains(t, err, tt.want.err.Error())
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.want.users, d.Users)
			assert.Equal(t, tt.want.remotes, d.Remotes)
			assert.Equal(t, tt.want.head, d.Head)
			assert.Equal(t, tt.want.branch, d.Branch)
		})
	}
}
