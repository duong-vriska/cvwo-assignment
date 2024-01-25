import {useState, useEffect} from 'react'
import axios from 'axios'
import { ViewPost } from "../Posts/viewPost"
import { ViewComment } from "../Comment/viewComment"
import { useNavigate, useParams } from 'react-router-dom'

interface Post {
    key_id: any;
    post_id: string;
    title: string;
    content: string;
    category: string;
}

export const ViewThread = () => {

    const {id} = useParams();
    const navigate = useNavigate(); 

    const [posts, setPosts] = useState<Post[]>([]);
  
    const client = axios.create({
      baseURL: `http://localhost:8005`
    });

    const fetchPosts = async() => {
        const response = await client.get(`/posts/${id}`);
        const postData = response.data;
        const postsArray = [postData];
        setPosts(postsArray);
    }

    useEffect(() => {
        fetchPosts()
     }, [id]);

    const deletePost = async(post_id: string) => {
        const response = await client.delete(`/posts/${id}`);
        setPosts(posts.filter((post) => {
          return post.post_id !== id;
        }))
    };

    const editPost = async(post_id: string) => {
        navigate(`/posts/${id}/edit`);
    }

    return (
        <div className="thread-wrapper">
            {(posts) && posts.map((post) =>
            <ViewPost
            key={post.key_id}
            id={post.post_id}
            category={post.category}
            title={post.title} 
            content={post.content} 
            deletePost={deletePost}
            editPost={editPost}
            />
            )}
        </div>
    )
}