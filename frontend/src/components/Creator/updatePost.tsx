import {useState, useEffect} from 'react';
import ReactDOM from 'react-dom/client';
import axios from 'axios'; 
import {useParams} from 'react-router-dom';
import EditPost from './editPost';
import "./Creator.css"

interface Post {
    id: string;
    title: string;
    content: string;
}

export default function PostUpdate() {
    const [posts, setPosts] = useState<Post[]>([]);
    const { id } = useParams();

    const client = axios.create({
        baseURL: "http://localhost:4000"
     });

     useEffect(() => {
        const fetchPost = async () => {
            try {
                const response = await axios.get(`http://localhost:4000/posts/${id}`);
                setPosts(response.data);
            } catch (error) {
                console.error('Error fetching post:', error);
            }
        };
        fetchPost();
    }, [id]);
 
    const editPost = async (title: string, content: string) => {
        try {
            const response = await axios.put(`http://localhost:4000/posts/${id}`, {
                title: title,
                content: content,
            });

            setPosts(response.data);
        } catch (error) {
            console.error('Error editing post:', error);
        }
    };
   
  return (
      <EditPost editPost={editPost}/>
  )
}

