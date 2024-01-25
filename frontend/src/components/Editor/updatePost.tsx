import {useState, useEffect} from 'react';
import ReactDOM from 'react-dom/client';
import axios from 'axios'; 
import {useParams, useNavigate} from 'react-router-dom';
import EditPost from './editPost';
import "./Editor.css"

interface Post {
    id: string;
    title: string;
    content: string;
    category: string;
}

function useRequiredParams<T extends Record<string, any>>() {
    const params = useParams<T>();
    return params as T;
}

export default function PostUpdate() {
    const [posts, setPosts] = useState<Post[]>([]);
    let {id} = useRequiredParams<{id: string}>();
    let navigate = useNavigate();

    const client = axios.create({
        baseURL: "http://localhost:8005"
     });

     useEffect(() => {
        const fetchPost = async () => {
            try {
                const response = await client.get(`/posts/${id}`);
                setPosts(response.data);
            } catch (error) {
                console.error('Error fetching post:', error);
            }
        };
        fetchPost();
    }, [id]);
 
    const editPost = async (title: string, content: string, category: string) => {
        try {
            const response = await client.put(`/posts/${id}`, {
                post_id: id, 
                title: title,
                content: content,
                category: category 
            });
            if (response.status === 200) {
                console.log("Post updated!");
                setTimeout(() => {
                navigate(`/posts/${id}`);
                }, 10); 
            } 
        } catch (error) {
            console.error('Error editing post:', error);
        }
    };
   
  return (
      <EditPost editPost={editPost}/>
  )
}

