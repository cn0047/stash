from flask import request
from flask_api import FlaskAPI
from flask_cors import CORS


app = FlaskAPI(__name__)
CORS(app)


@app.route('/')
def hello_world():
    return 'Hello, World!'


@app.route('/x', methods=['GET', 'POST'])
def x():
  if request.method == 'POST':
    return 'It is post.'
  return {'data': 'It is get.'}


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8080, debug=True)
