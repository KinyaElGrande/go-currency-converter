## Task

Implement a simple web service that converts currencies. This web service should only convert these currencies Nigerian 
**Naira(NGN), Ghanaian Cedis(GHS), and Kenyan Shillings (KSH)** from anyone to any other. This web service should not depend
on any external API for currency rates. A local conversion table should be used. Use the least number of dependencies as 
much as possible and any dependency used should be justifiable.

### Running the applicatiion

1. Open _cmd/server_ directory and run `go run main.go` command on your terminal
2. Visit your local server at `:8080/convert` and send your JSON request with the 
   following fields 
   `{
   "from": "CurrencyFrom",
   "to" : "CurrencyTo",
   "amount": 
   }`
   
**A request Example **

`{
"from": "KSh",
"to" : "ghs",
"amount": 800
}`

#### _NB_ : you can only convert 3 currencies (Nigerian Naira(NGN), Ghanaian Cedis(GHS), and Kenyan Shillings (KSH) ) against each other
