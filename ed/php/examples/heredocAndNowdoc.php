<?php

class foo
{
    public $name = 'Bar';
}
$code[1] = 200;
$f = new foo;

/**
 * Nowdoc
 */
$html = <<<'NOWDOC'
<!DOCTYPE html>
<html>
<head>
    <title class="ttl">test</title>
</head>
<body>
    $code[1]
    $f->name
</body>
</html>
NOWDOC;
print $html;

/**
 * Heredoc
 */
$html = <<<"HEREDOC"
<!DOCTYPE html>
<html>
<head>
    <title class="ttl">test</title>
</head>
<body>
    $code[1]
    $f->name
</body>
</html>
HEREDOC;
print $html;
