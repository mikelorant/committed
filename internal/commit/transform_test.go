package commit_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestMessageToEmoji(t *testing.T) {
	t.Parallel()

	type want struct {
		valid bool
		name  string
	}

	tests := []struct {
		name    string
		message string
		want    want
	}{
		{
			name:    "emoji_character_summary",
			message: "🎨 summary",
			want: want{
				valid: true,
				name:  "art",
			},
		},
		{
			name:    "emoji_shortcode_summary",
			message: ":art: summary",
			want: want{
				valid: true,
				name:  "art",
			},
		},
		{
			name:    "emoji_only",
			message: "🎨",
			want: want{
				valid: true,
				name:  "art",
			},
		},
		{
			name:    "emoji_multiline_message",
			message: "🎨 summary\n\nbody\n",
			want: want{
				valid: true,
				name:  "art",
			},
		},
		{
			name:    "unknown_emoji",
			message: "😀 summary",
		},
		{
			name:    "summary",
			message: "summary",
		},
		{
			name:    "empty",
			message: "",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := commit.MessageToEmoji(emoji.New(), tt.message)
			if !tt.want.valid {
				assert.False(t, tt.want.valid)
				assert.Empty(t, e.Emoji.Name)
				return
			}
			assert.True(t, tt.want.valid)
			assert.Equal(t, tt.want.name, e.Emoji.Name)
		})
	}
}

func TestMessageToSummary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		message string
		summary string
	}{
		{
			name:    "summary",
			message: "summary",
			summary: "summary",
		},
		{
			name:    "emoji_summary",
			message: "😀 summary",
			summary: "summary",
		},
		{
			name:    "emoji_only",
			message: "😀",
		},
		{
			name:    "short_summary",
			message: "a",
			summary: "a",
		},
		{
			name:    "long_summary",
			message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
		},
		{
			name:    "multiline",
			message: "summary\n\nbody",
			summary: "summary",
		},
		{
			name:    "multiline_newline",
			message: "summary\n",
			summary: "summary",
		},
		{
			name:    "multiline_no_newline",
			message: "summary\nbody\n",
		},
		{
			name: "empty",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := commit.MessageToSummary(tt.message)
			if tt.summary == "" {
				assert.Empty(t, s)
				return
			}
			assert.NotEmpty(t, s)
			assert.Equal(t, tt.summary, s)
		})
	}
}

func TestMessageToBody(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		message string
		body    string
	}{
		{
			name:    "summary_body",
			message: "summary\n\nbody",
			body:    "body",
		},
		{
			name:    "summary",
			message: "summary",
		},
		{
			name:    "multine_newline",
			message: "summary\n",
			body:    "",
		},
		{
			name:    "multine_no_newline",
			message: "summary\nbody",
			body:    "summary\nbody",
		},
		{
			name:    "empty",
			message: "",
		},
		{
			name:    "long_summary",
			message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
			body:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := commit.MessageToBody(tt.message)
			if tt.body == "" {
				assert.Empty(t, b)
				return
			}
			assert.NotEmpty(t, b)
			assert.Equal(t, tt.body, b)
		})
	}
}

func TestEmojiSummaryToSubject(t *testing.T) {
	t.Parallel()

	type args struct {
		emoji   string
		summary string
	}

	tests := []struct {
		name    string
		args    args
		subject string
	}{
		{
			name: "emoji_summary",
			args: args{
				emoji:   ":art:",
				summary: "summary",
			},
			subject: ":art: summary",
		},
		{
			name: "summary",
			args: args{
				summary: "summary",
			},
			subject: "summary",
		},
		{
			name: "empty",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := commit.EmojiSummaryToSubject(tt.args.emoji, tt.args.summary)
			if tt.subject == "" {
				assert.Empty(t, s)
				return
			}
			assert.NotEmpty(t, s)
			assert.Equal(t, tt.subject, s)
		})
	}
}

func TestUserToAuthor(t *testing.T) {
	t.Parallel()

	type args struct {
		name  string
		email string
	}

	tests := []struct {
		name   string
		args   args
		author string
	}{
		{
			name: "name_email",
			args: args{
				name:  "John Doe",
				email: "john.doe@example.com",
			},
			author: "John Doe <john.doe@example.com>",
		},
		{
			name: "name",
			args: args{
				name: "John Doe",
			},
		},
		{
			name: "email",
			args: args{
				email: "john.doe@example.com",
			},
		},
		{
			name: "empty",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := repository.User{
				Name:  tt.args.name,
				Email: tt.args.email,
			}

			a := commit.UserToAuthor(u)
			if tt.author == "" {
				assert.Empty(t, a)
				return
			}
			assert.NotEmpty(t, a)
			assert.Equal(t, tt.author, a)
		})
	}
}

