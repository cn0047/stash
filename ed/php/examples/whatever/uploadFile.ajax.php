<?php

/**
 * @example php -S localhost:8000 ed/php/examples/uploadFile.ajax.php
 * @example curl http://localhost:8000 -H "Content-Type: multipart/form-data" -F "file=@/Users/k/f.txt" -F "msg=MyFile"
 */

?>

<?php if (isset($_FILES['file'])): ?>
    <?php
    if (isset($_FILES['file']['error'])) {
        printf('<br>Error: %s', $_FILES['file']['error']);
    }
    if (isset($_POST['msg'])) {
        printf('<br>Msq with file: %s', $_POST['msg']);
    }
    if (isset($_FILES['file']['tmp_name'])) {
        $tmpName = $_FILES['file']['tmp_name'];
        $name = $_FILES['file']['name'];
        $uploadsDir = '/tmp/php-uploads/';
        $destinationFile = "$uploadsDir/$name";
        mkdir($uploadsDir);
        move_uploaded_file($tmpName, $destinationFile);
        printf('<br>File tmp name: %s', $tmpName);
        printf('<br>File name: %s', $name);
        printf('<br>File content: %s', file_get_contents($destinationFile));
    }
    ?>
<?php else: ?>
    <html>
    <body>
        <form>
            <input id="msg" type="text" name="msg">
            <input id="fileToUpload" type="file" name="file">
            <button id="upload" type="button">Upload</button>
        </form>
    </body>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
    <script>
        $("#upload").on("click", function() {
            var file_data = $('#fileToUpload').prop('files')[0];
            var form_data = new FormData();
            form_data.append('msg', $('#msg').val());
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
<?php endif; ?>
