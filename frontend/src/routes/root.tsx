import {useEffect, useState} from "react";
import {getAxiosClient} from "../api";
import Post, {PostProps} from "../components/Post";
import Composer from "../components/Composer";

export default function Root() {
    const [posts, setPosts] = useState<PostProps[]>([]);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const client = getAxiosClient();
                const res = await client.get("/posts");
                if (res.data === null) {
                    setPosts([]);
                } else {
                    setPosts(res.data);
                }
                setIsLoading(false);
            } catch (err) {
                console.error(err);
            }
        }
        fetchData();
    }, [])
    return (
        <>
            {isLoading ? (
                <div>Loading...</div>
            ) : (
                <div>
                    <Composer/>
                    {posts.map(post => (
                        <Post
                            title={post.title}
                            content={post.content}
                            creator={post.creator}
                            updatedAt={post.updatedAt}
                            createdAt={post.createdAt}/>
                    ))}
                </div>
            )}
        </>
    );
}
