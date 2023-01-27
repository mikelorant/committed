# Committed

[![Release](https://img.shields.io/github/v/release/mikelorant/committed)](https://github.com/mikelorant/committed/releases) [![Build](https://github.com/mikelorant/committed/actions/workflows/release.yaml/badge.svg)](https://github.com/mikelorant/committed/actions/workflows/release.yaml) [![License](https://img.shields.io/github/license/mikelorant/committed)](https://spdx.org/licenses/MIT.html) [![codecov](https://codecov.io/gh/mikelorant/committed/branch/main/graph/badge.svg?token=TW2XDZAJKT)](https://codecov.io/gh/mikelorant/committed)

Committed is a WYSIWYG Git commit editor that helps improve the quality of your
commits by showing you the layout in the same format as `git log`.

![demo](docs/demo.gif)

[![Highlights](https://img.shields.io/badge/-Highlights-ff0000)](#-highlights-) [![Purpose](https://img.shields.io/badge/-Purpose-ff8000)](#-purpose-) [![Limitations](https://img.shields.io/badge/-Limitations-ffff00)](#-limitations-) [![Installation](https://img.shields.io/badge/-Installation-80ff00)](#-installation-) [![Usage](https://img.shields.io/badge/-Usage-00aa00)](#-usage-) [![Configuration](https://img.shields.io/badge/-Configuration-00bb80)](#-configuration-) [![Best Practises](https://img.shields.io/badge/-Best_Practises-00ffff)](#-best-practises-) [![Shortcuts](https://img.shields.io/badge/-Shortcuts-0080ff)](#-shortcuts-) [![Views](https://img.shields.io/badge/-Views-0000ff)](#-views-) [![Authors](https://img.shields.io/badge/-Authors-8000ff)](#-authors-) [![License](https://img.shields.io/badge/-License-ff00ff)](#-license-) [![Thanks](https://img.shields.io/badge/-Thanks-ff0080)](#-thanks-)

#

## 💡 Highlights [⭡](#committed)

- Built-in multiline editor
- Emoji selector
- Author profile switching
- Inline text interface
- Subject line counter
- Appends sign-off
- Formats body width to 72 characters
- Best practise recommendations

## 💬 Purpose [⭡](#committed)

The benefits of high quality commits are well documented however the tooling to
follow these practises has been lacking. In most cases you are either providing
a single line commit message or forced into a full screen editor which has no
knowledge of recommended Git practises.

It is common to accidentally commit to the wrong branch or use the incorrect
author name. Improving clarity of commits with emojis or detailed messages is
often frustrating.

For many, knowing what makes a good commit is not even thought about.

Committed attempts to solve these problems by first educating on best practises.
It then helps guide and enforce these de factor standards while showing how the
commit will end up being displayed to other users.

The interface does not take over the sceen or force switching to another
application. All actions are done with the keyboard which is consistent with the
Git command which is often used before the actual commit. Having an editor which
allows for more advanced cursor movement and editing assists with revising
rather than accepting what has already been written.

These capabilities all contribute to helping create a commit message that is
useful.

## ⚠ Limitations [⭡](#committed)

Terminals render emojis differently and this makes alignment of borders
complicated and difficult. It is an ongoing process to improve the compatibility
with terminals. The following list are the terminals that have been tested.
Other terminals may display correctly and feedback would be appreciated to help
update the list.

| Terminal       | Status                  |
| :------------- | :---------------------- |
| macOS Terminal | Compatible              |
| iTerm2         | Compatible              |
| VS Code        | Compatible              |
| Hyper          | Compatible              |
| Alacritty      | Compatible              |
| WezTerm        | Compatible              |
| Kitty          | Border alignment issues |

## 💾 Installation [⭡](#committed)

Install Committed with Homebrew.

```bash
brew install mikelorant/taps/committed
```

## 🎛 Usage [⭡](#committed)

```text
Committed is a WYSIWYG Git commit editor

Usage:
  committed [flags]
  committed [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  version      Print the version information

Flags:
      --config string   Config file location (default "$HOME/.config/committed/config.yaml")
      --dry-run         Simulate applying a commit (default true)
  -a, --amend           Replace the tip of the current branch by creating a new commit
  -h, --help            help for committed
  -v, --version         version for committed

Use "committed [command] --help" for more information about a command.
```

To create and apply a commit run `committed` without any arguments. Shell or Git
aliases can be used to tailor this to your preferred workflow.

To amend an existing commit use `committed --amend`. There are certain
limitations when amending commits and it is recommended only for use with
commits created with Committed. The limitations are:

- Emoji character or shortcode must be in the existing data set.
- Trailers will be imported into the body.
- Summary will be truncated if more than 72 characters.
- Lines will not reflow when editing the body.

## ⚙ Configuration [⭡](#committed)

No configuration is necessary however there are some values that can be changed
based on preference.

Committed defaults to using a config file located at `$HOME/.config/committed/config.yaml`.

```yaml
view:
  # Starting component focus.
  # Values: author, emoji, summary
  # Default: emoji
  focus: emoji

  # Emoji selector placement in relation to subject.
  # Values: above, below
  # Default: below
  emojiSelector: below

  # Emoji set to use.
  # Values: gitmoji, devmoji, emojilog
  # Default: gitmoji
  emojiSet: gitmoji

  # Theme to display. Dark and light backgrounds have different themes.
  # Dark values:
  #   builtin_dark, dracula, gruvbox_dark, nord, retrowave,
  #   solarized_dark_higher_contrast, tokyo_night
  # Dark default: builtin_dark
  # Light values:
  #   builtin_light, gruvbox_light, builtin_solarized_light,
  #   builtin_tango_light, tokyo_night_light
  # Light default: builtin_light
  theme: builtin_dark

  # Colour profile for displaying themes.
  # Values: adaptive, dark, light
  # Default: adaptive
  colour: adaptive

commit:
  # Emoji format in commit.
  # Values: shortcode, character
  # Default: shortcode
  emojiType: shortcode

  # Enable author sign-off for commits.
  # Values: true, false
  # Default: false
  signoff: false

authors:
  # List of extra authors.
  - name: John Doe
    email: john.doe@example.com
```

## 🏆 Best Practises [⭡](#committed)

To create a well formed commit, these are some of the best practises that are
often cited.

> Capitalized, short (50 chars or less) summary
>
> More detailed explanatory text, if necessary.  Wrap it to about 72
> characters or so.  In some contexts, the first line is treated as the
> subject of an email and the rest of the text as the body.  The blank
> line separating the summary from the body is critical (unless you omit
> the body entirely); tools like rebase can get confused if you run the
> two together.
>
> Write your commit message in the imperative: "Fix bug" and not "Fixed bug"
> or "Fixes bug."  This convention matches up with commit messages generated
> by commands like git merge and git revert.
>
> Further paragraphs come after blank lines.
>
> - Bullet points are okay, too
>
> - Typically a hyphen or asterisk is used for the bullet, followed by a
>   single space, with blank lines in between, but conventions vary here
>
> - Use a hanging indent

Source: [Tim Pope](https://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html)

The placeholder text for the summary and body will show these recommendations.

Related links:

- [Joel Parker Henderson](https://github.com/joelparkerhenderson/git-commit-message)
- [Chris Beams](https://cbea.ms/git-commit/)

## ⌨ Shortcuts [⭡](#committed)

The global shortcuts can be used within any view.

| Command            | Key       |
| :----------------- | :-------- |
| Commit             | alt-enter |
| Toggle sign-off    | alt-s     |
| Toggle theme       | alt+t     |
| Help               | alt-/     |
| Focus author       | alt-1     |
| Focus emoji        | alt-2     |
| Focus summary      | alt-3     |
| Focus body         | alt-4     |
| Cancel             | control-c |
| Next component     | tab       |
| Previous component | shift-tab |

The emoji shortcuts are limited to the emoji view only.

| Command       | Key       |
| :------------ | :-------- |
| Clear emoji   | delete    |
| Reset filter  | escape    |
| Next page     | page down |
| Previous page | page up   |

## 🔍 Views [⭡](#committed)

### Main

The main view when entering either the summary or body.

```text
commit 1234567890abcdef1234567890abcdef1234567890
Author: John Doe <john.doe@example.com>
Date:   Mon Jan 2 15:04:05 2006 -0700

    ┌───┐ ┌──────────────────────────────────────────────────────────────┐
    │ X │ │ Capitalized, short (50 chars or less) summary                │ 47/50
    └───┘ └──────────────────────────────────────────────────────────────┘

    ┌──────────────────────────────────────────────────────────────────────────┐
    │ More detailed explanatory text, if necessary.  Wrap it to about 72       │
    │ characters or so.  In some contexts, the first line is treated as the    │
    │ subject of an email and the rest of the text as the body.  The blank     │
    │ line separating the summary from the body is critical (unless you omit   │
    │ the body entirely); tools like rebase can get confused if you run the    │
    │ two together.                                                            │
    │                                                                          │
    │ Write your commit message in the imperative: "Fix bug" and not "Fixed    │
    │ bug" or "Fixes bug."  This convention matches up with commit messages    │
    │ generated by commands like git merge and git revert.                     │
    │                                                                          │
    │ Further paragraphs come after blank lines.                               │
    │                                                                          │
    │ - Bullet points are okay, too                                            │
    │                                                                          │
    │ - Typically a hyphen or asterisk is used for the bullet, followed by a   │
    │   single space, with blank lines in between, but conventions vary here   │
    │                                                                          │
    │ - Use a hanging indent                                                   │
    └──────────────────────────────────────────────────────────────────────────┘

      Signed-off-by: John Doe <john.doe@example.com>

 Alt + <enter> Commit <s> Sign-off </> Help                Summary <tab>
Ctrl +     <c> Cancel                                       Author <tab> + Shift
```

### Emoji

The emoji view reduces the position and lines of the body section to make space
for a selector to appear. Filtering is available to narrow the choices down.

```text
commit 1234567890abcdef1234567890abcdef1234567890
Author: John Doe <john.doe@example.com>
Date:   Mon Jan 2 15:04:05 2006 -0700

    ┌───┐ ┌──────────────────────────────────────────────────────────────┐
    │ X │ │ Capitalized, short (50 chars or less) summary                │ 47/50
    └───┘ └──────────────────────────────────────────────────────────────┘

    ┌──────────────────────────────────────────────────────────────────────────┐
    │? Choose an emoji: █                                                      │
    │> x - Improve structure / format of the code.                             │
    │  x - Improve performance.                                                │
    │  x - Remove code or files.                                               │
    │  x - Fix a bug.                                                          │
    │  x - Critical hotfix.                                                    │
    │  x - Introduce new features.                                             │
    │  x - Add or update documentation.                                        │
    │  x - Deploy stuff.                                                       │
    │  x - Add or update the UI and style files.                               │
    └──────────────────────────────────────────────────────────────────────────┘

    ┌──────────────────────────────────────────────────────────────────────────┐
    │ More detailed explanatory text, if necessary.  Wrap it to about 72       │
    │ characters or so.  In some contexts, the first line is treated as the    │
    │ subject of an email and the rest of the text as the body.  The blank     │
    │ line separating the summary from the body is critical (unless you omit   │
    │ the body entirely); tools like rebase can get confused if you run the    │
    │ two together.                                                            │
    └──────────────────────────────────────────────────────────────────────────┘

      Signed-off-by: John Doe <john.doe@example.com>

 Alt + <enter> Commit <s> Sign-off </> Help                Summary <tab>
Ctrl +     <c> Cancel                                       Author <tab> + Shift
```

### Author

The author view moves the subject line down and reduces the height of the body
section. This provides space for a selector to choose the commit author.

```text
commit 1234567890abcdef1234567890abcdef1234567890
Author: John Doe <john.doe@example.com>
Date:   Mon Jan 2 15:04:05 2006 -0700

    ┌──────────────────────────────────────────────────────────────────────────┐
    │? Choose an author: █                                                     │
    │> John Doe <john.doe@example.com>                                         │
    │  John Doe <john.doe@example.org>                                         │
    │                                                                          │
    └──────────────────────────────────────────────────────────────────────────┘

    ┌───┐ ┌──────────────────────────────────────────────────────────────┐
    │ X │ │ Capitalized, short (50 chars or less) summary                │ 47/50
    └───┘ └──────────────────────────────────────────────────────────────┘

    ┌──────────────────────────────────────────────────────────────────────────┐
    │ More detailed explanatory text, if necessary.  Wrap it to about 72       │
    │ characters or so.  In some contexts, the first line is treated as the    │
    │ subject of an email and the rest of the text as the body.  The blank     │
    │ line separating the summary from the body is critical (unless you omit   │
    │ the body entirely); tools like rebase can get confused if you run the    │
    │ two together.                                                            │
    │                                                                          │
    │ Write your commit message in the imperative: "Fix bug" and not "Fixed    │
    │ bug" or "Fixes bug."  This convention matches up with commit messages    │
    │ generated by commands like git merge and git revert.                     │
    │                                                                          │
    │ Further paragraphs come after blank lines.                               │
    └──────────────────────────────────────────────────────────────────────────┘

      Signed-off-by: John Doe <john.doe@example.com>

 Alt + <enter> Commit <s> Sign-off </> Help                  Emoji <tab>
Ctrl +     <c> Cancel                                                    + Shift
```

### Commit

Accepting the commit shows the output that will closely match the `git log`
command.

```text
commit 1234567890abcdef1234567890abcdef1234567890
Author: John Doe <john.doe@example.com>
Date:   Mon Jan 2 15:04:05 2006 -0700

     X Capitalized, short (50 chars or less) summary

     More detailed explanatory text, if necessary.  Wrap it to about 72
     characters or so.  In some contexts, the first line is treated as the
     subject of an email and the rest of the text as the body.  The blank
     line separating the summary from the body is critical (unless you omit
     the body entirely); tools like rebase can get confused if you run the
     two together.

     Write your commit message in the imperative: "Fix bug" and not "Fixed
     bug" or "Fixes bug."  This convention matches up with commit messages
     generated by commands like git merge and git revert.

     Further paragraphs come after blank lines.

     - Bullet points are okay, too

     - Typically a hyphen or asterisk is used for the bullet, followed by a
       single space, with blank lines in between, but conventions vary here

     - Use a hanging indent

     Signed-off-by: John Doe <john.doe@example.com>

[master 1234567] Capitalized, short (50 chars or less) summary
 3 files changed, 2 insertions(+), 1 deletions(-)
```

## ✏️ Authors [⭡](#committed)

- [@mikelorant](https://www.github.com/mikelorant)

## 🎫 License [⭡](#committed)

[MIT](https://choosealicense.com/licenses/mit/)

## 👍 Thanks [⭡](#committed)

Thanks to [Carlos Cuesta](https://github.com/carloscuesta) for creating [gitmoji](https://gitmoji.dev/) and [gitmoji-cli](https://github.com/carloscuesta/gitmoji-cli) which was the
inspiration for this project.

Thanks to [Ahmad Awais](https://github.com/ahmadawais) for [Emoji-Log](https://github.com/ahmadawais/Emoji-Log) and [Folke Lemaitre](https://github.com/folke) for [Devmoji](https://github.com/folke/devmoji).

Many thanks to [David Ackroyd](https://github.com/dackroyd) and [Matt Hope](https://github.com/matthope) for all their guidance with Go.
Without their expertise I would never had the capability to build Committed.

Thanks to all the developers from [Charm](https://github.com/charmbracelet) for their amazing set of libraries.
Committed would never have looked the way it does without [Bubble Tea](https://github.com/charmbracelet/bubbletea), [Lipgloss](https://github.com/charmbracelet/lipgloss)
and [Bubbles](https://github.com/charmbracelet/bubbles).

Thanks to [Tim Pope](https://github.com/tpope) for his Git commit recommendations which was a core
component in the interface design.
