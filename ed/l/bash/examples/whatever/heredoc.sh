msg="Line [3]."

# example 1
cat << EndOfMessage
  $(echo -e "\033[32m") This is HEREDOC: $(echo -e "\033[0m")
    Line 1.
    Line 2.
    $msg
    $msg
    $msg
EndOfMessage

# example 2
read -r -d '' endStr << VAR
=== End ===
VAR
echo $endStr

# example 3
IFS=$'\n' # IMPORTANT!!!
read -r -d '' content << VAR
one
two
three
{"foo":"bar"}
{cron":"0 1 * * *"}
VAR
# print
for line in $content; do
  echo "==| ${line}"
done
