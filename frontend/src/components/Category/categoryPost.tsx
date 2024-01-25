import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, Link } from 'react-router-dom';
import { CardPost } from '../Home';
import "./Category.css";

interface Item {
    key_id: any;
    post_id: string;
    title: string;
    content: string;
    category: string;
}

export function CategorizedPosts({ value }: { value: string }) {
    const [items, setItems] = useState<Item[]>([]);
  
    const client = axios.create({
      baseURL: "http://localhost:8005"
    });

    useEffect(() => {
        const fetchPosts = async () => {
            const response = await client.get(`/posts`);
            setItems(response.data);
        }
        fetchPosts();
    }, []);    

    const filteredItems = items.filter((item: Item) => item.category === value);

    return (
        <div className="post-wrapper">
            {filteredItems.map((item) => 
                <Link to={`/posts/${item.post_id}`} key={item.key_id}>
                    <CardPost
                        id={item.post_id}
                        category={item.category}
                        title={item.title.substring(0,50)}
                        content={item.content.substring(0, 60) + "..."}>
                    </CardPost>
                </Link>
            )}
        </div>
    );
}
