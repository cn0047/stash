<?php
if (isset($_FILES['file']['error'])) {
    printf('<br>Error: %s', $_FILES['file']['error']);
}
if (isset($_FILES['file']['tmp_name'])) {
    printf('<br>File tmp name: %s', $_FILES['file']['tmp_name']);
}
if (isset($_FILES['file']['name'])) {
    printf('<br>File name: %s', $_FILES['file']['name']);
}
?>
<html>
<body>
    <form>
        <input id="fileToUpload" type="file" name="sortpic">
        <button id="upload" type="button">Upload</button>
    </form>
</body>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
<script>
    $("#upload").on("click", function() {
        var file_data = $('#fileToUpload').prop('files')[0];
        var form_data = new FormData();
        form_data.append('file', file_data);
        $.ajax({
            url: '/uploads',
            dataType: 'text',
            cache: false,
            contentType: false,
            processData: false,
            data: form_data,
            type: 'post',
            success: function (r) {
                $('body').append(r);
            },
            error: function (xhr, status, error) {
                console.error(xhr, status, error);
            },
         });
    });
</script>
</html>