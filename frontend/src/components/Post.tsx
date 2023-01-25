interface Creator {
    name: string
}

export interface PostProps {
    title: string,
    content: string,
    creator: Creator,
    updatedAt: string,
    createdAt: string
}

export default function Post(props: PostProps) {
    const {title, content, creator, updatedAt, createdAt} = props;
    return <div className="bg-gray-200">
        {title}
        {content}
        {creator.name}
        {updatedAt}
        {createdAt}
    </div>
}
