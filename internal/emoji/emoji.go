package emoji

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/forPelevin/gomoji"
	"github.com/goccy/go-yaml"
)

type Set struct {
	Name   string
	Emojis []Emoji

	profile   Profile
	rawEmojis string
}

type Emoji struct {
	Name        string `json:"name"`
	Character   string `json:"emoji"`
	Description string `json:"description"`
	Characters  int    `json:"characters"`
	Codepoint   string `json:"codepoint"`
	Hex         string `json:"hex"`
	Shortcode   string `json:"shortcode"`
	Variant     bool   `json:"variant"`
	ZWJ         bool   `json:"zwj"`
}

type NullEmoji struct {
	Valid bool
	Emoji Emoji
}

type Profile int

//go:embed gitmoji.yaml
var gitmoji string

//go:embed devmoji.yaml
var devmoji string

//go:embed emojilog.yaml
var emojiLog string

const (
	gitmojiName  = "gitmoji"
	devmojiName  = "devmoji"
	emojiLogName = "emojilog"
)

const (
	DefaultProfile Profile = iota
	GitmojiProfile
	DevmojiProfile
	EmojiLogProfile
	ProfileCount
)

func New(opts ...func(*Set)) *Set {
	var es Set

	for _, o := range opts {
		if o != nil {
			o(&es)
		}
	}

	es.load(es.profile)

	return &es
}

func (es *Set) Find(str string) NullEmoji {
	switch {
	case HasCharacter(str):
		return es.FindByCharacter(str)
	case HasShortcode(str):
		return es.FindByShortcode(str)
	default:
		return NullEmoji{}
	}
}

func (es *Set) FindByCharacter(str string) NullEmoji {
	for _, e := range es.Emojis {
		if e.Character == str {
			return NullEmoji{
				Valid: true,
				Emoji: e,
			}
		}
	}

	return NullEmoji{}
}

func ToString(p Profile) string {
	return []string{
		"default",
		"gitmoji",
		"devmoj",
		"emoji-log",
	}[int(p)]
}

func ToURL(p Profile) string {
	return []string{
		"https://gitmoji.dev/",
		"https://gitmoji.dev/",
		"https://github.com/folke/devmoji",
		"https://github.com/ahmadawais/emoji-log",
	}[int(p)]
}

func (es *Set) ListProfiles() []Profile {
	ps := make([]Profile, ProfileCount)
	for i := 0; i < int(ProfileCount); i++ {
		ps[i] = Profile(i)
	}

	return ps
}

func (es *Set) FindByShortcode(str string) NullEmoji {
	for _, e := range es.Emojis {
		if e.Shortcode == str {
			return NullEmoji{
				Valid: true,
				Emoji: e,
			}
		}
	}

	return NullEmoji{}
}

func (es *Set) load(p Profile) {
	switch p {
	case GitmojiProfile:
		es.Name = gitmojiName
		es.rawEmojis = gitmoji
	case DevmojiProfile:
		es.Name = devmojiName
		es.rawEmojis = devmoji
	case EmojiLogProfile:
		es.Name = emojiLogName
		es.rawEmojis = emojiLog
	default:
		es.Name = gitmojiName
		es.rawEmojis = gitmoji
	}

	if es.rawEmojis == "" {
		return
	}

	r := strings.NewReader(es.rawEmojis)

	if err := yaml.NewDecoder(r).Decode(&es.Emojis); err != nil {
		panic(fmt.Errorf("unable to decode emojis: %w", err))
	}
}

func WithEmojiSet(p Profile) func(*Set) {
	return func(e *Set) {
		e.profile = p
	}
}

func Has(str string) bool {
	return HasCharacter(str) || HasShortcode(str)
}

func HasCharacter(str string) bool {
	return gomoji.ContainsEmoji(str)
}

func HasShortcode(str string) bool {
	if len(str) <= 2 {
		return false
	}

	if strings.Count(str, ":") > 2 {
		return false
	}

	if string(str[0]) == ":" && string(str[len(str)-1]) == ":" {
		return true
	}

	return false
}
