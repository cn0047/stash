<?php

namespace Cogo;

trait ParamsHandler
{
    private $params;

    public function __construct(array $params)
    {
        $this->setParams($params);
    }

    public function setParams(array $params)
    {
        $this->params = $params;
    }

    public function getParams()
    {
        return $this->params;
    }

    public function setParam($key, $value)
    {
        $this->params[$key] = $value;
    }

    public function getParam($key)
    {
        return (isset($this->params[$key]) ? $this->params[$key] : null);
    }
}

abstract class Database
{
    use ParamsHandler {
        setParams as private;
        setParam as protected;
    }

    protected function someDatabaseRelatedMethod()
    {
    }

    abstract protected function someOtherDatabaseRelatedMethod();
}

abstract class Storage
{
    use ParamsHandler;

    protected function someFilesystemRelatedMethod()
    {
    }

    abstract protected function someOtherFilesystemRelatedMethod();
}

class MySQL extends Database
{
    protected function someOtherDatabaseRelatedMethod()
    {
    }
}

class Filesystem extends Storage
{
    protected function someOtherFilesystemRelatedMethod()
    {
    }
}

$db = new MySQL(['host' => 'localhost']);
$fs = new Filesystem(['dir' => '/tmp/']);
var_export([
    $db->getParam('host'),
    $fs->getParam('dir'),
]);
