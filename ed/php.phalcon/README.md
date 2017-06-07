Phalcon
-

2.0.9

````
php -S localhost:8000 -t ./public

php -r 'var_dump(Phalcon\Version::get());'
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

sprintf('%s://%s',$this->request->getScheme(), $this->request->getServerName());
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

// Join
/** @var \Phalcon\Mvc\Model\Query\Builder $qb */
$qb = Di::getDefault()->get('modelsManager')->createBuilder();
$qb->columns('DISTINCT(u.country) AS country');
$qb->from(['ue' => UserEventModel::class]);
$qb->leftJoin(UserModel::class, 'ue.user_id = u.user_id', 'u');
$qb->betweenWhere('ue.created_at', $this->request['dateFrom'], $this->request['dateTo']);
$qb->orderBy('u.country');
return $qb->getQuery()->execute()->toArray();

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
phalcon model video --get-set --output=models/v1/ --namespace=W3\\Ziipr\\Models\\v1

````

#### DB migrations

````
# create db migration
phalcon migration --action=generate --table=video --no-auto-increment --config=config/config.php

# run db migration
phalcon migration --action=run --version=1.0.0
phalcon migration --action=run --version=1.0.6 --config=config/config.php
phalcon migration --action=run --table=template --version=1.0.0
````

# db-migration insert
````php
public function up()
{
    self::$_connection->insert(
        'taxonomy',
        [11111, 'BLABLBALA', 'blabalbal'],
        ['TaxonomyID', 'Code', 'Name', ]
    );
    $column = new Column(
        'UpdatedAt',
        [
            'type' => Column::TYPE_TIMESTAMP,
            'size' => 1,
        ]
    );
    self::$_connection->addColumn('captions_new', 'ziipr', $column);
}
public function up()
{
    // create table
    $this->morphTable('moderatorComment', array(/* ... */));
}
public function down()
{
    self::$_connection->delete(
        'taxonomy',
        'TaxonomyID IN (201600, 201601, 201602, 201603, 201604, 201605)'
    );
    self::$_connection->dropTable('user_auth');
    self::$_connection->dropColumn('captions_new', 'ziipr', 'UpdatedAt');
}
````

#### Volt

````php
Tag::setDefault('type', $type);
$this->view->setVar('type', $type);
````

````twig
{{ select('email_template_name', email_templates_names, 'class': 'hidden') }}
{{ check_field(item.getLanguageId(), 'checked': item.getEnabled() === '1' ? 'true' : null) }}

{% if not (users is empty) %}{% endif %}
````

# Pagination:
````twig
{% for item in paginator.items %}
{% endfor %}
{{ paginator.total_items }}
{{ paginator.current }}
{{ paginator.total_pages }}
{% if paginator.total_items > 0 %}
  {% set url = '/reporting/multiple-user-search?begin_date=' ~ begin_date ~ '&end_date=' ~ end_date %}

  <div class="row-fluid" style="text-align:center">
    <nav class="pagination pagination-custom">
      <ul>
        {% if paginator.current > 1 %}
          <li><a href="{{ url ~ '&page=' ~ (paginator.current - 1) }}"><span aria-hidden="true">&laquo;</span></a></li>
        {% else %}
          <li><span class="disabled">&lt;</span></li>
        {% endif %}

        {% for page in 1..paginator.total_pages %}
          {% if page == paginator.current %}
            <li class='active'><a href="#">{{ page }}</a></li>
          {% else %}
            <li><a href="{{ url ~ '&page=' ~ page }}">{{ page }}</a></li>
          {% endif %}
        {% endfor %}

        {% if paginator.current < paginator.total_pages %}
          <li><a href="{{ url ~ '&page=' ~ (paginator.current + 1) }}"><span aria-hidden="true">&raquo;</span></a></li>
        {% else %}
          <li><span class="disabled">&raquo;</span></li>
        {% endif %}
      </ul>
    </nav>
  </div>
{% endif %}
````
