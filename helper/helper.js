$(function () {
    $('#interchange').click(function () {
        var tmp = $('#needles').val();
        $('#needles').val($('#haystack').val());
        $('#haystack').val(tmp);
    });
    $('.btn-primary').click(function () {
        $('#imgProgress').removeClass('hide');
        $('#action').val($(this).attr('id'));
        $.ajax({url: 'helper.php', type: 'post', dataType: 'json', data: $('section :input').serialize(),
            success: function (r) {
                var t = '';
                if ($.isArray(r.array)) {
                    t += r.array.join('\n');
                }
                t += '\n';
                if (r.text) {
                    t += '\n'+r.text;
                }
                $('#result').val(t);
            },
            error: function (x, s, e) {
                console.error(x);
                console.error(s);
                console.error(e);
            },
            complete: function () {
                $('#imgProgress').addClass('hide');
            }
        });
    });
});
