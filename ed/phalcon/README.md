Phalcon
-

2.0.9

````
php -S localhost:8000 -t ./public
````

````php
$di = \Phalcon\Di::getDefault();
$this->getDI()->getShared('config')->get('admin_module')->get('version')
````

#### Controller

````php
$this->request->getQuery(); // $_GET
$this->request->getPost(); // $_POST
$this->request->getQuery('_GET');
$this->request->getPost('_POST');
````
````
php app/cli.php main
````

#### Model
````php
Users::find([
    'created_at BETWEEN :from: AND :to:',
    'bind' => [
        'from' => $from,
        'to' => $to,
    ],
]);

// IN clause
$di = Di::getDefault();
/** @var QueryBuilder $qb */
$qb = $di->get('modelsManager')->createBuilder();
$qb->columns('v.*');
$qb->from(['v' => VideoModel::class]);
$qb->inWhere('v.userId', $userIds);
var_dump($qb->getQuery()->getSingleResult());

// Raw SQL
/** @var ResultSet $rs */
$rs = Di::getDefault()->get('dbSlave')->query('SELECT NOW()');
$rs->setFetchMode(\Phalcon\Db::FETCH_ASSOC);
return $rs->fetchAll($rs);
````

#### Devtools

````
sudo ln -s /var/www/html/public/phalcon-devtools/phalcon.php /usr/bin/phalcon
sudo chmod ugo+x /usr/bin/phalcon
````

````
# create project
phalcon project --name dbRelationships --type=cli

# generate model from db table
phalcon model languages --get-set --namespace=W3\\Ziipr\\Models\\v1

# create db migration
phalcon migration --action=generate --table=template --no-auto-increment

# run db migration
phalcon migration --action=run --table=template --version=1.0.0
````

#### Volt
````twig
{{ select('email_template_name', email_templates_names, 'class': 'hidden') }}

{% if not (users is empty) %}{% endif %}
````
