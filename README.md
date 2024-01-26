# sleeping_barber_dilemma
**Problem description:**
This is a simple demonstration of how to solve the Sleeping Barber dilemma, a classic computer science problem which illustrates the complexities that arise when there are multiple operating system processes. Here, we have a finite number of barbers, a finite number of seats in a waiting room, a fixed length of time the barbershop is open, and clients arriving at (roughly) regular intervals. When a barber has nothing to do, he or she checks the waiting room for new clients, and if one or more is there, a haircut takes place. Otherwise, the barber
goes to sleep until a new client arrives. 

**So the rules are as follows:**


  • if there are no customers, the barber falls asleep in the chair
  
  • a customer must wake the barber if he is asleep
  
  • if a customer arrives while the barber is working, the customer leaves if all chairs are occupied and sits in an empty chair if it's available
  
  • when the barber finishes a haircut, he inspects the waiting room to see if there are any waiting customers and falls asleep if there are none
  
  • shop can stop accepting new clients at closing time, but the barbers cannot leave until the waiting room is empty
  
  • after the shop is closed and there are no clients left in the waiting area, the barber goes home

**Step 1**
  * clone the repository **git clone https://github.com/hariprathap-hp/sleeping_barber_dilemma.git**
  * run **make run** in your terminal

**Step 2**
  * Pull the images cretaed from dockerHub using the command **docker pull deadpoet/hariprathap-sleeping-barber:v1**
  * Then run the command **docker run docker run deadpoet/hariprathap-sleeping-barber:v1**

**Sample Output**
  * please check the file "sample_output.txt" to see the sample output

