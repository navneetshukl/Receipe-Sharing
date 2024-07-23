import React from "react";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Dish_Image from "./Dish_Image";

const Home = () => {
  return (
    <Container style={{ border: "3px solid red" }}>
      <Row style={{ height: "180px" }}>
        <Col sm={5}>
          <Dish_Image />
        </Col>
      </Row>
      <Row style={{ height: "400px" }}>
        <Col sm={8} style={{ backgroundColor: "gray" }}>
          Receipe Section
        </Col>
        <Col sm={4} style={{ backgroundColor: "blue" }}>
          Ingredient Section
        </Col>
      </Row>
    </Container>
  );
};

export default Home;
