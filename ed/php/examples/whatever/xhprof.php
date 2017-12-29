<?php

if (extension_loaded('xhprof')) {
    $XHPROF_ROOT = '/home/bond/web/kovpak/gh/helper/xhprof';
    include_once $XHPROF_ROOT . '/xhprof_lib/utils/xhprof_lib.php';
    include_once $XHPROF_ROOT . '/xhprof_lib/utils/xhprof_runs.php';
    // start profiling
    xhprof_enable(XHPROF_FLAGS_CPU + XHPROF_FLAGS_MEMORY);
}

echo 'Hello world!';

if (extension_loaded('xhprof')) {
    //end profiling
    $runId = (new \XHProfRuns_Default())->save_run(xhprof_disable(), '{name}');
}
