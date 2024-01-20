import {useState, useEffect} from 'react';
import axios from 'axios'; 
import {AddPost} from './addPost';
import "./Creator.css"

interface Post {
    id: string;
    title: string;
    content: string;
}

export function PostCreator() {
    const [posts, setPosts] = useState<Post[]>([]);

    const client = axios.create({
        baseURL: "http://localhost:4000"
     });
 
    const addPost = async (title: any, content: any) => {
       let response = await client.post('/posts/new', {
          title: title,
          content: content,
       });
       setPosts((prevPosts) => [response.data, ...prevPosts]);
    }; 
   
  return (
      <AddPost addPost={addPost}/>
  )
}

