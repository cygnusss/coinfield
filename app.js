// this is primarily for testing

const express = require("express")
const path = require("path")

const dist = path.resolve(__dirname, "./client/dist")

const app = express()

app.use(express.static(dist))

app.get("/*", (req, resp) => {
  resp.sendFile(path.resolve(__dirname, "./client/dist/index.html"))
})

app.listen(8080, "localhost", () => {
  console.log("listening on 8080")
})