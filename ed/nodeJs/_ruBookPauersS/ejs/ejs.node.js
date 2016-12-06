// npm install ejs

var http = require('http')
    , ejs = require('ejs') ;
// создание http-сервера
http.createServer(function (req, res) {
    res.writeHead(200, {'content-type': 'text/html'});
    // визуализируемые данные
    var names = ['Joe', 'Mary', 'Sue', 'Mark'];
    var title = 'Testing EJS';
    // вывод данных или ошибки
    ejs.renderFile(
        __dirname + '/views/test.ejs',
        {title : 'testing', names : names},
        function(err, result) {
            if (!err) {
               res.end(result);
            } else {
                res.end('An error occurred accessing page');
                console.log(err);
            }
        }
    );
}).listen(8124);
console.log('Server running on 8124/');

/*
var names = ['Joe Brown', 'Mary Smith', 'Tom Thumb', 'Cinder Ella'];
var str = '<p><%=: users | first | downcase %></p>';
var html = ejs.render(str, {users : names });
Получается следующий результат:
<p>joe brown</p>

var people = [
{name : 'Joe Brown', age : 32},
{name : 'Mary Smith', age : 54},
{name : 'Tom Thumb', age : 21},
{name : 'Cinder Ella', age : 16}];
var str = "<p><%=: people | map:'name' | sort | join %></p>";
var html = ejs.render(str, {people : people });
Результат применения этой комбинации фильтров выглядит следующим образом:
Cinder Ella, Joe Brown, Mary Smith, Tom Thumb
*/
