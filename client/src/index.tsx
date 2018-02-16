import * as React from "react"
import * as ReactDOM from "react-dom"
import {BrowserRouter} from 'react-router-dom'
import { Provider } from 'react-redux'
import { createStore } from "redux";

import App from "./components/App/App"

ReactDOM.render((
  <BrowserRouter>
    <App />
  </BrowserRouter>
), document.getElementById('root'))