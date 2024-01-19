import axios from 'axios';
import React, { useState, useEffect } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './App.css';
import { Header } from "./components/Header"
import PostCreator from "./components/Creator/createPost"
import { ViewPost, ViewThread } from "./components/Posts"
import { CardPost } from "./components/Home"

function App() {

  return (
    <div className="App">
      <Header></Header>
      <Routes>
        <Route path = "/posts" element = {<ViewThread/>}
          ></Route>
        <Route path = "/new" element = { <PostCreator/>}></Route>
      </Routes>
    </div>
  );
};

export default App;
