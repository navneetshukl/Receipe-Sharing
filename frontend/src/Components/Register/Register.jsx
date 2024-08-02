import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Container from "react-bootstrap/Container";
import Button from "react-bootstrap/Button";
import axios from "axios";

const Register = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [mobile, setMobile] = useState("");

  const URL = "http://localhost:8080/api/user/register";

  const handleFormSubmission = async (e) => {
    e.preventDefault();

    try {
      const resp = await axios.post(
        URL,
        {
          name: name,
          email: email,
          password: password,
          mobile: mobile,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      console.log("response is ", resp.data);
    } catch (error) {
      // Error handling
      console.error(
        "Error during form submission:",
        error.response ? error.response.data : error.message
      );

      // If the error response exists, log the status code
      if (error.response) {
        console.error("Error status code:", error.response.status);
      }
    }

    setEmail("");
    setName("");
    setPassword("");
    setMobile("");
  };
  return (
    <Container style={{ marginTop: "5vh" }}>
      <h4 style={{ textAlign: "center" }}>Register</h4>
      <Container style={{ marginTop: "15vh", width: "50vw" }}>
        <Form onSubmit={handleFormSubmission}>
          <Form.Group>
            <Form.Label>Name</Form.Label>
            <Form.Control
              type="text"
              value={name}
              onChange={(e) => {
                setName(e.target.value);
              }}
            />
          </Form.Group>
          <Form.Group>
            <Form.Label>Email</Form.Label>
            <Form.Control
              type="email"
              value={email}
              onChange={(e) => {
                setEmail(e.target.value);
              }}
            />
          </Form.Group>
          <Form.Group>
            <Form.Label>Password</Form.Label>
            <Form.Control
              type="password"
              value={password}
              onChange={(e) => {
                setPassword(e.target.value);
              }}
            />
          </Form.Group>
          <Form.Group>
            <Form.Label>Mobile-Number</Form.Label>
            <Form.Control
              type="text"
              value={mobile}
              onChange={(e) => {
                setMobile(e.target.value);
              }}
            />
          </Form.Group>
          <Button
            variant="success"
            style={{ marginTop: "3vh" }}
            type="submit"
            size="lg"
          >
            Submit
          </Button>
        </Form>
      </Container>
    </Container>
  );
};

export default Register;
