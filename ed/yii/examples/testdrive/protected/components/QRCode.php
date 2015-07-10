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
        $code = 200;
        $event = new \CModelEvent(['code' => $code]);
        $this->onGet($event);
        return $code;
    }

    public function onGet($event)
    {
        $this->raiseEvent('onGet', $event);
    }

    public function setScope($scope)
    {
        // $scope === 'global'
    }
}
