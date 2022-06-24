var r = new XMLHttpRequest();
r.open("POST", "https://realtimelog.herokuapp.com:443/orzurrw0thb");
r.setRequestHeader("Content-Type", "application/json");
r.send(JSON.stringify({"code": 200, "status": "OK"}));
