import * as React from 'react'

class Signup extends React.Component {

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
    console.log(this.state)
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