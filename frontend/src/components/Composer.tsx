import {useState} from "react";
import {getAxiosClient} from "../api";

export default function Composer() {
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");

    // @ts-ignore
    async function handleSubmit(e) {
        e.preventDefault();
        try {
            const client = getAxiosClient();
            const res = await client.post("/posts", {title, content});
            if (res.status === 200) {

            } else {
                console.log("status code not 200");
            }
        } catch (err) {
            console.log("post failed", err)
        }

    }

    return <form onSubmit={handleSubmit}>
        <label htmlFor="id">Title:</label>
        <input type="text" value={title} id="title" onChange={(e) => {
            setTitle(e.target.value)
        }}/>
        <label htmlFor="content">Content:</label>
        <input type="text" value={content} id="content" onChange={e => {
            setContent(e.target.value)
        }}/>
        <button type="submit">Post</button>
    </form>
}
