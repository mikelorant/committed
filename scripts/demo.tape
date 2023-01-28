# Requirements:
# - Build release: go build -tags release

Output demo.gif

Set Width 794
Set Height 760
Set Padding 18
Set FontSize 16
Set Theme "Builtin Dark"
Set LetterSpacing 0

Hide
# Sets up a fake commit.
Type "cd /tmp"                                                  Enter Wait
Type "git init"                                                 Enter Wait
Type "git commit --allow-empty --allow-empty-message --no-edit" Enter Wait
Type "touch test"                                               Enter Wait
Type "git add test"                                             Enter Wait
Type "clear"                                                    Enter Wait

Show
Sleep 1s
Type "committed"
Sleep 1s
Enter

Sleep 3s
Shift+Tab
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
Alt+Enter
Sleep 5s

Hide
Ctrl+D
