Theory
-

````
┌─┬─┐
├─┼─┤
│ │ │   ╔═══╗
│ │ │   ║ + ║
│ │ │   ╚═══╝
└─┴─┘
````

<br>ACL     - Access Control Lists.
<br>API     - Application programming interface.
<br>B2B     - Business to business.
<br>BAAS    - Back-end as a service.
<br>BLOB    - Binary Large Object.
<br>CRM     - Customer relationship management.
<br>DRY     - Don't repeat yourself.
<br>GMT     - Greenwich Mean Time.
<br>GTIN    - Global Trade Item Number.
<br>GUI     - Graphical user interface.
<br>IDL     - Interface Definition Language.
<br>IPC     - Inter-process communication (processes manage shared data).
<br>JIT     - Just-in-time compilation.
<br>KISS    - Keep it simple, stupid (noted in 1960).
<br>MPN     - Manufacturer Part Number.
<br>NDA     - Non-disclosure agreement.
<br>ODM     - Object-document mapping.
<br>RAID    - Redundant Array of Inexpensive Disks.
<br>RBAC    - Role-based access control.
<br>REPL    - Read–Eval–Print Loop (interactive toplevel or language shell).
<br>RFC     - Request for Comments.
<br>RMI     - Remote Method Invocation.
<br>RnD     - Research and development.
<br>ROI     - Return on investment.
<br>RSS     - Rich Site Summary or Really Simple Syndication.
<br>RTB     - Real-time bidding.
<br>RTC     - Real-time communication.
<br>SaaS    - [Software as a service](https://raw.githubusercontent.com/cn007b/my/master/ed/t.theory/i/cloudComputingTypes.jpeg).
<br>SDK     - Software development kit.
<br>SDLC    - Software development life-cycle.
<br>UPI     - Unique Product Identifier.
<br>UTC     - Coordinated Universal Time.
<br>VOIP    - Voice over Internet Protocol.
<br>WORA    - Write once, run anywhere.
<br>WSDL    - Web Service Definition Language.
<br>WYSIWYG - What you see is what you get.
<br>XML     - Extensible Markup Language.
<br>YAGNI   - You ain't gonna need it.

Crawling - programmatically get web page, extract hyperlinks and follow them.
Scraping - programmatically get web page, extract very specific data.

Consensus - getting all of the nodes to agree on something.

Systems of record - aka source of truth.

Batch processing systems - offline systems.
Stream processing systems (near-real-time systems) -
a stream processor consumes inputs and produces outputs (rather than responding to requests).

Server-Sent Events - servers can initiate data transmission towards clients
once an initial client connection has been established.

IO throughput - data transfer speed in megabytes per second (MB/s or MBPS).

To deal with big complicated problems - have to perform breakdown or decomposition.
One of the best ways to deal with complexity is divide and conquer — split the problem
into smaller problems and solve each one separately.

Fault - when one component of the system deviating from its spec.
Failure - when the system as a whole stops providing the required service to the user.
Generally prefer tolerating faults over preventing faults.
For example, hardware faults:
disks may be set up in a RAID configuration,
servers may have dual power supplies and hot-swappable CPUs,
and datacenters may have batteries and diesel generators for backup power...

Telemetry - monitoring performance metrics and error rates.

Concurrency - single process use multiple threads.
But each thread uses the same resource concurrently.
Parallel - concurrency without using same resource.

Internationalization (i18n) - process of changing your software so that it isn't hardwired to one language.
Localization (l10n) - the process of adding the appropriate resources to your software so that a particular language/locale is supported.

Tell, Don’t Ask.
Law of Demeter - object A can call method of object B,
but object A should not "reach through" object B to access yet another object C, to request its services.

ACL (often simpler): List of users (IDs) and privileges to resource (what can do or see).
RBAC: Permissions based on role.

PC: CPU, Main Memory, Secondary Memory (SSD).

Weight Gross - with pack; Weight Net - without pack;

#### SL[AOI]

Service-level agreement (SLA) - commitment between a service provider and a client:
quality, availability, responsibility, performance, costs, etc.

Service-level objective (SLO) - (key element of SLA) specific measurable characteristics
of the SLA such as availability, throughput, frequency, response time, or quality.

Service-level indicator (SLI) - measure of the service level provided by a service provider to a customer.
Common SLIs include latency, throughput, availability, error rate, durability, correctness, etc.

#### !

* Read limits before use something.

#### Full stack:

* Presentation layer (html, css, js).
* Business layer (node, php, etc).
* Data access layer (mongo, mysql, etc).

#### 10 Questions Developers Should be Asking Themselves

* Is there a pattern here?
* How can I make this simpler?
* Why does it work like that?
  (Knowing that something works and knowing why it works that way are two very  different things.)
* Has somebody done this before?
* Who said it first?
  (Always try read the original source of a concept or theory.)
* Do I love what I’m doing?
* Where else could I use this?
* What did I fail at today?
* How can we make this possible?
  (Start from the assumption that whatever you want to do is possible.)
* Who can I learn from?
  (You should never work anywhere where you are the smartest person in the room.)

#### Simple stuff

`camelCase`
`kebab-case`
`snake_case`

````
Odd  (1, 3, 5...) -> x % 2 = 1.
Even (2, 4, 6...) -> x % 2 = 0.

3+4 = 7    # 3 - addend 1;     4 - addend 2;               7    - sum;
3-4 = -1   # 3 - minuend;      4 - subtrahend;             -1   - difference;
3*4 = 12   # 3 - multiplicand; 4 - multiplicator (factor); 12   - product;
3/4 = 0.75 # 3 - numerator;    4 - denominator;            0.75 - ;
3^4 = 81   # 3 - base;         4 - power/exponent;
````

````sh
11f  # (Float)

1e6  # 1 000 000
0b11 # (Binary)
0o77 # (Octal)
0x80 # (HEX) int 128
````

````sh
\0, \000, \x00, \z, \u0000, %00 # null character/terminator/byte
\r                              # carriage return
\b                              # backspace
````

````
byte           -    8 bits  ~ 1      ~ 1
kilobyte  (KB) - 1024 bytes ~ 1000   ~ 1 000
megabyte  (MB) - 1024 KB    ~ 1000^2 ~ 1 000 000
gigabyte  (GB) - 1024 MB    ~ 1000^3 ~ 1 000 000 000
terabyte  (TB) - 1024 GB    ~ 1000^4 ~ 1 000 000 000 000
petabyte  (PB) -            ~ 1000^5 ~
exabyte   (EB) -            ~ 1000^6 ~
zettabyte (ZB) -            ~ 1000^7 ~
yottabyte (YB) -            ~ 1000^8 ~

millisecond - 0,001
microsecond - 0,000 001
nanosecond  - 0,000 000 001
````

````
UTF-8  - from 1 to 4 bytes
UTF-16 - from 2 to 4 bytes
UTF-32 - exactly 32 bits (4 bytes)
````

````
uint8  (tinyint)  - 1byte, 0 to 255
uint16 (smallint) - 2byte, 0 to 65535
uint32 (int)      - 4byte, 0 to 4294967295
uint64 (bigint)   - 8byte, 0 to 18446744073709551615

int8              - 1byte, -128 to 127
int16             - 2byte, -32768 to 32767
int32             - 4byte, -2147483648 to 2147483647
int64             - 8byte, -9223372036854775808 to 9223372036854775807

boolean           - 1byte

float             - 4byte, 1.175494351E-38 to 3.402823466E+38
decimal           - 8byte, 2.2250738585072014E-308 to 1.7976931348623158E+308

ASCII character   - 1byte
UTF-8 character   - 1byte (for ASCII equivalents)
UTF-8 character   - 2byte (for special chars)
````

````
nil | null | none | blank | undefined | unknown | void | empty
anonymous |
environments: local; dev; stage|staging; prod|production;
````

````
( Open Parenthesis
) Close Parenthesis
[ Open Bracket
] Close Bracket
{ Open Curly Bracket
} Close Curly Bracket
< Open Angle Bracket
> Close Angle Bracket
! Exclamation Mark
? Question Mark
' Single Quote
" Double Quote
` Back quote (back tick)
/ Slash (forward slash)
\ Backward Slash
# Pound
$ Dollar
% Percent
& Ampersand
* Asterisk
— Dash (long dash)
- Hyphen
. Dot (period)
@ At Sign
^ Caret
_ Underscore
| Pipe
~ Tilde
≈ Almost Equal
: Colon
; Semicolon
````
