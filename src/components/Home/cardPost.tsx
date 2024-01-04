import React from 'react'
import { AiFillLike } from 'react-icons/ai'
import { FaCommentAlt } from 'react-icons/fa'
import "./Home.css";


export const CardPost = () => {
  return (
    <div className="card-post">
      <span className="text-wrapper">HEADING</span>
      <span className="heading-2"> This is a post </span>
      <span className="info">Posted on by </span>
      <span className="likes"> <AiFillLike></AiFillLike> </span>
      <span className="comments"> <FaCommentAlt></FaCommentAlt> </span>
    </div>
  );
};


