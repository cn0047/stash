// Shows time difference in format MM::ss.
const f = (start, end) => {
    const startedAt = new Date(start);
    const endedAt = new Date(end);

    const diffInSecconds = Math.floor((endedAt-startedAt) / 1000);
    const minutes = Math.floor(diffInSecconds / 60);
    const secconds = diffInSecconds % 60;

    let minutesStr =  minutes.toString();
    let seccondsStr = secconds.toString();
    minutesStr = minutesStr.length > 1 ? minutesStr : "0"+minutesStr;
    seccondsStr = seccondsStr.length > 1 ? seccondsStr : "0"+seccondsStr;

    return `${minutesStr}:${seccondsStr}`;
};

console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T17:14:19.6058921760Z'));
console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T17:15:09.6058921760Z'));
console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T17:14:10.6058921760Z'));
console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T17:14:09.6058921760Z'));
console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T17:15:08.6058921760Z'));
console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T17:15:09.6058921760Z'));
console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T17:15:10.6058921760Z'));
console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T17:25:09.6058921760Z')); // 11:00
console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T18:14:08.6058921760Z')); // 59:59
console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T18:14:09.6058921760Z')); // 60:00
console.log(f('2019-10-31T17:14:09.6058921760Z', '2019-10-31T18:14:10.6058921760Z')); // 60:01
