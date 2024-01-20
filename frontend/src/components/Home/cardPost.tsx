import React from 'react'
import { AiFillLike } from 'react-icons/ai'
import { FaCommentAlt } from 'react-icons/fa'
import "./Home.css";

export const CardPost = (props: any) => {
  return (
    <div className="card-post">
      <span className="title">{props.title}</span>
      <span className="content"> {props.content} </span>
      <span className="info"> Post {props.id} </span>
      <span className="likes"> <AiFillLike></AiFillLike> </span>
      <span className="comments"> <FaCommentAlt></FaCommentAlt> </span>
    </div>
  );
};


