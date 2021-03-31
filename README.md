# Golang Bootcamp

## Introduction

This poroject is about a consta list page. 

- http://localhost:10000/  Render the Home page. 
- http://localhost:10000/get-contacts  Render all contact list
- http://localhost:10000/get-contact-info?id=1 Render the contatc Info of a contact with id: 1


## How to run

run the command: go run main.go

## Requirements

These are the main requirements we will evaluate:

- Use all that you've learned in the course:
  - Best practices
  - Go basics
  - HTTP handlers
  - Error handling
  - Structs and interfaces
  - Clean architecture
  - Unit testing
  - CSV file fetching
  - Concurrency

## Getting Started

To get started, follow these steps:

1. Fork this project
1. Commit periodically
1. Apply changes according to the reviewer's comments
1. Have fun!




## Second Deliverable (due March 18th 23:59PM)

Based on the self-study material and mentorship covered until this deliverable, we suggest you perform the following:

- Create a client to consume an external API
- Add an endpoint to consume the external API client
- The information obtained should be stored in the CSV file
- Add unit testing
- Update the endpoint made in the first deliverable to display the result as a JSON
- Refator if needed

> Note: what’s listed in this deliverable is just for guidance and to help you distribute your workload; you can deliver more or fewer items if necessary. However, if you deliver fewer items at this point, you have to cover the remaining tasks in the next deliverable.

## Final Deliverable (due March 30th 23:59PM)

- Add a new endpoint
- The endpoint must read items from the CSV concurrently using a worker pool
- The endpoint must support the following query params:

```text
type: Only support "odd" or "even"
items: Is an Int and is the amount of valid items you need to display as a response
items_per_workers: Is an Int and is the amount of valid items the worker should append to the response
```

- Reject the values according to the query param ***type*** (you could use an ID column)
- Instruct the workers to shut down according to the query param ***items_per_workers*** collected
- The result should be displayed as a response
- The response should be displayed when:

  - The workers reached the limit
  - EOF
  - Valid items completed

> Important: this is the final deliverable, so all the requirements must be included. We will give you feedback on March 31th. You will have 3 days more to apply changes. On April 5th, we will stop receiving changes at 11:00 am.

## Submitting the deliverables

For submitting your work, you should follow these steps:

1. Create a pull request with your code, targeting the master branch of the repository academy-go-q12021.
2. Fill this [form](https://forms.gle/Kf9dvnth3hPe3Azg9) including the PR’s url
3. Stay tune for feedback
4. Do the changes according to the reviewer's comments

