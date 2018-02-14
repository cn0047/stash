yii
-
1.1.19
<br>1.1.15

[api](http://www.yiiframework.com/doc/api/1.1)
|
[guide](http://www.yiiframework.com/doc/guide/1.1/en/index)
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
\Yii::app()->request->getParam('id'); // $_GET['id'] or $_POST['id']

$args = json_decode(\Yii::app()->request->getRawBody(), true);

Yii::app()->getCache();

Yii::app()->getUser()->getId(); // from $_SESSION
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

Migrations:

````sh
php protected/yiic.php migrate create LastName_Desc
php protected/yiic.php migrate
php protected/yiic.php migrate redo 1
php protected/yiic.php migrate down 1
````

````php
TaskProgress::model()->updateByPk($taskName, ['done' => 'done + 3']);
TaskProgress::model()->updateAll(['done' => 'done + 3'], "action = '$taskName'");

# toArray()
$model->attributes

$where = [
    'date1' => $args['date1'],
    'like'  => $args['like'],
];
$data = $this->getDbConnection()->createCommand()
    ->select(['*'], 'SQL_CALC_FOUND_ROWS')
    ->from($this->tableName())
    ->join('userStat us', 'us.userId = uan.userId')
    ->where("`date` > :date1 `str` REGEXP :like", $where)
    ->andWhere(['in', 'id', [1, 5, 7]])
    ->andWhere('last_numbers = :ccl', compact('ccl'))
    ->group('user_id, group_id')
    ->having('rows > 7')
    ->order('q.id, q.type')
    ;
$count = $this->getDbConnection()->createCommand('SELECT FOUND_ROWS()')->queryScalar();

\Yii::app()->db->createCommand($sql)->execute();
\Yii::app()->db->createCommand($sql)->queryColumn($params); // queryAll|queryRow

$command = \Yii::app()->db->createCommand($sql);
$command->bindParam(':uId', \Yii::app()->getUser()->getId(), \PDO::PARAM_INT);
$command->bindParam(':action', $taskName);
$data = $command->queryAll();

$table = \Yii::app()->ext->getDbConnection()->schema->getTable('logCallMeBack');
````

Criteria:

````php
$sql = "
  DELETE FROM task_progress
  WHERE
      done >= total
      OR (
        expire <> 0
        AND EXTRACT(epoch FROM (created_at + (expire * interval '1 seconds'))) < EXTRACT(epoch FROM NOW())
      )
";
$subCriteria = new \CDbCriteria;
$subCriteria->addCondition('expire <> 0');
$subCriteria->addCondition("
    EXTRACT(epoch FROM (created_at + (expire * interval '1 seconds'))) < EXTRACT(epoch FROM NOW())
");
$criteria = new \CDbCriteria;
$criteria->addCondition('done >= total');
$criteria->addCondition($subCriteria->condition, 'OR');
TaskProgress::model()->deleteAll($criteria);
````

````
phpunit --colors --bootstrap=bootstrap.php unit/

cd ed/yii/examples/testdrive/protected/tests/
../../../../../../vendor/bin/phpunit  --colors --bootstrap=bootstrap.php unit/
````

#### Integration test frameworks

`phpspec` - won't work with static methods.
`codeception` - don't have monkey patches.
`kahlan` - won't `stub` (with layer) classes loaded not from composer. And won't work with yii.
