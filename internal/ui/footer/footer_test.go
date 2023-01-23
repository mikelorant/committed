package footer_test

import (
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/ui/footer"
	"github.com/mikelorant/committed/internal/ui/uitest"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	type args struct {
		author repository.User
		model  func(m footer.Model) footer.Model
	}

	type want struct {
		model func(m footer.Model)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			args: args{
				author: repository.User{
					Name:  "John Doe",
					Email: "john.doe@example.com",
				},
			},
			want: want{
				model: func(m footer.Model) {
					u := repository.User{
						Name:  "John Doe",
						Email: "john.doe@example.com",
					}

					assert.Equal(t, u, m.Author)
					assert.Equal(t, false, m.Signoff)
					assert.Equal(t, "", m.Value())
				},
			},
		},
		{
			name: "signoff",
			args: args{
				author: repository.User{
					Name:  "John Doe",
					Email: "john.doe@example.com",
				},
				model: func(m footer.Model) footer.Model {
					m.ToggleSignoff()
					m, _ = footer.ToModel(m.Update(nil))
					return m
				},
			},
			want: want{
				model: func(m footer.Model) {
					u := repository.User{
						Name:  "John Doe",
						Email: "john.doe@example.com",
					}

					assert.Equal(t, u, m.Author)
					assert.Equal(t, true, m.Signoff)
					assert.Equal(t, "Signed-off-by: John Doe <john.doe@example.com>", m.Value())
				},
			},
		},
		{
			name: "empty",
			want: want{
				model: func(m footer.Model) {
					u := repository.User{}

					assert.Equal(t, u, m.Author)
					assert.Equal(t, false, m.Signoff)
					assert.Equal(t, "", m.Value())
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c commit.Config
			if tt.args.author.Name != "" && tt.args.author.Email != "" {
				c.Repository.Users = []repository.User{tt.args.author}
			}

			m := footer.New(&c)

			if tt.args.model != nil {
				m = tt.args.model(m)
			}

			if tt.want.model != nil {
				tt.want.model(m)
			}

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}
