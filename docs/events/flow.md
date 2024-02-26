
# Flow of messages

The flagbearers will be sending messages to a queue implementation



flagbearers -> queue
queue -> scorers
scorers will score the services
the scorers will send the results via queue back to flagbearers -> queue -> flagbearer
the flagbearer will then store that check safely using an update key given to the service and note with scorer it gave it to
