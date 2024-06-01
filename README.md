# Committed

[![Release][release-badge]][release]
[![Build][build-badge]][build]
[![License][license-badge]][license]
[![Codecov][codecov-badge]][codecov]
[![CodeFactor][codefactor-badge]][codefactor]

Committed is a WYSIWYG Git commit editor that helps improve the quality of your
commits by showing you the layout in the same format as `git log`.

[release]: https://github.com/mikelorant/committed/releases
[release-badge]: https://img.shields.io/github/v/release/mikelorant/committed
[build]: https://github.com/mikelorant/committed/actions/workflows/release.yaml
[build-badge]: https://github.com/mikelorant/committed/actions/workflows/release.yaml/badge.svg
[license]: https://spdx.org/licenses/MIT.html
[license-badge]: https://img.shields.io/github/license/mikelorant/committed
[codecov]: https://codecov.io/gh/mikelorant/committed
[codecov-badge]: https://codecov.io/gh/mikelorant/committed/branch/main/graph/badge.svg?token=TW2XDZAJKT
[codefactor]: https://www.codefactor.io/repository/github/mikelorant/committed
[codefactor-badge]: https://www.codefactor.io/repository/github/mikelorant/committed/badge

## [![Highlights][highlights-badge]](#-highlights-) [![First Steps][first-steps-badge]](#-first-steps-) [![Installation][installation-badge]](#-installation-) [![Usage][usage-badge]](#-usage-) [![Configuration][configuration-badge]](#-configuration-) [![Best Practises][best-practises-badge]](#-best-practises-) [![Shortcuts][shortcuts-badge]](#-shortcuts-)

[highlights-badge]: https://img.shields.io/badge/-💡_Highlights-ff0000
[first-steps-badge]: https://img.shields.io/badge/-🐾_First_Steps-ffa500
[installation-badge]: https://img.shields.io/badge/-💾_Installation-ffff00
[usage-badge]: https://img.shields.io/badge/-🎛_Usage-008000
[configuration-badge]: https://img.shields.io/badge/-⚙️_Configuration-0000ff
[best-practises-badge]: https://img.shields.io/badge/-🏆_Best_Practises-4b0082
[shortcuts-badge]: https://img.shields.io/badge/-⌨️_Shortcuts-ee82ee

![demo](docs/demo.gif)

## 💡 Highlights [⭡](#committed)

- Built-in **multiline editor** with rich capabilities.
- Custom **emoji selector** providing popular sets to choose from.
- **Switch author** before applying the commit.
- Inline **text interface** mimics the Git log output.
- Dynamic **subject line counter**.
- Toggle appending **sign-off** required by many open source projects.
- Automatically **hard wraps** body to 72 characters.
- Best practise **recommendations**.
- Import and **amend** previous commit.
- **Adaptive colours** with **light** and **dark** themes.

## 🐾 First Steps [⭡](#committed)

1. Install using Homebrew.

   ```shell
   brew install mikelorant/committed/committed
   ```

2. Before creating and applying a commit you will need to stage the files you
   wish to add with the `git add` command.

3. Committed replaces the `git commit` command and all you need to do to commit
   your change is to run:

   ```shell
   committed
   ```

   It is also possible to amend your previous commit with:

   ```shell
   committed --amend
   ```

Once the UI has appeared take note of the keyboard shortcuts shown at the bottom
of the interface.

```text
 Alt + <enter> Commit <s> Sign-off </> Help                Summary <tab>
Ctrl +     <c> Cancel                                       Author <tab> + Shift
```

These shortcuts will help you apply or cancel a commit and navigate between the
different components. Pressing enter on most components will automatically
switch you to the next one.

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

## ⚠️ Limitations [⭡](#committed)

### Option Key

The option key needs to be set to send the `meta` or `esc+` keycode. Terminals
such as macOS Terminal or iTerm2 may not have this as default. If not set
correctly it will not be possible to apply a commit.

To make these changes following the instructions below.

| Terminal       | Setting                                               |
| :------------- | :---------------------------------------------------- |
| macOS Terminal | `Profiles` `Keyboard`<br />![macOS Terminal options]  |
| iTerm2         | `Preferences` `Profile` `Keys`<br />![iTerm2 options] |

