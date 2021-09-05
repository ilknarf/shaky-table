const API_ROOT = "https://localhost/api";

export const createAccount = async (data: FormData): Promise<Response> => (
    fetch(`${API_ROOT}/create_account`, {
        method: "POST",
        headers: {
            "Content-type": "application/x-www-form-urlencoded",
        },
        body: data,
}));