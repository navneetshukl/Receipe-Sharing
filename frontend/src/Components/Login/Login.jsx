import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Container from "react-bootstrap/Container";
import Button from "react-bootstrap/Button";
import axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css"; // Import CSS for toast

const Login = () => {
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");

  const URL = "http://localhost:8080/api/user/login";

  const handleFormSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await axios.post(
        URL,
        {
          email: email,
          password: password,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      // Logging response data and status code
      console.log("Response data:", res.data);
      console.log("Status code:", res.status);

      // Show success toast
      if (res.status === 200) {
        toast.success(res.data.message || "user login successfully!", {
          position: "top-right",
          autoClose: 10000,
          closeOnClick: true,
          pauseOnHover: true,
        });
      }
    } catch (error) {
      // Error handling
      console.error(
        "Error during form submission:",
        error.response ? error.response.data : error.message
      );

      // If the error response exists, log the status code
      if (error.response) {
        console.error("Error status code:", error.response.status);
        if (error.response.status === 500) {
          toast.error(error.response.data.error || "something went wrong", {
            position: "top-right",
            autoClose: 10000,
            closeOnClick: true,
            pauseOnHover: true,
          });
        }
      }
    }

    // Clear form fields
    // setEmail("");
    // setPassword("");
  };

  return (
    <Container style={{ marginTop: "5vh" }}>
      <h4 style={{ textAlign: "center" }}>Login</h4>
      <Container style={{ marginTop: "15vh", width: "50vw" }}>
        <Form onSubmit={handleFormSubmit}>
          {/* <Form.Group>
            <Form.Label>Name</Form.Label>
            <Form.Control
              type="text"
              onChange={(e) => {
                setName(e.target.value);
              }}
              value={name}
            />
          </Form.Group> */}
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
          >
            Submit
          </Button>
        </Form>
      </Container>
      <ToastContainer />

    </Container>
  );
};

export default Login;
