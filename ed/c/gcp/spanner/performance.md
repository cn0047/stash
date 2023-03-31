Performance
-

[configurations](https://cloud.google.com/spanner/docs/instance-configurations)

For optimal performance:
* Design a schema that prevents hotspots.
* For optimal write latency, place compute resources for write-heavy workloads within/close to the default leader region.
* For optimal read performance outside of the default leader region, use staleness of at least 15 seconds.
* To avoid single-region dependency for your workloads, place critical compute resources in at least two regions.
* Provision enough compute capacity to keep high priority total CPU utilization under 45% in each region.
* For the amount of throughput per Spanner node see: https://cloud.google.com/spanner/docs/performance#multi-region-performance
