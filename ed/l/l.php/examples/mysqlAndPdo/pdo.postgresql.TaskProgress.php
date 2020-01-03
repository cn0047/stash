<?php
/*


DROP TABLE task_progress;

CREATE TABLE task_progress (
  user_id INTEGER NOT NULL,
  action CHARACTER VARYING NOT NULL,
  total INTEGER NOT NULL DEFAULT 0,
  done INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expire INTEGER NOT NULL DEFAULT 0,
  PRIMARY KEY (user_id, action)
);

CREATE OR REPLACE FUNCTION task_progress_autodelete() RETURNS trigger AS $$
BEGIN
  IF NEW.total = NEW.done THEN
    DELETE from task_progress WHERE user_id = NEW.user_id AND action = NEW.action;
  END IF;
  RETURN NULL;
END
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER task_progress_autodelete
AFTER INSERT OR UPDATE ON task_progress FOR EACH ROW EXECUTE PROCEDURE task_progress_autodelete();


-- get
SELECT * FROM task_progress;

-- inc
INSERT INTO task_progress (user_id, action, total, expire) VALUES (1, 'test', 3, 10)
ON CONFLICT (user_id, action) DO UPDATE SET total = task_progress.total + 3, expire = task_progress.expire + 10;

-- dec
UPDATE task_progress SET done = task_progress.done + 1 WHERE user_id = 1 AND action = 'test';

-- expired
DELETE FROM task_progress
WHERE extract(epoch FROM (created_at + (expire * interval '1 seconds'))) < extract(epoch FROM NOW());



=======================================================================================================================

docker run -it --rm --net=xnet -v $PWD:/gh nphp \
    php /gh/ed/php/examples/mysqlAndPdo/pdo.postgresql.TaskProgress.php getState test

docker run -it --rm --net=xnet -v $PWD:/gh nphp \
    php /gh/ed/php/examples/mysqlAndPdo/pdo.postgresql.TaskProgress.php init test 5 10

docker run -it --rm --net=xnet -v $PWD:/gh nphp \
    php /gh/ed/php/examples/mysqlAndPdo/pdo.postgresql.TaskProgress.php tick test



docker run -it --rm --net=xnet -v $PWD:/gh nphp \
    php /gh/ed/php/examples/mysqlAndPdo/pdo.postgresql.TaskProgress.php getStateLoop test

docker run -it --rm --net=xnet -v $PWD:/gh nphp \
    php /gh/ed/php/examples/mysqlAndPdo/pdo.postgresql.TaskProgress.php tickLoop test

*/

class RecordNotFoundException extends Exception
{
}

class PostgreSQLStorage
{
    /**
     * @var PDO $dbh
     */
    private $dbh;

    public function __construct()
    {
        $this->dbh = new PDO('pgsql:host=xpostgres;port=5432;dbname=test;user=dbu;password=dbp');
        $this->expire();
    }

    private function expire(): void
    {
        $sql = "
          DELETE FROM task_progress
          WHERE extract(epoch FROM (created_at + (expire * interval '1 seconds'))) < extract(epoch FROM NOW())
        ";
        $s = $this->dbh->prepare($sql);
        if ($s->execute() === false) {
            var_dump($s->errorInfo());
        }
    }

    public function get(string $taskName)
    {
        $userId = 1;

        $s = $this->dbh->prepare('SELECT * FROM task_progress WHERE user_id = :uId AND action = :action');
        $s->bindParam(':uId', $userId, PDO::PARAM_INT);
        $s->bindParam(':action', $taskName);
        if ($s->execute() === false) {
            var_dump($s->errorInfo());
        }
        $result = $s->fetchAll(PDO::FETCH_ASSOC);
        if (empty($result)) {
            throw new RecordNotFoundException(404);
        }

        return $result;
    }

    public function increment(string $taskName, int $total, int $expire): void
    {
        $userId = 1;

        $sql = '
            INSERT INTO task_progress (user_id, action, total, expire)
            VALUES (:uId, :action, :total, :expire)
            ON CONFLICT (user_id, action) DO UPDATE
            SET total = task_progress.total + :upTotal, expire = task_progress.expire + :upExpire;
        ';
        $s = $this->dbh->prepare($sql);
        $s->bindValue(':uId', $userId, PDO::PARAM_INT);
        $s->bindValue(':action', $taskName);
        $s->bindValue(':total', $total, PDO::PARAM_INT);
        $s->bindValue(':expire', $expire, PDO::PARAM_INT);
        $s->bindValue(':upTotal', $total, PDO::PARAM_INT);
        $s->bindValue(':upExpire', $expire, PDO::PARAM_INT);
        if ($s->execute() === false) {
            var_dump($s->errorInfo());
        }
    }

    public function decrement(string $taskName, int $value): void
    {
        $userId = 1;

        $sql = '
            UPDATE task_progress SET done = task_progress.done + :done
            WHERE user_id = :uId AND action = :action
        ';
        $s = $this->dbh->prepare($sql);
        $s->bindValue(':done', $value, PDO::PARAM_INT);
        $s->bindValue(':uId', $userId, PDO::PARAM_INT);
        $s->bindValue(':action', $taskName);
        if ($s->execute() === false) {
            var_dump($s->errorInfo());
        }
    }
}

class TaskProgress
{
    /**
     * @var PostgreSQLStorage $storage
     */
    private $storage;

    public function __construct()
    {
        $this->storage = new PostgreSQLStorage();
    }

    public function init(string $taskName, int $total = 1, int $expire = 0): void
    {
        $this->storage->increment($taskName, $total, $expire);
    }

    public function getState(string $taskName)
    {
        try {
            return $this->storage->get($taskName);
        } catch (\Exception $e) {
            return -1;
        }
    }

    public function tick(string $taskName, int $value = 1): void
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
        while (true) {
            var_export((new TaskProgress())->getState($taskName));
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
