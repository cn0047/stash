<?php
/*

docker run -it --rm --net=xnet -v $PWD:/gh nphp \
    php /gh/ed/php/examples/whatever/memcache.TaskProgress.php getState test

docker run -it --rm --net=xnet -v $PWD:/gh nphp \
    php /gh/ed/php/examples/whatever/memcache.TaskProgress.php init test 5 10

docker run -it --rm --net=xnet -v $PWD:/gh nphp \
    php /gh/ed/php/examples/whatever/memcache.TaskProgress.php tick test



docker run -it --rm --net=xnet -v $PWD:/gh nphp \
    php /gh/ed/php/examples/whatever/memcache.TaskProgress.php getStateLoop test

docker run -it --rm --net=xnet -v $PWD:/gh nphp \
    php /gh/ed/php/examples/whatever/memcache.TaskProgress.php tickLoop test

*/

abstract class MemcacheStorage
{
    /**
     * @var bool $memcacheCompressed
     */
    protected $memcacheCompressed;

    /**
     * @var Memcache $storage
     */
    protected $storage;

    /**
     * @var bool $useLock
     */
    protected $useLock;

    public function __construct()
    {
        $this->storage = memcache_connect('xmemcached', 11211);
        $this->memcacheCompressed = false;
        $this->useLock = false;
    }

    protected function getKey(string $taskName): string
    {
        return "MySysPrefix:$taskName";
    }

    protected function lock(string $taskName): void
    {
        if ($this->useLock === false) {
            return;
        }

        $key = $this->getKey($taskName);
        while (!$this->storage->add("{$key}_lock", 1, false, 1000)) {
            var_dump('TRY LOCK '.microtime(1));
            usleep(10);
        }
    }

    protected function unlock(string $taskName): void
    {
        if ($this->useLock === false) {
            return;
        }

        $key = $this->getKey($taskName);
        while (!$this->storage->delete("{$key}_lock")) {
            var_dump('TRY UN-LOCK '.microtime(1));
            usleep(10);
        }
    }

    abstract protected function inc(string $taskName, int $value, int $expire): void;

    public function increment(string $taskName, int $value, int $expire): void
    {
        $this->lock($taskName);
        $this->inc($taskName, $value, $expire);
        $this->unlock($taskName);
    }

    abstract protected function dec(string $taskName, int $value): void;

    public function decrement(string $taskName, int $value): void
    {
        $this->lock($taskName);
        $this->dec($taskName, $value);
        $this->unlock($taskName);
    }
}

class SimpleMemcacheStorage extends MemcacheStorage
{
    public function get(string $taskName): int
    {
        return $this->storage->get($this->getKey($taskName));
    }

    /**
     * Won't refresh expiration.
     *
     * @param string $taskName
     * @param int $value
     * @param int $expire
     */
    protected function inc(string $taskName, int $value, int $expire): void
    {
        $key = $this->getKey($taskName);
        $r = $this->storage->increment($taskName, $value);
        if ($r === false) {
            $r = $this->storage->set($key, $value, $this->memcacheCompressed, $expire);
            if ($r === false) {
                var_dump('Exception during simpleIncrement.');
            }
        }
    }

    protected function dec(string $taskName, int $value): void
    {
        $key = $this->getKey($taskName);
        $r = $this->storage->decrement($key, $value);
        if ($r === 0) {
            var_dump("Delete $key.");
            $r = $this->storage->delete($key);
            if ($r === false) {
                var_dump('Decrement delete exception.');
            }
        }
        if ($r === false) {
            var_dump('Decrement exception.');
        }
    }
}

