package repository_test

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

type MockRepositoryRemote struct {
	remotes []*git.Remote
	err     error
}

func (m MockRepositoryRemote) Remotes() ([]*git.Remote, error) {
	return m.remotes, m.err
}

func TestRemote(t *testing.T) {
	type args struct {
		remotes []string
		err     error
	}

	type want struct {
		remotes []string
		err     error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "single",
			args: args{
				remotes: []string{"origin"},
			},
			want: want{
				remotes: []string{"origin"},
			},
		},
		{
			name: "multiple",
			args: args{
				remotes: []string{"origin", "upstream"},
			},
			want: want{
				remotes: []string{"origin", "upstream"},
			},
		},
		{
			name: "empty",
			args: args{
				remotes: []string{},
			},
			want: want{
				remotes: nil,
			},
		},
		{
			name: "error",
			args: args{
				err: errMockUser,
			},
			want: want{
				remotes: nil,
				err:     errMockUser,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r repository.Repository

			m := memory.NewStorage()

			rms := make([]*git.Remote, len(tt.args.remotes))
			for i, v := range tt.args.remotes {
				rc := config.RemoteConfig{
					Name: v,
				}
				rm := git.NewRemote(m, &rc)
				rms[i] = rm
			}

			r.Remoter = MockRepositoryRemote{
				remotes: rms,
				err:     tt.args.err,
			}

			rs, err := r.Remotes()
			if tt.want.err != nil {
				assert.ErrorContains(t, err, tt.want.err.Error())
				return
			}
			assert.Nil(t, err)

			assert.Equal(t, tt.want.remotes, rs)
		})
	}
}
