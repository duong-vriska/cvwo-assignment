import {useState, useEffect} from 'react';
import axios from 'axios'; 
import {AddPost} from './addPost';
import {useParams, useNavigate} from 'react-router-dom';
import "./Creator.css"

interface Post {
    id: string;
    title: string;
    content: string;
    category: string; 
}

export function PostCreator() {
    const [posts, setPosts] = useState<Post[]>([]);
    const { id } = useParams();
    const navigate = useNavigate();

    const client = axios.create({
        baseURL: "http://localhost:8005"
     });

    useEffect(() => {
        const fetchPosts = async () => {
            try {
                const response = await client.get('/posts');
                setPosts(response.data);
            } catch (error) {
                console.error('Error fetching posts:', error);
            }
        };
        fetchPosts();
    }, []);
 
    const addPost = async (title: any, content: any, category: any) => {
       let response = await client.post('/posts/new', {
          title: title,
          content: content,
          category: category,
       });
       setPosts((prevPosts) => [response.data, ...prevPosts]);
    }; 
   
  return (
      <AddPost addPost={addPost}/>
  )
}

