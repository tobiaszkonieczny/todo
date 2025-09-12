import http from "k6/http";
import { check, sleep } from "k6";



//HOW TO USE:
//k6 run --insecure-skip-tls-verify k6_test.js

export const options = {
    vus: 20,          //20 users
    duration: "5s",  //test duration
};

const BASE_URL = "https://127.0.0.1"
const USER = { username: "testuser", password: "testpassword123" };

export default function () {
    // 1. Login
    let loginRes = http.post(`${BASE_URL}/auth/login`, JSON.stringify(USER), {
        headers: { "Content-Type": "application/json" },
    });
    check(loginRes, { "login success": (r) => r.status === 200 });

    let token = loginRes.json("token");

    // 2. GET /tasks
    let tasksRes = http.get(`${BASE_URL}/tasks/`, {
        headers: { Authorization: `Bearer ${token}` },
    });
    check(tasksRes, { "get tasks ok": (r) => r.status === 200 });

    // 3. POST /tasks/new
    let newTask = http.post(
        `${BASE_URL}/tasks/new`,
        JSON.stringify({ title: "task123456 " }),
        {
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${token}`,
            },
        }
    );
    check(newTask, { "task created": (r) => r.status === 201 });

    sleep(1);
}