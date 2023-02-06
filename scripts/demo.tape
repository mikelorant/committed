# Requirements:
# - Use ttyd 1.7.2 which resolves missing cursor.
# - Use current build of vhs to resolve width issues.
# - Modify high-voltage with basic version. (emoji.yaml)
# - Copy $HOME/.config/committed/config.yaml to current directory.
# - Modify config to use:
#     view.compatibility = ttyd (config.yaml)
#     authors:
#       - name: John Doe
#         email: jdoe@example.org
# - Comment out user in Git config. (~/.gitconfig)
# - Build release: go build -tags release

Output demo.gif

Set Width 794
Set Height 760
Set Padding 18
Set FontSize 16
Set Theme "Builtin Dark"

Hide
# Sets up a fake commit.
# Sleep 1s
# Type "echo 1 >> test"
# Sleep 0.3 Enter Sleep 0.3
# Type "git add test"
# Sleep 0.3 Enter Sleep 0.3
# Type "export LIPGLOSS_TERMINAL=ttyd"
# Sleep 0.3 Enter Sleep 0.3
# Type "alias committed='./committed --config ./demo.yaml'"
# Sleep 0.3 Enter Sleep 0.3
# Type "clear"
# Sleep 0.3 Enter Sleep 0.3

Show
Sleep 1s
Type "committed"
Sleep 1s
Enter

Sleep 3s
ShiftTab
Sleep 3s
Enter

Sleep 1s
PageDown
Sleep 2s
PageDown
Sleep 2s
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
Alt+S

Sleep 3s
Alt+ENTER
Sleep 5s

Hide
Ctrl+D
