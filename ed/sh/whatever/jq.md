jq
-

[doc](https://stedolan.github.io/jq/manual/)
[play](https://jqplay.org/)

````sh
jq -c   # compact json (inline)
jq -C   # colorize
jq -M   # don't colorize
jq -r   # raw output
jq type # simple way to check JSON validity

--sort-keys / -S
````

````sh
.foo.bar
.["foo::bar"] # special chars
.foo?         # Optional object identifier
.[10:15]      # Array/string slice
,             # (comma) Two filters separated by a comma

.items[]|.id
... | jq '.items[]|.ImageLink' # ✅
... | jq '.items[]|.id,.name'
... | jq '.items[]|{id, name}'
... | jq '.items|join(",")'
... | jq '.[]|= keys'          # 1st level keys
... | jq -r '.taskArns[0]'     # string without quotes

o='{"foo": "f", "bar": "b", "items": [1, 2, 3] }'
echo $o | jq                                # prettify json output
echo $o | jq '.items | length'              # count
echo $o | jq '.items | last'                # 3
echo $o | jq '.items | map(select(. >= 2))' # [2,3]

echo '{"arr": [0, 1, "a", "b"]}' | jq '.arr[2:]'                    # ["a","b"]
echo '[{"n":"foo"},{"n":"bar"}]' | jq '.[].n'                       # foo \n bar
echo '[{"n":"foo"},{"n":"bar"}]' | jq '.[] | select(.n == ("foo"))' # {"n": "foo"}

cat data.json \
| jq '.gcp_price_list | del(.sustained_use_base,.sustained_use_tiers) | .[] | keys[]' \
| uniq

````

Select:
````sh
jq '.CVE_Items[] | select (.impact.baseMetricV3.cvssV3.baseSeverity == ("MEDIUM")) | .cve.CVE_data_meta.ID' <<JSON
{
  "CVE_data_type": "CVE",
  "CVE_data_format": "MITRE",
  "CVE_data_version": "4.0",
  "CVE_data_numberOfCVEs": "5465",
  "CVE_data_timestamp": "2018-07-04T07:00Z",
  "CVE_Items": [
    {
      "cve": {
        "data_type": "CVE",
        "data_format": "MITRE",
        "data_version": "4.0",
        "CVE_data_meta": {
          "ID": "CVE-2018-0001",
          "ASSIGNER": "cve@mitre.org"
        }
      },
      "impact": {
        "baseMetricV3": {
          "cvssV3": {
            "version": "3.0",
            "vectorString": "CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:H",
            "attackVector": "NETWORK",
            "attackComplexity": "HIGH",
            "privilegesRequired": "NONE",
            "userInteraction": "NONE",
            "scope": "UNCHANGED",
            "confidentialityImpact": "NONE",
            "integrityImpact": "NONE",
            "availabilityImpact": "HIGH",
            "baseScore": 5.9,
            "baseSeverity": "MEDIUM"
          },
          "exploitabilityScore": 2.2,
          "impactScore": 3.6
        }
      }
    }
  ]
}
JSON
````
