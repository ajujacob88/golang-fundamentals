/*
Part of exercise file for go lang course at
https://web.learncodeonline.in
*/

const express = require('express')
const app = express()
const port = 3000  // we can free to change to 3000, 6000, 7000, 8000 etc

app.use(express.json()); 
app.use(express.urlencoded({extended: true}));

app.get('/', (req, res) => {
  res.status(200).send("Welcome to Aju Jacob  server")
})

app.get('/get', (req, res) => {
    res.status(200).json({message: "Hello from learnCodeonline.in"})
  })


app.post('/postaju', (req, res) => {
    let myJson = req.body;      // your JSON
	
	res.status(200).send(myJson);
})

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
})
  

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})