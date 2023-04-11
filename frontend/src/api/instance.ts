import axios from "axios";

const BASE_URL = "http://192.168.1.20:8080";
const TIMEOUT_MS = 30000;

export const ServiceInstance = axios.create({
    baseURL: BASE_URL,
    timeout: TIMEOUT_MS,
});
