yii
-
1

[api](http://www.yiiframework.com/doc/api/1.1/CApcCache)
|
[guide](http://yiiframework.ru/doc/guide/en/index)
|
[cookbook](http://yiiframework.ru/doc/cookbook/ru/index)

Running the Yii Application
````
cd ed/yii/examples/testdrive
php -S localhost:8000 index.php
````

````
\Yii::app()->request->getQuery('id'); // $_GET['id']
\Yii::app()->request->getPost('id'); // $_POST['id']

['siteId, countryCode, userId, userMsisdn', 'required', 'on' => 'insert'],
['type'  , 'numerical'],
['type',     'match', 'pattern' => '/^(0|1|2)$/'],
['name', 'match', 'pattern' => '/^[\w\s]{2,64}$/'],
['email', 'email'],
['date', 'date', 'format' => 'yyyy-MM-dd hh:mm:ss'],
['userId', 'type', 'type' => 'string'],
['userId', 'length', 'min' => 32, 'max' => 32],
['type', 'in', 'range' => array_keys(\Yii::app()->controller->getTypes())],

php protected/yiic.php migrate create LastName_Desc
php protected/yiic.php migrate
php protected/yiic.php migrate redo 1
php protected/yiic.php migrate down 1
````
