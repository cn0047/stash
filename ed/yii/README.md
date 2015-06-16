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
````

#### Validation
````php
['siteId, countryCode, userId, userMsisdn', 'required', 'on' => 'insert'],
['status', 'numerical', 'min' => 0, 'max' => 255],
['type',   'match', 'pattern' => '/^(0|1|2)$/'],
['email',  'email'],
['date',   'date', 'format' => 'yyyy-MM-dd hh:mm:ss'],
['userId', 'type', 'type' => 'string'],
['userId', 'length', 'min' => 32, 'max' => 32],
['type',   'in', 'range' => array_keys(\Yii::app()->controller->getTypes())],
['userId', 'type', 'type' => 'string', 'except' => 'edit'],
````

#### DB
````php
php protected/yiic.php migrate create LastName_Desc
php protected/yiic.php migrate
php protected/yiic.php migrate redo 1
php protected/yiic.php migrate down 1

$where = [
    'date1' => $args['date1'],
    'like'  => $args['like'],
];
$data = $this->getDbConnection()->createCommand()
    ->select(['*'], 'SQL_CALC_FOUND_ROWS')
    ->from($this->tableName())
    ->where("`date` > :date1 `str` REGEXP :like", $where)
    ;
$count = $this->getDbConnection()->createCommand('SELECT FOUND_ROWS()')->queryScalar();
````
