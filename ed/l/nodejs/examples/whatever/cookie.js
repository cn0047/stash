var countdown = 360;
var matches = (document.cookie).match(/countdown=([\d]+)/);
if (matches) {
    countdown = matches[1];
} else {
    var d = new Date();
    d.setYear(d.getFullYear()+1);
    var expires = "expires=" + d.toUTCString();
    document.cookie = 'countdown=' + countdown + '; ' + expires + ': secure';
    document.cookie = 'ip_address=88.37.52.171; ' + expires;
}
// delete

// excited size
