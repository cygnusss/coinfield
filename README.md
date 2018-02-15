# coinfield

Coinfield is a _**unique**_ application that let's you get most recent prices for cryptocurrencies (never seen before, huh?). This app will make the world a better place.

# Team
  - Currently only me

# Instructions
 ### Basics

 1. Clone the repo from https://github.com/cygnusss/coinfield
 ```sh
  git clone https://github.com/cygnusss/coinfield
 ```
 2. Move into the repo after it has been clone onto your machine by typing the following command:
 ```sh
  cd coinfield/
 ```
 3. Follow the next section
 
 ### Running the servers
 
 #### To run the go server
 From the root directory run:
 ```sh
   cd backend/src/main && go install && go run main.go model.go redis.go
 ```
 #### To run the express server
 From the root directory run:
  ```sh
    npm start-dev
  ```


# Requirements

* Node 9.x 
* go 1.9.4
