import axios from "axios";
import React, { useState } from "react";
import { Button } from "react-bootstrap";

const Auth = () => {
  const URL = "http://localhost:8080/api/auth";

  const [resp, setResp] = useState(null);
  const handleClick = async () => {
    try {
      // Set withCredentials to true to ensure cookies are sent
      const response = await axios.get(URL, { withCredentials: true });
      console.log(response.data);
      setResp(response.data);
    } catch (error) {
      console.error("Error fetching data:", error.response ? error.response.data : error.message);
      setResp("Error fetching data");
    }
  };

  return (
    <div>
      <Button onClick={handleClick}>Press me</Button>
      <p>{resp}</p>
    </div>
  );
};

export default Auth;
