<?php

    $XHPROF_ROOT = "/home/bond/web/kovpak/gh/helper/xhprof";
    include_once $XHPROF_ROOT . "/xhprof_lib/utils/xhprof_lib.php";
    include_once $XHPROF_ROOT . "/xhprof_lib/utils/xhprof_runs.php";
    // start profiling
    xhprof_enable(XHPROF_FLAGS_CPU + XHPROF_FLAGS_MEMORY);

$result = [];
if (isset($_POST['action'])) {
    if (function_exists($_POST['action'])) {
        $result = $_POST['action']($_POST);
    } else {
        $result[] = 'Action not found.';
    }
} else {
    $result[] = "Don't passed parameter action.";
}
print json_encode($result);

function inNeedlesNotInHaystack($args)
{
    if (!isset($args['needles']) or empty($args['needles'])) {
        return ['Needles cannot be blank.'];
    }
    if (!isset($args['haystack']) or empty($args['haystack'])) {
        return ['Haystack cannot be blank.'];
    }
    $needles  = array_map('trim', explode("\n", $args['needles']));
    $haystack = array_map('trim', explode("\n", $args['haystack']));
    $result   = array_diff($needles, $haystack);
    return ['array' => array_values($result)];
}

function inNeedlesAndInHaystack($args)
{
    if (!isset($args['needles']) or empty($args['needles'])) {
        return ['Needles cannot be blank.'];
    }
    if (!isset($args['haystack']) or empty($args['haystack'])) {
        return ['Haystack cannot be blank.'];
    }
    $needles  = array_map('trim', explode("\n", $args['needles']));
    $haystack = array_map('trim', explode("\n", $args['haystack']));
    $result   = array_intersect($needles, $haystack);
    return ['array' => array_values($result)];
}

function urlEncod($args)
{
    $needles = [];
    if (!empty($args['needles'])) {
        $needles = explode("\n", $args['needles']);
    }
    $haystack = [];
    if (!empty($args['haystack'])) {
        $haystack = explode("\n", $args['haystack']);
    }
    $result = [];
    foreach ($needles + $haystack as $needle) {
        $result[] = urlencode($needle);
    }
    return ['array' => $result];
}

function urlDecod($args)
{
    $needles = [];
    if (!empty($args['needles'])) {
        $needles = explode("\n", $args['needles']);
    }
    $haystack = [];
    if (!empty($args['haystack'])) {
        $haystack = explode("\n", $args['haystack']);
    }
    $result = [];
    foreach ($needles + $haystack as $needle) {
        $result[] = urldecode($needle);
    }
    return ['array' => $result];
}

function regExpMatch($args)
{
    if (!isset($args['needles']) or empty($args['needles'])) {
        return ['Needles cannot be blank, it must contains reg.exp. pattern.'];
    }
    if (!isset($args['haystack']) or empty($args['haystack'])) {
        return ['Haystack cannot be blank, it must contains reg.exp. subject.'];
    }
    preg_match('/'.$args['needles'].'/', $args['haystack'], $matches);
    return ['array' => $matches, 'text' => var_export($matches, true)];
}

    //end profiling 
    $xhprof_data = xhprof_disable();
    $xhprof_runs = new XHProfRuns_Default();
    $run_id = $xhprof_runs->save_run($xhprof_data, 'helper');
