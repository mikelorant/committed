# Committed

[![Release](https://img.shields.io/github/v/release/mikelorant/committed)](https://github.com/mikelorant/committed/releases)[![Build](https://github.com/mikelorant/committed/actions/workflows/release.yaml/badge.svg)](https://github.com/mikelorant/committed/actions/workflows/release.yaml)[![License](https://img.shields.io/github/license/mikelorant/committed)](https://spdx.org/licenses/MIT.html)

Committed is a WYSIWYG Git commit editor that helps improve the quality of your
commits by showing you the layout in the same format as `git log`.

<img src="docs/demo.gif" align="left" style="zoom:100%;" />

## Highlights

- Built-in multiline editor
- Emoji selector
- Author profile switching
- Inline text interface
- Subject line counter
- Appends sign-off
- Formats body width to 72 characters
- Best practise recommendations

## Purpose

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

## Compatibility

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

## Installation

Install Committed with Homebrew.

```bash
brew install mikelorant/taps/committed
```

## Usage

```
Committed is a WYSIWYG Git commit editor

Usage:
  committed [flags]
  committed [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  version      Print the version information

Flags:
      --dry-run   Simulate applying a commit
  -a, --amend     Replace the tip of the current branch by creating a new commit
  -h, --help      help for committed
  -v, --version   version for committed

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

## Views

### Main

The main view when entering either the summary or body.

```
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

```
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

```
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

```
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

## Shortcuts

**Global**

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

**Emoji**

| Command       | Key       |
| :------------ | :-------- |
| Clear emoji   | delete    |
| Reset filter  | escape    |
| Next page     | page down |
| Previous page | page up   |

## Authors

- [@mikelorant](https://www.github.com/mikelorant)

## License

[MIT](https://choosealicense.com/licenses/mit/)
