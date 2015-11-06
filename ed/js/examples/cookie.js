var countdown = 360;
var matches = (document.cookie).match(/countdown=([\d]+)/);
if (matches) {
    countdown = matches[1];
} else {
    var d = new Date();
    d.setYear(d.getFullYear()+1);
    var expires = "expires=" + d.toUTCString();
    document.cookie = 'countdown=' + countdown + '; ' + expires;
}
