# Requirements:
# - Enable single space padding for all emojis (emoji.go)
# - Modify high-voltage emoji with text version (emoji.yaml)
# - Add keyboard alternatives: (ui.go)
#     alt+enter -> ctrl+f
#     alt+s -> ctrl+s
#     shift+tab -> ctrl+b
# - Replace Lipgloss with local package (go.mod)
# - Lipgloss emoji changes: (get.go)
#     Add package "github.com/rivo/uniseg"
#     Change "PrintableStringWidth" function to return:
#       uniseg.GraphemeClusterCount(string(sr))
# - Git config changes: (~/.gitconfig)
#     name = John Doe
#     email = john.doe@example.org

Output demo.gif

Set Width 720
Set Height 640
Set Padding 36
Set FontSize 12
Set FontFamily "JetBrains Mono"
Set Theme "Builtin Dark"

Hide
# Sets up a fake commit.
# Sleep 1s
# Type "echo 1 >> test"
# Sleep 0.3 Enter Sleep 0.3
# Type "git add test"
# Sleep 0.3 Enter Sleep 0.3

Sleep 1s
Type "./committed --yes"
Enter
Sleep 5s

Show
Sleep 3s
Ctrl+B
Sleep 3s
Enter

Sleep 1s
Type@200ms "bug"
Sleep 2s
Enter

Sleep 1s
Type "Prevent racing of requests"
Sleep 1s
Enter

Sleep 1s
Type "Introduce a request id and a reference to latest request. Dismiss incoming responses other than from latest request."
Enter
Sleep 0.5s
Enter
Sleep 0.5s
Type "Remove timeouts which were used to mitigate the racing issue but are obsolete now."
Sleep 0.5s
Enter

Sleep 1s
Ctrl+S

Sleep 3s
Ctrl+F
Sleep 5s

Hide
Ctrl+D
