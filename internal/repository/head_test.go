package repository_test

import (
	"errors"
	"testing"
	"time"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

type MockRepositoryHead struct {
	head            repository.Head
	headErr         error
	commitObjectErr error
}

func (m MockRepositoryHead) Head() (*plumbing.Reference, error) {
	hr := plumbing.NewHashReference(plumbing.HEAD, plumbing.NewHash(m.head.Hash))

	return hr, m.headErr
}

func (m MockRepositoryHead) CommitObject(h plumbing.Hash) (*object.Commit, error) {
	return &object.Commit{
		Hash: h,
		Author: object.Signature{
			Name:  m.head.Author.Name,
			Email: m.head.Author.Email,
			When:  m.head.When,
		},
		Message: m.head.Message,
	}, m.commitObjectErr
}

var errMockHead = errors.New("error")

var mockTime = time.Date(2022, time.January, 1, 1, 0, 0, 0, time.UTC)

func TestHead(t *testing.T) {
	type args struct {
		head            repository.Head
		headErr         error
		commitObjectErr error
	}

	type want struct {
		head repository.Head
		err  error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "main",
			args: args{
				head: repository.Head{
					Hash: "1234567890abcdef1234567890abcdef12345678",
					Author: repository.User{
						Name:  "John Doe",
						Email: "john.doe@example.com",
					},
					When:    mockTime,
					Message: "message",
				},
			},
			want: want{
				head: repository.Head{
					Hash: "1234567890abcdef1234567890abcdef12345678",
					Author: repository.User{
						Name:  "John Doe",
						Email: "john.doe@example.com",
					},
					When:    mockTime,
					Message: "message",
				},
			},
		},
		{
			name: "empty_hash",
			args: args{
				head: repository.Head{
					Hash: "",
				},
			},
			want: want{
				head: repository.Head{
					Hash: "0000000000000000000000000000000000000000",
				},
			},
		},
		{
			name: "head_error",
			args: args{
				headErr: errMockHead,
			},
			want: want{
				err: errMockHead,
			},
		},
		{
			name: "head_reference_not_found",
			args: args{
				headErr: plumbing.ErrReferenceNotFound,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "commit_object_error",
			args: args{
				commitObjectErr: errMockHead,
			},
			want: want{
				err: errMockHead,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r repository.Repository

			r.Header = MockRepositoryHead{
				head:            tt.args.head,
				headErr:         tt.args.headErr,
				commitObjectErr: tt.args.commitObjectErr,
			}

			h, err := r.Head()
			if tt.want.err != nil {
				assert.ErrorContains(t, err, tt.want.err.Error())
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tt.want.head, h)
		})
	}
}
