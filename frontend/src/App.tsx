import {useState} from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import createClient from "openapi-fetch";
import {paths} from "./api/schema";
import {useQuery} from "@tanstack/react-query";

function App() {
    const [count, setCount] = useState(0)

    const client = createClient<paths>({baseUrl: "http://localhost:3000/"});

    const query = useQuery({
        queryKey: ['getArticles'], queryFn: async ({signal}) => {
            const {data} = await client.GET("/articles", {
                params: {},
                signal
            });
            return data;
        }
    })


    if (query.isFetching) {
        return "Loading...";
    }

    if (query.error) {
        return "Error: " + query.error;
    }

    return (
        <>
            <div>
                <a href="https://vitejs.dev" target="_blank">
                    <img src={viteLogo} className="logo" alt="Vite logo"/>
                </a>
                <a href="https://react.dev" target="_blank">
                    <img src={reactLogo} className="logo react" alt="React logo"/>
                </a>
            </div>
            <h1>Vite + React</h1>
            <div className="card">
                <button onClick={() => setCount((count) => count + 1)}>
                    count is {count}
                </button>
                <p>
                    Edit <code>src/App.tsx</code> and save to test HMR
                </p>
            </div>
            <p className="read-the-docs">
                Click on the Vite and React logos to learn more
            </p>
            <>
                {JSON.stringify(query.data,undefined, 2)}
            </>
        </>
    )
}

export default App
