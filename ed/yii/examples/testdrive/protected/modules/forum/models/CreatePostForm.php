<?php

class CreatePostForm extends CFormModel
{
    public $username;
    public $post;

    public function rules()
    {
        return [
            ['username, post', 'required'],
            ['username', 'length', 'min' => 1],
            ['post', 'length', 'min' => 10],
        ];
    }
}
