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

This repository contains the serverside code for the Funds to Votes project. It contains a gateway for our custom API backend, the dockerfiles and configuration for our custom mySQL database, and a variety of scripts for fetching and exploring the APIs we acquire data from. 

This repository is primarily written in Go. Each continuously running component of the code also includes a Dockerfile, and is intended to be run inside a Docker container. These containers are then deployed to our Amazon AWS instance using the helper .sh scripts. 

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

## Hosting <TODO - significant revision on this section>

The Gateway folder in this project is hosted on an AWS EC2 instance inside a Docker container. The Dattabase folder is hosted as a seperate Docker container on that same instance. 

Upon `ssh`ing into the EC2 instance, you can run the command `docker ps -a` to see whether or not these containers are running. 

The project contains a variety of Helper Scripts intended to make building and deploying these containers easier. If unfamilliar with Docker, see the Tutorials section for more information. 

Helper scripts can be found in any folder that gets built and deployed to the EC2 server. To run these scripts, navigate to them using a Bash command terminal, then enter the command `./scriptName.sh`. For example, to run the `build.sh` script in the gateway folder, navigate to `{whever_you_saved_the_repo}/src/gateway/`, and type `./build.sh`

The helper scripts: 
- `build.sh` - builds a Go executable, then builds a Docker container in the current directory, according to Dockerfile
- `deploy.sh` - builds a Docker container via the `build.sh` script, pushes it to Dockerhub, then `ssh`s into the EC2 server. It then runs the `inside_aws_script.sh` script. 
- `inside_aws_script.sh` - Stops the existing Docker Container, pulls the new container from Dockerhub, then runs it. This script is *only* intended to be called as part of `deploy.sh`, it should not be run on its' own. 
-  `run_docker_locally` - a debugging tool. Rather than building and deploying to the remote server, this script runs a built Docker container on the local development computer. 
    - Note: You will need to create self-signed certificates on your computer for this to work. 
    - Additional warning: This script is not kept up to date with the current version of the Gateway script. It is only updated when the primary developer has a need to test on their local machine. Therefore, it may or may not be working. 


## API Keys - TODO UPDATE THIS FOR THE SERVERSIDE API KEYS

To use this application, you need to obtain API Keys for: 

- [Google Civic Information API](https://developers.google.com/civic-information/docs/using_api#APIKey)
- [ProPublica Congress API](https://www.propublica.org/datastore/api/propublica-congress-api)
- [ProPublica Campaign Finance API](https://www.propublica.org/datastore/api/campaign-finance-api)

Store them in a file called `.env` at the root of your cloned repository. It should look like this before you fill in your keys:
```
REACT_APP_GOOGLE_API_KEY=
REACT_APP_PROPUBLICA_CAMPAIGN_FINANCE_API_KEY=
REACT_APP_PROPUBLICA_CONGRESS_API_KEY=
```

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