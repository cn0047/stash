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
