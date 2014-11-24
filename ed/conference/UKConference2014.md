PHP UK Conference 2014
-

Nate Abele - Weild AngularJS Like A Pro
Ole Michaelis - Service Oriented Architecture for Robust & Scalable Systems
Ricard Clau - Scaling with Symfony2
Rowan Merewood - Algorithm, Review, Sorting
Stuart Herbert - PHP at the Firehose Scale



####Morgan Tocker - Locking And Concurrency Control
* When write to two tables (create relation records) wrap it by transaction.

####Stephan Hochdorfer - The 7 Deadly Sins of Dependency Injection
* Inversion of control?
* Service locator pattern?
* Reflection is very slow, do not use it at production.

####Beth Tucker Long - Normalisation: Friend Or Foe
* Normalisation not need when we have a very big count of join tables.
* When data doesn't update.

####Glen Campbell - FAIL: The Best Ways To Bring Down Your Website
Developer should to know how to get down own website.
If something really very wrong with site or it hacked - change password to server to db and ssh...

####Marcello Duarte - Test, Transform, Refactor
* Ping pong model of programming?
* TBD?
* Code transformation!!!

####Lorna Mitchell - Debugging HTTP
* Curl!
* Wireshark!

####Glen Campbell - The Future of PHP Is In The Clouds
* *WE DON'T HAVE TO WRITE CODE ANYMORE, WE COULD REUSE CODE!!!*

####Beth Tucker Long - I've Been Hacked, Now What

####Derick Rethans - PHP In Space

####Davey Shafik - PHP Under The Hood

####Keynote - Juozas 'Joe' Kaziukenas - What Makes Technology Work
* Try to look at something from a different angle.
* API first! Because everything must be connected!
* Before real work create PROTOTYPE that will show what you really creating, and how it will be work, and show it as presentation.
* Focus on core part of application!!!
* If you not solve real problem nobody buy your idea/application.

####Andrei Zmievski - Intl Me This, Intl me That

####Bastian Hofmann - Profiling PHP Applications
* Logstash - is a tool for managing events and logs.
* Graphite - Scalable Realtime Graphing.
* xhprof!! can use at production!
* http://www.perfplanet.com/.

####Davey Shafik - PHP 5.NEXT: The New Bits

####Ben Mathews - Facebook's Approach To Common Web Vulnerabilities
* xss:
  * htmlspecialchars();
  * print at browser "Ben" and "B&#101;n" if it loocks the same - here are B&#101;n.
* sql injection: don't write sql.
* csrf:
  * token.
  * read controllers.
  * write controllers.
* cryptography.

####Eli White - Caching Best Practices
1. Whole page cache.
2. Partial page cache.
3. Database queries cache.
4. Biggest-Smallest reusable object.

* Varnish.
* Pre-generation.
* Beware FALSE (use NULL instead FALSE).
* Cache expiration mast be configurable (have own config file).

####Erika Heidi - Vagrant Provisioners In A Nutshell
* Ansible.
* Puppet.
* Chef.
* PuPHPet.

####Eli White - Web Security And You
* Return public images, or md files only through php (in such way apache can check permission to file).
* Don't just use MD5.
* Don't allow to upload an executable file.
* Don't allow use session id from url. The solution:
  * Use `session.use_cookies =1` or better: `session.use_only_cookies = 1`.
  * Use `session_regenerate_id` after user login, logout.
* FIEO - filter input escape output.
* Clickjacking (click on iframe). The solution:
  * `header('X-Frame Options: DENY');`
  * `header('X-Frame Options: SAMEORIGIN');`
* **Keep all your software up to date!**
* Man in the middle?

####Gary Hockin - Maximising Performance With Zend Framework
* **Siege** (An HTTP/HTTPS stress tester). `siedge -c 10 -t 1m -b http://mysite.dev`
* xhprof (can be installed in production).
* https://github.com/EvanDotPro/EdpSuperluminal.
* If site don't have favicon - 2 overhead request per every request.

####Ian Barber/Mandy Waite - Building Scalable PHP Apps With Google App Engine
* How many queries per second you have?
* Google Cloud!

####Jeremy Quinton - Gathering Metrics With StatsD And Graphite
* Measurement is the first step that leads to control and eventually to improvement. If you can't measure something, you can't understand it. If you can't understand it, you can't control it. If you can't control it, you can't improve it.
* StatsD (on pacagist).
* Graphite.

####Joshua Thijssen - (Re)Discovering The SPL (Standard PHP Library)
* lxr.php.net/xref/PHP_5_5/ext/spl/internal.
* Interface traversable connot be implemented.
* Use iterators!
* Never trhow exception, always catch exception!

####Joshua Thijssen - RPMing Your Apps And Tools

####Julien - PHP Opcache Explained
* How does PHP wokr:
  * Parsing
  * Compiling
  * Executing

####Mathias Verraes - Unbreakable Domain Models
* DDD.

####Michelle Sanver - 'ProTalking' Your Way into Open Source
* GitHub workflow:
  * Clone repe.
  * Read contribute file.
  * Read README.
  * Look test.
  * Communicate before your code!
  * Pick an issue.
  * Make changes, and push it.
  * Pull request.
  * Code review.
  * Correct or discuss.
  * Contribute documentation first!

####Morgan Tocker - My SQL 5 6 - Online Operations and Improved Diagnostics
* At 5.6 deleting index don't block read write.
* Explain UPDATE/DELETE statements.
* index usage statistics.