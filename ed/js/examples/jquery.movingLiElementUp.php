<html>
<body>
<a href="http://runnable.com/VSJ7CD9V2OQxpFQY/move-element-up-for-jquery-and-javascript">tweak it online</a>
<ul>
    <li>
        <div class="rows">
            <a id='position1' data-seq='id1'  class="noDownloadFormat" href="#" target="_blank">title1</a>
            <div style="float:right;margin-top:-15px;">
                <input type="button" class="btn" value="UP 1">
            </div>
        </div>
    </li>
    <li>
        <div class="rows">
            <a id='position2' data-seq='id2'  class="noDownloadFormat" href="#" target="_blank">title2</a>
            <div style="float:right;margin-top:-15px;">
                <input type="button" class="btn" value="UP 2">
            </div>
        </div>
    </li>
    <li>
        <div class="rows">
            <a id='position3' data-seq='id3'  class="noDownloadFormat" href="#" target="_blank">title3</a>
            <div style="float:right;margin-top:-15px;">
                <input type="button" class="btn" value="UP 3">
            </div>
        </div>
    </li>
</ul>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
<script>
    var handler = function (e) {
        $li = $(e.target).parent().parent().parent();
        // Avoid delete element if it's at top of list.
        if ($li.prev().size() !== 0) {
            // Move element.
            $li.prev().before($li.remove());
            // Bind click handler to new button.
            $('.btn', $($li)).click(handler);
        }
    };
    $('.btn').on('click', handler);
</script>
</body>
</html>
