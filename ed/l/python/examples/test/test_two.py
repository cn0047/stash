import pytest
import subprocess

from unittest.mock import patch, MagicMock

def pwd() -> str:
  cmd = 'pwd'
  proc = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True)
  (out, error_code) = proc.communicate()
  return out

def test_pwd_a():
  r = pwd()
  assert b'web' in r

@patch("subprocess.Popen")
def test_pwd_b(mock_popen):
  m = MagicMock()
  m.communicate.return_value = (b"", b"")
  m.returncode = 0
  mock_popen.return_value = m

  r = pwd()
  cli_command = mock_popen.call_args_list[0][0][0]

  assert cli_command == 'pwd'
