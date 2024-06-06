import React from "react";
import Counter from "./components/Counter";
import TextField from "./components/TextField";

function App(props) {
  console.log("APP rendered", props);
  return (
    <div>
      <h1>POC name: {props.Name}</h1>
      <div>
        <TextField />
      </div>
      <Counter defaultNum={props.InitialNumber} />
    </div>
  );
}

export default App;
