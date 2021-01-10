import { useState } from "react";
import {Button, Card, Col, Row, Container, Form} from "react-bootstrap"
import axios from "axios"

function App() {
  const [signupUser, setSignupUser] = useState("")
  const [signupPass, setSignupPass] = useState("")
  const [signinUser, setSigninUser] = useState("")
  const [signinPass, setSigninPass] = useState("")
  const [srvResp, setSrvResp] = useState("")
  const [insertAnagram, setInsertAnagram] = useState("")
  const [retrieveAnagram, setRetrieveAnagram] = useState("")

  const submitInsert = (e) => {
    e.preventDefault()
    axios.post('/api/anagram/insert', {
      Query: insertAnagram
    }).then((resp) => {setSrvResp(JSON.stringify(resp.data))})
      .catch((err) => setSrvResp(JSON.stringify(err.response.data)))
  }

  const submitRetrieve = (e) => {
    e.preventDefault()
    axios.post('/api/anagram/retrieve', {
      Query: retrieveAnagram
    }).then((resp) => setSrvResp(JSON.stringify(resp.data)))
      .catch((err) => setSrvResp(JSON.stringify(err.response.data)))
  }

  const submitRetrieveAll = (e) => {
    e.preventDefault()
    axios.get('/api/anagram/retrieve')
      .then((resp) => setSrvResp(JSON.stringify(resp.data)))
      .catch((err) => setSrvResp(JSON.stringify(err.response.data)))
  }

  const submitSignin = (e) => {
    e.preventDefault()
    axios.post('/api/signin', {
      Username: signinUser,
      Password: signinPass
    }).then((resp) => {setSrvResp(JSON.stringify(resp.data))})
      .catch((err) => {setSrvResp(JSON.stringify(err.response.data))})
  }

  const submitSignup = (e) => {
    e.preventDefault()
    axios.post('/api/signup', {
      Username: signupUser,
      Password: signupPass
    }).then((resp) => {setSrvResp(JSON.stringify(resp.data))})
      .catch((err) => {setSrvResp(JSON.stringify(err.response.data))})
  }

  const signout = (e) => {
    e.preventDefault()
    axios.get('/api/signout')
      .then(resp => setSrvResp(JSON.stringify(resp.data)))
      .catch(err => setSrvResp(JSON.stringify(err.response.data)))
  }

  return (
    <div className="App">
        <div className="d-flex justify-content-center p-2">
          
            <Form.Control value={srvResp} placeholder="server responses will be here" plaintext readOnly />
          
        </div>
        <Container>  
        <Row>
          <Col>
          <Card style={{width: '18rem'}}>
            <h3>Sign up form</h3>
            <Form onSubmit={submitSignup}>
              <Form.Group controlId="signup.Email">
                <Form.Label>Username</Form.Label>
                <Form.Control onChange={(e) => setSignupUser(e.target.value)} type="text" placeholder="EvansGambit" />
              </Form.Group>
              <Form.Group controlId="signup.Password">
                <Form.Label>Password</Form.Label>
                <Form.Control onChange={(e) => setSignupPass(e.target.value)} type="password" placeholder="********" />
              </Form.Group>
              <Button variant="primary" type="submit">
                Sign Up
              </Button>
            </Form>
          </Card>
          </Col>
          <Col>
          <Card style={{width: '18rem'}}>
            <h3>Sign in form</h3>
            <Form onSubmit={submitSignin}>
              <Form.Group controlId="signin.Email">
                <Form.Label>Username</Form.Label>
                <Form.Control onChange={(e) => setSigninUser(e.target.value)} type="text" placeholder="EvansGambit" />
              </Form.Group>
              <Form.Group controlId="signin.Password">
                <Form.Label>Password</Form.Label>
                <Form.Control onChange={(e) => setSigninPass(e.target.value)} type="password" placeholder="********" />
              </Form.Group>
              <Button variant="primary" type="submit">
                Sign In
              </Button>
            </Form>
          </Card>
          </Col>
          <Col>
          <Button variant="primary" onClick={signout}>Sign out</Button>
          </Col>
          </Row>
          <Row>
          <Col>
          <Card style={{width: '18rem', marginTop: 20}}>
          <h3>Insert anagram</h3>
            <Form onSubmit={submitInsert}>
              <Form.Group controlId="insertAnagram">
                <Form.Label>Anagram</Form.Label>
                <Form.Control onChange={(e) => setInsertAnagram(e.target.value)} type="text" placeholder="dog"/>
              </Form.Group>
              <Button variant="primary" type="submit">
                Insert
              </Button>
            </Form>
          </Card>
          </Col>
        <Col>
          <Card style={{width: '18rem', marginTop: 20}}>
          <h3>Retrieve anagram</h3>
            <Form onSubmit={submitRetrieve}>
              <Form.Group controlId="retrieveAnagram">
                <Form.Label>Anagram</Form.Label>
                <Form.Control onChange={(e) => setRetrieveAnagram(e.target.value)} type="text" placeholder="god" />
              </Form.Group>
              <Button variant="primary" type="submit">
                Retrieve
              </Button>
            </Form>
          </Card>
          </Col>
        <Col>
          <Button variant="primary" type="submit" onClick={submitRetrieveAll}>
            Retrieve All
          </Button>
          </Col>
        </Row>
        </Container>
      </div>
    
  );
}

export default App;
