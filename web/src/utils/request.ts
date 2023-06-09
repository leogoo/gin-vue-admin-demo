import axios from 'axios';

const request = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 1000
});
export { request };