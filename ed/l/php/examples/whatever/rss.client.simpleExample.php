<?php if (isset($_GET["q"]) and !empty($_GET["q"])): ?>
    <?php
    $q = $_GET["q"];
    if ($q=="Google") {
        $xml=("http://news.google.com/news?ned=us&topic=h&output=rss");
    } elseif($q=="NBC") {
        $xml=("http://rss.msnbc.msn.com/id/3032091/device/rss/rss.xml");
    }
    $xmlDoc = new DOMDocument();
    $xmlDoc->load($xml);
    //get elements from "<channel>"
    $channel=$xmlDoc->getElementsByTagName('channel')->item(0);
    $channel_title = $channel->getElementsByTagName('title')
        ->item(0)->childNodes->item(0)->nodeValue;
    $channel_link = $channel->getElementsByTagName('link')
        ->item(0)->childNodes->item(0)->nodeValue;
    $channel_desc = $channel->getElementsByTagName('description')
        ->item(0)->childNodes->item(0)->nodeValue;
    echo("<p><a href='" . $channel_link . "'>" . $channel_title . "</a>");
    echo("<br>");
    echo($channel_desc . "</p>");
    $x = $xmlDoc->getElementsByTagName('item');
    for ($i=0; $i<=2; $i++) {
        $item_title=$x->item($i)->getElementsByTagName('title')
            ->item(0)->childNodes->item(0)->nodeValue;
        $item_link=$x->item($i)->getElementsByTagName('link')
            ->item(0)->childNodes->item(0)->nodeValue;
        $item_desc=$x->item($i)->getElementsByTagName('description')
            ->item(0)->childNodes->item(0)->nodeValue;
        echo ("<p><a href='" . $item_link . "'>" . $item_title . "</a>");
        echo ("<br>");
        echo ($item_desc . "</p>");
    }
    ?>
<?php else: ?>
    <html>
    <head>
    <script>
        function showRSS (str) {
            if (str.length==0) {
                document.getElementById("rssOutput").innerHTML="";
                return;
            }
            if (window.XMLHttpRequest) {
                // code for IE7+, Firefox, Chrome, Opera, Safari
                xmlhttp=new XMLHttpRequest();
            } else {    // code for IE6, IE5
                xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
            }
            xmlhttp.onreadystatechange=function() {
                if (xmlhttp.readyState==4 && xmlhttp.status==200) {
                    document.getElementById("rssOutput").innerHTML=xmlhttp.responseText;
                }
            }
            xmlhttp.open("GET", "getrss.php?q="+str, true);
            xmlhttp.send();
        }
    </script>
    </head>
    <body>
        <form>
            <select onchange="showRSS(this.value)">
            <option value="">Select an RSS-feed:</option>
            <option value="Google">Google News</option>
            <option value="NBC">NBC News</option>
            </select>
        </form>
        <br>
        <div id="rssOutput">RSS-feed will be listed here...</div>
    </body>
    </html>
<?php endif; ?>
