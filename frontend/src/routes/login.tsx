import axios from "axios";
import {BACKEND_URL} from "../constants";
import {useState} from "react";


export default function Login() {
    const [name, setName] = useState<string>("");

    // @ts-ignore
    async function handleSubmit(e) {
        e.preventDefault();
        try {
            const tokenResponse = await axios.post(`${BACKEND_URL}/auth`, {name});
            if (tokenResponse.status === 200) {
                localStorage.setItem("token", tokenResponse.data.token);
            } else {
                console.log("status code not 200");
            }
        } catch (err) {
            console.log("login failed", err)
        }

    }

    return <div>
        <form onSubmit={handleSubmit}>
            <input type="text" placeholder="Username" value={name} onChange={(e) => setName(e.target.value)}/>
            <button type="submit">Log in</button>
        </form>
    </div>
}
