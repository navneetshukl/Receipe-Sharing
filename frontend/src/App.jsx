import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import Home from "./Components/Receipe/Home";
import Descripition from "./Components/Add-Receipe/Descripition";
import Login from "./Components/Login/Login";

function App() {
  return (
    <>
      {/* <Home/> */}
      {/* <Descripition /> */}
      <Login/>
    </>
  );
}

export default App;
