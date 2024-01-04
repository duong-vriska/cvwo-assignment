import React, { useState } from 'react';
import './App.css';
import { Header } from "./components/Header"
import { ViewThread } from "./components/Posts"
import { CardPost } from "./components/Home"
import { Article } from "./model";

const App: React.FC = () => {

  const [article, setArticle] = useState<string>("");
  const [articles, setArticles] = useState<Article[]>([]);

  const handleAdd = (e: React.FormEvent) => {
    e.preventDefault(); 
    if (article){
      setArticles([...articles,{ 
        id:Date.now(), 
        article,
        isPublic:false }])
        setArticle("");
    }
  }

  return (
    <div className="App">
      <Header></Header>
      <ViewThread></ViewThread>
    </div>
  );
};



export default App;
