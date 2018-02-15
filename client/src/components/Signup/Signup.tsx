import * as React from "react"
import axios from "axios"

class Signup extends React.Component <any, any> {

  constructor(props: any) {
    super(props)

    this.handleChange = this.handleChange.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this)

    this.state = {
      firstName: "",
      lastName: "",
      email: "",
      password: "",
    }
  }

  handleChange(e: any) {
    const q: string = e.target.value
    const t: string = e.target.placeholder

    switch(t) {
      case "First Name":
        this.setState({
          firstName: q
        })
        break
      case "Last Name":
        this.setState({
          lastName: q
        })
        break
      case "Email":
        this.setState({
          email: q
        })
        break
      case "Password":
        this.setState({
          password: q
        })
        break
    }

  }

  handleSubmit() {
    const body = JSON.stringify({
      firstname: this.state.firstName,
      lastname: this.state.lastName,
      email: this.state.email,
      password: this.state.password,
    })

    axios.post('/signup', body)
    // .then(function (response) {
    //   console.log(response);
    // })
    // .catch(function (error) {
    //   console.log(error);
    // });
  }

  render() {
    return (
      <div className="signup_box">
        
        <div className="name_field">
          <label htmlFor="username">Name</label>
          <div>
            <input type="text" placeholder="First Name" onChange={this.handleChange}/>
            <input type="text" placeholder="Last Name" onChange={this.handleChange}/>
          </div>
        </div>
  
        <div>
          <label htmlFor="email">Email</label>
          <div>
            <input type="text" placeholder="Email" onChange={this.handleChange}/>
          </div>
        </div>
  
        <div>
          <label htmlFor="password">Password</label>
          <div>
            <input type="text" placeholder="Password" onChange={this.handleChange}/>
          </div>
        </div>
  
        <div>
          <input type="checkbox"/>
          <span>
            I certify that I am 18 years of age or older, and I agree to the User Agreement and Privacy Policy.
          </span>
        </div>
  
        <div>
          <input type="submit" value="Create an Account" onClick={this.handleSubmit}/>
        </div>
  
      </div>
    )
  }

}

export default Signup