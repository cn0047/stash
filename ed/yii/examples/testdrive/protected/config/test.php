<?php

return CMap::mergeArray(
	require(dirname(__FILE__).'/main.php'),
	array(
		'components'=>array(
			'QRCode' => ['class' => 'application\tests\mock\components\QRCode'],
			'fixture'=>array(
				'class'=>'system.test.CDbFixtureManager',
			),
			'db'=>array(
				'connectionString' => 'mysql:host=localhost;dbname=testdrive_unit_test',
			),
		),
	)
);
