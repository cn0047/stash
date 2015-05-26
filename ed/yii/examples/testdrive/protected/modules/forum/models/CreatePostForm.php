<?php

class CreatePostForm extends CFormModel
{
    public $username;
    public $type;
    public $post;
    public $tags;

    public function rules()
    {
        return [
            ['username, type, post, tags', 'required'],
            ['username', 'length', 'min' => 1],
            ['type', 'numerical'],
            ['post', 'length', 'min' => 2],
            ['tags', 'type', 'type' => 'array'],
        ];
    }
}
