<?php

$xmlDoc = new DOMDocument();
$xmlDoc->load('http://jobs.dou.ua/vacancies/feeds/?cities=%D0%9A%D0%B8%D0%B5%D0%B2&category=PHP');
$items = $xmlDoc->getElementsByTagName('item');
$foundUrls = [];
foreach ($items as $i => $el) {
    echo "item: $i\r";
    $desc = $el
        ->getElementsByTagName('description')->item(0)
        ->childNodes->item(0)
        ->nodeValue
    ;
    $isInteresting = preg_match('/remote|удален|відд?ален/ms', $desc);
    if ($isInteresting) {
        $foundUrls[] = $el
            ->getElementsByTagName('link')->item(0)
            ->childNodes->item(0)
            ->nodeValue
        ;
    }
}
if (count($foundUrls) > 0) {
    $urls = implode(' ', $foundUrls);
    `/usr/bin/chromium-browser $urls`;
}
