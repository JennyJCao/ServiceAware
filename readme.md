
# ServiceAware

The project is a monitoring service. It can monitor the changes of service status in real time. 

Besides, it can send emails or SMS to notify the changes of a specific service. 

## Build

Build in the normal way on Mac:
1. make sure ipe is running
  
    ```
    cd ipe
    ./ipe
    ```

2. create database using postgreSQL, which name is vigilate. 

    ```
    // create tables.
    soda migrate 
    // insert a user into the table users, and use it to sign in
    ```
    
3. run the service

    ```
    cd ServiceAware
      ./run.sh
    ```
  
4. visit localhost:4000
