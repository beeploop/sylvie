import { useLoaderData } from "react-router"

export default function Playback() {
    const data = useLoaderData();

    return (
        <div>
            <h1>Playback page</h1>
            <p>{data.id}</p>
        </div>
    )
}
