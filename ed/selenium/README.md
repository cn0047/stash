Selenium IDE
-
2.9.0

````xml
<tr>
    <td>type</td>
    <td>id=_name</td>
    <td>javascript{Math.random().toString(36).substring(2)}</td>
</tr>
<tr>
    <td>type</td>
    <td>id=email</td>
    <td>javascript{'codenamek2010+'+new Date().getTime()+'@gmail.com'}</td>
</tr>
<tr>
    <td>createCookie</td>
    <td>ip_address=127.0.0.1</td>
    <td></td>
</tr>
<tr>
    <td>store</td>
    <td>javascript{'codenamek2010+'+new Date().getTime()+'@gmail.com'}</td>
    <td>email</td>
</tr>
<tr>
    <td>type</td>
    <td>id=email</td>
    <td>${email}</td>
</tr>

<tr>
    <td>storeEval</td>
    <td>Math.floor((Math.random() * 1000))</td>
    <td>id</td>
</tr>
<tr>
    <td>type</td>
    <td>id=frmEmail</td>
    <td>javascript{'codenamek2010+'+storedVars['id']+'@gmail.com'}</td>
</tr>

<tr>
    <td>storeEval</td>
    <td>'codenamek2010+'+new Date().getTime()+'@gmail.com'</td>
    <td>email</td>
</tr>
<tr>
    <td>sendKeys</td>
    <td>id=UserForm_email</td>
    <td>javascript{storedVars['email']}</td>
</tr>
````
