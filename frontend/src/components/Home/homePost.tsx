import {useState, useEffect} from 'react';
import axios from 'axios'
import { CardPost } from "./cardPost"
import "./Home.css"

interface Post {
    id: string;
    title: string;
    content: string;
}

export const HomePosts = () => {

    const [posts, setPosts] = useState<Post[]>([]);
  
    const client = axios.create({
      baseURL: "http://localhost:4000"
    });

    const fetchPosts = async() => {
        const response = await client.get('/posts');
        setPosts(response.data);
    }

    useEffect(() => {
        fetchPosts()
     }, []);    

    return (
        <div className="post-wrapper">
            {posts.map((post) => 
            <CardPost
            key={post.id} 
            id={post.id}
            title={post.title} 
            content={post.content} 
            />
            )}
        </div>
    )
}