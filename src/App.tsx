import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import Header from "./components/Header";
import ManaForm from "./components/ManaForm";
import OutputBox from "./components/OutputBox";

function App() {
  return (
    <>
      <Header />
      <main>
        <div className="flex justify-center items-center">
          <ManaForm />
          <OutputBox />
        </div>
      </main>
    </>
  );
}

export default App;
