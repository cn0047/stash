<?php

use Kahlan\plugin\Stub;
use Phalcon\Db\Adapter\Pdo\Mysql;
use Phalcon\Db\AdapterInterface;
use Phalcon\Db\Result\Pdo as PdoResult;
use Phalcon\Di;
use W3\Ziipr\Services\v1\AdminReport\KPI\Daily;

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
