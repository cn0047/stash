class W:
    def __enter__(self):
        print('enter')

    def __exit__(self, type, val, traceback):
        print('exit')


with W() as w:
    print(w)

"""
enter
None
exit
"""
