# Synctera Tech challenge

Context: 
This challenge is running in a simulated environment and is using a simple one file approach to understand how you think about and solve problems. The boundaries are that the results need to be accessible via curl, the code needs to compile and run within this environment, and the instructions to get the results should be clear. 

## Objectives

This challenge has two components: written and code. For the written, create a well formatted markdown document outlining your responses. Point form and succinct responses are valued.

### Written
- Review the existing code, what is the work that needs to be done in your opinion? How does this mock up / challenge exercise differ from what you would expect to see in the 'real world'? 
- Also use this document to highlight what you have done in the code (and why)

## Code
Please complete the following and build/test in this simulated environment
- Right now the project uses mock data, please externalize this so that the program can injest a JSON file as a source of mock transactions
- Add a command line option to be able to specify the data source, e.g: 
```
./main --transactions transactions.json
```
- Instead of displaying the PAN with GetTransactions it is preferred to only display the last four digits and replace the rest of the PAN with `*`; create a function to achieve this and ensure that all output is handled in this way
- Create an endpoint that returns the transactions ordered by descending posted_timestamp 
- Create a test for GetTransactions and your new functions
- Provide documentation for how to build, test, run your project


