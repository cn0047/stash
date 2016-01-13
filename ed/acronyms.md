Acronyms
-

[ACID](#ACID)
[AMQP](#AMQP)
[B2B](#B2B)
[CDN](#CDN)
[CRM](#CRM)
[CSRF](#CSRF)
[DAO](#DAO)
[DBAL](#DBAL)
[DDD](#DDD)
[DDL](#DDL)
[DRY](#DRY)
[GMT](#GMT)
[GUI](#GUI)
[HTML](#HTML)
[HTTP](#HTTP)
[IMAP](#IMAP)
[IoC](#IoC)
[LDAP](#LDAP)
[MIME](#MIME)
[MTA](#MTA)
[ORM](#ORM)
[PEAR](#PEAR)
[PECL](#PECL)
[RBAC](#RBAC)
[RCS](#RCS)
[REST](#REST)
[RFC](#RFC)
[ROI](#ROI)
[SaaS](#SaaS)
[SFTP](#SFTP)
[SMTP](#SMTP)
[SOA](#SOA)
[SOAP](#SOAP)
[SOLID](#SOLID)
[SPL](#SPL)
[SSH](#SSH)
[SSL](#SSL)
[TCP](#TCP)
[TDD](#TDD)
[TLS](#TLS)
[UDP](#UDP)
[UTC](#UTC)
[WSDL](#WSDL)
[XML](#XML)

<dl>

<dt><a name="ACID"></a>ACID</dt>
<dd>
    Atomicity, Consistency, Isolation, Durability.
    <ul>
        <li><i>Atomicity</i> - all or nothing.</li>
        <li><i>Consistency</i> ensures that any transaction will bring the database from one valid state to another (constraints, cascades, triggers).</li>
        <li><i>Isolation</i> ensures that the concurrent execution of transactions will executed serially, i.e., one after the other.</li>
        <li><i>Durability</i> ensures that once a transaction has been committed, it will remain so, even in the event of power loss, crashes, or errors...</li>
    </ul>
</dd>

<dt><a name="AMQP"></a>AMQP</dt>
<dd>Advanced Message Queuing Protocol.</dd>

<dt><a name="B2B"></a>B2B</dt>
<dd>Business to business.</dd>

<dt><a name="CDN"></a>CDN</dt>
<dd>Content distribution network.</dd>

<dt><a name="CRM"></a>CRM</dt>
<dd>Customer relationship management.</dd>

<dt><a name="CSRF"></a>CSRF</dt>
<dd>Cross-site request forgery.</dd>

<dt><a name="DAO"></a>DAO</dt>
<dd>Data access object.</dd>

<dt><a name="DBAL"></a>DBAL</dt>
<dd>Database Abstraction Layer.</dd>

<dt><a name="DDD"></a>DDD</dt>
<dd>Domain-driven design.</dd>

<dt><a name="DDL"></a>DDL</dt>
<dd>Data description language.</dd>

<dt><a name="DRY"></a>DRY</dt>
<dd>Don't repeat yourself.</dd>

<dt><a name="GMT"></a>GMT</dt>
<dd>Greenwich Mean Time.</dd>

<dt><a name="GUI"></a>GUI</dt>
<dd>Graphical user interface.</dd>

<dt><a name="HTML"></a>HTML</dt>
<dd>HyperText Markup Language.</dd>

<dt><a name="HTTP"></a>HTTP</dt>
<dd>Hypertext Transfer Protocol.</dd>

<dt><a name="IMAP"></a>IMAP</dt>
<dd>Internet Message Access Protocol.</dd>

<dt><a name="IoC"></a>IoC</dt>
<dd>
    Inversion of control - is used to increase modularity of the program and make it extensible.
    Software frameworks, callbacks, schedulers, event loops and dependency injection
    are examples of design patterns that follow the inversion of control principle.
    <br>IoC serves the following design purposes:
    <ul>
        <li>To decouple the execution of a task from implementation.</li>
        <li>To focus a module on the task it is designed for.</li>
        <li>
            To free modules from assumptions about how other systems do
            what they do and instead rely on contracts.
        </li>
        <li>To prevent side effects when replacing a module.</li>
    </ul>
    ("Hollywood Principle: Don't call us, we'll call you".)
    <br>There are several basic techniques to implement inversion of control:
    <ul>
        <li>Using a factory pattern</li>
        <li>Using a service locator pattern</li>
        <li>
            Using a dependency injection, for example
            (constructor injection, parameter injection, setter injection, interface injection).
        </li>
        <li>Using a contextualized lookup</li>
        <li>Using template method design pattern</li>
        <li>Using strategy design pattern</li>
    </ul>
</dd>

<dt><a name="LDAP"></a>LDAP</dt>
<dd>Lightweight Directory Access Protocol.</dd>

<dt><a name="MIME"></a>MIME</dt>
<dd>Multipurpose Internet Mail Extensions.</dd>

<dt><a name="MTA"></a>MTA</dt>
<dd>Message Transfer Agent or Mail Transfer Agent.</dd>

<dt><a name="ORM"></a>ORM</dt>
<dd>
Object-relational mapping - is a programming technique
for converting data between incompatible type systems in object-oriented programming languages.
</dd>

<dt><a name="PEAR"></a>PEAR</dt>
<dd>PHP Extension and Application Repository.</dd>

<dt><a name="PECL"></a>PECL</dt>
<dd>Php extension community library.</dd>

<dt><a name="RBAC"></a>RBAC</dt>
<dd>Role-based access control.</dd>

<dt><a name="RCS"></a>RCS</dt>
<dd>Revision control system.</dd>

<dt><a name="REST"></a>REST</dt>
<dd>Representational State Transfer.</dd>

<dt><a name="RFC"></a>RFC</dt>
<dd>Request for Comments.</dd>

<dt><a name="ROI"></a>ROI</dt>
<dd>Return on investment.</dd>

<dt><a name="SaaS"></a>SaaS</dt>
<dd>Software as a service.</dd>

<dt><a name="SFTP"></a>SFTP</dt>
<dd>SSH File Transfer Protocol (also Secure File Transfer Protocol).</dd>

<dt><a name="SMTP"></a>SMTP</dt>
<dd>Simple Mail Transfer Protocol.</dd>

<dt><a name="SOA"></a>SOA</dt>
<dd>Service-oriented architecture (architectural pattern).</dd>

<dt><a name="SOAP"></a>SOAP</dt>
<dd>Simple Object Access protocol.</dd>

<dt><a name="SOLID"></a>SOLID</dt>
<dd>
    <ul>
        <li>
            (SRP) Single responsibility principle - states that every class should have responsibility over a single part of the functionality provided by the software.
        </li>
        <li>
            (OCP) Open/closed principle - (classes, modules, functions, etc.) should be open for extension, but closed for modification.
        </li>
        <li>
            (LSP) Liskov substitution principle - if S is a subtype of T, then objects of type T may be replaced with objects of type S without altering any of the desirable properties of the program.
        </li>
        <li>
            (ISP) Interface segregation principle - splits interfaces which are very large into smaller and more specific ones.
        </li>
        <li>
            (DIP) Dependency inversion principle - refers to a specific form of decoupling software modules.
            <ul>
                <li>High-level modules should not depend on low-level modules. Both should depend on abstractions.</li>
                <li>Abstractions should not depend on details. Details should depend on abstractions.</li>
                <li>Ownership inversion - both high- and lower-level layers should depend on abstractions that draw the behavior.</li>
                <li>
                    Abstraction dependency:
                    <ul>
                        <li>All member variables in a class must be interfaces or abstracts.</li>
                        <li>All concrete class packages must connect only through interface/abstract classes packages.</li>
                        <li>No class should derive from a concrete class.</li>
                        <li>No method should override an implemented method.</li>
                        <li>All variable instantiation requires the implementation of a Creational pattern as the Factory Method or the Factory pattern, or the more complex use of a Dependency Injection framework.</li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
</dd>

<dt><a name="SPL"></a>SPL</dt>
<dd>Standard PHP Library.</dd>

<dt><a name="SSH"></a>SSH</dt>
<dd>Secure Shell.</dd>

<dt><a name="SSL"></a>SSL</dt>
<dd>Secure Sockets Layer.</dd>

<dt><a name="TCP"></a>TCP</dt>
<dd>Transmission Control Protocol.</dd>

<dt><a name="TDD"></a>TDD</dt>
<dd>
    Test-driven development.
    <ul>
        <li><b>Acceptance</b>: Does the whole system work?</li>
        <li><b>Integration</b>: Does our code work against code we can't change?</li>
        <li><b>Unit</b>: Do our objects do the right thing, are they convenient to work with?</li>
    </ul>
</dd>

<dt><a name="TLS"></a>TLS</dt>
<dd>Transport Layer Security.</dd>

<dt><a name="UDP"></a>UDP</dt>
<dd>User Datagram Protocol.</dd>

<dt><a name="UTC"></a>UTC</dt>
<dd>Coordinated Universal Time</dd>

<dt><a name="WSDL"></a>WSDL</dt>
<dd>Web Service Definition Language.</dd>

<dt><a name="XML"></a>XML</dt>
<dd>Extensible Markup Language.</dd>

</dl>
