import {useState, useEffect} from 'react';
import {Link} from 'react-router-dom';
import axios from 'axios'
import { CardPost } from "./cardPost"
import "./Home.css"
import { CategoryButtons } from "../Category/buttonCategory"

interface Post {
    key_id: any;
    post_id: string;
    title: string;
    content: string;
    category: string; 
}

export const HomePosts = () => {

    const [posts, setPosts] = useState<Post[]>([]);
  
    const client = axios.create({
      baseURL: "http://localhost:8005"
    });

    const fetchPosts = async() => {
        const response = await client.get(`/posts`);
        setPosts(response.data);
    }

    useEffect(() => {
        fetchPosts()
     }, []);    

    return (
        <div className="post-wrapper">
            <CategoryButtons/>
            {posts.map((post) => 
            <Link to={`/posts/${post.post_id}`}>
                <CardPost
                    key={post.key_id}
                    id={post.post_id}
                    category={post.category}
                    title={post.title.substring(0,50)}
                    content={post.content.substring(0, 60) + "..."}>
                </CardPost>
                </Link>)}
        </div>
    )
}