Mapping
-

````php
<?php
/**
 * @Entity
 * @Table(name="my_persistent_class")
 */
class MyPersistentClass {
    /**
     * @Id @Column(type="integer")
     * @GeneratedValue(strategy="AUTO")
     */
    private $id;
    /** @Column(type="integer") */
    private $intId;
    /** @Column(type="string") */
    private $name;
    /** @Column(name="db_name", type="string") */
    private $dbName;
    /** @Column(name="`number`", type="integer") */
    private $number;
}
````

Built-in mapping types:

* string: Type that maps an SQL VARCHAR to a PHP string.
* integer: Type that maps an SQL INT to a PHP integer.
* smallint: Type that maps a database SMALLINT to a PHP integer.
* bigint: Type that maps a database BIGINT to a PHP string.
* boolean: Type that maps an SQL boolean to a PHP boolean.
* decimal: Type that maps an SQL DECIMAL to a PHP double.
* date: Type that maps an SQL DATETIME to a PHP DateTime object.
* time: Type that maps an SQL TIME to a PHP DateTime object.
* datetime: Type that maps an SQL DATETIME/TIMESTAMP to a PHP DateTime object.
* text: Type that maps an SQL CLOB to a PHP string.

Column annotation attributes:

* type: The mapping type to use for the column.
* name: (optional, defaults to field name) The name of the column in the database.
* length: (optional, default 255) The length of the column in the database. (Applies only if a string-valued column is used).
* unique: (optional, default FALSE) Whether the column is a unique key.
* nullable: (optional, default FALSE) Whether the database column is nullable.
* precision: (optional, default 0) The precision for a decimal (exact numeric) column. (Applies only if a decimal column is used.)
* scale: (optional, default 0) The scale for a decimal (exact numeric) column. (Applies only if a decimal column is used.)

Custom Mapping Types:
````php
<?php
namespace My\Project\Types;
use Doctrine\DBAL\Types\Type;
use Doctrine\DBAL\Platforms\AbstractPlatform;

/**
 * My custom datatype.
 */
class MyType extends Type
{
    public function getSqlDeclaration(array $fieldDeclaration, AbstractPlatform $platform)
    {
    // return the SQL used to create your column type.
    // To create a portable column type, use the $platform.
    }

    public function convertToPHPValue($value, AbstractPlatform $platform)
    {
    // This is executed when the value is read from the database.
    // Make your conversions here, optionally using the $platform.
    }

    public function convertToDatabaseValue($value, AbstractPlatform $platform)
    {
    // This is executed when the value is written to the database.
    // Make your conversions here, optionally using the $platform.
    }
}
?>

<?php
use Doctrine\DBAL\Types\Type;
// Register my type
$config->setCustomTypes(array('mytype' => 'My\Project\Types\MyType'));
?>

<?php
class MyPersistentClass
{
    /** @Column(type="mytype") */
    private $field;
}
````

#### Inheritance Mapping

````php
<?php
/** @MappedSuperclass */
class MappedSuperclassBase
{
    /** @Column(type="integer") */
    private $mapped1;
    /** @Column(type="string") */
    private $mapped2;
    /**
     * @OneToOne(targetEntity="MappedSuperclassRelated1")
     * @JoinColumn(name="related1_id", referencedColumnName="id")
     */
    private $mappedRelated1;
    // ... more fields and methods
}

/** @Entity */
class EntitySubClass extends MappedSuperclassBase
{
    /** @Id @Column(type="integer") */
    private $id;
    /** @Column(type="string") */
    private $name;
    // ... more fields and methods
}
````

````sql
CREATE TABLE EntitySubClass (
    mapped1 INTEGER NOT NULL,
    mapped2 TEXT NOT NULL,
    id INTEGER NOT NULL,
    name TEXT NOT NULL,
    related1_id INTEGER DEFAULT NULL,
    PRIMARY KEY(id)
);
````

Single Table Inheritance:

````php
<?php
namespace MyProject\Model;
/**
 * @Entity
 * @SubClasses({"MyProject\Model\Employee"})
 * @InheritanceType("SINGLE_TABLE")
 * @DiscriminatorColumn(name="discr", type="string")
 * @DiscriminatorValue("person")
 */
class Person
{
}

/**
 * @Entity
 * @DiscriminatorValue("employee")
 */
class Employee extends Person
{
}
````

Class Table Inheritance:

````php
<?php
namespace MyProject\Model;
/**
 * @Entity
 * @SubClasses({"MyProject\Model\Employee"})
 * @InheritanceType("JOINED")
 * @DiscriminatorColumn(name="discr", type="string")
 * @DiscriminatorValue("person")
 */

class Person
{
}

/**
 * @Entity
 * @DiscriminatorValue("employee")
 */
class Employee extends Person
{
}
````
