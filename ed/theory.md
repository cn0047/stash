Theory
-

#### Programming:

Imperative programming is a programming paradigm that uses statements that change a program's state.
Imperative program consists of commands for the computer to perform.

Declarative programming focuses on what the program should accomplish
without specifying how the program should achieve the result.

#### Relational database:

A relation is in **1NF** if the domain of each attribute contains only atomic values.
1NF enforces these criteria:
* Eliminate repeating groups in individual tables.
* Create a separate table for each set of related data.
* Identify each set of related data with a primary key.

(most basic: each cell in a table must contain only one piece of information,
and there can be no duplicate rows).

Table is in **2NF** if it is in 1NF
and every non-prime attribute of the table (in main table) is dependent on the whole of every candidate key.

R1

| Employee  | Skill          | Current Work Location         |
|-----------|----------------|-------------------------------|
| Brown     | Light Cleaning | 73 Industrial Way             |
| Brown     | Typing         | 73 Industrial Way             |
| Harrison  | Light Cleaning | 73 Industrial Way             |
| Jones     | Shorthand      | 114 Main Street               |
| Jones     | Typing         | 114 Main Street               |
| Jones     | Whittling      | 114 Main Street               |

<table>
<tr>
<td>
    R2
    <table>
        <thead>
            <th>Employee</th>
            <th>Current Work Location</th>
        </thead>
        <tr>
            <td>Brown</td>
            <td>73 Industrial Way</td>
        </tr>
        <tr>
            <td>Harrison</td>
            <td>73 Industrial Way</td>
        </tr>
        <tr>
            <td>Jones</td>
            <td>114 Main Street</td>
        </tr>
    </table>
</td>
<td>
    R3
    <table>
        <thead>
            <th>Employee</th>
            <th>Skill</th>
        </thead>
        <tr>
            <td>Brown</td>
            <td>Light Cleaning</td>
        </tr>
        <tr>
            <td>Brown</td>
            <td>Typing</td>
        </tr>
        <tr>
            <td>Harrison</td>
            <td>Light Cleaning</td>
        </tr>
        <tr>
            <td>Jones</td>
            <td>Shorthand</td>
        </tr>
        <tr>
            <td>Jones</td>
            <td>Typing</td>
        </tr>
        <tr>
            <td>Jones</td>
            <td>Whittling</td>
        </tr>
    </table>
</td>
</tr>
</table>

Table is in **3NF** if it is in 2NF
and each attribute must represent a fact about the key, the whole key, and nothing but the key
(every non-prime attribute of table is non-transitively dependent on every key of table).

R1

| employee | Department | Tel      |
|----------|------------|----------|
| Grishin  | Accounts   | 11-22-33 |
| Vasilyev | Accounts   | 11-22-33 |
| Petrov   | Logistic   | 44-55-66 |

<table>
<tr>
<td>
    R2
    <table>
        <thead>
          <tr>
              <th>Department</th>
              <th>Tel</th>
          </tr>
        </thead>
        <tr>
            <td>Accounts</td>
            <td>Logistic</td>
        </tr>
        <tr>
            <td>11-22-33</td>
            <td>44-55-66</td>
        </tr>
    </table>
</td>
<td>
    R3
    <table>
        <thead>
            <tr>
                <th>Employee</th>
                <th>Department</th>
            </tr>
        </thead>
        <tr>
            <td>Grishin</td>
            <td>Accounts</td>
        </tr>
        <tr>
            <td>Vasilyev</td>
            <td>Accounts</td>
        </tr>
        <tr>
            <td>Petrov</td>
            <td>Logistic</td>
        </tr>
    </table>
</td>
</tr>
</table>