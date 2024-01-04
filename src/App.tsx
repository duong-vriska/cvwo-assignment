import React, { useState } from 'react';
import './App.css';
import { Header } from "./components/Header"
import { CardPost } from "./components/Home"

const App: React.FC = () => {
  return (
    <div className="App">
      <Header></Header>
      <CardPost></CardPost>
    </div>
  );
};



export default App;
