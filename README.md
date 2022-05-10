# Stats Collector home assignment (Golang)

## Table of contents

* [Assignment](#assignment)
* [Testing](#testing)
* [Returning the assignment](#returning-the-assignment)

## Motivation of the exercise
The motivation for the exercise is to demonstrate how you approach application development in your day-to-day worklife, excluding normal PR reviews and calibration sessions which typically exist in normal working environment.

For us the only measure is not the speed of completing the work, but also the quality of the code and the end result.

## Assignment
Application servers receive approximately 20,000 http requests per second. Response timeout is 19,000ms.

Implement a statistics collector that calculates the median and average request response time for the whole data set that has been recorded.

The skeleton implementation for StatsCollector is a starting point and the signatures can be modified as long as the outcome is a functional application.

The HR partner will tell you in what language or languages you should complete the assignment. Make the implementation in the corresponding directory:
- [golang](golang) for Go
- [nodejs](nodejs) for NodeJS / Typescript
If you're asked to implement the assignment in only one language, please do **only** that.

### Goals / expectations:
- While writing your implementation use git as versioning control for your workflow.
- Working implementation for StatsCollector including median and average calculation.
- Ensuring the quality of the code.
- Ensuring the performance of the solution.

### Non goals / not expected:
- Implementing any time series functionality.
- Implementing any data retention policy, e.g., removing data older than N days.
- Persisting data to DB. While this is interesting, it is not expected.
- Providing an executable binary of the application.

## Testing
- Please add unit tests for you implementation, without them the assignment will be rejected.
- There are also benchmark tests included in the tests. Please do ensure the performance of the implementation is sufficient given the number of requests.

## Returning the assignment
Return tar or zip -file containing the whole project as email attachment or other method agreed with recruitment. They will hand it over for engineering review.
