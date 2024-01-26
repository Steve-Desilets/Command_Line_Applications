#!/usr/bin/env python
# coding: utf-8

# # Command Line Application - Benchmarking Experiment - Python Code

# #### Steve Desilets

# #### January 26, 2024

# This project focuses on conducting a benchmarking experiment that compares the computational efficiency of Python, R, and Go when executing data science programs from the command prompt application.  For this experiment, we focus on the ability of these three languages to quickly calculate summary statistics for the variables in a dataset.  This file focuses on the Python portion of this experiment.

# In[1]:


#import the necessary packages
import pandas as pd
import datetime
import matplotlib.pyplot as plt 

# Define necessary inputs
N = 100 
runtimes_list = []

# Run the Benchmarking Test function 100 times and compile the runtimes 
with open('housesOutputPy.txt', 'wt') as outfile:
    for i in range(N):
        start_time = datetime.datetime.now()
        
        houses = pd.read_csv("housesInput.csv")
        outfile.write(houses.describe().to_string(header=True, index=True))
        outfile.write("\n")
        
        end_time = datetime.datetime.now()
        runtime = end_time - start_time
        runtimes_list.append(runtime)
            
    microseconds_list = []

    for i in runtimes_list:
        microseconds_list.append(i.microseconds)
    
    microseconds_series = pd.Series(microseconds_list) 
    
    outfile.write("\n")
    outfile.write("Summary Statistics for Distribution of Experimental Trial Runtimes in Python (in Microseconds)")
    outfile.write("\n")
    
    outfile.write(microseconds_series.describe().to_string(header=True, index=True))
    
    outfile.write("\n")
    

