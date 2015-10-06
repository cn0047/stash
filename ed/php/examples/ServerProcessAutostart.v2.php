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
    final private function logToFile($message)
    {
        $d = date('Y-m-d H:i:s');
        file_put_contents(
            '/tmp/ServerProcessAutostart.log',
            "[$d] $message\n",
            FILE_APPEND
        );
    }

    final private function logToWeb($message)
    {
        exec('curl -silent '.escapeshellarg("https://log.onthe.io/?app=71qnqvd1qeb&msg=".urlencode($message)).' > /dev/null &');
    }

    public function log($message)
    {
        $this->logToWeb($message);
    }

    public function notify($message)
    {
        $m = urlencode("server_process_autostart_v2.$message");
        exec('curl -silent '.escapeshellarg("https://tapi.onthe.io/?k=9818:$m&s=e9223ad24767f9583d39515511cbc0d5").' > /dev/null &');
        // $result = urlencode('11:signup_additional['.implode(',', $paramsPerson).']');
        // exec('curl -silent '.escapeshellarg("https://tapi.onthe.io/?k=$result&s=b0fcd37bdfa230974b4f35870e8617ac").' > /dev/null &');
    }

    public function alert($message)
    {
        $this->notify("alert[message:$message}]");
        $this->log($message);
    }
}

class ServerProcessAutostart
{
    use Notifier;

