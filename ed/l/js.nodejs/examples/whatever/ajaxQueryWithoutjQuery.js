// prepare post parameters
var form = document.getElementById('create');
var parameters = [];
for (i = form.elements.length - 1; i >= 0; i = i - 1) {
    var type = form.elements[i].nodeName;
    if (type.toLowerCase() === 'input') {
        parameters.push(
            form.elements[i].name+'='+encodeURIComponent(form.elements[i].value)
        );
    }
}
parameters = parameters.join('&');
// process request
var xmlHttp = new XMLHttpRequest();
var url = 'http://dev.server.com/?page=reg';
xmlHttp.open('POST', url, true);
xmlHttp.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
xmlHttp.setRequestHeader('Content-length', parameters.length);
xmlHttp.setRequestHeader('Connection', 'close');
xmlHttp.onreadystatechange = function () {
    if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
        var r = JSON.parse(xmlHttp.responseText);
        if (typeof r.status !== 'undefined' && r.status === 'fail') {
            if (typeof r.errors !== 'undefined') {
                for (inputName in r.errors) {
                    var el = document.getElementsByName(inputName)[0];
                    el.classList.add('error');
                    setTimeout(
                        function () {
                            el.classList.remove('error');
                        },
                        4000
                    );
                    addNotification('error', r.errors[inputName]);
                }
            }
        }
        if (r.status === 'ok') {
            console.log(200);
            hide();
        }
    }
}
xmlHttp.send(parameters);
