package snapshot_test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/snapshot"
	"github.com/stretchr/testify/assert"
)

type readWriteCloser struct {
	bytes.Buffer
}

func (wc *readWriteCloser) Close() error {
	return nil
}

type errorReadWriteCloser struct {
	readWriteCloser
}

func (t *errorReadWriteCloser) Write(p []byte) (n int, err error) {
	return 0, errMock
}

var errMock = errors.New("error")

func TestLoad(t *testing.T) {
	t.Parallel()

	type args struct {
		reader io.Reader
	}
	type want struct {
		snapshot snapshot.Snapshot
		err      string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "data",
			args: args{
				reader: strings.NewReader(heredoc.Doc(`
					emoji: ":art:"
					summary: summary
					body: body
					footer: footer
					author:
					  name: John Doe
					  email: john.doe@example.com
					amend: true
				`)),
			},
			want: want{
				snapshot: snapshot.Snapshot{
					Emoji:   ":art:",
					Summary: "summary",
					Body:    "body",
					Footer:  "footer",
					Author: repository.User{
						Name:  "John Doe",
						Email: "john.doe@example.com",
					},
					Amend: true,
				},
			},
		},
		{
			name: "empty",
			args: args{
				reader: strings.NewReader(""),
			},
		},
		{
			name: "error_reader",
			want: want{
				err: "empty reader",
			},
		},
		{
			name: "error_eof",
			args: args{
				reader: io.LimitReader(strings.NewReader("summary: summary"), 0),
			},
		},
		{
			name: "error_decode",
			args: args{
				reader: io.LimitReader(strings.NewReader("summary: summary"), 1),
			},
			want: want{
				err: "unable to decode snapshot",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var s snapshot.Snapshot

			snap, err := s.Load(tt.args.reader)

			if tt.want.err != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.want.err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want.snapshot, snap)
		})
	}
}

func TestSave(t *testing.T) {
	t.Parallel()

	type args struct {
		writer   io.ReadWriteCloser
		snapshot snapshot.Snapshot
	}

	type want struct {
		data string
		err  string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "data",
			args: args{
				writer: new(readWriteCloser),
				snapshot: snapshot.Snapshot{
					Emoji:   ":art:",
					Summary: "summary",
					Body:    "body",
					Footer:  "footer",
					Author: repository.User{
						Name:  "John Doe",
						Email: "john.doe@example.com",
					},
					Amend: true,
				},
			},
			want: want{
				data: heredoc.Doc(`
					emoji: ':art:'
					summary: summary
					body: body
					footer: footer
					author:
					    name: John Doe
					    email: john.doe@example.com
					amend: true
				`),
			},
		},
		{
			name: "empty",
			args: args{
				writer: new(readWriteCloser),
			},
			want: want{
				data: "{}\n",
			},
		},
		{
			name: "error_encode",
			args: args{
				writer: new(errorReadWriteCloser),
			},
			want: want{
				err: "unable to encode snapshot",
			},
		},
		{
			name: "error_writer",
			want: want{
				err: "empty writer",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var s snapshot.Snapshot

			err := s.Save(tt.args.writer, tt.args.snapshot)
			if tt.want.err != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.want.err)
				return
			}
			assert.NoError(t, err)

			got, _ := io.ReadAll(tt.args.writer)
			assert.Equal(t, tt.want.data, string(got))
		})
	}
}
