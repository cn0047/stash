msg="Line 3."
cat << EndOfMessage
  $(echo -e "\033[32m") This is HEREDOC: $(echo -e "\033[0m")
    Line 1.
    Line 2.
    $msg
EndOfMessage