func TestSortUsersByDefault(t *testing.T) {
	t.Parallel()

	type args struct {
		repositoryUsers []repository.User
		configUsers     []repository.User
	}

	tests := []struct {
		name  string
		args  args
		users []repository.User
	}{
		{
			name: "empty",
		},
		{
			name: "nil",
			args: args{
				repositoryUsers: nil,
				configUsers:     nil,
			},
		},
		{
			name: "one_repository_user",
			args: args{
				repositoryUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com"},
			},
		},
		{
			name: "two_repository_user",
			args: args{
				repositoryUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
					{Name: "John Doe", Email: "jdoe@example.org"},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com"},
				{Name: "John Doe", Email: "jdoe@example.org"},
			},
		},
		{
			name: "one_config_users",
			args: args{
				configUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com"},
			},
		},
		{
			name: "two_config_users",
			args: args{
				configUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
					{Name: "John Doe", Email: "jdoe@example.org"},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com"},
				{Name: "John Doe", Email: "jdoe@example.org"},
			},
		},
		{
			name: "repository_config_users",
			args: args{
				repositoryUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
				},
				configUsers: []repository.User{
					{Name: "John Doe", Email: "jdoe@example.org"},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com"},
				{Name: "John Doe", Email: "jdoe@example.org"},
			},
		},
		{
			name: "repository_config_users_multiple",
			args: args{
				repositoryUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
					{Name: "John Doe", Email: "jdoe@example.org"},
				},
				configUsers: []repository.User{
					{Name: "John Doe", Email: "jd@example.net"},
					{Name: "John Doe", Email: "j@example.id"},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com"},
				{Name: "John Doe", Email: "jdoe@example.org"},
				{Name: "John Doe", Email: "jd@example.net"},
				{Name: "John Doe", Email: "j@example.id"},
			},
		},
		{
			name: "one_config_user_default",
			args: args{
				configUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com", Default: true},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com", Default: true},
			},
		},
		{
			name: "two_config_users_default",
			args: args{
				configUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
					{Name: "John Doe", Email: "jdoe@example.org", Default: true},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "jdoe@example.org", Default: true},
				{Name: "John Doe", Email: "john.doe@example.com"},
			},
		},
		{
			name: "two_config_users_default_both",
			args: args{
				configUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com", Default: true},
					{Name: "John Doe", Email: "jdoe@example.org", Default: true},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "john.doe@example.com", Default: true},
				{Name: "John Doe", Email: "jdoe@example.org", Default: true},
			},
		},
		{
			name: "repository_config_users_default",
			args: args{
				repositoryUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
				},
				configUsers: []repository.User{
					{Name: "John Doe", Email: "jdoe@example.org", Default: true},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "jdoe@example.org", Default: true},
				{Name: "John Doe", Email: "john.doe@example.com"},
			},
		},
		{
			name: "repository_config_users_default_multiple",
			args: args{
				repositoryUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
					{Name: "John Doe", Email: "jdoe@example.org"},
				},
				configUsers: []repository.User{
					{Name: "John Doe", Email: "jd@example.net"},
					{Name: "John Doe", Email: "j@example.id", Default: true},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "j@example.id", Default: true},
				{Name: "John Doe", Email: "john.doe@example.com"},
				{Name: "John Doe", Email: "jdoe@example.org"},
				{Name: "John Doe", Email: "jd@example.net"},
			},
		},
		{
			name: "repository_config_users_default_all",
			args: args{
				repositoryUsers: []repository.User{
					{Name: "John Doe", Email: "john.doe@example.com"},
					{Name: "John Doe", Email: "jdoe@example.org"},
				},
				configUsers: []repository.User{
					{Name: "John Doe", Email: "jd@example.net", Default: true},
					{Name: "John Doe", Email: "j@example.id", Default: true},
				},
			},
			users: []repository.User{
				{Name: "John Doe", Email: "jd@example.net", Default: true},
				{Name: "John Doe", Email: "j@example.id", Default: true},
				{Name: "John Doe", Email: "john.doe@example.com"},
				{Name: "John Doe", Email: "jdoe@example.org"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			us := concatSlice(tt.args.repositoryUsers, tt.args.configUsers)
			users := commit.SortUsersByDefault(us...)

			assert.Equal(t, tt.users, users)
		})
	}
}

func concatSlice[T any](first []T, second []T) []T {
	n := len(first)
	return append(first[:n:n], second...)
}
