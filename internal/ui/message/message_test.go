package message_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/message"
	"github.com/mikelorant/committed/internal/ui/uitest"

	"github.com/hexops/autogold/v2"
)

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		emoji   string
		summary string
		body    string
		footer  string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "default",
		},
		{
			name: "summary",
			args: args{
				summary: "summary",
			},
		},
		{
			name: "emoji_summary",
			args: args{
				emoji:   ":art:",
				summary: "summary",
			},
		},
		{
			name: "summary_body",
			args: args{
				summary: "summary",
				body:    "body",
			},
		},
		{
			name: "summary_body_multiline",
			args: args{
				summary: "summary",
				body:    "line 1\nline 2\nline 3",
			},
		},
		{
			name: "summary_footer",
			args: args{
				summary: "summary",
				footer:  "footer",
			},
		},
		{
			name: "all",
			args: args{
				emoji:   ":art:",
				summary: "summary",
				body:    "body",
				footer:  "footer",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := message.State{
				Theme:   theme.New(config.ColourAdaptive),
				Emoji:   tt.args.emoji,
				Summary: tt.args.summary,
				Body:    tt.args.body,
				Footer:  tt.args.footer,
			}

			m := message.New(c)

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(tt.name))
		})
	}
}
