<?php

$pdo = new PDO('mysql:dbname=test;host=127.0.0.1', 'root');
$sql = <<<"SQL"
DROP TABLE IF EXISTS types
;
CREATE TABLE `types` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `intT` tinyint(4) NOT NULL DEFAULT '0',
  `intS` smallint(6) NOT NULL DEFAULT '0',
  `intM` mediumint(9) NOT NULL DEFAULT '0',
  `intI` int(11) NOT NULL DEFAULT '0',
  `intB` bigint(20) NOT NULL DEFAULT '0',
  `floatF` float(3,1) NOT NULL DEFAULT '0.0',
  `doubleD` double(25,24) NOT NULL DEFAULT '0.000000000000000000000000',
  `bitB` bit(1) NOT NULL DEFAULT b'0',
  `dateD` date NOT NULL DEFAULT '0000-00-00',
  `dateDT` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `dateTS` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `dateT` time NOT NULL DEFAULT '00:00:00',
  `dateY` year(4) NOT NULL DEFAULT '0000',
  `dateY2` year(2) NOT NULL DEFAULT '00',
  `dateY4` year(4) NOT NULL DEFAULT '0000',
  `stringC` char(5) NOT NULL DEFAULT '',
  `stringVC` varchar(7) NOT NULL DEFAULT '',
  `stringB` binary(5) NOT NULL DEFAULT '\0\0\0\0\0',
  `stringVB` varbinary(7) NOT NULL DEFAULT '',
  `blobT` tinyblob NOT NULL,
  `blobB` blob NOT NULL,
  `blobM` mediumblob NOT NULL,
  `blobL` longblob NOT NULL,
  `textTT` tinytext NOT NULL,
  `textT` text NOT NULL,
  `textM` mediumtext NOT NULL,
  `textL` longtext NOT NULL,
  `e` enum('foo','boo') NOT NULL DEFAULT 'foo',
  `s` set('a','b','c') NOT NULL DEFAULT 'a',
  `geometryG` geometry NOT NULL DEFAULT '',
  `pointP` point NOT NULL DEFAULT '',
  `linestringL` linestring NOT NULL DEFAULT '',
  `polygonP` polygon NOT NULL DEFAULT '',
  `multipointM` multipoint NOT NULL DEFAULT '',
  `multilinestringM` multilinestring NOT NULL DEFAULT '',
  `multipolygonM` multipolygon NOT NULL DEFAULT '',
  `geometrycollectionG` geometrycollection NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
;
INSERT INTO types SET intT = 1
;
SQL;
$s = $pdo->prepare($sql);
if (!$s->execute()) {
    throw new \RuntimeException(var_export($s->errorInfo(), true));
}
$s = $pdo->prepare('SELECT * FROM types');
if (!$s->execute()) {
    throw new \RuntimeException(var_export($s->errorInfo(), true));
}
$r = $s->fetchAll(\PDO::FETCH_ASSOC);
var_dump($r);

/*
array(1) {
  [0] =>
  array(38) {
    'id'                  => string(1) "1"
    'intT'                => string(1) "1"
    'intS'                => string(1) "0"
    'intM'                => string(1) "0"
    'intI'                => string(1) "0"
    'intB'                => string(1) "0"
    'floatF'              => string(3) "0.0"
    'doubleD'             => string(26) "0.000000000000000000000000"
    'bitB'                => string(1) "\000"
    'dateD'               => string(10) "0000-00-00"
    'dateDT'              => string(19) "0000-00-00 00:00:00"
    'dateTS'              => string(19) "0000-00-00 00:00:00"
    'dateT'               => string(8) "00:00:00"
    'dateY'               => string(4) "0000"
    'dateY2'              => string(2) "00"
    'dateY4'              => string(4) "0000"
    'stringC'             => string(0) ""
    'stringVC'            => string(0) ""
    'stringB'             => string(5) "\000\000\000\000\000"
    'stringVB'            => string(0) ""
    'blobT'               => string(0) ""
    'blobB'               => string(0) ""
    'blobM'               => string(0) ""
    'blobL'               => string(0) ""
    'textTT'              => string(0) ""
    'textT'               => string(0) ""
    'textM'               => string(0) ""
    'textL'               => string(0) ""
    'e'                   => string(3) "foo"
    's'                   => string(1) "a"
    'geometryG'           => string(0) ""
    'pointP'              => string(0) ""
    'linestringL'         => string(0) ""
    'polygonP'            => string(0) ""
    'multipointM'         => string(0) ""
    'multilinestringM'    => string(0) ""
    'multipolygonM'       => string(0) ""
    'geometrycollectionG' => string(0) ""
  }
}
*/
