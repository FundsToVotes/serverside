# Developer notes
### Document purpose - keeping track of to-do lists, shortcuts, and other things that will be helpful during the development process 

## To Do List
* Serverside stuff - not my problem
* Basic setup - done
* Visualizations I'm contributing backend to
    + Votes on Bills by Representative
        > Details: representative votes on bill and category of bill is different. This would be two api calls, so, it should be in the backend. Our goal is a list of bills, with category and how they voted

        > From: Propublica

        >Status: ???

        > Via: Custom Back End
    + Top Sectors and Pacs
        > Top sectors and Pacs - Jay

        > From: Open Secrets

        > Done: Not Done 

        > Via: Custom Back end
* Finance API plan: 
    + Inside APi folder 
        > function that calls the relevant APIs and spits out data
            >> Inserts new things in the database
            >> Updates not-new things
        > Function for emptying database (developer use only) 
    + Inside gateway
        > Gethandler for getting Data from the database
    + On server
        > Chron job to call the API .exe every day
* Daily job to call all relevant APIs
* Load data into a SQL relational database
    > Use the categories of propublica to open secrets to compare bills to finance. Apparently there’s a list on congress.gov
* Create endpoints: 
    + Given a representative Return Open Secrets top 10 sector information
        > Requires a CRP ID for that representative - comes from Propublica
    + Given a representative name, return a list of bills they’ve voted on, as well as the category for those bills
        > Propublica

        > We specifically want the vote for the final approval of the bill, not the vote for adjourning for the day or other garbage
* Cleanup and Optimization
    + Opensecrets top 10: 
        - Fetch CRP ID programatically rather than manual .csv file creation

## Gameplan
* The opensecrets only one for the top 10 sector information seems easiest
* I'm going to do that one first
* Then, once I've figured that out, I'll start looking into the one using both propublica and opensecrets
* First step, research 
    * the endpoint to fetch from: 
    * notes on the fetching process: 

## Tips, Code Snippets, and Shortcuts
* In Visual Studio Code, `ctrl+k v` will render the preview of this page and any other markdown page