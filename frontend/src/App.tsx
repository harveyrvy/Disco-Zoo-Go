import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);
  const [animalsList, setAnimalsList] = useState<string[]>([]);
  const [tiles, setTiles] = useState();

  const addRabbit = () => {
    setAnimalsList((animalsList) => [...animalsList, "Rabbit"]);
  };
  const calculate = async () => {
    console.log("called");
    const response = await fetch("http://localhost:8080/postcalculate/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ animals: animalsList }),
    });
    console.log(response);
    const data = await response.json();
    setTiles(data.tiles);
  };

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <button onClick={addRabbit}>add Rabbit</button>
      <p>Animal List = {animalsList}</p>
      <button onClick={calculate}>send request</button>
      <p>{tiles}</p>
    </>
  );
}

export default App;
