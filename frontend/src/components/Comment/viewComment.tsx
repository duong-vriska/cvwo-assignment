import React from 'react'
import "./Comment.css"
import { CommentCreator } from '../Creator/index'
import { Comments } from './singleComments'

export const ViewComment = () => {
    return (
    <div className = "comment-wrapper">
        <CommentCreator></CommentCreator>
        <Comments></Comments>
    </div>
    )
}