import axios from "axios";
import React, { useState } from "react";

const Auth = () => {
  const [msg, setMsg] = useState("");
  const [headers, setHeaders] = useState("");
  const URL = "http://localhost:8080/api/auth";

  const click = async (e) => {
    e.preventDefault();
    try {
      const res = await axios.get(URL, { withCredentials: true });
      console.log("Response headers:", res.headers);
      console.log("Response body:", res.data);
      setHeaders(JSON.stringify(res.headers, null, 2));
      setMsg(res.data.email);
    } catch (error) {
      console.error("There was an error!", error);
      if (error.response) {
        console.log("Error headers:", error.response.headers);
        setHeaders(JSON.stringify(error.response.headers, null, 2));
      } else {
        setHeaders("No response from server");
      }
      setMsg("error in email");
    }
  };

  return (
    <div>
      <button onClick={click}>Click me</button>
      <h1>Email: {msg}</h1>
      <pre>Headers: {headers}</pre>
    </div>
  );
};

export default Auth;