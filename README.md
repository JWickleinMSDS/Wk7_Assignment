***Instructions:*** 

To generate the original output from Week 3 "Testing Go For Statistics" do the following:  Execute the "anscombe_data.R" file from the terminal.  Execute the "anscombe_data.py" file from the terminal.  Execute the "anscombe_data.go" file from the terminal. Execute the "anscombe_data_test.go" file from the terminal using the command "go test".

***Automated Code Generation with Jennifer***

I created a new program called "Go_Jennifer.go" that uses the Jennifer Go library.  When this file is excecuted it creates a new program entitled "generated_program.go".  That program attempts to capture the output times from the Go, Python, and R files from the Instructions above.  (Note:  It fails and I have yet to determine why but I was able to successfully leverage Jennifer to automatically generate a new Go program).

***AI Assisted Programming***

I temporarily licensed GitHub CoPilot and tried it out in the file "CoPilot_expriment.go". CoPilot suggested several steps that resembled a file I already produced.  It suggested the same libraries to import and suggested many of the same lines of code.  I then gave it something new by asking it to create a new function called "igloo" that prints odd numbers from 1 to 10.  I then accepted it's suggestion and made it part of the code base.

***AI Generated Code***

I used ChatGPT 3.5 (https://chat.openai.com/).  I began by asking it to generate a go program that compared the outputs from the execution times required for the "Testing Go for Statistics" assignment. ChatGPT generated nearly working code very quickly.  I amended my prompt to ask it to use the actual file names my prior program (see Instructions above) generated.  I used quotes around the file names.  ChatGPT, however, interpreted the quotes as part of the file names.  This was easy enough for me to change myself.  I then got a couple different errors and simply pasted in the messages from the terminal directly into ChatGPT.  It came back with some useful suggestions but nothing that ultimately allowed me to successfully execute the code it produced.  The full exchange can be found in the file "ChatGPT_Convo.txt". The ChatGPT output can be found in "ChatGPT_Experiment.go".

***Overall Thoughts***  
I could see how using a Go code generator may have application.  That said, writing the go program using the Jennifer library to generate a go program was relatively slow.  I think in this case it may make more sense to simply generate a go program yourself.  That said, I do see how it could allow one to create more independent test cases.  But, overall, I don't see it as a time saving effort. 

GitHub CoPilot is likely one of the best ways to go.  It also allows for better code security in that you don't have your engineers risk sharing code with a more public LLM model like ChatGPT.  It allows engineers to write their code faster.  It effectively models pair programming.  

ChatGPT is also pretty good.  Especially when looking for an LLM to generate the bulk of your code.  That said, it's important to understand what the code that is being written actually does.  It can also require a good amount of prompt engineering to produce the output that you are looking for.  Additionally, I would be concerned about putting proprietary code out on a public LLM like ChatGPT.  

Overall, GitHub CoPilot is probably the best choice of the three for most professional settings. 
