import React from "react";
import Col from "react-bootstrap/Col";
import Image from "react-bootstrap/Image";

const Dish_Image = () => {
  return (
    <Col className="d-flex justify-content-center align-items-center" style={{ height: "150px" }}>
      <Image src="https://via.placeholder.com/300x200" rounded style={{ maxHeight: "100%", maxWidth: "100%" }} />
    </Col>
  );
};

export default Dish_Image;
