# simple_bank
Zywa Initial Assignment -> a simple bank service


### Project Structure

**cmd/** for all publically exposed services
>> main.go, (api package) api/

**internal/** for all private services that are references by public services 
>> (db package) database/, (email package) email/, (pdf package) pdf/
