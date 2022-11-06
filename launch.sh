#!/bin/bash

# Start the second process
./users &

# Start the first process
./jokes &
  
# Start the second process
./ratings &
  
# Wait for any process to exit
wait -n
  
# Exit with status of process that exited first
exit $?
