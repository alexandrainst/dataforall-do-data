### Token exchange
A superuser token must be exchanged and embedded in python script to call pocketbase.  
  
Alternatives: 
- Auth with password for every python request (https://pocketbase.io/docs/api-records/#auth-with-password) using predefined superuser.
- Reverse proxy protect certain api endpoints.
- Use `JSON` data source
    - Implement custom basic auth endpoint for data
- ...?

### Query restrictions
Restrict redash queries to only be callable by members of the organization/group which should have access to see them.  

Ways to do this?: 
- Keep query level api keys in pocketbase
    - New queries must be created for each organization
        - Could this be done through redash API?
- ...?

### Redash authentication
How do we ensure users authenticated and not just in possesion on api_key?
- We cannot revoke authorization?
- Is this necessary?
- Possible to proxy requests through pocketbase for authenticated users

### Validate imported Klimapas data
More validation based on current use case.
Such as:
- Only allow imports for a whole year
- No upload of same year twice
- ... 

### A sensible reverse proxy setup?
Client has to connect to both Pocketbase API and Redash to fetch data
