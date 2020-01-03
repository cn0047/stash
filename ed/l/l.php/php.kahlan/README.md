kahlan
-

4.0
<br>2.4

````sh
vendor/bin/kahlan --istanbul="coverage.json"
istanbul report
/usr/bin/chromium-browser coverage/lcov-report/index.html

vendor/bin/kahlan --coverage='One\Text::get()'
vendor/bin/kahlan --coverage='One\Text'
vendor/bin/kahlan --coverage='One'
````

````bash
vendor/bin/kahlan --coverage="One"

vendor/bin/kahlan --pattern=spec/suite/services/v1/CaptionsSpec.php  --coverage="W3\Zii\Services\v1\Captions"
vendor/bin/kahlan --coverage="W3\Zii\Services\v1" --istanbul="coverage.json"

vendor/bin/kahlan --coverage=4 --istanbul="coverage.json"
istanbul report
/usr/bin/chromium-browser coverage/lcov-report/index.html
````

````php
$rs = kahlan\plugin\Stub::create([
    'extends' => SimpleResultSet::class,
    'layer' => true,
    'params' => [[], new VideoModel(), null]
]);
Stub::on($rs)->method('toArray', function () {
    return [
        ['id' => 1],
    ];
});

Stub::on(SproutVideoModel::class)->method('::findFirst')->andReturn(false);
expect(function () {
    SproutVideo::getByVideoId(1);
})->toThrow(new \DomainException('SproutVideo not found.'));
````

````php
<?php

use Kahlan\plugin\Stub;
use Phalcon\Db\Adapter\Pdo\Mysql;
use Phalcon\Db\AdapterInterface;
use Phalcon\Db\Result\Pdo as PdoResult;
use Phalcon\Di;
use W3\Zii\Services\v1\AdminReport\KPI\Daily;

describe('KPI report daily v1', function () {

    it('get data', function () {
        /** @var AdapterInterface $dbSlave */
        $dbSlave = Stub::create(['extends' => Mysql::class, 'layer' => true, 'params' => [[]]]);
        /** @var \PDOStatement $st */
        $st = Stub::create(['extends' => \PDOStatement::class, 'layer' => true]);
        Stub::on($dbSlave)->method('query')->andReturn(new PdoResult($dbSlave, $st));
        Stub::on($st)->method('fetchAll')->andReturn(['ok']);
        Di::getDefault()->set('dbSlave', $dbSlave);

        $actualResult = (new Daily())->getData([
            'type' => 'DailyCountry',
            'dateFrom' => '2016-08-30',
            'dateTo' => '2016-08-31',
        ]);
        expect($actualResult)->toBeAn('array');
    });

});

````
