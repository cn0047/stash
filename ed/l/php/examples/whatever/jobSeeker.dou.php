<?php

define('LOG_FILE', '/tmp/jobSeeker.dou.log');

$baseURL = 'https://jobs.dou.ua/vacancies/feeds/?';
$arr = [
    'city=Kyiv&category=Golang',
    'remote&category=Golang',
    // 'city=Kyiv&category=PHP',
    // 'remote&category=PHP',
    // 'city=Kyiv&category=Front+End',
    // 'remote&category=Front+End',
    // 'city=Kyiv&category=Node.js',
    // 'remote&category=Node.js',
    // 'city=Kyiv&category=Blockchain',
    // 'remote&category=Blockchain',
    'city=Kyiv&category=Other',
    'remote&category=Other',
];
foreach ($arr as $el) {
    run($baseURL, $el);
}

function run(string $baseURL, string $qs)
{
    $url = $baseURL . $qs;
    $xmlDoc = new DOMDocument();
    $xmlDoc->load($url);
    $items = $xmlDoc->getElementsByTagName('item');
    $res = find($items);
    printf("\n%s items: \n\n %s", colorize($qs), implode('', $res));
}

function colorize(string $str): string
{
    $res = "\033[36m$str\033[0m";

    return $res;
}

function find(object $items): array
{
    $res = [];
    foreach ($items as $i => $el) {
        $desc = $el->getElementsByTagName('description')->item(0)->childNodes->item(0)->nodeValue;
        $title = $el->getElementsByTagName('title')->item(0)->childNodes->item(0)->nodeValue;
        $link = $el->getElementsByTagName('link')->item(0)->childNodes->item(0)->nodeValue;
        $sign = notInteresting($title, $desc);
        $out = sprintf("\t%s \t%s \t %s\n", $link, $sign, $title);
        if (strlen($sign) === 0) {
            file_put_contents(LOG_FILE, deleteNoiseWords($out), FILE_APPEND);
            file_put_contents(LOG_FILE.'.desc', $desc, FILE_APPEND);
            continue;
        }
        $res[] = $out;
    }

    return $res;
}

function deleteNoiseWords(string $str): string
{
    $str = preg_replace('/golang/msi', '', $str);
    $str = preg_replace('/go/msi', '', $str);
    $str = preg_replace('/php/msi', '', $str);
    $str = preg_replace('/laravel/msi', '', $str);
    $str = preg_replace('/node/msi', '', $str);
    $str = preg_replace('/javascript/msi', '', $str);
    $str = preg_replace('/react/msi', '', $str);
    $str = preg_replace('/angular/msi', '', $str);
    $str = preg_replace('/vue/msi', '', $str);
    $str = preg_replace('/full.?stack/msi', '', $str);
    $str = preg_replace('/back.?end/msi', '', $str);
    $str = preg_replace('/front.?end/msi', '', $str);
    $str = preg_replace('/full.?time/msi', '', $str);
    $str = preg_replace('/develop(er|ment)/msi', '', $str);
    $str = preg_replace('/engineer/msi', '', $str);
    $str = preg_replace('/software/msi', '', $str);
    $str = preg_replace('/junior/msi', '', $str);
    $str = preg_replace('/middle/msi', '', $str);
    $str = preg_replace('/senior/msi', '', $str);

    return $str;
}

function notInteresting(string $title, string $desc): string
{
    $inBlackList = isInBlackList($title) === true || isInBlackList($desc) === true;
    $notCTO = isCTO($title) === false && isCTO($desc) === false;
    $notArchitect = isArchitect($title) === false && isArchitect($desc) === false;
    $notLead = isLead($title) === false && isLead($desc) === false;
    $notBackEnd = isBackEnd($title) === false;
    $notRust = isRust($title) === false && isRust($desc) === false;
    $noNumbers = withAnyNumber($title) === false && withAnyNumber($desc) === false;
    $noMoney = withMoney($title) === false && withMoney($desc) === false;

    $res = '';
    // $res .= $inBlackList === true ? '' : '';
    $res .= $notCTO === true ? '' : '✳️';
    $res .= $notArchitect === true ? '' : '👔';
    $res .= $notLead === true ? '' : '🙎‍♂️';
    $res .= $notBackEnd === true ? '' : '📺';
    $res .= $notRust === true ? '' : '🅡';
    $res .= $noMoney === true ? '' : '💵';
    $res .= $noNumbers === true ? '' : '🔢';

    return $res;
}

function withAnyNumber(string $str): bool
{
    $ok = preg_match('/\d{4,}/msi', $str);
    $no = preg_match('/\#\d{4,}|2019|jobs\.dou\.ua\/companies\/.*\/vacancies\/\d+/msi', $str);
    $res = $ok === 1 && $no === 0;

    return $res;
}

function withMoney(string $str): bool
{
    $ok = preg_match('/\$\d+–\d+/msi', $str);
    $res = $ok === 1;

    return $res;
}

function isInBlackList(string $str): bool
{
    $no = preg_match('/wordpress|magento/msi', $str);
    $res = $no === 0;

    return $res;
}

function isCTO(string $str): bool
{
    $ok = preg_match('/cto|chief technology|chief tech/msi', $str);
    $res = $ok === 1;

    return $res;
}

function isArchitect(string $str): bool
{
    $ok = preg_match('/architect/msi', $str);
    $no = preg_match('/architecture/msi', $str);
    $res = $ok === 1 && $no === 0;

    return $res;
}

function isLead(string $str): bool
{
    $ok = preg_match('/lead\b/msi', $str);
    $res = $ok === 1;

    return $res;
}

function isBackEnd(string $str): bool
{
    $ok = preg_match('/back.?end\b/msi', $str);
    $res = $ok === 1;

    return $res;
}

function isRemote(string $str): bool
{
    $ok = preg_match('/remote|удален|відд?ален/msi', $str);
    $no = preg_match(
        '/удаленную работу не рассматрива|удаленно не рассматриваем/msi',
        $str
    );
    $res = $ok === 1 && $no === 0;

    return $res;
}

function isRust(string $str): bool
{
    $ok = preg_match('/rust/msi', $str);
    $res = $ok === 1;

    return $res;
}
