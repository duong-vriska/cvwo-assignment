import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './App.css';
import { Header } from "./components/Header"
import { PostCreator } from "./components/Creator/index"
import { ViewThread } from "./components/Thread/index"
import { HomePosts } from "./components/Home/index"
import PostUpdate from "./components/Editor/updatePost"
import { CategorizedPosts } from "./components/Category/categoryPost"
import { LoginBox } from "./components/Login/loginBox"

function App() {

  return (
    <div className="App">
      <Header></Header>
      <Routes>
        <Route path = "/" element = {<HomePosts/>}></Route>
        <Route path = "/login" element = {<LoginBox/>}></Route>
        <Route path = "/posts"> 
          <Route index element = {<HomePosts/>}/>
          <Route path = ":id" element = {<ViewThread/>}/>
          <Route path = ":id/edit" element = {<PostUpdate/>}/>
        </Route>
        <Route path = "new" element = { <PostCreator/>}></Route>
        {categories.map(category => (
          <Route
            key={category.path}
            path={`/posts/category/${category.value}`}
            element={<CategorizedPosts value={category.path} />}/>
        ))} 
      </Routes>
    </div>
  );
};

const categories = [
  { path: 'News', value: 'news' },
  { path: 'Entertainment', value: 'entertainment' },
  { path: 'Food', value: 'food' },
  { path: 'Study', value: 'study' },
  { path: 'Sports', value: 'sports' },
];

export default App;