    public function __construct($hostIp, $hostname)
    {
        $time = microtime(true);
        $this->notify('in');
        try {
            if (!function_exists('ssh')) {
                throw new RuntimeException("Not exists necessary function 'ssh'.");
            }
            preg_match('/^(bulk|shard|search)/', $hostname, $matches);
            if (!isset($matches[0])) {
                throw new RuntimeException("Unknown host type: $hostname.");
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
        $spentTime = round(microtime(true) - $time, 1);
        $this->notify("time[time:$spentTime]");
        $this->notify('out');
    }
}

trait Transport
{
    final private function exec($ip, array $commands)
    {
        $ssh = 'ssh -o StrictHostKeyChecking=no -o ConnectTimeout=60 -q -i /root/io.rsa root@onthe.io ssh -o StrictHostKeyChecking=no -o ConnectTimeout=60 -q -i /var/www/keys/io.rsa root@' . $ip . ' \'' . implode(';', $commands) . '\'';
        exec($ssh, $output);
        if (empty($output) or !is_array($output)) {
            throw new UnexpectedValueException('Unexpected ssh response.');
        }
        return $output;
    }
}

class Checker
{
    use Notifier;
    use Transport;

    public function __construct(Host $host, array $nodes)
    {
        $level = 0;
        do {
            $nextLevel = false;
            $action = 'action'.$level;
            // Filter by available actions in node.
            $commands = [];
            $filtredNodes = [];
            foreach ($nodes as $entity) {
                if (method_exists($entity, $action)) {
                    $filtredNodes[] = $entity;
                    $commands[] = $entity->$action();
                } else {
                    $parentNodeName = get_parent_class($entity);
                    $nodeName = $parentNodeName ? $parentNodeName : get_class($entity);
                    $this->alert("$nodeName on {$host->name} fell down.");
                    $this->notify(
                        strtolower($nodeName)."[action:send_alert,host:{$host->name},level:$level]"
                    );
                }
            }
            $nodes = $filtredNodes;
            if (empty($commands) and empty($nodes)) {
                return;
            }
            // Temporary code.
            // When we decide to apply any action upon server process
            // we make screenshot of process status for log reasons.
            if ($level > 0) {
                $screenShotCommands = [];
                foreach ($nodes as $entity) {
                    $screenShotCommands[] = $entity->getScreenShotCommand();
                    //
                    $parentNodeName = get_parent_class($entity);
                    $nodeName = $parentNodeName ? $parentNodeName : get_class($entity);
                    $this->notify(
                        strtolower($nodeName)."[action:$action,host:{$host->name},level:$level]"
                    );
                }
                $output = $this->exec($host->ip, $screenShotCommands);
                $this->log(
                    __METHOD__.' Screen shot '.var_export([$screenShotCommands, $output], true)
                );
            }
            $output = $this->exec($host->ip, $commands);
            // Filter by response.
            $filtredNodes = [];
            foreach ($output as $i => $response) {
                if ($response === '0') {
                    $nextLevel = true;
                    $filtredNodes[] = $nodes[$i];
                }
            }
            $nodes = $filtredNodes;
            $level++;
        } while ($nextLevel);
    }
}

class BulkServerProcessAutostart
{
    public function __construct(Host $host)
    {
        // $nodes = [
        //     new Nginx(),
        //     new Supervisor('bulk.conf'),
        // ];
        // new Checker($host, $nodes);
        new Checker($host, [new Nginx()]);
        // new Checker($host, [new Supervisor('insert.conf')]);
    }
}

class ShardServerProcessAutostart
{
    public function __construct(Host $host)
    {
        // $nodes = [
        //     new Mysql(),
        //     new Nginx(),
        //     new PhpFpm(),
        //     new Supervisor('insert.conf'),
        // ];
        // new Checker($host, $nodes);
        new Checker($host, [new Nginx()]);
        // new Checker($host, [new Supervisor('insert.conf')]);
        new Checker($host, [new PhpFpm()]);
        new Checker($host, [new Mysql()]);
    }
}

class SearchServerProcessAutostart
{
    public function __construct(Host $host)
    {
        // $nodes = [
        //     new MysqlRestartable(),
        //     new Sphinx(),
        //     new Nginx(),
        //     new PhpFpm(),
        //     new Supervisor('insert.conf'),
        // ];
        // new Checker($host, $nodes);
        new Checker($host, [new Nginx()]);
        // new Checker($host, [new Supervisor('insert.conf')]);
        new Checker($host, [new PhpFpm()]);
        new Checker($host, [new Mysql()]);
        // new Checker($host, [new MysqlRestartable()]);
        // new Checker($host, [new Sphinx()]);
    }
}

class Nginx
{
    public function getScreenShotCommand()
    {
        return 'service nginx status ';
    }

    public function getCheckIsOkCommand()
    {
        return 'service nginx status | grep "(nginx is running.|Loaded: loaded.*\n.*Active: active \(running\))" -czP';
    }

    public function action0()
    {
        return $this->getCheckIsOkCommand();
    }

    public function action1()
    {
        return $this->action2();
        $action = '/etc/init.d/nginx reload > /dev/null && ';
        return $action.$this->getCheckIsOkCommand();
    }

    public function action2()
    {
        $action = 'service nginx restart > /dev/null && ';
        return $action.$this->getCheckIsOkCommand();
    }
}

class Supervisor
{
    private $configFile = '';

    public function __construct($configFile)
    {
        $this->configFile = $configFile;
    }

    public function getScreenShotCommand()
    {
        return '
            for i in $(cat /etc/supervisor/conf.d/'.$this->configFile.' | grep "command=\/usr\/bin\/php \/var\/www\/.*\.php" -oE | grep "\/var\/www\/.*\.php" -oE);
            do
                result=$((`ps aux | grep -v grep | grep $i -c`))
                echo "$i $result";
            done
        ';
    }

    public function getCheckIsOkCommand()
    {
        return '
            fail=0;
            for i in $(cat /etc/supervisor/conf.d/'.$this->configFile.' | grep "command=\/usr\/bin\/php \/var\/www\/.*\.php" -oE | grep "\/var\/www\/.*\.php" -oE);
            do
                result=$((`ps aux | grep -v grep | grep $i -c`))
                if [ $result -gt 0 ]
                then
                    fail=$((fail+0));
                else
                    fail=$((fail+1));
                fi
            done
            if [ $fail -gt 0 ];
            then
                echo 0;
            else
                echo 1;
            fi
        ';
    }

    public function action0()
    {
        return $this->getCheckIsOkCommand();
    }

    public function action1()
    {
        $action = '/etc/init.d/supervisor restart > /dev/null && ';
        return $action.$this->getCheckIsOkCommand();
    }

    public function action2()
    {
        return $this->action1();
    }
}

class PhpFpm
{
    public function getScreenShotCommand()
    {
        return 'service php5-fpm status';
    }

    public function getCheckIsOkCommand()
    {
        return 'service php5-fpm status | grep "(php5-fpm is running.|Loaded: loaded.*\n.*Active: active \(running\))" -czP';
    }

    public function action0()
    {
        return $this->getCheckIsOkCommand();
    }

    public function action1()
    {
        $action = 'service php5-fpm restart > /dev/null && ';
        return $action.$this->getCheckIsOkCommand();
    }
}

class Mysql
{
    public function getScreenShotCommand()
    {
        return 'service mysql status';
    }

    public function getCheckIsOkCommand()
    {
        return 'service mysql status | grep "(MySQL running|Loaded: loaded.*\n.*Active: active \(running\)|Threads: \d+  Questions: \d+  Slow queries: \d+  Opens: \d+  Flush tables: \d+  Open tables: \d+  Queries per second avg:.*)" -czP';
        return 'service mysql status | grep "(MySQL running|Loaded: loaded.*\n.*Active: active \(running\)|\/usr\/bin\/mysqladmin/.*\n.*UNIX socket\s+\/var\/run\/mysqld\/mysqld\.sock/.*\n.*Threads: \d+  Questions: \d+  Slow queries: \d+  Opens: \d+  Flush tables: \d+  Open tables: \d+  Queries per second avg:)" -czP';
    }

    public function action0()
    {
        return $this->getCheckIsOkCommand();
    }
}

class MysqlRestartable
{
    public function action1()
    {
        $action = 'service mysql restart > /dev/null && ';
        return $action.$this->getCheckIsOkCommand();
    }
}

class Sphinx
{
    public function getScreenShotCommand()
    {
        return 'service sphinxsearch status';
    }

    public function getCheckIsOkCommand()
    {
        return 'service sphinxsearch status | grep "(sphinxsearch is running.|Loaded: loaded.*\n.*Active: active \(running\))" -czP';
    }

    public function action0()
    {
        return $this->getCheckIsOkCommand();
    }

    public function action1()
    {
        $action = 'service sphinxsearch restart > /dev/null && ';
        return $action.$this->getCheckIsOkCommand();
    }
}
