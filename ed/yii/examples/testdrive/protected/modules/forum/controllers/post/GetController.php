<?php

class GetController extends CController
{
    public function actionIndex()
    {
        $m = new CreatePostForm;
        if (isset($_POST['CreatePostForm'])) {
            $m->attributes = $_POST['CreatePostForm'];
            if ($m->validate()) {
                var_dump($m);
            }
        }
        $this->render('get2', ['model' => $m]);
    }
}
