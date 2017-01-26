<?php

newrelic_set_appname('test');
newrelic_name_transaction('cli-test');
set_error_handler('newrelic_notice_error');

$e = new LogicException('case-1');
newrelic_notice_error('ERR-0', $e);

throw new BadMethodCallException('BadMethodCallException-1');

newrelic_custom_metric('myMetric', 1);
