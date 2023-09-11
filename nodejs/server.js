const express = require('express')
const server = express()
const port = 8000


server.use(express.json())

const tasks = []

server.get('/tasks', (req, res) => {
    res.json(tasks)
})

server.get('/healthz', (req, res) => {
    res.json("OK")
})

server.post('/tasks', (req, res) => {
    const {title, description} = req.body
     if(!title) {
        return res.status(400).json({error: 'Title is required!'})
     }

     if(!description) {
        return res.status(400).json({error: 'Description is required!'})
     }
     const newTask = {title, description}
     tasks.push(newTask)
     res.status(201).json(newTask)
})  

server.listen(port,  () => {
    console.log(`Server is running on http://localhost:${port}`)
})

module.exports = server; 