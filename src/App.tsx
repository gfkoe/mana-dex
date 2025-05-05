import { useState } from "react";
import "./App.css";
import Header from "./components/Header";
import ManaForm from "./components/ManaForm";
import OutputBox from "./components/OutputBox";

function App() {
  const [output, setOutput] = useState<string>("");
  return (
    <>
      <Header />
      <main>
        <div className="flex justify-center items-stretch">
          <div className="w-1/2">
            <ManaForm setOutput={setOutput} />
          </div>
          <div className="w-1/2 flex">
            <OutputBox text={output} />
          </div>
        </div>
      </main>
    </>
  );
}

export default App;
