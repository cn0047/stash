<?php

namespace application\components;

class QRCode extends \CComponent
{
    private $scope;

    public function init()
    {
    }

    public function get()
    {
        return 200;
    }

    public function setScope($scope)
    {
        // $scope === 'global'
    }
}
