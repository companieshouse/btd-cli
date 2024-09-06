# Configuration file for generating btd-cli.gif image using VHS
# command-line tool (https://github.com/charmbracelet/vhs).
#
# To run:
#  vhs < btd-cli.tape

Output btd-cli.gif

Set Framerate 30
Set Width 1400
Set Height 880

Type "btd-cli parse string "

Copy "'10000008AB01234520070009Crown Way20270007Cardiff50020008CF14 3UZ20060007Cardiff50010012Non Existent50050013Image Systems200200007004000911/22/33/70240001A70060003acc70070011ABCDEFGH12370080001170090003123501200051234570000003123702200011702300010702200012'"

Sleep 1s

Paste

Sleep 2s

Enter

Sleep 8s
