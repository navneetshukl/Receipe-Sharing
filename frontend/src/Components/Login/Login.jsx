import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Container from "react-bootstrap/Container";
import Button from "react-bootstrap/Button";
import axios from "axios";

const Login = () => {
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [response, setResponse] = useState(null);

  const URL = "http://localhost:8080/api/user/login";

  const handleFormSubmit = (e) => {
    e.preventDefault();
    axios
      .post(
        URL,
        {
          name: name,
          email: email,
          password: password,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      )
      .then((res) => {
        console.log("Response is ", res);
      });
    setResponse(res);
  };
  return (
    <Container style={{ marginTop: "5vh" }}>
      <h4 style={{ textAlign: "center" }}>Login</h4>
      <Container style={{ marginTop: "15vh", width: "50vw" }}>
        <Form>
          <Form.Group>
            <Form.Label>Name</Form.Label>
            <Form.Control
              type="text"
              onChange={(e) => {
                setName(e.target.value);
              }}
              value={name}
            />
          </Form.Group>
          <Form.Group>
            <Form.Label>Email</Form.Label>
            <Form.Control
              type="email"
              onChange={(e) => {
                setEmail(e.target.value);
              }}
              value={email}
            />
          </Form.Group>
          <Form.Group>
            <Form.Label>Password</Form.Label>
            <Form.Control
              type="password"
              onChange={(e) => {
                setPassword(e.target.value);
              }}
              value={password}
            />
          </Form.Group>
          <Button
            variant="success"
            style={{ marginTop: "3vh" }}
            type="submit"
            size="lg"
            onClick={handleFormSubmit}
          >
            Submit
          </Button>{" "}
        </Form>
      </Container>
    </Container>
  );
};

export default Login;
