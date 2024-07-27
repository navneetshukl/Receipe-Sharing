import React from "react";
import { Form } from "react-bootstrap";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Button from "react-bootstrap/Button";

const Description = () => {
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
          <Button variant="success" type="submit">
            Submit
          </Button>
          <>
            <p>upload Image</p>
            <input type="file"></input>
          </>
        </Col>
      </Row>
    </Form.Group>
  );
};

export default Description;
