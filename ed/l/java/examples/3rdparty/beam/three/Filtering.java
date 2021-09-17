package com.p.apache;

import org.apache.beam.sdk.Pipeline;
import org.apache.beam.sdk.options.PipelineOptions;
import org.apache.beam.sdk.options.PipelineOptionsFactory;
import org.apache.beam.sdk.io.TextIO;
import org.apache.beam.sdk.transforms.*;

public class Filtering {

    private static final String CSV_HEADER = 
            "car,price,body,mileage,engV,engType,registration,year,model,drive";

    public static void main(String[] args) {
        PipelineOptions options = PipelineOptionsFactory.create();
        Pipeline pipeline = Pipeline.create(options);

        pipeline.apply("ReadAds", TextIO.read().from("src/main/resources/source/car_ads*.csv"))
                .apply("FilterHeader", ParDo.of(new FilterHeaderFn(CSV_HEADER)))
                .apply("FilterSedanHatchback", ParDo.of(new FilterSedanHatchbackFn()))
                .apply("FilterPrice", ParDo.of(new FilterPriceFn(2000)))
                .apply("PrintToConsole", ParDo.of(new DoFn<String, Void>() {
                    @ProcessElement
                    public void processElement(ProcessContext c) {
                        System.out.println(c.element());
                    }
                }));

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

    private static class FilterSedanHatchbackFn extends DoFn<String, String> {

        @ProcessElement
        public void processElement(ProcessContext c) {
            String[] fields = c.element().split(",");

            String body = fields[2];

            if (body.equals("sedan") || body.equals("hatch")) {
                c.output(c.element());
            }
        }
    }

    private static class FilterPriceFn extends DoFn<String, String> {

        private Double priceThreshold = 0.0;

        public FilterPriceFn(double priceThreshold) {
            this.priceThreshold = priceThreshold;
        }

        @ProcessElement
        public void processElement(@Element String line, OutputReceiver<String> out) {
            String[] fields = line.split(",");

            double price = Double.parseDouble(fields[1]);

            if (price != 0 && price < priceThreshold) {
                out.output(line);
            }
        }
    }
}
