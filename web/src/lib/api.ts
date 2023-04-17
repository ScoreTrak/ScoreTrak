import axios, {AxiosInstance} from "axios";

export const api: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_SERVER_URL || "http://localhost:3000",
  timeout: 3000,
  withCredentials: true
})

