import * as axios from 'axios';

const instance: axios.AxiosInstance = axios.default.create({
    baseURL: import.meta.env.VITE_API_ENDPOINT,
    timeout: 1000,
    headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    },
});

export { instance };
