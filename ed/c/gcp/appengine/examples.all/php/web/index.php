<?php

function fibonacci($n) {
    if ($n < 2) {
        return $n;
    }

    return fibonacci($n-1) + fibonacci($n-2);
}

switch ($_SERVER['PATH_INFO']) {
    case '/fib':
        $params = [];
        parse_str($_SERVER['QUERY_STRING'], $params);
        $result = fibonacci($params['n']);
        printf('fibonacci %d = %d', $params['n'], $result);
        break;

    case '/x':
        break;

    default:
        echo 'php - ok';
        break;
}
