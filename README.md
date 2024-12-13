# notes
a simple, lightweight command-line application to help you record, organize, and recall new terminal commands, Vim key bindings, and other useful shortcuts you discover over time. With notes, you can quickly save commands along with descriptions, tags, and categories so you can easily reference them when needed.


Available Commands: 


  help       : Help about any command
  
  init    :    initializes the database and get everything ready
  
  note     :   Note a command you just learned along with a small description of what it does .
            
  Note new : creates new note

  Note list : lists all notes
  
  Flags:
  
   -n: new note
   
   -c: category of new note
   
   -d: definition of new note
   
   -s: use in list to search by a word in definition       

usage: ./go-cli note -n "word" -c "category" -d "definition"

search flags: 

-c: search by category

-d: search by definition 

usage: ./go-cli note list -c or -d "search word "

edit : this is nowhere close to being finished, I am working on basic functionalities right now like autocomplete for the commands, fuzzy finding perhaps.

here's the roadmap I've come up with for now :

1. make all the basic functionalities like search,save and various usefull flags.

2. make the go binary into a CLI command so we don't have to do a ./filename note to run and can just do a note directly

3. try to put that on package managers and other places ppl can download from, for this I might have to write a init shell file(I'm not sure)

4. Change the local SQL db to a cloud one(I want it to be flexible where if anyone wants to save third notes locally they can, otherwise they can do it on the cloud by default)

5. make a mobile app so if you ever read or learn anything and want to write it down. this will be synced to the CLI tool too obviously. (I will have to create a user based system for this but don't want to use oauth or any third party so I will have a unique I'd per user that they will input on the app to log into their notes.)


this is as far as I have thought about this project, I don't know if this is the best route or the best way to do this but this is what I have come up with and I'll adapt and change as I grow and as needed. 
also I want to make this a neovim plugin too, that's probably the last step. 