# Developer notes
### Document purpose - keeping track of to-do lists, shortcuts, and other things that will be helpful during the development process 
#### Tasks for the future (ie to do lists) now live in the trello, and will not be added here

## To Do List
* Visualizations I'm contributing backend to
    + Votes on Bills by Representative
        > Details: representative votes on bill and category of bill is different. This would be two api calls, so, it should be in the backend. Our goal is a list of bills, with category and how they voted

        > From: Propublica

        >Status: ???

        > Via: Custom Back End
    + Top Sectors and Pacs
        > Top sectors and Pacs - Jay

        > From: Open Secrets

        > Status: As done as it's going to be until I have dummy data for Bills by Representatives 

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
* Finance - sloppy code or unfinished things to fix later
    + in insert.go, instead of putting current date in ftv_updated in a reasonable format, I just did time.now()
    + CRP IDs - find a better way to get them, cause my code currently breaks on ids for people who are no longer in office
    + handler to serve the data
* Daily job to call all relevant APIs
* Load data into a SQL relational database
    > Use the categories of propublica to open secrets to compare bills to finance. Apparently there’s a list on congress.gov
* Create endpoints: 
    + Given a representative Return Open Secrets top 10 sector information
        > Requires a CRP ID for that representative - comes from Propublica
    + Given a representative name, return a list of bills they’ve voted on, as well as the category for those bills
        > Propublica

        > We specifically want the vote for the final approval of the bill, not the vote for adjourning for the day or other garbage


### Bills endpoint rambling and thoughts
Okay, so
I need to have several things
* API call to the propublica database that I want - which?? (chase this down soon) 
    +  Call by bill, get category
* API call to the Opensecrets database
    + Call by person, get recent bills voted on? 
* I think I'm slightly confused about what data comes from where - I want bill information, I want opensecrets information on the industry related to this, i want bills voted on by politician
* Goal: correlate donor industries (opensecrets) with bill categories (propublica), and use this to make a visualization
* Putting this together will require looking at a list of donor industries and a list of bill categories

### Correlation 
Lets correlate "Finance and Financial Sector" 
there's more opensecrets categories than legislative areas
Plan: prefix of the sector for each thing
Taxation -> ???
Health -> H, Health
Government Operations and Politics -> ????
Armed Forces and National Security -> D, Defense
Congress -> ??
International Affairs -> Q, Ideological/Single-Issue
Foreign Trade and International Finance -> N, Misc Buisness
Public Lands and Natural Resources -> E, Energy & Natural Resources
Crime and Law Enforcement -> ? 
Transportation and Public Works -> M, Transportation
Social Welfare -> ? 
Education -> W, Other
Energy ->  E, Energy & Natural Resources
Agriculture and Food -> A, Agribuisness 
Labor and Employment -> P, Labor
Finance and Financial Sector -> F, Finance, Insurance & Real Estate
Enviromental Protection ->  Q, Ideological/Single-Issue
Economics and Public Finance -> F, Finance, Insurance & Real Estate
Commerce -> N, Misc Buisness
Science, Technology, Communications -> B, Communications/Electronics
Immigration -> Q, Ideological/Single-Issue
Housing and Community Devolopment -> 
Law -> K, Lawyers & Lobbiests
Water Resources Development ->  E, Energy & Natural Resources
Civil Rights and Liberties, Minority Issues -> Q, Ideological/Single-Issue
Native Americans -> Q, Ideological/Single-Issue
Emergency Management -> 
Families -> W, Other
Animals -> A, Agribuisness
Arts, Culture, Religion -> W, Other
Sports and Recreation -> N, Misc Buisness
Social Sciences and History ->  W, Other

missing: construction sector


## Gameplan
* The opensecrets only one for the top 10 sector information seems easiest
* I'm going to do that one first
* Then, once I've figured that out, I'll start looking into the one using both propublica and opensecrets
* First step, research 
    * the endpoint to fetch from: 
    * notes on the fetching process: 

## Tips, Code Snippets, and Shortcuts
* In Visual Studio Code, `ctrl+k v` will render the preview of this page and any other markdown page