The alternative keyboard shortcut <kbd>⌥ Option</kbd> + <kbd>\\</kbd> can also
be used to apply a commit.

[macos terminal options]: docs/keyboard-options-macos-terminal.png
[iterm2 options]: docs/keyboard-options-iterm2.png

### Rendering Borders

Terminals render emojis differently and this makes alignment of borders
complicated and difficult. It is an ongoing process to improve the compatibility
with terminals. The following list are the terminals that have been tested.
Other terminals may display correctly and feedback would be appreciated to help
update the list.

| Terminal                   | Status         | Compatibility Mode |
| :------------------------- | :------------- | :----------------- |
| macOS Terminal[¹](#iTerm2) | ![✅][confirm] | Unicode 9          |
| iTerm2                     | ![✅][confirm] | Unicode 14         |
| Visual Studio Code         | ![✅][confirm] | Unicode 9          |
| Hyper                      | ![✅][confirm] | Unicode 9          |
| Alacritty                  | ![✅][confirm] | Unicode 9          |
| WezTerm[²](#WezTerm)       | ![✅][confirm] | Unicode 9          |
| ttyd                       | ![✅][confirm] | Unicode 9          |
| kitty                      | ![✅][confirm] | Unicode 14         |

#### iTerm2

iTerm2 versions lower than `3.5.0` will need to use Unicode 9 compatibility.

It is also important to make sure that options related to [variation selector 16][iterm2 vs16]
use full width when using Unicode 14 compatibility. This requires setting
`Support variation selector 16 making emoji fullwidth outside of alternate
screen mode?` to `Yes`.

To correctly align emojis using variation select 16 within `git log`, the
setting `Support variation selector 16 making emoji fullwidth in all modes?`
must be set to `Yes`.

[iterm2 vs16]: https://gitlab.com/gnachman/iterm2/-/issues/10480

#### WezTerm

WezTerm has [configurable][wezterm unicode] Unicode version support. This
currently defaults to Unicode 9.

[wezterm unicode]: https://wezfurlong.org/wezterm/config/lua/config/unicode_version.html

## 💾 Installation [⭡](#committed)

Install Committed with Homebrew.

```shell
brew install mikelorant/taps/committed
```

## ⚙ Usage [⭡](#committed)

```text
Committed is a WYSIWYG Git commit editor

Usage:
  committed [flags]
  committed [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  hook         Install and uninstall Git hook
  list         List settings with profiles or IDs
  version      Print the version information

Flags:
      --config string     Config file location (default
                          "$HOME/.config/committed/config.yaml")
      --snapshot string   Snapshot file location (default
                          "$HOME/.local/state/committed/snapshot.yaml")
      --dry-run           Simulate applying a commit (default false)
  -a, --amend             Replace the tip of the current branch by creating a new commit
  -h, --help              help for committed
  -v, --version           version for committed

Use "committed [command] --help" for more information about a command.
```

### List

```text
Usage:
  committed list [command]

Available Commands:
  emojis       List emoji profiles
  themes       List theme IDs
```

### Hook

```text
Usage:
  committed hook [flags]

Flags:
      --install     Install Git hook
      --uninstall   Uninstall Git hook
```

## 🎛 Configuration [⭡](#committed)

No configuration is necessary however there are some values that can be changed
based on preference.

Committed defaults to using a config file located at
`$HOME/.config/committed/config.yaml`.

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

  # Terminal compatibility.
  # Values: unicode14, unicode9
  # Default: unicode14
  compatibility: unicode14

  # Highlight active component.
  # Value: true, false
  # Default: false
  highlightActive: false

  # Ignore Git global author.
  # Value: true, false
  # Default: false
  ignoreGlobalAuthor: false

commit:
  # Emoji format in commit.
  # Values: shortcode, character
  # Default: character
  emojiType: character

  # Enable author sign-off for commits.
  # Values: true, false
  # Default: false
  signoff: false

authors:
  # List of extra authors.
  - name: John Doe
    email: john.doe@example.com
```

### Themes

There are a number of themes available that modify the colours. By default, the
background colour is detected which changes the choices of themes. This
detection can be disabled by setting the colour profile in the configuration.
The first theme of each set is the default theme applied.

#### Dark Themes

| Name                                        | ID                             |
| :------------------------------------------ | :----------------------------- |
| Builtin Dark                                | builtin_dark                   |
| [Dracula][dracula]                          | dracula                        |
| [Gruvbox Dark][gruvbox]                     | gruvbox_dark                   |
| [Nord][nord]                                | nord                           |
| [Retrowave][retrowave]                      | retrowave                      |
| [Solarized Dark Higher Contrast][solarized] | solarized_dark_higher_contrast |
| [Tokyo Night][tokyo-night]                  | tokyo_night                    |

#### Light Theme

| Name                                 | ID                      |
| :----------------------------------- | :---------------------- |
| Builtin Light                        | builtin_light           |
| [Builtin Solarized Light][solarized] | builtin_solarized_light |
| [Builtin Tango Light][tango]         | builtin_tango_light     |
| [Gruvbox Light][gruvbox]             | gruvbox_light           |
| [Tokyo Night Light][tokyo-night]     | tokyo_night_light       |

[dracula]: https://draculatheme.com/
[gruvbox]: https://github.com/morhetz/gruvbox
[nord]: https://www.nordtheme.com/
[retrowave]: https://github.com/juanmnl/vs-1984
[solarized]: https://ethanschoonover.com/solarized/
[tokyo-night]: https://github.com/enkia/tokyo-night-vscode-theme
[tango]: http://tango.freedesktop.org/Tango_Desktop_Project

### Emoji Profiles

Popular emoji sets can be set as the default profile:

- [Gitmoji](https://gitmoji.dev/)
- [Devmoji](https://github.com/folke/devmoji)
- [Emoji-Log](https://github.com/ahmadawais/emoji-log)

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

| Key Binding                              | Command            |
| :--------------------------------------- | :----------------- |
| <kbd>⌥ Option</kbd> + <kbd>⏎ Enter</kbd> | Commit             |
| <kbd>⌥ Option</kbd> + <kbd>\\</kbd>      | Commit             |
| <kbd>⌥ Option</kbd> + <kbd>S</kbd>       | Toggle sign-off    |
| <kbd>⌥ Option</kbd> + <kbd>T</kbd>       | Toggle theme       |
| <kbd>⌥ Option</kbd> + <kbd>/</kbd>       | Help               |
| <kbd>⌥ Option</kbd> + <kbd>1</kbd>       | Focus author       |
| <kbd>⌥ Option</kbd> + <kbd>2</kbd>       | Focus emoji        |
| <kbd>⌥ Option</kbd> + <kbd>3</kbd>       | Focus summary      |
| <kbd>⌥ Option</kbd> + <kbd>4</kbd>       | Focus body         |
| <kbd>⌃ Control</kbd> + <kbd>C</kbd>      | Cancel             |
| <kbd>⇥ Tab</kbd>                         | Next component     |
| <kbd>⇧ Shift</kbd> + <kbd>⇥ Tab</kbd>    | Previous component |

The emoji shortcuts are limited to the emoji view only.

| Key Binding            | Command       |
| :--------------------- | :------------ |
| <kbd>⌫ Delete</kbd>    | Clear emoji   |
| <kbd>⎋ Escape</kbd>    | Reset filter  |
| <kbd>⇟ Page Down</kbd> | Next page     |
| <kbd>⇞ Page Up</kbd>   | Previous page |

## 📚 Tips [⭡](#committed)

### Aliases

Shell or Git aliases can be used to tailor Committed to your preferred workflow.
An example Git alias is as follows:

```shell
git config --global alias.co '! committed'
```

You can then commit changes with:

```shell
git co
```

### Prepare Message Hook

Committed can be installed as a Git prepare message hook. Be aware that any
existing `prepare-commit-msg` hook will not be replaced and it is necessary to
remove this hook before installing.

Installation:

```shell
committed hook --install
```

Removal:

```shell
committed hook --uninstall
```

### Editor

Committed can replace the default Git editor which allows commits to be applied
using `git commit`. Most of the standard Git command arguments can be used.

```shell
git config --global core.editor "committed --editor"
```

This can be removed with:

```shell
git config --global --unset-all core.editor
```

There are some limitations related to acting as an editor.

- Comment lines will be truncated to the width of the editor.
- Interactive rebasing and other operations which require edit commands may have
  visual issues. The first command may be part of the summary.
- Author cannot be set. The configured Git author will be used and will be
  selected using the default Git method (repository followed by global).
- When amending, subject line may be part of the body.
- When amending, emoji character or shortcode must be in the existing data set.
- When amending, summary will be truncated if more than 72 characters.
- When amending, trailers will be imported into the body.

### Amend

There are certain limitations when amending commits and it is recommended only
for use with commits created with Committed. The limitations share similarities
with using Git as an editor.

- Emoji character or shortcode must be in the existing data set.
- Trailers will be imported into the body.
- Summary will be truncated if more than 72 characters.
- Lines will not reflow when editing the body.

## ⚖️ Comparison

### Git Functions

| Feature                      | Committed      | Gitmoji CLI       |
| :--------------------------- | :------------- | :---------------- |
| Git hooks                    | ![✅][confirm] | ![✅][confirm]    |
| Git editor replacement       | ![✅][confirm] | ![❌][cancel]     |
| Amend commit                 | ![✅][confirm] | ![❌][cancel]     |
| Add files                    | ![❌][cancel]  | ![✅][confirm]    |
| Signed commits               | ![❌][cancel]  | ![✅][confirm]    |
| Sign-off commits             | ![✅][confirm] | ![❌][cancel]     |
| Switch author                | ![✅][confirm] | ![❌][cancel]     |
| Save and load failed commits | ![✅][confirm] | ![❌][cancel][^1] |

[^1]: [Print Git command on failure](https://github.com/carloscuesta/gitmoji-cli/pull/681).

### Visual Style

| Feature                 | Committed      | Gitmoji CLI    |
| :---------------------- | :------------- | :------------- |
| Subject counter         | ![✅][confirm] | ![✅][confirm] |
| Custom emojis           | ![❌][cancel]  | ![✅][confirm] |
| Mulitple emoji profiles | ![✅][confirm] | ![❌][cancel]  |
| Multiline editor        | ![✅][confirm] | ![❌][cancel]  |
| Scope prompt            | ![❌][cancel]  | ![✅][confirm] |
| Themes                  | ![✅][confirm] | ![❌][cancel]  |
| Hard wrap body width    | ![✅][confirm] | ![❌][cancel]  |

## ✏️ Authors [⭡](#committed)

- [@mikelorant](https://www.github.com/mikelorant)

## 🎫 License [⭡](#committed)

[MIT](https://choosealicense.com/licenses/mit/)

## 👍 Thanks [⭡](#committed)

Thanks to [Carlos Cuesta] for creating [gitmoji] and [gitmoji-cli] which was the
inspiration for this project.

Thanks to [Ahmad Awais] for [Emoji-Log] and [Folke Lemaitre] for [Devmoji].

Many thanks to [David Ackroyd] and [Matt Hope] for all their guidance with Go.
Without their expertise I would never had the capability to build Committed.

Thanks to all the developers from [Charm] for their amazing set of libraries.
Committed would never have looked the way it does without [Bubble Tea],
[Lipgloss] and [Bubbles].

Thanks to [Tim Pope] for his Git commit recommendations which was a core
component in the interface design.

[carlos cuesta]: https://github.com/carloscuesta
[ahmad awais]: https://github.com/ahmadawais
[folke lemaitre]: https://github.com/folke
[david ackroyd]: https://github.com/dackroyd
[matt hope]: https://github.com/matthope
[charm]: https://github.com/charmbracelet
[tim pope]: https://github.com/tpope

[gitmoji]: https://gitmoji.dev/
[gitmoji-cli]: https://github.com/carloscuesta/gitmoji-cli
[emoji-log]: https://github.com/ahmadawais/Emoji-Log
[devmoji]: https://github.com/folke/devmoji
[bubble tea]: https://github.com/charmbracelet/bubbletea
[lipgloss]: https://github.com/charmbracelet/lipgloss
[bubbles]: https://github.com/charmbracelet/bubbles

[//]: # (shared link reference definitions)

[confirm]: docs/assets/confirm-24x24.svg
[cancel]: docs/assets/cancel-24x24.svg
