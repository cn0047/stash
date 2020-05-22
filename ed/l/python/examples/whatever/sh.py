import subprocess


def f1():
  out = subprocess.call(["ls", "-lha"])
  print(out)


def f2():
  out = subprocess.check_call(["ls", "-lha"]) # throw exception if error
  print(out)


def f3():
    cmd = 'echo ok'
    proc = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True)
    (out, err) = proc.communicate()
    if err:
        raise BlurringError(f"error: {err}")
    print(out)


# f1()
f2()
# f3()
