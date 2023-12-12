# Zywa Initial Assignment -> a simple bank service


### Project Structure

**cmd/** for all publically exposed services
>> main.go, (api package) api/

**internal/** for all private services that are references by public services 
>> (db package) database/, (email package) email/, (pdf package) pdf/

### API's supported
- /addtransaction is for adding a single transaction record to the database.csv
- /filtertransactions is for filtering the transactions based on (user_email, startDate, endDate) filtering.

#### API Calls in Action:
- <img width="607" alt="image" src="https://github.com/hktrib/simple_bank/assets/116051160/021bff9b-788c-4ecb-925a-f05726ff40eb">
- <img width="589" alt="image" src="https://github.com/hktrib/simple_bank/assets/116051160/90f016ec-3a8a-4f86-94ac-bee8146683f7">

#### How routes are wired on Golang!
- <img width="518" alt="image" src="https://github.com/hktrib/simple_bank/assets/116051160/12b56d0b-a4aa-48c0-8c4f-5fabfff40aef">



# Starting server
- Clone the repo, and `cd simple_bank/`
- Run `go mod tidy` to fetch dependencies
- Run `go build cmd/main.go` to create a binary
- **Run `./main` to start server**
