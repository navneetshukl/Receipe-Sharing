import axios from "axios";
import React, { useState } from "react";
import { Button } from "react-bootstrap";

const Auth = () => {
  const URL = "http://localhost:8080/api/auth";

  const [resp, setResp] = useState(null);
  const handleClick = () => {
    const resp = axios.get(URL);
    console.log(resp.data);
    setResp(resp.data);
  };
  return (
    <div>
      <Button onClick={handleClick}> Press me</Button>
      <p>{resp}</p>
    </div>
  );
};

export default Auth;
