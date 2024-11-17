library(boot)

median_function <- function(data, indices) {
  return(median(data[indices]))
}

set.seed(123)
sample_data <- rnorm(1000, mean = 50, sd = 10)

bootstrap_results <- boot(data = sample_data, statistic = median_function, R = 2000)

print(bootstrap_results)

bootstrap_ci <- boot.ci(bootstrap_results, type = "perc")
print(bootstrap_ci)
