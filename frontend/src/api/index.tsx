import axios from "axios";
import {BACKEND_URL} from "../constants";

export const getAxiosClient = () => {
    const token = localStorage.getItem("token");
    if (!token) {
        return axios.create({baseURL: BACKEND_URL});
    } else {
        return axios.create({baseURL: BACKEND_URL, headers: {"Authorization": `Bearer ${token}`}});
    }
}

