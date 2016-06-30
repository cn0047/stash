<?php

$config_aws = (object)[
        'region' => 'eu-west-1',
        'credentials' => [
            'key'    => '',
            'secret' => '',
        ],
        'cognito' => [
            'identity_pool_id' => '/',
            'identity_name'    => '',
            'token_duration'   => ,
            'roles' => [
                'unauth' => '',
                'auth'   => ''
            ]
        ],
        's3' => (object)[
            'bucket' => ''
        ]
];
