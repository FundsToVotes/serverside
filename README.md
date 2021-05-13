# Funds to Votes Serverside
_Informatics Capstone Project 2021 | University of Washington Information School_

## Team Members
- Grady Thompson (Project Manager, Generalist)
- Jay Houppermans (Backend Developer, Server Administrator)
- Haykal Mubin (Frontend Developer)
- Reyan Haji (Data Analyst)

## Contact Information

Email the project team at hello@fundstovotes.info. Or, create an issue in our GitHub repository.

## Repository Purpose

This repository contains the serverside code for the Funds to Votes project. It contains a gateway for our [custom API backend](https://api.fundstovotes.info/hello), the dockerfiles and configuration for our custom mySQL database, and a variety of scripts for fetching and exploring the APIs we acquire data from. 

This repository is primarily written in Go. Each continuously running component of the code also includes a Dockerfile, and is intended to be run inside a Docker container. These containers are then deployed to our Amazon AWS instance using the helper .sh scripts. 

## Project Architecture 

Go expects to see all Go projects on your computer to live in one directory, your `$GOPATH`/src directory. This repository is currently configured with the expectation that your `$GOPATH` is the location where the repository has been cloned. 

Each folder within `src` is a different project. They have their own main.go files, build to their own .exe files, and generally run independently. 

The projects, as well as significant subfiles, are listed below. All subfolders are considered "packages" of the main project. 

- **gateway** - the most significant project. This is the code for the fundstovotes custom API backend. It contains: 
    - helper scripts, as described in the *Hosting* section
    - a Dockerfile for building the ftvgateway container
    - `main.go` - A script that routes queries for the various API endpoints, loads enviroment variables, and connects to the database 
    - `handlers` - a package containing the logic for each API endpoint
        + `billsAndIndustries.go` - code for the Bills endpoint
        + `topTen.go` - code for the topten endpoint
        + `dummybillsAndIndustries.go` - hardcoded example data for testing the bills and industries endpoint
            - Note: This is currently out of date, due to restructuring
        + `dummytop10Industries.go` - hardcoded example data for testing the top10 endpoint
        + `data_structures.go` - type declarations for every Struct used in this folder
        + `cors.go` - handler code for CORS preflight request security. 
        + `hello.go` - a basic "Hello World" endpoint used to confirm whether the gateway is online
- **db** - This contains the configuration files for the repo used by the topten endpoint. 
    - Dockerfile for building the docker container
    - `schema.sql` - database schema
    - `runDocker.sh` - helper command to run the database locally for testing purposes
    - `seeTables.sh` - helper command to view all tables inside a currently running database
- **archived_test_code** - this repository contains scripts that call the various APIs we pull data from. This code has all been either abandoned, or encorporated into one of the gateway handlers. It can be used to help in debugging or designing handlers. 
    - bills_api_call - the code used for designing and testing the `bills` endpoint
        + Can be useful if you want to view the responses from the Propublica Members and Bills API endpoints
    - crp_id_calls - an abandoned project for fetching the CRP_id of a representative given their name. 
        + Can be useful if you want to view the responses from the Propublica Members endpoint
- **api-call** - Code for acquiring and storing the data used in the topTen endpoint
    - `main.go` - For all members of congress, fetches openSecrets data on their top ten contributing industries, and stores it in the database
    - assorted csv files - lists of CRP ids for the current members of Congress, for use in `main.go`
    - `fetchdata` - package containing the functions involved in fetching the data from opensecrets 
    - `useDB` - package containing the functions involved in communication with the database
        + `creation.go` - creates a database if one doesn't exist, opens a connection to a database
        + `insert.go` - inserts data into a database
        + `dataStructures.go` - type declarations for every Struct used in this folder

## API Endpoints

The custom backend currently supports two endpoints. They both accept GET requests over HTTPS, and use query parameters. They do not support HTTP requests, or any other request types or methods. 

### Bills - https://api.fundstovotes.info/billstest?member_id=K000388

Purpose: Given a member of Congress's congressional ID, return a list of all bills they've recently voted on. Correlate the Primary Subject of each bill with an OpenSecrets industry, and return all data as a JSON object. 

Data sources: 
- [Propublica Members API](https://projects.propublica.org/api-docs/congress-api/members/) - Get a Specific Member’s Vote Positions
- [Propublica Bills API](https://projects.propublica.org/api-docs/congress-api/bills/) - Get a Specific Bill
- [Library of Congress](https://www.congress.gov/search?q={%22source%22:%22legislation%22,%22congress%22:117}&searchResultViewType=expanded) - List of Primary Subjects on left sidebar, hardcoded in
- [OpenSecrets Industry Codes](https://www.opensecrets.org/open-data/api-documentation) - List of Industries as recognized by OpenSecrets, hardcoded in. Data found in the CRP_IDs.xls spreadsheet. 

Query Parameter: `member_id`. An ID corresponding to a Member of Congress. This parameter is also known as a representative's "Biography Number". It can be found as the "id" parameter in Propublica API data responses. 

### Topten - https://api.fundstovotes.info/topten?name=Pelosi,Nancy

Purpose: Given a member of Congress's name, return a list of the top ten industries that contributed to their campaign, as reported by OpenSecrets. OpenSecrets has a very low limit on the amount of queries it accepts per day (200), so these queries should be run through a backend. 

The backend code in the api-call folder periodically fetches this data, and stores it in the Funds to Votes TopTen database. This endpoint then queries the internal Funds to Votes database, rather than querying OpenSecrets directly. 

Query Parameter `name`. A representative's name, in "Lastname,Firstname" format. 

## Contributing

If you'd like to contribute to this repository as an individual, please follow the following steps. 

 1. Clone the repository to your local machine
 2. Install Go and Docker, as described [in this tutorial](https://drstearns.github.io/tutorials/server-side-setup/)
    - This tutorial includes instructions for installing Node Js and other tools. These are unnecessary.
3. Create API keys as described in the API Keys section of this Readme
4. Create a seperate branch for your work. All new commits should occur on branches other than main. 
5. Push these commits, then create a pull request
6. One of the project maintainers will review and merge your pull request as soon practical
    - If you don't hear from us within a few weeks, feel free to reach out at hello@fundstovotes.info and ask about an expected timeline

If you're a member of a future iSchool Capstone team who intends to improve upon this project, please follow the following steps

1. Contact us at hello@fundstovotes.info for transfer of hosting credentials, and other accounts
    - On the serverside, this is the Docker Hub account, the domain name, and the AWS account.
    - We'll also share with you the server IP and credentials for the AWS instance, so you can ssh into it. 
2. Configure the domain name email services to forward to your own emails as well
3. Make sure all team members have access to both project Github repos
3. Review the existing code and documentation, and send any questions any questions you have
4. Follow the individual contributer steps 1 - 3
5. Get started! Come up with a list of new features or improvements you want to add, and begin implementing them. 

## Hosting

The Gateway folder in this project is hosted on an AWS EC2 instance inside a Docker container. The Docker container built from the db folder is hosted as a seperate Docker container on that same instance. 

Upon `ssh`ing into the EC2 instance, you can run the command `docker ps -a` to see whether or not these containers are running. 

The project contains a variety of Helper Scripts intended to make building and deploying these containers easier. If unfamilliar with Docker, see the Tutorials section for more information. 

Helper scripts can be found in any folder that gets built and deployed to the EC2 server. To run these scripts, navigate to them using a Bash command terminal, then enter the command `./scriptName.sh`. For example, to run the `build.sh` script in the gateway folder, navigate to `{whever_you_saved_the_repo}/src/gateway/`, and type `./build.sh`

The helper scripts: 
- `build.sh` - builds a Go executable, then builds a Docker container in the current directory, according to Dockerfile
- `deploy.sh` - builds a Docker container via the `build.sh` script, pushes it to Dockerhub, then `ssh`s into the EC2 server. It then runs the `inside_aws_script.sh` script. 
- `inside_aws_script.sh` - Stops the existing Docker Container, pulls the new container from Dockerhub, then runs it. This script is *only* intended to be called as part of `deploy.sh`, it should not be run on its own. 
-  `run_docker_locally` - a debugging tool. Rather than building and deploying to the remote server, this script runs a built Docker container on the local development computer. 
    - Note: You will need to create self-signed certificates on your computer for this to work. 
    - Additional warning: This script is not kept up to date with the current version of the Gateway script. It is only updated when the primary developer has a need to test on their local machine. Therefore, it may or may not be working. 


## API Keys and Enviroment Variables

To use this application, you need to obtain API Keys for: 

- [OpenSecrets](https://www.opensecrets.org/api/admin/index.php)
    + We have one registered with the hello@fundstovotes.info account. Future capstone teams, ask us for the credentials!
- [ProPublica Congress API](https://www.propublica.org/datastore/api/propublica-congress-api)
    + You can use the same one that you're using for the Clientside code

Store them in a file called `.env` at the root of your cloned repository. It should look like this before you fill in your keys:
```
GATEWAY_PORT=:443
REMOTE_SERVER_LOGIN=
SERVERSIDE_APP_PROPUBLICA_CONGRESS_API_KEY=
SERVERSIDE_OPENSECRETS_API_KEY=
MYSQL_ROOT_PASSWORD=
MYSQL_DATABASE=ftvBackEnd
```
Variable explanations 
- `GATEWAY_PORT` - Port that the Gateway listens for requests on
- `REMOTE_SERVER_LOGIN` - Login information for the Amazon EC2 Remote server. Should be in format `ec2-user@{{{IP_ADDRESS}}}.us-west-2.compute.amazonaws.com`
- `SERVERSIDE_APP_PROPUBLICA_CONGRESS_API_KEY` - API key for Propublica Congress API
- `SERVERSIDE_OPENSECRETS_API_KEY` - API key for Opensecrets API
- `MYSQL_ROOT_PASSWORD` - Root password for the internal MySql database used for the TopTen endpoint
- `MYSQL_DATABASE` - Name of the internal MySQL database

## Tips, Tutorials, and Tools

If you haven't written in Go before, it sometimes has a steep learning curve! These resources could prove helpful in learning the language

- UW iSchool Info 441 Tutorials - https://drstearns.github.io/tutorials/ 
- The official Go documentation - https://golang.org/doc/
- Automatic JSON to Go Struct converter - https://mholt.github.io/json-to-go/ 
- Go & MySQL tutorials - https://golangbot.com/mysql-create-table-insert-row/ 
    - These instructions are easier and simpler than those covered in the iSchool tutorials

Similarly, these resources provide tutorials on setting up an AWS EC2 Server to host our Gateway and Database
- Pointing your NameCheap Domain Name to an EC2 Instance -  https://u.osu.edu/walujo.1/2016/07/07/
- Alternative tutorial - https://techgenix.com/namecheap-aws-ec2-linux/associate-namecheap-domain-to-amazon-ec2-instance/ 
- Routing Traffic to an AWS Instance - https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-to-ec2-instance.html
    - This is a more detailed companion tutorial to the two above. However, you should ignore the Route 53 part, and use Elastic IPs instead. For a project of this size, Elastic IPs are cheap, while Route 53 costs a small amount of money
- Cheap Domain Names - https://nc.me/
- Installing Docker on an AWS instance - https://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-basics.html#install_docker

Docker tutorials 
- iSchool tutorial on what Docker is and how it works - https://drstearns.github.io/tutorials/docker/
- Dockerfile Reference - https://docs.docker.com/engine/reference/builder/
- Best practices for writing Dockerfiles - https://docs.docker.com/develop/develop-images/dockerfile_best-practices/ 