class SimpleRefreshMemcacheStorage extends SimpleMemcacheStorage
{
    /**
     * Will refresh expiration but not increment expiration.
     *
     * @param string $taskName
     * @param int $value
     * @param int $expire
     */
    protected function inc(string $taskName, int $value, int $expire): void
    {
        $key = $this->getKey($taskName);
        $oldValue = $this->get($taskName);
        if ($oldValue === false) {
            var_dump("Set $key.");
            $r = $this->storage->set($key, $value, $this->memcacheCompressed, $expire);
        } else {
            var_dump("Replace $key.");
            $newValue = $oldValue + $value;
            $r = $this->storage->replace($key, $newValue, $this->memcacheCompressed, $expire);
        }
        if ($r === false) {
            var_dump('Increment exception.');
        }
    }
}

class SophisticatedMemcacheStorage extends MemcacheStorage
{
    public function get(string $taskName): array
    {
        return $this->storage->get($this->getKey($taskName));
    }

    /**
     * @todo Finish it.
     * @param string $taskName
     * @param int $value
     * @param int $expire
     */
    protected function inc(string $taskName, int $value, int $expire): void
    {
        $key = $this->getKey($taskName);
        $data = $this->get($taskName);

        if ($data === false) {
            var_dump("Set $key.");
            $params = ['value' => $value, 'createdAt' => time()];
            $r = $this->storage->set($key, $params, $this->memcacheCompressed, $expire);
        } else {
            var_dump("Replace $key.");
//            $now = time();
//            $params = [
//                'value' => $data['value'] + $value,
//                'createdAt' => $now,
//            ];
//            $expireAt = $now - $data['createdAt'] + $expire;
//            $r = $this->storage->replace($key, $params, $this->memcacheCompressed, $expireAt);
        }
        if ($r === false) {
            var_dump('Increment exception.');
        }
    }

    /**
     * @todo Finish it.
     * @param string $taskName
     * @param int $value
     */
    protected function dec(string $taskName, int $value): void
    {
        $key = $this->getKey($taskName);
        $data = $this->get($taskName);

        if ($data === false) {
            var_dump('Decrement exception.');
            return;
        }

        if ($data['value'] === 0) {
            var_dump("Delete $key.");
            $r = $this->storage->delete($key);
            if ($r === false) {
                var_dump('Decrement delete exception.');
            }
        } else {
            var_dump("Decrement replace $key.");
//            $now = time();
//            $params = [
//                'value' => $data['value'] - $value,
//                'createdAt' => $now,
//            ];
//            $expireAt = $now - $data['createdAt'];
//            $this->storage->replace($key, $params, $this->memcacheCompressed, $expireAt);
        }
    }
}

class TaskProgress
{
    /**
     * @var MemcacheStorage $storage
     */
    private $storage;

    public function __construct()
    {
        $this->storage = new SimpleMemcacheStorage();
    }

    public function init(string $taskName, int $total = 1, int $expire = 0)
    {
        $this->storage->increment($taskName, $total, $expire);
    }

    public function getState(string $taskName)
    {
        return $this->storage->get($taskName);
    }

    public function tick(string $taskName, int $value = 1)
    {
        $this->storage->decrement($taskName, $value);
    }
}

class Command
{
    public function __construct(array $args)
    {
        $action = $args[1];
        $total = $args[3] ?? 1;
        $expire = $args[4] ?? 0;
        $taskName = $args[2];
        $this->$action($taskName, $total, $expire);
    }

    public function init(string $taskName, int $total, int $expire)
    {
        (new TaskProgress())->init($taskName, $total, $expire);
    }

    public function getState(string $taskName)
    {
        var_dump((new TaskProgress())->getState($taskName));
    }

    public function getStateLoop(string $taskName)
    {
        $tp = new TaskProgress();
        while (true) {
            var_dump($tp->getState($taskName));
            sleep(1);
        }
    }

    public function tick(string $taskName, int $value)
    {
        (new TaskProgress())->tick($taskName, $value);
    }

    public function tickLoop(string $taskName, int $value)
    {
        $tp = new TaskProgress();
        while (true) {
            printf("\r%s", microtime(true));
            $tp->tick($taskName, $value);
            sleep(1);
        }
    }
}

new Command($argv);
