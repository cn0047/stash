/*
Task:
For companies in our collection founded in 2004 and having 5 or more rounds of funding,
calculate the average amount raised in each round of funding.
*/

db.companies.aggregate([
    {$match: {"founded_year": {$eq: 2004}}},
    {$project: {_id: 1, name: 1, founded_year: 1, funding_rounds: 1, countFundingRounds: {$size: "$funding_rounds"}}},
    {$match: { "countFundingRounds": {$gte: 5}}},
    {$unwind: "$funding_rounds"},
    {$project: {_id: 0, name: 1, roundCode: "$funding_rounds.round_code", amount: "$funding_rounds.raised_amount"}},
    {$group: {_id: {companyName: "$name", r: "$roundCode"}, amountAVG: {$avg: "$amount"}}},
    {$group: {_id: "$_id.companyName", minAVGAmount: {$min: "$amountAVG"}}},
    {$sort: {minAVGAmount: 1}},
]).pretty();

// or

db.companies.aggregate([
    {$match: {"founded_year": {$eq: 2004}}},
    {$project: {_id: 1, name: 1, founded_year: 1, funding_rounds: 1, countFundingRounds: {$size: "$funding_rounds"}}},
    {$match: { "countFundingRounds": {$gte: 5}}},
    {$project: {_id: 0, name: 1, minAVGAmount: {$avg: "$funding_rounds.raised_amount"}}},
    {$sort: {minAVGAmount: 1}},
]).pretty();
