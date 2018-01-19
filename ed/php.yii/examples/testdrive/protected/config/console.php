<?php

// This is the configuration for yiic console application.
// Any writable CConsoleApplication properties can be configured here.
return array(
	'basePath'=>dirname(__FILE__).DIRECTORY_SEPARATOR.'..',
	'name'=>'My Console Application',

	// preloading 'log' component
	'preload'=>array('log'),

	// application components
	'components'=>array(
		'db'=>array(
			'connectionString' => 'sqlite:'.dirname(__FILE__).'/../data/testdrive.db',
		),
		'db'=>array(
			'connectionString' => 'mysql:host=xmysql;dbname=testdrive',
			'emulatePrepare' => true,
			'username' => 'user',
			'password' => 'pass',
			'charset' => 'utf8',
		),
		'dbUnitTest'=>array(
			'class' => 'system.db.CDbConnection',
			'connectionString' => 'mysql:host=xmysql;dbname=testdrive_unit_test',
			'emulatePrepare' => true,
			'username' => 'user',
			'password' => 'pass',
			'charset' => 'utf8',
		),
		'log'=>array(
			'class'=>'CLogRouter',
			'routes'=>array(
				array(
					'class'=>'CFileLogRoute',
					'levels'=>'error, warning',
				),
			),
		),
	),
);