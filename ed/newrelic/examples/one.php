<?php

newrelic_background_job(true);

newrelic_set_appname('test');
newrelic_name_transaction('cli-test');
set_error_handler('newrelic_notice_error');

$e = new LogicException('case-1');
newrelic_notice_error('ERR-0', $e);
$e2 = new BadMethodCallException('2ndException-BadMethodCallException');
newrelic_notice_error('Exception: ' . get_class($e), $e2);

throw new BadMethodCallException('BadMethodCallException-1');

newrelic_custom_metric('myMetric', 1);

newrelic_end_transaction();
