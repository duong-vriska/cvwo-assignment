import {useState, useEffect} from 'react';
import axios from 'axios'
import "./Posts.css"
import { ViewPost } from "./viewPost"
import { ViewComment } from "./viewComment"

interface Post {
    id: string;
    title: string;
    content: string;
}

export const ViewThread = () => {

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
        <div className="thread-wrapper">
            {posts.map((post) => 
            <ViewPost
            key={post.id} 
            id={post.id}
            title={post.title} 
            content={post.content} 
            />
            )}
            <ViewComment></ViewComment>
        </div>
    )
}