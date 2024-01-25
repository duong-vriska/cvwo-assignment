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
import { CategoryButtons } from "./components/Category/buttonCategory"
import { CategorizedPosts } from "./components/Category/categoryPost"
import { LoginBox } from "./components/Login/loginBox"

function App() {

  return (
    <div className="App">
      <Header></Header>
      <Routes>
        <Route path = "/" element = {<HomePosts/>}></Route>
        <Route path = "/login" element = {<LoginBox/>}></Route>
        <Route path = "/posts" element = {<HomePosts/>}></Route>
        <Route path = "/posts/:id" element = {<ViewThread/>}></Route>
        <Route path = "/posts/:id/edit" element = {<PostUpdate/>}></Route>
        <Route path = "new" element = { <PostCreator/>}></Route>
        <Route path= "/posts/category/news" element={CategorizedPosts({ value: "news" })}></Route>
        <Route path= "/posts/category/entertainment" element={CategorizedPosts({ value: "entertainment" })}></Route>
        <Route path = "/posts/category/food" element={CategorizedPosts({ value: "food" })}></Route>
        <Route path = "/posts/category/study" element={CategorizedPosts({ value: "study" })}></Route> 
        <Route path = "/posts/category/sports" element={CategorizedPosts({ value: "sports" })}></Route>
      </Routes>
    </div>
  );
};

export default App;
