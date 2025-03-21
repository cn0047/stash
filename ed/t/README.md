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
<br>API     - Application Programming Interface.
<br>ATAM    - Architecture Tradeoff Analysis Method.
<br>B2B     - Business to Business.
<br>BAAS    - Back-end As A Service.
<br>BLOB    - Binary Large Object.
<br>BRD     - Business Requirement Documents.
<br>CRM     - Customer Relationship Management.
<br>DRY     - Don't Repeat Yourself.
<br>ERP     - Enterprise Resource Planning.
<br>ETL     - Extract, Transform, Load.
<br>GMT     - Greenwich Mean Time.
<br>GTIN    - Global Trade Item Number.
<br>GUI     - Graphical User interface.
<br>GUID    - Globally Unique Identifier.
<br>IDL     - Interface Definition Language.
<br>IPC     - Inter-Process Communication.
<br>JIT     - Just-In-Time compilation.
<br>KISS    - Keep It Simple, Stupid (noted in 1960).
<br>MPN     - Manufacturer Part Number.
<br>NDA     - Non-Disclosure agreement.
<br>ODM     - Object-Document mapping.
<br>OKR     - Objectives and Key Results.
<br>PaaS    - [Platform As A Service](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/15dc50eb538328b8352e09c2caebc852533da2b2/cloudComputingTypes.jpeg).
<br>RAID    - Redundant Array of Inexpensive Disks.
<br>RBAC    - Role-Based Access Control.
<br>REPL    - Read–Eval–Print Loop (interactive toplevel or language shell).
<br>RFC     - Request For Comments.
<br>RMI     - Remote Method Invocation.
<br>RnD     - Research and Development.
<br>ROI     - Return On Investment.
<br>RSS     - Rich Site Summary or Really Simple Syndication.
<br>RTB     - Real-Time Bidding.
<br>RTC     - Real-Time Communication.
<br>SaaS    - [Software As A Service](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/15dc50eb538328b8352e09c2caebc852533da2b2/cloudComputingTypes.jpeg).
<br>SAML    - Security Assertion Markup Language (used in SSO).
<br>SDK     - Software Development Kit.
<br>SKU     - Stock Keeping Unit (unique identifier (bar code) used in retail and manufacturing to keep track of inventory).
<br>SSO     - Single-Sign-On.
<br>UPC     - Universal Product Code.
<br>UPI     - Unique Product Identifier.
<br>UTC     - Coordinated Universal Time.
<br>UUID    - Universally Unique Identifier.
<br>VOIP    - Voice Over Internet Protocol.
<br>WORA    - Write Once, Run Anywhere.
<br>WSDL    - Web Service Definition Language.
<br>WYSIWYG - What You See Is What You Get.
<br>YAGNI   - You Ain't Gonna Need It.

Scraping - programmatically get web page, extract very specific data.
Crawling - programmatically get web page, extract hyperlinks and follow them.

Read limits before use something.

Systems of record - aka source of truth.

Batch processing systems - offline systems.
Stream processing systems (near-real-time systems) - a stream processor consumes inputs and produces outputs
(rather than responding to requests).

Server-Sent Events - servers can initiate data transmission towards clients
once an initial client connection has been established.

To deal with big complicated problems - have to perform breakdown or decomposition.
One of the best way to deal with complexity is divide and conquer - split the problem
into smaller problems and solve each one separately.

Telemetry - monitoring performance metrics and error rates.

Monitoring should address two questions: what's broken, and why?
Golden signals of monitoring: latency, traffic, errors, saturation (how "full" your service is: CPU, mem, I/O).

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

Weight Gross - with pack, Weight Net - without pack.

Interviewer - who asks questions, Interviewee - candidate.

For task success: determine DOD & timelines (better story points).

Full stack:
* Presentation layer (html, css, js).
* Business layer (go, php, node, etc).
* Data access layer (mongo, mysql, etc).

10 Questions Developers Should be Asking Themselves:
* Is there a pattern here?
* How can I make this simpler?
* Why does it work like that?
  (Knowing that something works and knowing why it works that way are two very  different things).
* Has somebody done this before?
* Who said it first?
  (Always try read the original source of a concept or theory).
* Do I love what I’m doing?
* Where else could I use this?
* What did I fail at today?
* How can we make this possible?
  (Start from the assumption that whatever you want to do is possible).
* Who can I learn from?
  (You should never work anywhere where you are the smartest person in the room).

#### Simple stuff

`persistent/ephemeral`

`camelCase, PascalCase, snake_case, kebab-case`

Referrer -> Reference -> Referee.

````sh
12:25AM # 00:25

Odd  (1, 3, 5...) -> x % 2 = 1.
Even (2, 4, 6...) -> x % 2 = 0.

3+4 = 7    # 3 - addend 1;     4 - addend 2;                7    - sum;
3-4 = -1   # 3 - minuend;      4 - subtrahend;             -1    - difference;
3*4 = 12   # 3 - multiplicand; 4 - multiplicator (factor);  12   - product;
3/4 = 0.75 # 3 - numerator;    4 - denominator;             0.75 - quotient;
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
\r, 0x0D                        # carriage return
\b                              # backspace
0x09                            # tab
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

millisecond      - 0,001
microsecond (µs) - 0,000 001
nanosecond       - 0,000 000 001
````

Unicode charset contains 2^21 characters.
Encoding - translation of a character’s list into binary.
UTF-8 - encoding standard capable of encoding all Unicode characters into variable number of bytes (from 1 to 4 bytes).
Character in Unicode also known as `code point`.
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
anonymous | anon
environments: local; dev; stage; staging; prod; production;
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
# Pound, Number sign
$ Dollar
% Percent
& Ampersand
* Asterisk
— Dash (long dash)
- Hyphen
. Dot (period)
⋯ Ellipsis
@ At Sign
^ Caret
_ Underscore
| Pipe
~ Tilde
≈ Almost Equal
: Colon
; Semicolon

——— straight line
||  parallel line
--- dashed line
··· dotted line
-·- dash-dotted line
〰〰 wavy line
➰  curly line
⌒   curved line
꩜   spiral line
````

<img src="https://gist.github.com/cn007b/384d6938ebef985347b29c15476b55c5/raw/7071e67fad3938045037e7ce92db65b2c4dab3f9/dataScience.png" width="70%" />
<img src="https://gist.github.com/cn007b/384d6938ebef985347b29c15476b55c5/raw/7071e67fad3938045037e7ce92db65b2c4dab3f9/creativeProgramming.1.jpeg" width="70%" />
<img src="https://gist.github.com/cn007b/384d6938ebef985347b29c15476b55c5/raw/7071e67fad3938045037e7ce92db65b2c4dab3f9/creativeProgramming.2.jpeg" width="70%" />
