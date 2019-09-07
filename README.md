# Downsampling Algorthim

The Golang implemtation for downsampling time series data algorthim 

## Background

While monitoring the online system, there are so many metrics's time series data could be store into the Elasticsearch for analysis. When the time passed, the histrical data is not very effective, and it could impact the system performance and the cost.

However, remove the histrical data some time is not a option, but we can compressing them. 

Here is an demo shows how to downsamping the time series data from 7500 points to 500 points.

## Acknowledgement

- All of the algorthims are based on Sveinn Steinarsson's 2013 paper [Downsampling Time Series for Visual Representation]( 
https://skemman.is/bitstream/1946/15343/3/SS_MSthesis.pdf)

- The implmentation is base on [Ján Jakub Naništa's implementation by Typescript](https://github.com/janjakubnanista/downsample)

- The test data I borrow from one of python implmentation which is [here](https://github.com/devoxi/lttb-py/)


## Usage

[Sveinn Steinarsson's paper]( 
https://skemman.is/bitstream/1946/15343/3/SS_MSthesis.pdf) mentioned 3 types of algorithm:

- Largest triangle three buckets (LTTB)
- Largest triangle one bucket (LTOB)
- Largest triangle dynamic (LTD)

You can file all of the implmentation under `src/downsampling` directory.


Following the below instuction to compile and run this repo.

```
make vget 
make 
./build/bin/main
```

If everything goes fine, you will see the following message

```
2019/09/07 18:34:42 Reading the testing data...
2019/09/07 18:34:42 Downsampling the data from 7501 to 500...
2019/09/07 18:34:42 Downsampling data - LTOB algorithm done!
2019/09/07 18:34:42 Downsampling data - LTTB algorithm done!
2019/09/07 18:34:42 Downsampling data - LTD algorithm done!
2019/09/07 18:34:42 Creating the diagram file...
2019/09/07 18:34:43 Successfully created the diagram - ..../build/data/downsampling.chart.png
```

You can go to the `./build/data/` directory to check the diagram and the cvs files.

The diagram picture as below
- The first black chart at the top is the raw data with 7500 points
- The second, third, and fourth respectively are LTOB, LTTB and LTD downsampling data with 500 points
- The last one at the bottom just put all together.

![](./data/downsampling.chart.png?raw=true)

Enjoy it!