# @example: python3 ed/l/python/examples/http/server.simple.py
# @example: curl localhost:8080

from http.server import HTTPServer, BaseHTTPRequestHandler


class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):

    def do_GET(self):
        self.send_response(200)
        self.end_headers()
        self.wfile.write(b'Hello world!')


def web():
    httpd = HTTPServer(('localhost', 8080), SimpleHTTPRequestHandler)
    httpd.serve_forever()


if __name__ == '__main__':
    web()
