const API_ROOT = "https://localhost/api";

export const createAccount = async (data: URLSearchParams): Promise<Response> => (
    fetch(`${API_ROOT}/create_account`, {
        method: "POST",
        headers: {
            "Accept": "application/json",
            "Content-type": "application/x-www-form-urlencoded",
        },
        body: data,
}));
