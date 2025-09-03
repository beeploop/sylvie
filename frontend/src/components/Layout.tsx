import { Outlet } from "react-router";

export default function Layout() {
    return (
        <div>
            <div>
                <h1>Header</h1>
            </div>
            <Outlet />
        </div>
    )
}
