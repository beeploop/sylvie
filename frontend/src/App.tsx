import { RouterProvider } from "react-router/dom"
import { createBrowserRouter } from "react-router"
import Layout from "./components/Layout";
import Homepage from "./pages/Homepage";
import Playback from "./pages/Playback";

const router = createBrowserRouter([
    {
        path: "/",
        Component: Layout,
        children: [
            {
                index: true, Component: Homepage,
            },
            {
                path: "playback/:id",
                loader: async ({ params }) => {
                    const id = params.id;
                    return { id: id };
                },
                Component: Playback,
            }
        ],
    }
]);

export default function App() {
    return (
        <RouterProvider router={router} />
    )
}
