# A basic guide to the steps, since I do this so often 

1. x___Create an Instance using the "launch Instance" wizard. Choose the AWS version 2, don't bother with any other settings, and add/create a keypair
2. x___Make sure the keypair is shared with everyone you're collaborating with, and is saved
3. x___Create an SSH shortcut to the remote server
    + `ssh-add new_funds_to_votes_aws.pem`
    + `alias sshwebc=ssh ec2-user@ec2-54-218-168-107.us-west-2.compute.amazonaws.com`
4. x___Test that I can ssh in
5. x___Install certbot and docker on the remote server
    + https://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-basics.html#install_docker
    + https://drstearns.github.io/tutorials/https/#secrunletsencryptonamazonlinux2
4. x___Allow HTTP, HTTPs traffic in via the security groups
5. x___Configure DNS and connect to namecheap
7. x___create certificates 
    + https://drstearns.github.io/tutorials/https/#secrunletsencryptonamazonlinux2
8. x___Copy these files into my development directory
10. Create a docker hub account if necessary
9. run and deploy the Hello Handler to test it
10. Refine these instructions as necessary for future me's benefit