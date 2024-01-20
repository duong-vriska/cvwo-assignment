import {useState, useEffect} from 'react'
import axios from 'axios'
import { ViewPost } from "../Posts/viewPost"
import { ViewComment } from "../Comment/viewComment"
import { useParams } from 'react-router-dom'

interface Post {
    id: string;
    title: string;
    content: string;
}

export const ViewThread = () => {

    const { id } = useParams();

    const [posts, setPosts] = useState<Post[]>([]);
  
    const client = axios.create({
      baseURL: `http://localhost:4000`
    });

    const fetchPosts = async() => {
        const response = await client.get(`/posts/${id}`);
        setPosts(response.data);
    }

    useEffect(() => {
        fetchPosts()
     }, [id]);

    const deletePost = async(id: string) => {
        const response = await client.delete(`/posts/${id}`);
        setPosts(posts.filter((post) => {
          return post.id !== id;
        }))
    };

    return (
        <div className="thread-wrapper">
            {posts.map((post) =>
            <ViewPost
            key={post.id} 
            id={post.id}
            title={post.title} 
            content={post.content} 
            deletePost={deletePost}
            />
            )}
            <ViewComment></ViewComment>
        </div>
    )
}