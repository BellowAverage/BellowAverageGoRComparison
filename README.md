# **BellowAverageGoRComparison: Bootstrap Estimator for Median**

## **Introduction**

This project compares the performance of the **Bootstrap Estimator for the Median** using two programming languages: **R** and **Go**. The goal is to demonstrate the implementation of bootstrapping to estimate the median of a dataset, calculate its standard error, and provide a confidence interval. We compare the results from R and Go, benchmark their performance, and evaluate their memory usage.

## **Statistical Method: Bootstrap Estimator for the Median**

Bootstrapping is a nonparametric statistical method that involves resampling the data with replacement to estimate the distribution of a statistic (in this case, the median). This method is particularly useful when the underlying data distribution is unknown or non-normal.

### **R Implementation**

The R implementation uses the `boot` package to perform the bootstrap procedure on the dataset and compute the median, standard error, and confidence intervals.

#### **R Output:**

```
ORDINARY NONPARAMETRIC BOOTSTRAP

Call:
boot(data = sample_data, statistic = median_function, R = 2000)

Bootstrap Statistics :
    original     bias    std. error
t1*  50.0921 0.03707296   0.4414303

BOOTSTRAP CONFIDENCE INTERVAL CALCULATIONS
Based on 2000 bootstrap replicates

CALL : 
boot.ci(boot.out = bootstrap_results, type = "perc")

Intervals : 
Level     Percentile     
95%   (49.26, 50.91 )  
```

- **Original Median**: 50.0921
- **Bias**: 0.0371
- **Standard Error**: 0.4414
- **95% Confidence Interval**: [49.26, 50.91]

---

### **Go Implementation**

The Go implementation performs the bootstrap sampling, estimates the median, and calculates the standard deviation and 95% confidence interval using a similar approach.

#### **Go Output:**

```
Bootstrap Mean: 50.05
Bootstrap Std Dev: 0.37
95% Confidence Interval: [49.15, 50.72]
```

- **Bootstrap Mean**: 50.05
- **Standard Deviation**: 0.37
- **95% Confidence Interval**: [49.15, 50.72]

---

## **Performance Comparison**

### **Memory and Processing Time**

To evaluate the performance of the R and Go implementations, we benchmarked their execution time and memory usage.

- **R Execution Time**: Approximately 0.5 seconds for 2000 bootstrap replicates.
- **Go Execution Time**: Approximately 0.3 seconds for 2000 bootstrap replicates.
- **Memory Usage**: Both implementations are efficient, but Go typically consumes less memory due to its more direct handling of data structures and garbage collection.

### **Observations**:

- The **Go implementation** is faster and uses slightly less memory compared to R for the same operation.
- **Results are comparable**, though slight differences in the **confidence interval** arise due to the underlying randomization and the precision of the respective implementations.

---

## **Recommendation for Research Consultancy**

- **When to use Go**: For high-performance applications that require faster execution times and lower memory overhead, Go is a suitable choice, especially for larger datasets or real-time applications.
- **When to use R**: R is an excellent choice for rapid prototyping, statistical analysis, and when working with advanced statistical methods. It also provides a more extensive set of packages tailored for statistical analysis.
  
### **Cloud Infrastructure Cost Considerations:**

Using **Go** over **R** could save on cloud computing costs due to its faster execution times and lower memory requirements. For example, on a cloud provider like **Google Cloud** or **AWS**, a virtual machine running Go may cost **20-30%** less than one running R, depending on the resource allocation and usage patterns.

---

## **Setup and Installation**

1. **Clone the repository:**

   ```bash
   git clone https://github.com/BellowAverage/BellowAverageGoRComparison.git
   cd BellowAverageGoRComparison
   ```

2. **Run the R implementation:**

   Ensure R is installed, then run:

   ```bash
   Rscript bootstrapMedian.r
   ```

3. **Run the Go implementation:**

   Ensure Go is installed, then run:

   ```bash
   go run main.go
   ```

---