# @example: python3 ed/l/python/examples/http/server.prof.py
# @example: curl localhost:8080

import time
import cProfile
from http.server import HTTPServer, BaseHTTPRequestHandler


class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):

    def do_GET(self):
        pr = cProfile.Profile()
        pr.enable()

        self.send_response(200)
        self.end_headers()
        v = fib(15)
        msg = 'My Fib is {} \n'.format(v)
        self.wfile.write(msg.encode())

        pr.disable()
        pr.dump_stats('callgrind.cprof.{}'.format(time.time()))


def fib(n: int) -> int:
    if n < 2:
        return 1
    return fib(n - 1) + fib(n - 2)


def web():
    httpd = HTTPServer(('localhost', 8080), SimpleHTTPRequestHandler)
    httpd.serve_forever()


if __name__ == '__main__':
    web()
