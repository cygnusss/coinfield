import * as React from 'react'
import { Switch, Route } from 'react-router-dom'

import NavBar from '../NavBar/NavBar'
import Home from '../Home/Home'
import Signup from '../Signup/Signup'

const App = () => {
  return (
    <div>
      <NavBar />
      <Switch>
        <Route exact path="/" component={Home}/>
        <Route path="/signup" component={Signup}/>
      </Switch>
    </div>
  )
}

export default App