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
      <main className="gap-y-6">
        <div className="flex justify-center items-stretch gap-x-2">
          <div className="w-1/2">
            <ManaForm setOutput={setOutput} />
          </div>
          <div className="w-1/2 flex flex-col">
            <OutputBox text={output} />
          </div>
        </div>
      </main>
    </>
  );
}

export default App;
