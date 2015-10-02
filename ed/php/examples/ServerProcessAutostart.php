<?php

class Host
{
    private $ip = '';
    private $name = '';

    public function __construct($ip, $name)
    {
        $this->ip = $ip;
        $this->name = $name;
    }

    public function __get($name)
    {
        if (isset($this->$name)) {
            return $this->$name;
        }
        throw new DomainException("Unknown property $name");
    }
}

trait Notifier
{
    public function log($message)
    {
        $d = date('Y-m-d H:i:s');
        file_put_contents(
            '/tmp/ServerProcessAutostart.log',
            "[$d] $message\n",
            FILE_APPEND
        );
    }

    public function notify($message)
    {
        var_export($message);
    }

    public function alert($message)
    {
        var_export($message);
    }
}

class ServerProcessAutostart
{
    use Notifier;

    public function __construct($hostIp, $hostname)
    {
        $this->notify('in');
        try {
            if (!function_exists('ssh')) {
                throw new RuntimeException("Not exists necessary function 'ssh'.");
            }
            preg_match('/^(bulk|shard|search)/', $hostname, $matches);
            if (!isset($matches[0])) {
                throw new RuntimeException('Unknown host type.');
            }
            $className = ucfirst(strtolower($matches[0])).__CLASS__;
            if (!class_exists($className)) {
                throw new RuntimeException("Unknown class: $className");
            }
            new $className(new Host($hostIp, $hostname));
        } catch (UnexpectedValueException $e) {
            $m = $e->getMessage();
            $this->notify("exception[type:unexpected_value,message:$m]");
        } catch (RuntimeException $e) {
            $m = $e->getMessage();
            $this->notify("exception[type:runtime,message:$m]");
        } catch (Exception $e) {
            $m = $e->getMessage();
            $this->notify("exception[type:exception,message:$m]");
        }
        $this->notify('out');
    }
}

class BulkServerProcessAutostart
{
    public function __construct(Host $host)
    {
        new NginxAutostart($host);
        new SupervisorAutostart($host, 'b.conf');
    }
}

class ShardServerProcessAutostart
{
    public function __construct(Host $host)
    {
        new NginxAutostart($host);
        new SupervisorAutostart($host, 'insert.conf');
        new PhpFpmAutostart($host);
        new MysqlAutostart($host);
    }
}

class SearchServerProcessAutostart
{
    public function __construct(Host $host)
    {
        new NginxAutostart($host);
        new SupervisorAutostart($host, 'insert.conf');
        new PhpFpmAutostart($host);
        new MysqlRestartableAutostart($host);
        new SphinxAutostart($host);
    }
}

class NginxAutostart
{
    use Notifier;

    public function __construct(Host $host)
    {
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->reload($host);
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->serviceRestart($host);
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->sendAlert($host);
    }

    final private function checkIsOk(Host $host)
    {
        $output = ssh($host->ip, ['service nginx status']);
        if (empty($output) or !is_array($output)) {
            throw new UnexpectedValueException('Nginx status.');
        }
        if (count($output) === 1
            and isset($output[0])
            and $output[0] === 'nginx is running.'
        ) {
            return true;
        }
        if (isset($output[1])
            and isset($output[2])
            and preg_match('/Loaded: loaded/', $output[1])
            and preg_match('/Active: active \(running\)/', $output[2])
        ) {
            return true;
        }
        $this->log(__METHOD__.' '.var_export($output, true));
        return false;
    }

    final private function reload(Host $host)
    {
        $output = ssh($host->ip, ['/etc/init.d/nginx reload']);
        $this->notify("nginx[action:reload,host:{$host->name}]");
    }

    final private function serviceRestart(Host $host)
    {
        $output = ssh($host->ip, ['service nginx restart']);
        $this->notify("nginx[action:service_restart,host:{$host->name}]");
    }

    final private function sendAlert(Host $host)
    {
        $this->alert("Nginx on {$host->name} fell down.");
        $this->notify("nginx[action:send_alert,host:{$host->name}]");
    }
}

class SupervisorAutostart
{
    use Notifier;

    private $commands = [];

    public function __construct(Host $host, $configFile)
    {
        $this->init(
            $host,
            "/etc/supervisor/conf.d/$configFile"
        );
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->restart($host);
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->restart($host);
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->sendAlert($host);
    }

    final private function init(Host $host, $config)
    {
        $output = ssh($host->ip, ["cat $config"]);
        if (empty($output) or !is_array($output)) {
            throw new UnexpectedValueException('Supervisor config.');
        }
        $commands = [];
        array_walk($output, function ($v) use (&$commands) {
            preg_match('/command=\/usr\/bin\/php (\/var\/www\/.*\.php)/', $v, $m);
            if (isset($m[1])) {
                $commands[] = $m[1];
            }
        });
        if (empty($commands)) {
            throw new UnexpectedValueException('Supervisor commands.');
        }
        $this->commands = $commands;
    }

