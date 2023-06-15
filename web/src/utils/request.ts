import axios from 'axios';

const {VITE_BASE_API, VITE_BASE_PATH} = import.meta.env;
const request = axios.create({
  baseURL: `${VITE_BASE_PATH}${VITE_BASE_API}`,
  timeout: 1000
});
export { request };