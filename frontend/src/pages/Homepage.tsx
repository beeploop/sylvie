import { Link } from "react-router";

export default function Homepage() {
    return (
        <div>
            <h1>Hello world</h1>
            <Link to={"/playback/foobar"}>playback</Link>
        </div>
    )
}
