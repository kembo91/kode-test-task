import { useState } from "react";
import {Button, Form} from "react-bootstrap"
import axios from "axios"

function App() {
  const [signupUser, setSignupUser] = useState("")
  const [signupPass, setSignupPass] = useState("")
  const [signinUser, setSigninUser] = useState("")
  const [signinPass, setSigninPass] = useState("")
  const [srvResp, setSrvResp] = useState("")

  const submitSignin = (e) => {
    e.preventDefault()
    axios.post('/api/signin', {
      Username: signinUser,
      Password: signinPass
    }).then((resp) => {setSrvResp(resp.data)})
      .catch((err) => {setSrvResp(err)})
  }

  const submitSignup = (e) => {
    e.preventDefault()
    axios.post('/api/signup', {
      Username: signupUser,
      Password: signupPass
    }).then((resp) => {setSrvResp(resp.data)})
      .catch((err) => {setSrvResp(err)})
  }

  const logout = (e) => {
    e.preventDefault()
    axios.get('/api/logout')
      .then(resp => setSrvResp(resp.data))
      .catch(err => setSrvResp(err))
  }

  return (
    <div className="App">
      <div className="d-flex flex-column justify-content-center">
        <div className="p-2">
          <Form.Control value={srvResp} placeholder="server responses will be here" plaintext readOnly />
        </div>
        <div className="p-2">
          <h3>Sign up form</h3>
          <Form onSubmit={submitSignin}>
            <Form.Group controlId="signup.Email">
              <Form.Label>Username</Form.Label>
              <Form.Control onChange={(e) => setSigninUser(e.target.value)} type="text" placeholder="Mr. Evans Gambit" />
            </Form.Group>
            <Form.Group controlId="signup.Password">
              <Form.Label>Password</Form.Label>
              <Form.Control onChange={(e) => setSigninPass(e.target.value)} type="password" placeholder="*****" />
            </Form.Group>
            <Button variant="primary" type="submit">
              Sign Up
            </Button>
          </Form>
          <h3>Sign in form</h3>
          <Form onSubmit={submitSignup}>
            <Form.Group controlId="signin.Email">
              <Form.Label>Username</Form.Label>
              <Form.Control onChange={(e) => setSignupUser(e.target.value)} type="text" placeholder="Mr. Evans Gambit" />
            </Form.Group>
            <Form.Group controlId="signin.Password">
              <Form.Label>Password</Form.Label>
              <Form.Control onChange={(e) => setSignupPass(e.target.value)} type="password" placeholder="*****" />
            </Form.Group>
            <Button variant="primary" type="submit">
              Sign In
            </Button>
          </Form>
          <Button variant="primary" onClick={logout}>Log out</Button>
        </div>
      </div>
    </div>
  );
}

export default withCookies(App);
