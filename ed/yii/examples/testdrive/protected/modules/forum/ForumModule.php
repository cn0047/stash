<?php

class ForumModule extends CWebModule
{
    public function init()
    {
        parent::configure([
            'import' => [
                'application.modules.forum.models.*',
            ],
        ]);
    }
}
