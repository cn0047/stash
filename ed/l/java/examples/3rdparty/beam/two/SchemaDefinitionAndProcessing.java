package com.p.apache;

import org.apache.beam.sdk.Pipeline;
import org.apache.beam.sdk.options.PipelineOptions;
import org.apache.beam.sdk.options.PipelineOptionsFactory;
import org.apache.beam.sdk.schemas.transforms.Select;
import org.apache.beam.sdk.transforms.MapElements;
import org.apache.beam.sdk.io.TextIO;
import org.apache.beam.sdk.transforms.*;
import org.apache.beam.sdk.values.KV;
import org.apache.beam.sdk.values.Row;
import org.apache.beam.sdk.values.TypeDescriptors;

import java.util.Collections;

public class SchemaDefinitionAndProcessing {

    private static final String CSV_HEADER = "Date,Product,Card,Country";

    public static void main(String[] args) {
        PipelineOptions options = PipelineOptionsFactory.create();
        Pipeline pipeline = Pipeline.create(options);

        pipeline.apply("ReadLines", TextIO.read().from("src/main/resources/source/in.csv"))
                .apply(ParDo.of(new FilterHeaderFn(CSV_HEADER)))
                .apply(ParDo.of(new ParseSalesRecord()))
                .apply("Extract Payment Type", FlatMapElements
                        .into(TypeDescriptors.strings())
                        .via(row -> Collections.singletonList(row.paymentType)))
                .apply("Count Payment Type", Count.perElement())
                .apply("FormatResult", MapElements
                        .into(TypeDescriptors.strings())
                        .via((KV<String, Long> kv) ->
                                kv.getKey() + "," + kv.getValue()))
                .apply("WriteResult",
                        TextIO.write()
                                .to("src/main/resources/sink/payment_type_count")
                                .withSuffix(".csv")
                                .withShardNameTemplate("-SSS")
                                .withHeader("PaymentType,Count"));

        pipeline.run().waitUntilFinish();
    }

    private static class FilterHeaderFn extends DoFn<String, String> {

        private final String header;

        public FilterHeaderFn(String header) {
            this.header = header;
        }

        @ProcessElement
        public void processElement(ProcessContext c) {
            String row = c.element();

            if (!row.isEmpty() && !row.equals(this.header)) {
                c.output(row);
            }
        }
    }

    private static class ParseSalesRecord extends DoFn<String, SalesRecord> {

        @ProcessElement
        public void processElement(@Element String line, OutputReceiver<SalesRecord> out) {
            String[] data = line.split(",");

            SalesRecord record = new SalesRecord(data[0], data[1],
                    Integer.parseInt(data[2]), data[3], data[4]);

            out.output(record);
        }
    }
}
