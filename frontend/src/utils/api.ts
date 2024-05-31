import axios from "axios";

const apiUrl = import.meta.env.VITE_API_URL;

export const API = axios.create({
  baseURL: apiUrl,
  withCredentials: false,
});

export default apiUrl;
