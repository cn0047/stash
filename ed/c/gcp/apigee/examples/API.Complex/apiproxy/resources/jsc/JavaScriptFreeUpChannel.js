body = context.getVariable('in.body');
payload = JSON.parse(decodeURIComponent(body));

lId = payload.length > 0 && payload[0].args && payload[0].args.length > 0 ? payload[0].args[0] : "";
cmd = payload.length > 0 && payload[0].cmd ? payload[0].cmd : "";

if (cmd !== "xxx") {
  throw new Error("got unsupported command: "+cmd);
}

IDs = [];
for (i = 0; i < payload.length; i++) {
  if (!payload[i].cmd || !payload[i].args || !payload[i].args.length) {
    throw new Error("each object must have next properties: cmd, args");
  }
  if (payload[i].cmd !== cmd) {
    throw new Error("commands mismatch: "+cmd+" and "+payload[i].cmd);
  }
  if (payload[i].args[0] !== lId) {
    throw new Error("locations mismatch: "+lId+" and "+payload[i].args[0]);
  }
  if (!payload[i].args[1]) {
    throw new Error("missing required param: 1");
  }
  IDs.push(payload[i].args[1]);
}
IDQuery = IDs.map(function(id) { return "id="+id; }).join("&");

context.setVariable('out.lId', lId);
context.setVariable('out.IDs', IDs);
context.setVariable('out.IDQuery', IDQuery);
