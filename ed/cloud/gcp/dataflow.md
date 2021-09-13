Dataflow
-

[docs](https://cloud.google.com/dataflow/docs/quickstarts)
[resources](https://cloud.google.com/dataflow/docs/resources)

Dataflow - service for creating and evaluating data processing pipelines.
Streaming analytics for stream and batch processing.

Pipeline: PCollection (immutable), PTransform.

Window - windowing functions divide unbounded collections into logical components, or windows.
Watermark - threshold that indicates when Dataflow expects all of data in a window to have arrived.
Trigger - determine when to emit aggregated results as data arrives.
Source.
Sink.
Schema.
State.
Timer (event time, processing time).

ParDo - transform for generic parallel processing.
