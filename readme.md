
# ServiceAware

The project is a monitoring service. It can monitor the changes of service status in real time. 

Besides, it can send emails or SMS to notify the changes of a specific service.

The system can monitor services in real time by using websocket.

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

## Screenshot

1. The overview page shows the count of services in different status, and all the hosts.

![img_1.png](img_1.png)

2. The Login and Logout page

![img_16.png](img_16.png)

![img_17.png](img_17.png)

3. Adding the host which you want to monitor

![img_3.png](img_3.png)

4. Monitoring the specific services on the host.

![img_4.png](img_4.png)

5. The event page shows the changes of service status.

![img_9.png](img_9.png)

6. The schedule page displays all the services we are monitoring when the monitoring button opens.

![img_10.png](img_10.png)

7. The settings page allows us to get notifications about the changes of service status through email or SMS.

![img_11.png](img_11.png)

   Email notifications (mailhog):

   ![img_12.png](img_12.png)

   SMS notifications (twilio):

   ![img_14.png](img_14.png)

8. The users page show all the users.

![img_15.png](img_15.png)


## Real-time Monitoring System: 

There are four status of services, including healthy, warning, problem, and pending. When monitoring, the changes of status will be displayed **automatically**.

eg. The Host Page: /admin/host/hostID

Before we open the service on port 8080, these two services are in the problem status:

![img_6.png](img_6.png)

After we open this service, their status changes to be in healthy status:

![img_7.png](img_7.png)

![img_8.png](img_8.png)
     
   
Besides, All the pages are in real time, including the counts of services with different status on overview page, events and schedule.
   
   