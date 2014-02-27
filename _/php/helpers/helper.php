<?

if (!empty($_POST['action'])) {
    $message = array();
    if ($_POST['action'] == 'in_needles_not_in_haystacks') {
        if (!empty($_POST['t1']) and !empty($_POST['t2'])) {
            $t1 = explode("\n", trim($_POST['t1']));
            $t2 = explode("\n", trim($_POST['t2']));
            if (is_array($t1) and !empty($t1)) {
                foreach ($t1 as $needle) {
                    if (!strstr($_POST['t2'], $needle)) {
                        $diff[] = $needle;
                    }
                }
            }
            $message[] = implode('<br />', $diff);
        } else {
            $message[] = 'Check input fields: date.';
        }
    }
    if ($_POST['action'] == 'url_decode') {
        if (!empty($_POST['t1'])) {
            $t1 = explode("\n", trim($_POST['t1']));
            if (is_array($t1)) {
                foreach ($t1 as $needle) {
                    $diff[] = urldecode($needle);
                }
            }
            $message[] = implode('<br />', $diff);
        } else {
            $message[] = 'Check input fields: date.';
        }
    }
    if ($_POST['action'] == 'sum') {
        if (!empty($_POST['t1'])) {
            $message[] = 'Needles: '.array_sum(explode("\n", str_replace(' ', '', $_POST['t1'])));
        }
        if (!empty($_POST['t2'])) {
            $message[] = 'Haystacks: '.array_sum(explode("\n", str_replace(' ', '', $_POST['t2'])));
        }
    }
    if ($_POST['action'] == '1_line') {
        if (!empty($_POST['t1'])) {
            $message[] = str_replace("\n", ',', str_replace(' ', '', $_POST['t1']));
        }
        if (!empty($_POST['t2'])) {
            $message[] = str_replace("\n", ',', str_replace(' ', '', $_POST['t2']));
        }
    }
    die(json_encode(join('<br />', $message)));
}
?>

<html>
    <head>
        <link rel="stylesheet" href="../js/jquery-ui/development-bundle/themes/base/jquery.ui.all.css">
        <script type="text/javascript" src="../js/jquery-ui/js/jquery-1.4.2.min.js"></script>
        <script type="text/javascript" src="../js/jquery-ui/js/jquery-ui-1.8.6.custom.min.js"></script>

        <script>
            $(function(){
                $('.datepicker').datepicker({'autoSize': true, 'dateFormat': 'yy-mm-dd'});
                $('input[name=btn]').click(function(){
                    $('input[name=action]').val($(this).val());
                    $('#result').html('Result:<br /><img src="/i/ajax-loader.gif" />');
                    $.ajax({type: 'post', dataType: 'json', data: $('body :input').serialize(),
                        success: function(response) {
                            $('#result').html('Result:<br />' + response);
                        },
                        error: function(xhr, status, error) {$('#result').html('Result:<br />' + error);}
                    });
                });
            });
        </script>
    </head>
    <body>
        <table cellspacing="7">
            <tr valign="top" align="center">
                <td>
                    Needles: <br />
                    <textarea name="t1" rows="15"></textarea>
                </td>
                <td>
                    Haystacks: <br />
                    <textarea name="t2" rows="15"></textarea>
                </td>
                <td>
                    Action:<br />
                    <input type="hidden" name="action" value="" />
                    <input type="button" name="btn" value="in_needles_not_in_haystacks" /><br />
                    <input type="button" name="btn" value="url_decode" /><br />
                    <input type="button" name="btn" value="sum" /><br />
                    <input type="button" name="btn" value="1_line" /><br />
                </td>
                <td id="result">Result:<br /></td>
            </tr>
        </table>
    </body>
</html>
