const API_ROOT_V1 = "https://localhost/api/v1";

export const createAccount = async (data: URLSearchParams): Promise<Response> => (
    fetch(`${API_ROOT_V1}/create_account`, {
        method: "POST",
        headers: {
            "Accept": "application/json",
            "Content-type": "application/x-www-form-urlencoded",
        },
        body: data,
}));

export const loginUser = async (data: URLSearchParams): Promise<Response> => (
    fetch(`${API_ROOT_V1}/login`, {
        method: "POST",
        // allow cookies to attach
        credentials: "include",
        headers: {
            "Accept": "application/json",
            "Content-type": "application/x-www-form-urlencoded",
        },
        body: data,
}));
