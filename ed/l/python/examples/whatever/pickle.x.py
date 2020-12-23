import pickle


FILE_NAME = "/tmp/py.pickle.x"

def pack():
    data = {"foo": "bar"}
    f = open(FILE_NAME, "wb")
    pickle.dump(data, f)
    f.close()


def unpack():
    f = open(FILE_NAME, "rb")
    data = pickle.load(f)
    f.close()
    print(f"data = {data}")


pack()
unpack()
