<?php

class IndexController extends CController
{
    public function actionIndex()
    {
        var_dump(YII_DEBUG);
        var_dump('admin');
        $connection=Yii::app()->db;
        $command=$connection->createCommand('select now()');
        $rows=$command->queryAll();
        var_export($rows);
    }

    public function actionRd()
    {
        throw new CHttpException(404);
    }
}
