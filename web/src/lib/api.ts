import axios, {AxiosRequestConfig} from "axios";
import {AxiosConfig} from "./scoretrak-queries";

const axiosConfig: AxiosRequestConfig & AxiosConfig = {
  baseURL: import.meta.env.VITE_API_SERVER_URL || "http://localhost:3000",
  timeout: 3000,
  withCredentials: true,
}

export const scoretrakAxiosInstance = axios.create(axiosConfig)

