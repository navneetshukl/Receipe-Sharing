import React, { useState } from "react";
import { Form } from "react-bootstrap";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Button from "react-bootstrap/Button";
import axios from "axios";

const Description = () => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [ingredients, setIngredients] = useState("");

  const URL = "http://localhost:8080/api/receipe/add";

  // splitIngredients function will split the ingredients after comma
  const splitIngredients = () => {
    const arr = ingredients.split(",");
    return arr;
  };

  // handleFormSubmit function perform some action on submit of receipe
  const handleFormSubmit = (e) => {
    e.preventDefault();

    const ingre = splitIngredients();

    axios
      .post(
        URL,
        {
          name: name,
          description: description,
          ingredients: ingre,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      )
      .then((response) => {
        console.log("Response is ", response);
        alert("Receipe added successfully!");
      });

    setDescription("");
    setName("");
    setIngredients("");
  };

  return (
    <Form.Group>
      <Row
        style={{
          justifyContent: "center",
          margin: "auto",
          maxWidth: "100%",
        }}
      >
        <Col
          sm={5}
          style={{
            marginTop: "7vh",
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <h6>Receipe Name</h6>
          <Form.Control
            as="textarea"
            style={{
              height: "8vh",
              width: "40vw",
              resize: "none",
            }}
            onChange={(e) => {
              setName(e.target.value);
            }}
            value={name}
          />
        </Col>
      </Row>
      <Row
        style={{
          justifyContent: "center",
          marginTop: "12vh",
          marginLeft: "1px",
          marginRight: "1px",
          maxWidth: "100%",
        }}
      >
        <Col
          sm={5}
          style={{
            marginTop: "10vh",
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            margin: "auto",
          }}
        >
          <h6>Receipe Description</h6>
          <Form.Control
            as="textarea"
            style={{
              width: "45vw",
              height: "30vh",
              resize: "none",
            }}
            onChange={(e) => {
              setDescription(e.target.value);
            }}
            value={description}
          />
        </Col>
        <Col
          sm={4}
          style={{
            marginTop: "10vh",
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            margin: "auto",
          }}
        >
          <h6>Receipe Ingredients</h6>
          <Form.Control
            as="textarea"
            style={{
              width: "25vw",
              height: "30vh",
              resize: "none",
            }}
            onChange={(e) => {
              setIngredients(e.target.value);
            }}
            value={ingredients}
          />
        </Col>
      </Row>

      <Row
        style={{
          justifyContent: "center",
          marginTop: "5vh",
          maxWidth: "100%",
        }}
      >
        <Col
          style={{
            display: "flex",
            justifyContent: "center",
            gap: "2vw",
          }}
        >
          <Button variant="success" type="submit" onClick={handleFormSubmit}>
            Submit
          </Button>
          <>
            <p>upload Image</p>
            <input type="file"></input>
          </>

          <Button variant="danger">Back</Button>
        </Col>
      </Row>
    </Form.Group>
  );
};

export default Description;
