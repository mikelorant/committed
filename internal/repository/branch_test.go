package repository_test

import (
	"errors"
	"testing"

	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

type MockRepositoryBranch struct {
	local     string
	remote    string
	refs      []string
	configErr error
	headErr   error
	refsErr   error
}

var errMockBranch = errors.New("error")

var mockHash = "1234567890abcdef1234567890abcdef12345678"

func (m MockRepositoryBranch) Config() (*config.Config, error) {
	var cfg config.Config

	if m.configErr != nil {
		return &cfg, m.configErr
	}

	if m.local == "" {
		return &cfg, nil
	}

	bs := make(map[string]*config.Branch)
	bs[m.local] = &config.Branch{
		Name:   m.local,
		Remote: m.remote,
	}

	cfg.Branches = bs

	return &cfg, nil
}

func (m MockRepositoryBranch) Head() (*plumbing.Reference, error) {
	if m.local == "" {
		var ref plumbing.Reference

		return &ref, nil
	}

	hr := plumbing.NewHashReference(plumbing.NewBranchReferenceName(m.local), plumbing.NewHash(mockHash))

	return hr, m.headErr
}

//nolint:ireturn
func (m MockRepositoryBranch) References() (storer.ReferenceIter, error) {
	if m.refsErr != nil {
		return nil, m.refsErr
	}

	rs := make([]*plumbing.Reference, len(m.refs))
	for i, r := range m.refs {
		hr := plumbing.NewHashReference(plumbing.NewBranchReferenceName(r), plumbing.NewHash(mockHash))
		rs[i] = hr
	}

	return storer.NewReferenceSliceIter(rs), nil
}

func TestBranch(t *testing.T) {
	type args struct {
		local     string
		remote    string
		refs      []string
		configErr error
		headErr   error
		refsErr   error
	}

	type want struct {
		local  string
		remote string
		refs   []string
		err    error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "master",
			args: args{
				local:  "master",
				remote: "origin",
				refs:   []string{"v1.0.0"},
			},
			want: want{
				local:  "master",
				remote: "origin/master",
				refs:   []string{"v1.0.0"},
			},
		},
		{
			name: "local_in_refs",
			args: args{
				local:  "test",
				remote: "origin",
				refs:   []string{"v1.0.0", "test"},
			},
			want: want{
				local:  "test",
				remote: "origin/test",
				refs:   []string{"v1.0.0"},
			},
		},
		{
			name: "config_error",
			args: args{
				configErr: errMockBranch,
			},
			want: want{
				err: errMockBranch,
			},
		},
		{
			name: "head_error",
			args: args{
				local:   "master",
				headErr: errMockBranch,
			},
			want: want{
				err: errMockBranch,
			},
		},
		{
			name: "head_reference_not_found",
			args: args{
				local:   "master",
				headErr: plumbing.ErrReferenceNotFound,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "no_local",
			want: want{
				err: repository.ErrLocalBranchNotFound,
			},
		},
		{
			name: "no_remote",
			args: args{
				local: "master",
			},
			want: want{
				local: "master",
			},
		},
		{
			name: "refs_error",
			args: args{
				local:   "master",
				refsErr: errMockBranch,
			},
			want: want{
				local: "master",
				err:   errMockBranch,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r repository.Repository

			r.Brancher = MockRepositoryBranch{
				local:     tt.args.local,
				remote:    tt.args.remote,
				refs:      tt.args.refs,
				configErr: tt.args.configErr,
				headErr:   tt.args.headErr,
				refsErr:   tt.args.refsErr,
			}

			branch, err := r.Branch()
			if tt.want.err != nil {
				assert.ErrorContains(t, err, tt.want.err.Error())
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tt.want.local, branch.Local)
			assert.Equal(t, tt.want.remote, branch.Remote)
			assert.Equal(t, tt.want.refs, branch.Refs)
		})
	}
}
