// this is primarily for testing

const express = require("express")
const path = require("path")

const dist = path.resolve(__dirname, "./client/dist/")

const app = express()
const port = process.env.PORT || 8080

app.use("/dist/", express.static(dist))

app.get("/*", (req, resp) => {
  resp.sendFile(path.resolve(__dirname, "./client/index.html"))
})

app.listen(port, () => {
  console.log("Running express app on port:", port)
})