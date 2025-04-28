import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import Header from "./components/Header";
import ManaForm from "./components/ManaForm";

function App() {
  return (
    <>
      <Header />
      <main>
        <ManaForm />
      </main>
    </>
  );
}

export default App;
