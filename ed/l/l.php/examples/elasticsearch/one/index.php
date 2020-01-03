<!DOCTYPE html>
<html>
<head>
    <title>ES</title>
</head>
<body>
    <form method="post">
        Title:<input type="text" name="title">
        Body:<input type="text" name="body">
        Keywords:<input type="text" name="keywords">
        <button>submit</button>
    </form>
    <form>
        Search:
        <input type="text" name="q">
        <button>submit</button>
    </form>
</body>
</html>

<?php

require_once 'vendor/autoload.php';

$es = new Elasticsearch\Client(['hosts' => ['127.0.0.1:9200']]);

/*
curl -XPUT localhost:9200/articles/article/1 -d '
{
"title" : "This is a test article",
"body" : "This is the body of my article. I hope you like it."
}
'
curl -XPUT localhost:9200/articles/article/2 -d '
{
"title" : "This is a test article",
"body" : "This is the body of my article. I hope you like it."
}
'
curl -XPUT localhost:9200/articles/article/3 -d '
{
"title" : "This is a test article",
"body" : "This is the body of my article. I hope you like it.",
"keywords": ["test", "php"]
}
'
curl -XPUT localhost:9200/articles/article/4 -d '
{
"title" : "This is a another test article",
"body" : "This is the body of my another article. I hope you like it.",
"keywords": ["test", "es"]
}
'
*/

if (!empty($_POST)) {
    if (isset($_POST['title'])
        and isset($_POST['body'])
        and isset($_POST['keywords'])
    ) {
        $title = $_POST['title'];
        $body = $_POST['body'];
        $keywords = explode(',', $_POST['keywords']);
        $indexed = $es->index([
            'index' => 'articles',
            'type' => 'article',
            'body' => [
                'title' => $title,
                'body' => $body,
                'keywords' => $keywords,
            ],
        ]);
        if ($indexed) {
            echo '<pre>';
            var_export($indexed);
            echo '</pre>';
        }
    }
}
if (isset($_GET['q'])) {
    $q = $_GET['q'];
    $query = $es->search([
        'body' => [
            'query' => [
                'bool' => [
                    'should' => [
                        'match' => ['title' => $q],
                        'match' => ['body' => $q],
                    ],
                ],
            ],
        ],
    ]);
    if ($query['hits']['total'] >= 1) {
        $result = $query['hits']['hits'];
        echo 'Result:<hr><pre>';
        var_export($result);
        echo '</pre><hr>';
    }
}
