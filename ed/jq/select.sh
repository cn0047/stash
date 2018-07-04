
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
