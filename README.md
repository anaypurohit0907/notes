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

