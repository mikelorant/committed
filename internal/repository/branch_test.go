package repository_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/mikelorant/committed/internal/repository"

	"github.com/go-git/go-billy/v5/memfs"
	fixtures "github.com/go-git/go-git-fixtures/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/stretchr/testify/assert"
)

type MockRepositoryBranch struct {
	repo       *git.Repository
	local      string
	remote     string
	localRefs  []string
	remoteRefs []string
	tagRefs    []string
	idx        int

	configErr error
	headErr   error
	refsErr   error
}

var errMockBranch = errors.New("error")

var mockHash = plumbing.NewHash("6ecf0ef2c2dffb796033e5a02219af86ec6584e5")

func (m *MockRepositoryBranch) Config() (*config.Config, error) {
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

func (m *MockRepositoryBranch) Head() (*plumbing.Reference, error) {
	if m.local == "" {
		var ref plumbing.Reference

		return &ref, nil
	}

	hr := plumbing.NewHashReference(plumbing.NewBranchReferenceName(m.local), mockHash)

	return hr, m.headErr
}

func (m *MockRepositoryBranch) TagObject(hash plumbing.Hash) (*object.Tag, error) {
	var ref *plumbing.Reference
	var err error

	ref, err = m.repo.CreateTag(m.tagRefs[m.idx], hash, &git.CreateTagOptions{
		Tagger: &object.Signature{
			Name:  "John Doe",
			Email: "john.doe@example.com",
			When:  time.Now(),
		},
		Message: m.tagRefs[m.idx],
	})
	if err != nil {
		ref, err = m.repo.Tag(m.tagRefs[m.idx])
		if err != nil {
			return nil, fmt.Errorf("unable to create tag or get existing tag: %w", err)
		}
	}
	m.idx++

	tag, err := m.repo.TagObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	return tag, nil
}

//nolint:ireturn
func (m MockRepositoryBranch) References() (storer.ReferenceIter, error) {
	if m.refsErr != nil {
		return nil, m.refsErr
	}

	var rs []*plumbing.Reference

	for _, r := range m.localRefs {
		rs = append(rs, plumbing.NewHashReference(plumbing.NewBranchReferenceName(r), mockHash))
	}

	for _, r := range m.remoteRefs {
		rs = append(rs, plumbing.NewHashReference(plumbing.NewRemoteReferenceName(r, m.local), mockHash))
	}

	for _, r := range m.tagRefs {
		rs = append(rs, plumbing.NewHashReference(plumbing.NewTagReferenceName(r), mockHash))
	}

	return storer.NewReferenceSliceIter(rs), nil
}

func TestBranch(t *testing.T) {
	t.Parallel()

	type args struct {
		local      string
		remote     string
		localRefs  []string
		remoteRefs []string
		tagRefs    []string
		configErr  error
		headErr    error
		refsErr    error
	}

	type want struct {
		local  string
		remote string
		refs   repository.Refs
		err    string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		// local
		{
			name: "local",
			args: args{
				local: "master",
			},
			want: want{
				local: "master",
			},
		},
		{
			name: "no_local",
			want: want{
				err: "unable to get local branch: local branch not found",
			},
		},

		// remote
		{
			name: "remote",
			args: args{
				local:  "master",
				remote: "origin",
			},
			want: want{
				local:  "master",
				remote: "origin/master",
			},
		},

		// refs local
		{
			name: "local_refs",
			args: args{
				local:     "master",
				localRefs: []string{"test"},
			},
			want: want{
				local: "master",
				refs: repository.Refs{
					Locals: []string{"test"},
				},
			},
		},
		{
			name: "local_refs_multiple",
			args: args{
				local:     "master",
				localRefs: []string{"test1", "test2"},
			},
			want: want{
				local: "master",
				refs: repository.Refs{
					Locals: []string{"test1", "test2"},
				},
			},
		},
		{
			name: "local_refs_duplicate",
			args: args{
				local:     "master",
				localRefs: []string{"master"},
			},
			want: want{
				local: "master",
			},
		},
		{
			name: "local_refs_multiple_duplicate",
			args: args{
				local:     "master",
				localRefs: []string{"test", "master"},
			},
			want: want{
				local: "master",
				refs: repository.Refs{
					Locals: []string{"test"},
				},
			},
		},

		// refs remote
		{
			name: "remote_refs",
			args: args{
				local:      "master",
				remote:     "origin",
				remoteRefs: []string{"test"},
			},
			want: want{
				local:  "master",
				remote: "origin/master",
				refs: repository.Refs{
					Remotes: []string{"test/master"},
				},
			},
		},
		{
			name: "remote_refs_multiple",
			args: args{
				local:      "master",
				remote:     "origin",
				remoteRefs: []string{"test1", "test2"},
			},
			want: want{
				local:  "master",
				remote: "origin/master",
				refs: repository.Refs{
					Remotes: []string{"test1/master", "test2/master"},
				},
			},
		},
		{
			name: "remote_refs_duplicate",
			args: args{
				local:      "master",
				remote:     "origin",
				remoteRefs: []string{"test", "origin"},
			},
			want: want{
				local:  "master",
				remote: "origin/master",
				refs: repository.Refs{
					Remotes: []string{"test/master"},
				},
			},
		},

		// refs tags
		{
			name: "tag_refs",
			args: args{
				local:   "master",
				tagRefs: []string{"test"},
			},
			want: want{
				local: "master",
				refs: repository.Refs{
					Tags: []string{"test"},
				},
			},
		},
		{
			name: "tag_refs_multiple",
			args: args{
				local:   "master",
				tagRefs: []string{"test1", "test2"},
			},
			want: want{
				local: "master",
				refs: repository.Refs{
					Tags: []string{"test1", "test2"},
				},
			},
		},
		{
			name: "tag_refs_duplicate",
			args: args{
				local:   "master",
				tagRefs: []string{"test1", "test2", "test1"},
			},
			want: want{
				local: "master",
				refs: repository.Refs{
					Tags: []string{"test1", "test2", "test1"},
				},
			},
		},

		// error
		{
			name: "config_error",
			args: args{
				configErr: errMockBranch,
			},
			want: want{
				err: "unable to get repository config: error",
			},
		},
		{
			name: "head_error",
			args: args{
				local:   "master",
				headErr: errMockBranch,
			},
			want: want{
				err: "unable to get head reference: error",
			},
		},
		{
			name: "head_reference_not_found",
			args: args{
				local:   "master",
				headErr: plumbing.ErrReferenceNotFound,
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
				err:   "unable to get head references: unable to get references: error",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dotgit := fixtures.Basic().One().DotGit()
			st := filesystem.NewStorage(dotgit, cache.NewObjectLRUDefault())
			wt := memfs.New()
			repo, _ := git.Open(st, wt)

			var r repository.Repository

			r.Brancher = &MockRepositoryBranch{
				repo:       repo,
				local:      tt.args.local,
				remote:     tt.args.remote,
				localRefs:  tt.args.localRefs,
				remoteRefs: tt.args.remoteRefs,
				tagRefs:    tt.args.tagRefs,
				configErr:  tt.args.configErr,
				headErr:    tt.args.headErr,
				refsErr:    tt.args.refsErr,
			}

			branch, err := r.Branch()
			if tt.want.err != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.want.err)
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tt.want.local, branch.Local)
			assert.Equal(t, tt.want.remote, branch.Remote)
			assert.Equal(t, tt.want.refs, branch.Refs)
		})
	}
}
