import {useState, useEffect} from 'react';
import ReactDOM from 'react-dom/client';
import axios from 'axios'; 
import AddPost from '../Creator/addPost';
import {ViewPost} from '../Posts/viewPost';
import { SelectCategory } from './selectCategory';
import "./Creator.css"

interface Post {
    id: string;
    title: string;
    content: string;
}

function PostCreator() {
    const [posts, setPosts] = useState<Post[]>([]);

    const client = axios.create({
        baseURL: "http://localhost:4000"
     });
 
    // GET with Axios
 
    // Delete with Axios
 
    // Post with Axios
    const addPost = async (title: any, content: any) => {
       let response = await client.post('/posts/new', {
          title: title,
          content: content,
       });
       setPosts((prevPosts) => [response.data, ...prevPosts]);
    }; 
   
  return (
    <main>
      <AddPost addPost={addPost}/>
   </main>
  )
}

export default PostCreator;
