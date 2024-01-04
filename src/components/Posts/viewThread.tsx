import "./Posts.css"
import { ViewPost } from "./viewPost"
import { ViewComment } from "./viewComment"

export const ViewThread = () => {
    return (
        <div className="thread-wrapper">
            <ViewPost></ViewPost>
            <ViewComment></ViewComment>
        </div>
    )
}