    final private function checkIsOk(Host $host)
    {
        foreach ($this->commands as $cmd) {
            $output = ssh(
                $host->ip,
                ["ps aux | grep $cmd | grep -v grep | wc -l"]
            );
            if (empty($output) or !is_array($output) or !isset($output[0])) {
                throw new UnexpectedValueException('Supervisor command status.');
            }
            $count = (int)$output[0];
            if ($count < 1) {
                return false;
            }
        }
        return true;
    }

    final private function restart(Host $host)
    {
        $output = ssh($host->ip, ['/etc/init.d/supervisor restart']);
        $this->notify("supervisor[action:restart,host:{$host->name}]");
    }

    final private function sendAlert(Host $host)
    {
        $this->alert("Supervisor on {$host->name} fell down.");
        $this->notify("supervisor[action:send_alert,host:{$host->name}]");
    }
}

class MysqlAutostart
{
    use Notifier;

    public function __construct(Host $host)
    {
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->sendAlert($host);
    }

    final protected function checkIsOk(Host $host)
    {
        $output = ssh($host->ip, ['service mysql status']);
        if (empty($output) or !is_array($output)) {
            throw new UnexpectedValueException('Mysql status.');
        }
        if (count($output) === 1
            and isset($output[0])
            and preg_match('/^MySQL running/', $output[0])
        ) {
            return true;
        }
        if (isset($output[0])
            and isset($output[10])
            and isset($output[13])
            and preg_match('/\/usr\/bin\/mysqladmin/', $output[0])
            and preg_match('/UNIX socket\s+\/var\/run\/mysqld\/mysqld\.sock/', $output[10])
            and preg_match('/Threads: \d+  Questions: \d+  Slow queries: \d+  Opens: \d+  Flush tables: \d+  Open tables: \d+  Queries per second avg:/', $output[13])
        ) {
            return true;
        }
        $this->log(__METHOD__.' '.var_export($output, true));
        return false;
    }

    final protected function sendAlert(Host $host)
    {
        $this->alert("Mysql on {$host->name} fell down.");
        $this->notify("mysql[action:send_alert,host:{$host->name}]");
    }
}

class MysqlRestartableAutostart extends MysqlAutostart
{
    public function __construct(Host $host)
    {
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->restart($host);
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->sendAlert($host);
    }

    final private function restart(Host $host)
    {
        $output = ssh($host->ip, ['service mysql restart']);
        $this->notify("mysql[action:restart,host:{$host->name}]");
    }
}

class PhpFpmAutostart
{
    use Notifier;

    public function __construct(Host $host)
    {
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->restart($host);
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->sendAlert($host);
    }

    final private function checkIsOk(Host $host)
    {
        $output = ssh($host->ip, ['service php5-fpm status']);
        if (empty($output) or !is_array($output)) {
            throw new UnexpectedValueException('Php5-fpm status.');
        }
        if (count($output) === 1
            and isset($output[0])
            and $output[0] === 'php5-fpm is running.'
        ) {
            return true;
        }
        $this->log(__METHOD__.' '.var_export($output, true));
        return false;
    }

    final private function restart(Host $host)
    {
        $output = ssh($host->ip, ['service php5-fpm restart']);
        $this->notify("php5_fpm[action:restart,host:{$host->name}]");
    }

    final private function sendAlert(Host $host)
    {
        $this->alert("Php5-fpm on {$host->name} fell down.");
        $this->notify("php5_fpm[action:send_alert,host:{$host->name}]");
    }
}

class SphinxAutostart
{
    use Notifier;

    public function __construct(Host $host)
    {
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->restart($host);
        if ($this->checkIsOk($host)) {
            return;
        }
        $this->sendAlert($host);
    }

    final private function checkIsOk(Host $host)
    {
        $output = ssh($host->ip, ['service sphinxsearch status']);
        if (empty($output) or !is_array($output)) {
            throw new UnexpectedValueException('Sphinx status.');
        }
        if (count($output) === 1
            and isset($output[0])
            and $output[0] === 'sphinxsearch is running.'
        ) {
            return true;
        }
        $this->log(__METHOD__.' '.var_export($output, true));
        return false;
    }

    final private function restart(Host $host)
    {
        $output = ssh($host->ip, ['service sphinxsearch restart']);
        $this->notify("sphinx[action:restart,host:{$host->name}]");
    }

    final private function sendAlert(Host $host)
    {
        $this->alert("Sphinx on {$host->name} fell down.");
        $this->notify("sphinx[action:send_alert,host:{$host->name}]");
    }
}
