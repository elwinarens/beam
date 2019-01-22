package org.apache.beam.runners.spark.structuredstreaming.translation.batch;

import java.io.Serializable;
import org.apache.beam.runners.spark.structuredstreaming.SparkPipelineOptions;
import org.apache.beam.runners.spark.structuredstreaming.SparkRunner;
import org.apache.beam.sdk.Pipeline;
import org.apache.beam.sdk.options.PipelineOptions;
import org.apache.beam.sdk.options.PipelineOptionsFactory;
import org.apache.beam.sdk.transforms.Create;
import org.apache.beam.sdk.transforms.Flatten;
import org.apache.beam.sdk.values.PCollection;
import org.apache.beam.sdk.values.PCollectionList;
import org.junit.BeforeClass;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

/**
 * Test class for beam to spark source translation.
 */
@RunWith(JUnit4.class)
public class FlattenTest implements Serializable {
  private static Pipeline pipeline;

  @BeforeClass
  public static void beforeClass(){
    PipelineOptions options = PipelineOptionsFactory.create().as(SparkPipelineOptions.class);
    options.setRunner(SparkRunner.class);
    pipeline = Pipeline.create(options);
  }


  @Test
  public void testFlatted(){
    PCollection<Integer> input1 = pipeline.apply(Create.of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10));
    PCollection<Integer> input2 = pipeline.apply(Create.of(11, 12, 13, 14, 15, 16, 17, 18, 19, 20));
    PCollectionList<Integer> pcs = PCollectionList.of(input1).and(input2);
    PCollection<Integer> merged = pcs.apply(Flatten.<Integer>pCollections());
    pipeline.run();
  }

}