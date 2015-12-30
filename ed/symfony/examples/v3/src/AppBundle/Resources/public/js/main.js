$(function () {

    $('#comment').keypress(function (e) {
        if (e.which == 13) {
            console.log(window.location.pathname);
        }
    });
});
