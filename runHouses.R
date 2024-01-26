# Define inputs and import needed libraries
N = 100 
#install.packages("RCurl")
library(RCurl)

# Define the file to which we'd like R to write summary statistics
sink("C:/Users/steve/command_line_app_benchmarking/housesOutputR_v3.txt")

# Create vector to stor runtimes for benchmarking experiment
runtimes_vector <- c()

# Conduct 100 iterations of our benchmarking experiment in R
# For each iteration, we import the data, calculate & print summary statistics, and record the trial runtime
for (i in 1:N) {
    start_time <- Sys.time()
  
    houses <- read.csv("C:/Users/steve/command_line_app_benchmarking/housesInput.csv")
    print(summary(houses)) 
    
    end_time <- Sys.time()
    
    trial_run_time <- end_time - start_time
    
    runtimes_vector <- c(runtimes_vector, as.numeric(trial_run_time))
    
}

# Let's take a look at the runtimes for each of the experimental trials.

print(runtimes_vector)


# Let's print summary statistics for the trial runtimes gathered

summary(runtimes_vector)

sink()



