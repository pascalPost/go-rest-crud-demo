import {useState} from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import createClient from "openapi-fetch";
import {paths} from "./oapi/schema";
import {useQuery} from "@tanstack/react-query";

function App() {
    const [count, setCount] = useState(0)

    const client = createClient<paths>({baseUrl: "http://localhost:8011/"});

    const query = useQuery<string, Error>({
        queryKey: ['test'], queryFn: () => {
            return client.GET("/hello/{name}", {
                params: {
                    path: {name: "Pascal"},
                    query: {locale: "en-US"},
                },
            }).then((res) => {
                    return res.data?.message ?? "";
                }
            );
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
                {query.data}
            </>
            {/*<form onSubmit={}>*/}
            {/*    <label>*/}
            {/*        Name:*/}
            {/*        <input type="text" name="name"/>*/}
            {/*    </label>*/}
            {/*    <input type="submit" value="Submit"/>*/}
            {/*</form>*/}
        </>
    )
}

export default App
