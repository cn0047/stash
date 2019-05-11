$(function () {
    if (document.location.hash) {
        // Hide all pages.
    $('.page').addClass('hide');
    // Show selected page.
    $('.page[data-id="'+document.location.hash+'"]').removeClass('hide');
    }
    $('.showPage').click(function () {
        // Hide dropdown main menu.
        $('#navbar').removeClass('in');
        // Hide all pages.
        $('.page').addClass('hide');
        // Show selected page.
        $('.page[data-id="'+$(this).attr('href')+'"]').removeClass('hide');
    });
});
