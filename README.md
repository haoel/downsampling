# Downsampling Algorithm

The Golang implementation for downsampling time series data algorithm

- [Downsampling Algorithm](#downsampling-algorithm)
  - [Background](#background)
  - [Acknowledgment](#acknowledgment)
  - [Algorithms Library](#algorithms-library)
  - [Demo](#demo)
  - [Performance](#performance)
    - [Profiling](#profiling)
    - [Benchmark](#benchmark)
  - [Further Reading](#further-reading)
  - [License](#license)

## Background

While monitoring the online system, there could be so many metrics' time series data will be stored in the ElasticSearch or NoSQL database for analysis. As time passed, storing every piece of historical data is not a very effective way, and those huge data could impact the analysis performance and the cost of storage.

One solution just simply deletes the aged historical data(e.g. only keep the latest 6 months' data), but there is a solution we can compressing those data to a small size with good resolution.

Here is the Go library to demonstrate how to downsamping the time series data from 7500 points to 500 points (Actually, you can downsample it to 200 or 300 points).


## Acknowledgment

- All of the algorithms are based on Sveinn Steinarsson's 2013 paper [Downsampling Time Series for Visual Representation](
https://skemman.is/bitstream/1946/15343/3/SS_MSthesis.pdf)

- This implementation refers to Ján Jakub Naništa's [implementation by Typescript](https://github.com/janjakubnanista/downsample)

- The test data I borrow from one of Python implementation which is [here](https://github.com/devoxi/lttb-py/)


## Algorithms Library

[Sveinn Steinarsson's paper](
https://skemman.is/bitstream/1946/15343/3/SS_MSthesis.pdf) mentioned 3 types of algorithms:

- Largest triangle three buckets (LTTB)
- Largest triangle one bucket (LTOB)
- Largest triangle dynamic (LTD)

You can find all of these implementations under `core` directory.

And you can import the library by:

```go
import "github.com/haoel/downsampling/core"
```

## Demo


Following the below instructions compile and run this repo.

```shell
make
./demo/build/bin/main
```

If everything goes fine, you will see the following message

```
2019/09/07 18:34:42 Reading the testing data...
2019/09/07 18:34:42 Downsampling the data from 7501 to 500...
2019/09/07 18:34:42 Downsampling data - LTOB algorithm done!
2019/09/07 18:34:42 Downsampling data - LTTB algorithm done!
2019/09/07 18:34:42 Downsampling data - LTD algorithm done!
2019/09/07 18:34:42 Creating the diagram file...
2019/09/07 18:34:43 Successfully created the diagram - ....../data/downsampling.chart.png
```

You can go to the `./demo/build/data/` directory to check the diagram and the CVS files.

The diagram picture as below
- The first black chart at the top is the raw data with 7500 points
- The second, third, and fourth respectively are LTOB, LTTB, and LTD downsampling data with 500 points
- The last one at the bottom is just put all together.

![](./demo/data/downsampling.chart.png?raw=true)

## Performance

You can use the following makefile target to analyze the performance of these algorithms.

### Profiling

```
make prof
```

### Benchmark

```
make bench
```

## Further Reading

* [The Billion Data Point [Challenge](https://eng.uber.com/billion-data-point-challenge/) by the Uber Engineering team
* [Visualize Big Data on Mobile](http://dduraz.com/2019/04/26/data-visualization-mobile/) by dduraz
* [Sampling large datasets in d3fc](http://blog.scottlogic.com/2015/11/16/sampling-large-data-in-d3fc.html) by William Ferguson
* [Downsampling algorithms](http://www.adrian.idv.hk/2018-01-24-downsample/) by Adrian S. Tam


Enjoy it!

## License

[MIT License](./LICENSE)