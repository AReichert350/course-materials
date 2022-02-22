In this directory:  
go build main.go  
SHODAN_API_KEY=YOURAPIKEY ./main <S | Q>  // Use S to run a host/search API call or use Q to run a /query API call

If you select S:
- You will be prompted to enter a query.
- Press Enter.
- Press "Y" and Enter for each additional page of results you'd like to see. Press anything else and Enter to exit the program.
- Example flow: 
    - SHODAN_API_KEY=63yhTuV2IAenrJNUZVO2x4a5vYe4KaTJ ./main S
    - Apache
    - N  

If you select Q:
- The /query search will run
- Press "Y" and Enter for each additional page of results you'd like to see. Press anything else and Enter to exit the program.
- Example flow:
    - SHODAN_API_KEY=63yhTuV2IAenrJNUZVO2x4a5vYe4KaTJ ./main Q
    - Y
    - Y
    - N
