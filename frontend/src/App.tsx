import axios from 'axios';
import React, { useState, useEffect } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './App.css';
import { Header } from "./components/Header"
import { PostCreator } from "./components/Creator/index"
import { ViewThread } from "./components/Thread/index"
import { CardPost } from "./components/Home"
import { HomePosts } from "./components/Home/index"
import PostUpdate from "./components/Editor/updatePost"

function App() {

  return (
    <div className="App">
      <Header></Header>
      <Routes>
        <Route path = "/" element = {<HomePosts/>}></Route>
        <Route path = "/posts" element = {<HomePosts/>}></Route>
        <Route path = "/posts/:id" element = {<ViewThread/>}></Route>
        <Route path = "new" element = { <PostCreator/>}></Route>
        <Route path = "posts/:id/edit" element = {<PostUpdate/>}></Route>
      </Routes>
    </div>
  );
};

export default App;
